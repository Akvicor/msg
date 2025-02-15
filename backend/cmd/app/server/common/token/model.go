package token

import (
	"encoding/json"
	"github.com/Akvicor/util"
	"msg/cmd/app/server/common/encrypt"
	"time"
)

type Model struct {
	ID     int64  `json:"i"`
	Type   Type   `json:"p"`
	Time   int64  `json:"t"`
	Random string `json:"r"`
}

func (t *Model) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	encData, err := encrypt.Encrypt(string(data))
	if err != nil {
		return ""
	}
	return encData
}

func NewLoginToken(id int64) string {
	token := &Model{
		ID:     id,
		Time:   time.Now().Unix(),
		Type:   TypeLoginToken,
		Random: util.RandomString(8),
	}
	return token.String()
}

func NewAccessToken(id int64) string {
	token := &Model{
		ID:     id,
		Time:   time.Now().Unix(),
		Type:   TypeAccessToken,
		Random: util.RandomString(8),
	}
	return token.String()
}

func Parse(token string) *Model {
	decData, err := encrypt.Decrypt(token)
	if err != nil {
		return nil
	}
	tokenV := new(Model)
	err = json.Unmarshal([]byte(decData), tokenV)
	if err != nil {
		return nil
	}
	return tokenV
}
