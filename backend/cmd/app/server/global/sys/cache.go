package sys

import (
	"msg/cmd/app/server/common/cache"
	"msg/cmd/app/server/global/auth"
)

type TokenManagerItemModel struct {
	Authorization *auth.Authorization `json:"authorization"`
	Expiration    int64               `json:"expiration"`
	Expired       bool                `json:"expired"`
}

func TokenManagerItems() map[string]*TokenManagerItemModel {
	result := make(map[string]*TokenManagerItemModel)
	items := cache.TokenManager.Items()
	for k, v := range items {
		result[k] = &TokenManagerItemModel{
			Authorization: v.Object.(*auth.Authorization),
			Expiration:    v.Expiration,
			Expired:       v.Expired(),
		}
	}
	return result
}

type LoginFailedManagerItemModel struct {
	Count      int   `json:"count"`
	Expiration int64 `json:"expiration"`
	Expired    bool  `json:"expired"`
}

func LoginFailedManagerItems() map[string]*LoginFailedManagerItemModel {
	result := make(map[string]*LoginFailedManagerItemModel)
	items := cache.LoginFailedManager.Items()
	for k, v := range items {
		result[k] = &LoginFailedManagerItemModel{
			Count:      v.Object.(int),
			Expiration: v.Expiration,
			Expired:    v.Expired(),
		}
	}
	return result
}
