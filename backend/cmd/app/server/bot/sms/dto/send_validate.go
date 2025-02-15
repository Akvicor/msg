package dto

import (
	"errors"
)

func (q *SendModel) Validate() error {
	if len(q.Phone) == 0 {
		return errors.New("phone is empty")
	}
	if len(q.Msg) == 0 {
		return errors.New("msg is empty")
	}
	return nil
}
