package dto

import (
	"errors"
)

func (q *Verify) Validate() error {
	if len(q.MsgSignature) == 0 {
		return errors.New("msg_signature is empty")
	}
	if len(q.TimeStamp) == 0 {
		return errors.New("timestamp is empty")
	}
	if len(q.Nonce) == 0 {
		return errors.New("nonce is empty")
	}
	if len(q.EchoStr) == 0 {
		return errors.New("echostr is empty")
	}

	return nil
}

func (q *Callback) Validate() error {
	if len(q.MsgSignature) == 0 {
		return errors.New("msg_signature is empty")
	}
	if len(q.TimeStamp) == 0 {
		return errors.New("timestamp is empty")
	}
	if len(q.Nonce) == 0 {
		return errors.New("nonce is empty")
	}

	return nil
}
