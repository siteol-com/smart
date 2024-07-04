package platService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
)

// getAccountsRolesAndDept 获取账号列表的角色列表
func getAccountsRolesAndDept(list []*platDB.Account) (rolesIds [][]uint64, deptIds []uint64, err error) {
	rolesIds = make([][]uint64, len(list))
	deptMap := make(map[uint64]bool, 0)
	for i, item := range list {
		roleIds, errI := platDB.AccountRole{}.GetRoleIds(item.Id)
		if errI != nil {
			err = errI
			return
		}
		rolesIds[i] = roleIds
		deptMap[item.DeptId] = true
	}
	deptIds = make([]uint64, 0)
	for k, _ := range deptMap {
		deptIds = append(deptIds, k)
	}
	return
}

// syncAccountRoles 编辑角色对应的角色
func syncAccountRoles(traceID string, accountId uint64, roleIds []uint64, editFlag bool) (err error) {
	if editFlag {
		// 移除当前权限的路由
		err = platDB.AccountRole{}.DeleteByAccountId(accountId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByAccountId By %d Fail . Err Is : %v", accountId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(roleIds) > 0 {
		accountRoles := make([]platDB.AccountRole, len(roleIds))
		for i, item := range roleIds {
			accountRoles[i] = platDB.AccountRole{
				AccountId: accountId,
				RoleId:    item,
			}
		}
		err = platDB.AccountRoleTable.InsertBatch(&accountRoles)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchAccountRole By AccountId %d Fail . Err Is : %v", accountId, err)
			return
		}
	}
	if editFlag {
		// TODO 通知账号产生变化

	}
	return
}
