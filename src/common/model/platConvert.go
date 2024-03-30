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

// ToDictPageRes 字典转为分页对象
func ToDictPageRes(list []*platDb.Dict) []*DictPageRes {
	res := make([]*DictPageRes, len(list))
	for i, r := range list {
		res[i] = &DictPageRes{
			DictGetRes: DictGetRes{
				Id:       r.Id,
				GroupKey: r.GroupKey,
				Val:      r.Val,
				Label:    r.Label,
				LabelEn:  r.LabelEn,
				Choose:   r.Choose,
				Remark:   r.Remark,
			},
			Sort: r.Sort,
			Mark: r.Mark,
		}
	}
	return res
}

// ToDictGetRes 字典转为查询对象
func ToDictGetRes(r *platDb.Dict) *DictGetRes {
	return &DictGetRes{
		Id:       r.Id,
		GroupKey: r.GroupKey,
		Val:      r.Val,
		Label:    r.Label,
		LabelEn:  r.LabelEn,
		Choose:   r.Choose,
		Remark:   r.Remark,
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
func ToSysConfigGetRes(r *platDb.SysConfig) *SysConfigGetRes {
	return &SysConfigGetRes{
		LoginSwitch:       r.LoginSwitch,
		LoginNum:          r.LoginNum,
		LoginFailSwitch:   r.LoginFailSwitch,
		LoginFailUnit:     r.LoginFailUnit,
		LoginFailNum:      r.LoginFailNum,
		LoginFailLockUnit: r.LoginFailLockUnit,
		LoginFailLockNum:  r.LoginFailLockNum,
		LoginFailTryNum:   r.LoginFailTryNum,
		LogoutSwitch:      r.LogoutSwitch,
		LogoutUnit:        r.LogoutUnit,
		LogoutNum:         r.LogoutNum,
	}
}

// ToDbReq 响应码创建对象转字典对象
func (r *ResponseAddReq) ToDbReq() *platDb.ResponseCode {
	now := time.Now()
	return &platDb.ResponseCode{
		Id:          0,
		ServiceCode: r.ServiceCode,
		Type:        r.Type,
		ZhCn:        r.ZhCn,
		EnUs:        r.EnUs,
		Remark:      r.Remark,
		Common: platDb.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// ToResponsePageRes 响应码转为分页对象
func ToResponsePageRes(list []*platDb.ResponseCode) []*ResponsePageRes {
	res := make([]*ResponsePageRes, len(list))
	for i, r := range list {
		res[i] = &ResponsePageRes{
			ResponseGetRes: ResponseGetRes{
				Id:          r.Id,
				Code:        r.Code,
				ServiceCode: r.ServiceCode,
				Type:        r.Type,
				ZhCn:        r.ZhCn,
				EnUs:        r.EnUs,
				Remark:      r.Remark,
			},
			Mark: r.Mark,
		}
	}
	return res
}

// ToResponseGetRes 响应码转为查询对象
func ToResponseGetRes(r *platDb.ResponseCode) *ResponseGetRes {
	return &ResponseGetRes{
		Id:          r.Id,
		Code:        r.Code,
		ServiceCode: r.ServiceCode,
		Type:        r.Type,
		ZhCn:        r.ZhCn,
		EnUs:        r.EnUs,
		Remark:      r.Remark,
	}
}

// ToDbReq 响应码更新对象转字典对象
func (r *ResponseEditReq) ToDbReq(d *platDb.ResponseCode) {
	now := time.Now()
	d.ZhCn = r.ZhCn
	d.EnUs = r.EnUs
	d.Remark = r.Remark
	d.UpdateAt = &now
}
