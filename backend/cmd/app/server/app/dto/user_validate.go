package dto

import (
	"errors"
	"strings"
)

func (b *UserCreate) Validate() error {
	b.Username = strings.TrimSpace(b.Username)
	if len(b.Username) == 0 {
		return errors.New("请输入用户名")
	}
	b.Password = strings.TrimSpace(b.Password)
	if len(b.Password) == 0 {
		return errors.New("请输入密码")
	}
	if len(b.Password) > 128 {
		return errors.New("您输入的密码过长")
	}
	b.Nickname = strings.TrimSpace(b.Nickname)
	if len(b.Nickname) == 0 {
		b.Nickname = b.Username
	}
	b.Mail = strings.TrimSpace(b.Mail)
	b.Phone = strings.TrimSpace(b.Phone)
	if !b.Role.Valid() {
		return errors.New("错误角色类型")
	}
	return nil
}

func (b *UserFind) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *UserLoginLogFind) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *UserUpdate) Validate() error {
	b.Username = strings.TrimSpace(b.Username)
	if len(b.Username) == 0 {
		return errors.New("请输入用户名")
	}
	b.Password = strings.TrimSpace(b.Password)
	if len(b.Password) > 128 {
		return errors.New("您输入的密码过长")
	}
	b.Nickname = strings.TrimSpace(b.Nickname)
	if len(b.Nickname) == 0 {
		b.Nickname = b.Username
	}
	b.Mail = strings.TrimSpace(b.Mail)
	b.Phone = strings.TrimSpace(b.Phone)
	if !b.Role.Valid() {
		return errors.New("错误角色类型")
	}
	return nil
}

func (b *UserLogin) Validate() error {
	b.Username = strings.TrimSpace(b.Username)
	if len(b.Username) == 0 {
		return errors.New("请输入用户名")
	}
	b.Password = strings.TrimSpace(b.Password)
	if len(b.Password) == 0 {
		return errors.New("请输入密码")
	}
	return nil
}

func (b *UserFindAccessToken) Validate() error {
	b.Search = strings.TrimSpace(b.Search)
	return nil
}

func (b *UserCreateAccessToken) Validate() error {
	b.Name = strings.TrimSpace(b.Name)
	if len(b.Name) == 0 {
		return errors.New("请输入名称")
	}
	return nil
}

func (b *UserUpdateAccessToken) Validate() error {
	if b.ID == 0 {
		return errors.New("请输入ID")
	}
	b.Name = strings.TrimSpace(b.Name)
	if len(b.Name) == 0 {
		return errors.New("请输入名称")
	}
	return nil
}

func (b *UserDeleteAccessToken) Validate() error {
	if b.ID == 0 {
		return errors.New("请输入ID")
	}
	return nil
}

func (b *UserDisableEnable) Validate() error {
	b.Username = strings.TrimSpace(b.Username)
	return nil
}

func (b *UserInfoUpdate) Validate() error {
	b.Username = strings.TrimSpace(b.Username)
	if len(b.Username) == 0 {
		return errors.New("请输入用户名")
	}
	b.Password = strings.TrimSpace(b.Password)
	if len(b.Password) > 128 {
		return errors.New("您输入的密码过长")
	}
	b.Nickname = strings.TrimSpace(b.Nickname)
	if len(b.Nickname) == 0 {
		b.Nickname = b.Username
	}
	b.Mail = strings.TrimSpace(b.Mail)
	b.Phone = strings.TrimSpace(b.Phone)
	return nil
}
