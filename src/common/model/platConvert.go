package model

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/mysql/platDb"
	"time"
)

// ToDbReq 字典创建对象转字典对象
func (r *DictAddReq) ToDbReq() *platDb.Dict {
	now := time.Now()
	return &platDb.Dict{
		Id:       0,
		GroupKey: r.GroupKey,
		Label:    r.Label,
		LabelEn:  r.LabelEn,
		Choose:   r.Choose,
		Val:      r.Val,
		Pid:      1,
		Sort:     0,
		Remark:   r.Remark,
		Common: platDb.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// ToDbReq 字典更新对象转字典对象
func (r *DictEditReq) ToDbReq(d *platDb.Dict) {
	now := time.Now()
	d.Label = r.Label
	d.LabelEn = r.LabelEn
	d.Choose = r.Choose
	d.Remark = r.Remark
	d.UpdateAt = &now
}

// ToDictPageRes 字典转为分页对象
func ToDictPageRes(list []*platDb.Dict) []*DictPageRes {
	res := make([]*DictPageRes, len(list))
	for i, r := range list {
		res[i] = &DictPageRes{
			Id:       r.Id,
			GroupKey: r.GroupKey,
			Label:    r.Label,
			LabelEn:  r.LabelEn,
			Choose:   r.Choose,
			Val:      r.Choose,
			Sort:     r.Sort,
			Remark:   r.Remark,
			Mark:     r.Mark,
		}
	}
	return res
}

// ToDictGetRes 字典转为查询对象
func ToDictGetRes(res *platDb.Dict) *DictGetRes {
	return &DictGetRes{
		Id:       res.Id,
		GroupKey: res.GroupKey,
		Label:    res.Label,
		LabelEn:  res.LabelEn,
		Choose:   res.Choose,
		Remark:   res.Remark,
	}
}

// ToDbReq 字典更新对象转字典对象
func (r *SysConfigEditReq) ToDbReq(d *platDb.SysConfig) {
	d.LoginSwitch = r.LoginSwitch
	d.LoginNum = r.LoginNum
	d.LoginFailSwitch = r.LoginFailSwitch
	d.LoginFailUnit = r.LoginFailUnit
	d.LoginFailNum = r.LoginFailNum
	d.LoginFailLockUnit = r.LoginFailLockUnit
	d.LoginFailLockNum = r.LoginFailLockNum
	d.LoginFailTryNum = r.LoginFailTryNum
	d.LogoutSwitch = r.LogoutSwitch
	d.LogoutUnit = r.LogoutUnit
	d.LogoutNum = r.LogoutNum
}

// ToSysConfigGetRes 系統配置转为查询对象
func ToSysConfigGetRes(res *platDb.SysConfig) *SysConfigGetRes {
	return &SysConfigGetRes{
		LoginSwitch:       res.LoginSwitch,
		LoginNum:          res.LoginNum,
		LoginFailSwitch:   res.LoginFailSwitch,
		LoginFailUnit:     res.LoginFailUnit,
		LoginFailNum:      res.LoginFailNum,
		LoginFailLockUnit: res.LoginFailLockUnit,
		LoginFailLockNum:  res.LoginFailLockNum,
		LoginFailTryNum:   res.LoginFailTryNum,
		LogoutSwitch:      res.LogoutSwitch,
		LogoutUnit:        res.LogoutUnit,
		LogoutNum:         res.LogoutNum,
	}
}
