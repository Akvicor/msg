package send

type Status string

const (
	StatusOK     Status = "ok"
	StatusWait   Status = "wait"
	StatusCancel Status = "cancel"
	StatusFailed Status = "failed"
)

type StatusModel struct {
	Status Status `json:"status"`
	SendAt int64  `json:"send_at"`
	SentAt int64  `json:"sent_at"`
	ErrMsg string `json:"err_msg"`
}

func NewStatusOK(sendAt, sentAt int64) *StatusModel {
	return &StatusModel{
		Status: StatusOK,
		SendAt: sendAt,
		SentAt: sentAt,
		ErrMsg: "ok",
	}
}

func NewStatusWait(sendAt, sentAt int64) *StatusModel {
	return &StatusModel{
		Status: StatusWait,
		SendAt: sendAt,
		SentAt: sentAt,
		ErrMsg: "wait",
	}
}

func NewStatusCancel(sendAt, sentAt int64) *StatusModel {
	return &StatusModel{
		Status: StatusCancel,
		SendAt: sendAt,
		SentAt: sentAt,
		ErrMsg: "cancel",
	}
}

func NewStatusFailed(sendAt, sentAt int64, errMsg string) *StatusModel {
	return &StatusModel{
		Status: StatusFailed,
		SendAt: sendAt,
		SentAt: sentAt,
		ErrMsg: errMsg,
	}
}
