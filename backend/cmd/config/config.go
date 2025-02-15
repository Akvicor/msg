package config

import (
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/Akvicor/util"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path"
)

var Global *Model
var FileData []byte

func Load(p string) {
	if util.FileStat(p).NotFile() {
		glog.Fatal("missing config [%s]!", p)
	}

	// Load config.toml
	data, err := os.ReadFile(p)
	if err != nil {
		glog.Fatal("failed to read file %v", err)
	}
	FileData = data
	Global = new(Model)
	err = toml.Unmarshal(FileData, Global)
	if err != nil {
		glog.Fatal("failed to unmarshal file %v", err)
	}

	// Set mask
	mask := uint32(0)
	for _, v := range Global.Log.Mask {
		switch v {
		case "unknown":
			mask |= glog.MaskUNKNOWN
		case "debug":
			mask |= glog.MaskDEBUG
		case "trace":
			mask |= glog.MaskTRACE
		case "info":
			mask |= glog.MaskINFO
		case "warning":
			mask |= glog.MaskWARNING
		case "error":
			mask |= glog.MaskERROR
		case "fatal":
			mask |= glog.MaskFATAL
		}
	}
	glog.SetMask(mask)

	// Set flag
	flg := uint32(0)
	for _, v := range Global.Log.Flag {
		switch v {
		case "date":
			flg |= glog.FlagDate
		case "time":
			flg |= glog.FlagTime
		case "long_file":
			flg |= glog.FlagLongFile
		case "short_file":
			flg |= glog.FlagShortFile
		case "func":
			flg |= glog.FlagFunc
		case "prefix":
			flg |= glog.FlagPrefix
		case "suffix":
			flg |= glog.FlagSuffix
		}
	}
	glog.SetFlag(flg)

	// Set file
	if Global.Log.EnableFile {
		logDir := path.Dir(Global.Log.File)
		if util.FileStat(logDir).NotExist() {
			if err = os.MkdirAll(logDir, os.ModePerm); err != nil {
				glog.Fatal("failed to set log file [%s]", err.Error())
			}
		}
		err = glog.SetLogFile(Global.Log.File)
		if err != nil {
			glog.Fatal("failed to set log file [%s]", err.Error())
		}
	}

	keyLen := len([]byte(Global.Encrypt.Key))
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		glog.Fatal("wrong key length, must 16/24/32")
	}
	if len([]byte(Global.Encrypt.Iv)) != 16 {
		glog.Fatal("wrong iv length, must 16")
	}
}

func GenerateExample(basePath string) {
	a := &Model{
		AppName: "msg",
		Debug:   false,
		Server: ServerModel{
			Domain:      "example.com",
			BaseUrl:     "https://example.com",
			HttpIp:      "0.0.0.0",
			HttpPort:    3000,
			WebPath:     "build",
			EnableHttps: false,
			CrtFile:     path.Join(basePath, "cert/example.com.crt"),
			KeyFile:     path.Join(basePath, "cert/example.com.key"),
		},
		Database: DatabaseModel{
			Type:     "sqlite",
			File:     path.Join(basePath, "msg.db"),
			Host:     "db",
			Port:     5432,
			Database: "msg",
			Username: "postgres",
			Password: "password",
		},
		Encrypt: EncryptModel{
			Key: util.RandomStringAtLeastOnce(32),
			Iv:  util.RandomStringAtLeastOnce(16),
		},
		Log: LogModel{
			EnableFile: false,
			File:       path.Join(basePath, "msg.log"),
			Mask:       []string{"unknown", "debug", "trace", "info", "warning", "error", "fatal"},
			Flag:       []string{"date", "time", "short_file", "prefix", "suffix"},
			Debug:      []string{"database", "echo"},
		},
		Bot: BotModel{
			Maid: &BotSenderModel{
				SMS: &SMSBotModel{
					EnableReceiver: false,
					EnableSender:   false,
					Api:            "https://example.com",
					Token:          "",
				},
				Mail: &MailBotModel{
					EnableImap: false,
					HostImap:   "https://imap.example.com",
					PortImap:   993,
					EnableSmtp: false,
					HostSmtp:   "https://smtp.example.com",
					PortSmtp:   465,
					From:       "MSG Backend <maid@example.com>",
					Username:   "maid@example.com",
					Password:   "",
				},
				Telegram: &TelegramBotModel{
					EnableReceiver: false,
					EnableSender:   false,
					API:            "https://api.telegram.org/bot%s/%s",
					Token:          "",
				},
				Wechat: &WechatBotModel{
					EnableReceiver: false,
					EnableSender:   false,
					CorpId:         "",
					Secret:         "",
					AgentId:        0,
					Token:          "",
					AesKey:         "",
				},
			},
			Reminder: &BotSenderModel{
				Wechat: &WechatBotModel{
					EnableReceiver: false,
					EnableSender:   false,
					CorpId:         "",
					Secret:         "",
					AgentId:        0,
					Token:          "",
					AesKey:         "",
				},
			},
		},
		Cron: CronModel{
			Debug: &CronDebugModel{
				Enable: false,
				Hour:   0,
				Minute: 0,
				Second: 0,
			},
		},
	}
	data, err := toml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
