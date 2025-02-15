package config

type Model struct {
	AppName  string        `toml:"app-name"`
	Debug    bool          `toml:"debug"`
	Server   ServerModel   `toml:"server"`
	Database DatabaseModel `toml:"database"`
	Encrypt  EncryptModel  `toml:"encrypt"`
	Log      LogModel      `toml:"log"`
	Bot      BotModel      `toml:"bot"`
	Cron     CronModel     `toml:"cron"`
}

type ServerModel struct {
	Domain      string `toml:"domain"`
	BaseUrl     string `toml:"base-url"`
	HttpIp      string `toml:"http-ip"`
	HttpPort    int    `toml:"http-port"`
	WebPath     string `toml:"web-path"`
	EnableHttps bool   `toml:"enable-https"`
	CrtFile     string `toml:"crt-file"`
	KeyFile     string `toml:"key-file"`
}

type DatabaseModel struct {
	Type     string `toml:"type" comment:"sqlite, postgres"`
	File     string `toml:"file"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type EncryptModel struct {
	Key string `toml:"key" comment:"16/24/32 byte"`
	Iv  string `toml:"iv" comment:"16 byte"`
}

type LogModel struct {
	EnableFile bool     `toml:"enable-file"`
	File       string   `toml:"file"`
	Mask       []string `toml:"mask" comment:"unknown, debug, trace, info, warning, error, fatal"`
	Flag       []string `toml:"flag" comment:"date, time, long_file, short_file, func, prefix, suffix"`
	Debug      []string `toml:"debug" comment:"database, echo"`
}

type BotModel struct {
	Maid     *BotSenderModel `toml:"maid"`
	Reminder *BotSenderModel `toml:"reminder"`
}

type BotSenderModel struct {
	SMS      *SMSBotModel      `toml:"sms"`
	Mail     *MailBotModel     `toml:"mail"`
	Telegram *TelegramBotModel `toml:"telegram"`
	Wechat   *WechatBotModel   `toml:"wechat"`
}

type SMSBotModel struct {
	Debug bool `toml:"debug"`

	EnableReceiver bool   `toml:"enable-receiver"`
	EnableSender   bool   `toml:"enable-sender"`
	Api            string `toml:"api"`
	Token          string `toml:"token"`
}

type MailBotModel struct {
	Debug bool `toml:"debug"`

	EnableImap bool   `toml:"enable-imap"`
	HostImap   string `toml:"host-imap"`
	PortImap   int    `toml:"port-imap"`

	EnableSmtp bool   `toml:"enable-smtp"`
	HostSmtp   string `toml:"host-smtp"`
	PortSmtp   int    `toml:"port-smtp"`

	From     string `toml:"from"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type TelegramBotModel struct {
	Debug bool `toml:"debug"`

	EnableReceiver bool   `toml:"enable-receiver"`
	EnableSender   bool   `toml:"enable-sender"`
	API            string `toml:"api"`
	Token          string `toml:"token"`
}

type WechatBotModel struct {
	Debug bool `toml:"debug"`

	EnableReceiver bool   `toml:"enable-receiver"`
	EnableSender   bool   `toml:"enable-sender"`
	CorpId         string `toml:"corp-id"`
	Secret         string `toml:"secret"`
	AgentId        int64  `toml:"agent-id"`
	Token          string `toml:"token"`
	AesKey         string `toml:"aes-key"`
}

type CronModel struct {
	Debug *CronDebugModel `toml:"debug"`
}

type CronDebugModel struct {
	Enable bool `toml:"enable"`

	Hour   uint `toml:"hour"`
	Minute uint `toml:"minute"`
	Second uint `toml:"second"`
}
