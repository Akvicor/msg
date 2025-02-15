package dto

import (
	"errors"
	"github.com/wneessen/go-mail"
)

func (q *SendModel) Validate() error {
	if len(q.To) == 0 {
		return errors.New("to is empty")
	}
	if len(q.Msg) == 0 {
		return errors.New("msg is empty")
	}
	if len(q.Type) == 0 {
		q.Type = mail.TypeTextPlain
	}
	if len(q.Subject) == 0 {
		q.Subject = q.Msg[:min(24, len(q.Msg))]
	}
	return nil
}
