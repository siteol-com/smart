package authService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/common/utils/security"
)

// AuthLogin 账号密码登陆
func AuthLogin(traceID string, req *baseModel.AccountLoginReq) *baseModel.ResBody {
	// 获取风控缓存
	sysConf, err := cacheModel.GetSysConfigCache(traceID)
	if err != nil {
		return baseModel.Fail(constant.AuthLoginNG)
	}
	// 风控判定
	re := checkLoginFailRule(req.Acc, sysConf)
	if re {
		log.ErrorTF(traceID, "LoginFailRule %s Reject", req.Acc)
		return baseModel.Fail(constant.AuthLoginRuleNG)
	}
	// 查询账号
	account, err := platDB.AccountTable.GetOneByObject(&platDB.Account{Account: req.Acc})
	if err != nil {
		log.ErrorTF(traceID, "GetAccount %s Fail . Err Is : %v", req.Acc, err)
		return baseModel.Fail(constant.AuthLoginNG)
	}
	// 账号状态
	if account.Status != constant.StatusOpen {
		log.ErrorTF(traceID, "Account %s Status Is %s . Not Open", req.Acc, account.Status)
		return baseModel.Fail(constant.AuthLoginNG)
	}
	// 对比密码
	reqPwdC, err := security.AESEncrypt(req.Pwd, account.SaltKey)
	if err != nil {
		log.ErrorTF(traceID, "Account %s EncryptPwd Fail . Err Is : %v", req.Acc, err)
		return baseModel.Fail(constant.AuthLoginNG)
	}
	// 密码错误进入风控
	if reqPwdC != account.Encryption {
		log.ErrorTF(traceID, "Account %s EncryptPwd Fail . Err Is : %v", req.Acc, err)
		// 失败风控添加
		syncLoginFailRule(traceID, req.Acc, sysConf)
		return baseModel.Fail(constant.AuthLoginNG)
	}
	// 生成随机Token
	token := utils.Token()
	// 登陆记录，踢出账号的其他登陆（根据终端上限）
	syncLoginRecord(traceID, token, account.Id, sysConf)
	// 组装权限对象并写入Redis，如果密码过期则增加密码重置提示
	cacheAuth, err := makeAuthCache(traceID, token, account, sysConf)
	if err != nil {
		log.ErrorTF(traceID, "SetAuthCache %d Fail . Err Is : %v", account.Id, err)
		return baseModel.Fail(constant.AuthLoginNG)
	}
	return baseModel.Success(constant.AuthLoginSS, &baseModel.AccountLoginRes{
		Tk: token,
		Re: cacheAuth.NeedResetPwd,
	})
}

// AuthReset 密码重置
func AuthReset(traceID string, req *baseModel.AccountResetReq, authUser *cacheModel.CacheAuth) *baseModel.ResBody {
	_, err := platDB.AccountTable.GetOneById(authUser.AccountId)
	if err != nil {
		log.ErrorTF(traceID, "GetAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	saltKey := utils.SaltKey()
	// 初始密码
	pwdC, err := security.AESEncrypt(req.Pwd, saltKey)
	// 重置密码
	err = platDB.AccountTable.Executor().ResetAccount(authUser.AccountId, saltKey, pwdC, true)
	if err != nil {
		log.ErrorTF(traceID, "AuthReset %d Fail . Err Is : %v", authUser.AccountId, err)
		return baseModel.Fail(constant.AuthResetNG)
	}
	// 重置成功，请重新登陆
	return baseModel.Success(constant.AuthResetSS, true)
}
