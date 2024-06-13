package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
)

// getPermissionRouters 获取权限路由集 withRouter 是否提取路由信息
func getRolePermissions(traceID string, roleId uint64) (rolePermissions []*platDB.RolePermission, err error) {
	// 补充查询关联的路由和路由ID
	rolePermissions, err = platDB.RolePermissionTable.FindByObject(&platDB.RolePermission{RoleId: roleId})
	if err != nil {
		log.WarnTF(traceID, "GetPermissionsByRoleId By %d Fail . Err Is : %v", roleId, err)
	}
	return
}

// syncRolePermissions 编辑路由对应的权限
func syncRolePermissions(traceID string, roleId uint64, permissionIds, halfPermissionIds []uint64, editFlag bool) (err error) {
	if editFlag {
		// 移除当前角色选定的权限集
		err = platDB.RolePermission{}.DeleteByRoleId(roleId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByRoleId By %d Fail . Err Is : %v", roleId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(permissionIds) > 0 || len(halfPermissionIds) > 0 {
		rolePermissions := make([]platDB.RolePermission, len(permissionIds)+len(halfPermissionIds))
		i := 0
		for _, item := range permissionIds {
			rolePermissions[i] = platDB.RolePermission{
				RoleId:       roleId,
				PermissionId: item,
				CheckType:    constant.StatusOpen, // 选中
			}
			i++
		}
		for _, item := range halfPermissionIds {
			rolePermissions[i] = platDB.RolePermission{
				RoleId:       roleId,
				PermissionId: item,
				CheckType:    constant.StatusLock, // b半选
			}
			i++
		}
		err = platDB.PermissionRouterTable.InsertBatch(&rolePermissions)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchRolePermissions By RoleId %d Fail . Err Is : %v", roleId, err)
			return
		}
	}
	if editFlag {
		// TODO 如果角色被选择，则反向通知账号需要权限刷新
	}
	return
}
