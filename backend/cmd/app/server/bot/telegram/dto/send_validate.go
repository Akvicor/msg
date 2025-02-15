package dto

import (
	"errors"
)

func (q *SendModel) Validate() error {
	if q.Chat == 0 {
		return errors.New("chat is empty")
	}
	if len(q.Msg) == 0 {
		return errors.New("msg is empty")
	}
	return nil
}
