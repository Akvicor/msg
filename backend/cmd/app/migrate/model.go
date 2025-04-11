package migrate

import (
	"msg/cmd/app/server/model"
)

var list = []any{
	&model.User{}, &model.UserAccessToken{}, &model.LoginLog{}, &model.UserBindHomeTips{},
	&model.Channel{}, &model.Schedule{}, &model.Send{},
}
