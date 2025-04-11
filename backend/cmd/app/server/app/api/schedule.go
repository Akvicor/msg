package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"msg/cmd/app/server/app/dto"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/global/auth"
	"msg/cmd/app/server/model"
	"msg/cmd/app/server/service"
)

var Schedule = new(scheduleApi)

type scheduleApi struct{}

// Create 创建
func (a *scheduleApi) Create(c echo.Context) (err error) {
	input := new(dto.ScheduleCreate)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	_, err = service.Schedule.Create(user.ID, input.Category, input.Type, input.Title, input.Message, input.ChannelID, input.PeriodType, input.StartAt, input.Year, input.Quarter, input.Month, input.Week, input.Day, input.Hour, input.Minute, input.Second, input.ExpirationDate, input.ExpirationTimes)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "创建失败: "+err.Error())
	}
	return resp.Success(c)
}

// Find 获取通知渠道的信息 (全部, 单个, 模糊搜索)
func (a *scheduleApi) Find(c echo.Context) (err error) {
	input := new(dto.ScheduleFind)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	if input.ID == 0 {
		var schedules = make([]*model.Schedule, 0)
		schedules, err = service.Schedule.FindAllByUIDLike(&input.PageModel, !input.All, model.NewPreloaderSchedule().Channel(), user.ID, input.Search)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithPageData(c, &input.PageModel, schedules)
	} else {
		var schedule *model.Schedule
		schedule, err = service.Schedule.FindByID(!input.All, model.NewPreloaderSchedule().Channel(), user.ID, input.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.FailWithMsg(c, resp.Failed, "系统错误")
		}
		return resp.SuccessWithData(c, schedule)
	}
}

// Update 更新
func (a *scheduleApi) Update(c echo.Context) (err error) {
	input := new(dto.ScheduleUpdate)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.Update(user.ID, input.ID, input.Category, input.Type, input.Title, input.Message, input.ChannelID, input.PeriodType, input.StartAt, input.Year, input.Quarter, input.Month, input.Week, input.Day, input.Hour, input.Minute, input.Second, input.ExpirationDate, input.ExpirationTimes)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败: "+err.Error())
	}
	return resp.Success(c)
}

// UpdateNext 更新下一个时间
func (a *scheduleApi) UpdateNext(c echo.Context) (err error) {
	input := new(dto.ScheduleUpdateNext)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.UpdateNext(user.ID, input.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "更新失败: "+err.Error())
	}
	return resp.Success(c)
}

// UpdateSequence 更新序号
func (a *scheduleApi) UpdateSequence(c echo.Context) (err error) {
	input := new(dto.ScheduleUpdateSequence)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	self := auth.GetUser(c)
	if self == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.UpdateSequenceByUID(self.ID, input.ID, input.Target)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "序号更新失败: "+err.Error())
	}
	return resp.Success(c)
}

// Disable 停用
func (a *scheduleApi) Disable(c echo.Context) (err error) {
	input := new(dto.ScheduleDisable)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.UpdateDisabledByUID(user.ID, input.ID, true)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "停用失败: "+err.Error())
	}
	return resp.Success(c)
}

// Enable 启用
func (a *scheduleApi) Enable(c echo.Context) (err error) {
	input := new(dto.ScheduleDisable)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.UpdateDisabledByUID(user.ID, input.ID, false)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "启用失败: "+err.Error())
	}
	return resp.Success(c)
}

// Delete 删除
func (a *scheduleApi) Delete(c echo.Context) (err error) {
	input := new(dto.ScheduleDelete)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, err.Error())
	}
	user := auth.GetUser(c)
	if user == nil {
		return resp.FailWithMsg(c, resp.UnAuthorized, "请登录")
	}
	err = service.Schedule.Delete(user.ID, input.ID)
	if err != nil {
		return resp.FailWithMsg(c, resp.Failed, "删除失败: "+err.Error())
	}
	return resp.Success(c)
}
