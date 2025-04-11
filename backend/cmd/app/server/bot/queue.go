package bot

import (
	"container/heap"
	"context"
	"errors"
	"github.com/Akvicor/glog"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/types/send"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/repository"
	"msg/cmd/app/server/service"
	"sync"
	"time"
)

/*
Item
*/

type queueHeapItemModel struct {
	index    int   // 在堆中的索引
	id       int64 // 唯一标识符 -> send.ID
	priority int64 // 元素的优先级 -> send.SendAt
	value    *model.Send
}

func newQueueHeapItem(value *model.Send) *queueHeapItemModel {
	return &queueHeapItemModel{
		index:    -1,
		id:       value.ID,
		priority: value.SendAt,
		value:    value,
	}
}

/*
Heap
*/

type queueHeapModel struct {
	items []*queueHeapItemModel
	index map[int64]*queueHeapItemModel // item.id -> item
}

func newQueueHeap() *queueHeapModel {
	return &queueHeapModel{
		items: make([]*queueHeapItemModel, 0, 16),
		index: make(map[int64]*queueHeapItemModel),
	}
}

func (q *queueHeapModel) Len() int {
	return len(q.items)
}

func (q *queueHeapModel) Less(i, j int) bool {
	// 小顶堆，优先级数字越小越优先
	return q.items[i].priority < q.items[j].priority
}

func (q *queueHeapModel) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
	q.items[i].index = i
	q.items[j].index = j
}

func (q *queueHeapModel) Push(x any) {
	n := len(q.items)
	item := x.(*queueHeapItemModel)
	item.index = n
	q.items = append(q.items, item)
	q.index[item.id] = item
}

func (q *queueHeapModel) Pop() any {
	n := len(q.items)
	old := q.items
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // 标记已移除
	q.items = old[:n-1]
	delete(q.index, item.id)
	return item
}

// Peek 查看队首元素，但不删除
func (q *queueHeapModel) Peek() *queueHeapItemModel {
	// 小顶堆的顶部元素就是最小优先级的元素
	return q.items[0]
}

// Remove 从优先队列中移除指定 ID 的元素
func (q *queueHeapModel) Remove(uid, id int64) *queueHeapItemModel {
	item, exists := q.index[id]
	if !exists {
		glog.Warning("Remove: !exists")
		return nil
	}
	if item.value.UID != uid {
		glog.Warning("Remove: item.value.UID != uid")
		return nil
	}
	return heap.Remove(q, item.index).(*queueHeapItemModel)
}

// Update 更新元素的优先级
func (q *queueHeapModel) Update(id int64, newValue *model.Send, newPriority int64) {
	item, exists := q.index[id]
	if !exists {
		return
	}
	item.value = newValue
	item.priority = newPriority
	heap.Fix(q, item.index)
}

/*
Queue
*/

type queueModel struct {
	lock  *sync.RWMutex
	queue *queueHeapModel
}

func newQueue() (*queueModel, error) {
	que := &queueModel{
		lock:  &sync.RWMutex{},
		queue: newQueueHeap(),
	}
	heap.Init(que.queue)
	// 从数据库初始化
	sends, err := repository.Send.FindAllPending(context.Background(), nil, false, model.NewPreloaderSend())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		for _, item := range sends {
			glog.Debug("init queue: ID[%d] UID[%d]", item.ID, item.UID)
			que.Push(item)
		}
	}
	return que, nil
}

func (q *queueModel) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.queue.Len()
}

func (q *queueModel) Peek() *model.Send {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.queue.Len() == 0 {
		return nil
	}
	item := q.queue.Peek()
	return item.value
}

func (q *queueModel) Push(item *model.Send) {
	q.lock.Lock()
	defer q.lock.Unlock()
	heap.Push(q.queue, newQueueHeapItem(item))
}

func (q *queueModel) PushNew(uid, channelID, scheduleID int64, sendAt int64, ip string, cType send.Type, title, msg string) (*model.Send, error) {
	// 写入数据库
	item, err := service.Send.Create(time.Now().Unix(), sendAt, uid, channelID, scheduleID, ip, cType, title, msg)
	if err != nil {
		return nil, err
	}
	// 放入队列
	q.Push(item)
	return item, nil
}

func (q *queueModel) Pop() *model.Send {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.queue.Len() == 0 {
		return nil
	}
	item := heap.Pop(q.queue).(*queueHeapItemModel)
	return item.value
}

func (q *queueModel) Remove(uid, id int64) *model.Send {
	q.lock.Lock()
	item := q.queue.Remove(uid, id)
	q.lock.Unlock()
	if item == nil {
		return nil
	}
	return item.value
}
