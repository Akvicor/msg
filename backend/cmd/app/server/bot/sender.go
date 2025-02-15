package bot

import (
	"context"
	"github.com/Akvicor/glog"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/common/types/channel"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
	"sync"
	"time"
)

type SenderCommon interface {
	Key() string                                                  // 名称
	Name() string                                                 // 名称
	Status(p channel.Type) (sender, receiver status.SenderStatus) // 查看当前Sender的某个渠道的状态
	Send(p channel.Type, target string, msg *model.Send) error    // 发送通知
}

type SenderModel struct {
	bot   *Model
	queue *queueModel
	new   chan struct{}
}

func newSender() *SenderModel {
	return &SenderModel{
		bot:   Bot,
		queue: nil,
		new:   make(chan struct{}, 4),
	}
}

var Sender = newSender()

func (b *SenderModel) Run(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	queue, err := newQueue()
	if err != nil {
		glog.Warning("queue error %v", err)
		return
	}
	b.queue = queue

	b.bot.Wait()

	bwg := &sync.WaitGroup{}

	glog.Info("Sender OK")
	wait := time.Second
	for {
		select {
		case <-ctx.Done():
			close(b.new)
			bwg.Wait()
			glog.Info("Sender exist")
			return
		case <-b.new: // 手动触发
		case <-time.After(wait): // 等待下一个信息的发送时间
		}
		wait = time.Hour
		msg := b.queue.Peek()
		if msg == nil {
			continue
		}

		timestamp := time.Now().Unix()
		if msg.SendAt > timestamp {
			wait = time.Second * time.Duration(msg.SendAt-timestamp)
			continue
		}

		{ // 发送消息
			bwg.Add(1)
			tmsg := b.queue.Pop()
			go func(msg *model.Send) {
				defer bwg.Done()
				if msg == nil {
					return
				}
				var err error
				msg.Channel, err = service.Channel.FindByUID(true, nil, msg.UID, msg.ChannelID)
				if err != nil {
					glog.Warning("find channel error: UID[%d] CID[%d] %v", msg.UID, msg.ChannelID, err)
					return
				}
				retry := 3
				for retry > 0 {
					retry--
					err = b.bot.Send(msg.Channel.Bot, msg.Channel.Type, msg.Channel.Target, msg)
					if err != nil {
						glog.Warning("send sms channel[%d] msg[%d] failed: %v", msg.Channel.ID, msg.ID, err)
						time.Sleep(3 * time.Second)
						continue
					}
					err = service.Send.UpdateSentByIDFinished(msg.ID)
					if err != nil {
						glog.Warning("send sms channel[%d] msg[%d] update finished failed: %v", msg.Channel.ID, msg.ID, err)
					}
					break
				}
				if err != nil {
					err = service.Send.UpdateSentByIDFailed(msg.ID, err.Error())
					if err != nil {
						glog.Warning("send sms channel[%d] msg[%d] update failed failed: %v", msg.Channel.ID, msg.ID, err)
					}
				}
			}(tmsg)
		}

		msg = b.queue.Peek()
		if msg == nil {
			continue
		}
		wait = time.Second * time.Duration(msg.SendAt-time.Now().Unix())
	}
}

func (b *SenderModel) Send(uid int64, channelID int64, sendAt int64, ip string, cType send.Type, title string, msg string) (int64, error) {
	defer func() {
		b.new <- struct{}{}
	}()
	item, err := b.queue.PushNew(uid, channelID, sendAt, ip, cType, title, msg)
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (b *SenderModel) Cancel(uid int64, sendID int64) error {
	defer func() {
		b.new <- struct{}{}
	}()
	item := b.queue.Remove(uid, sendID)
	if item == nil {
		return ErrorMsgNotFound
	}
	if item.ID == 0 {
		return ErrorMsgNotFound
	}
	err := service.Send.UpdateSentByIDCancel(item.ID)
	if err != nil {
	}
	return nil
}
