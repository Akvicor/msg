package dto

import (
	"errors"
	"strings"
)

func (q *SendModel) Validate() error {
	q.Touser = strings.TrimSpace(q.Touser)
	q.Type = strings.TrimSpace(q.Type)
	q.Title = strings.TrimSpace(q.Title)
	q.Msg = strings.TrimSpace(q.Msg)
	q.Url = strings.TrimSpace(q.Url)
	q.Btn = strings.TrimSpace(q.Btn)
	if len(q.Touser) == 0 {
		return errors.New("touser is empty")
	}
	if len(q.Msg) == 0 {
		return errors.New("msg is empty")
	}
	return nil
}
