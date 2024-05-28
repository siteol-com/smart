package platService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
)

// getPermissionRouters 获取权限路由集 withRouter 是否提取路由信息
func getPermissionRouters(permissionId uint64, withRouter bool, traceID string) (routerIds []uint64, routers []*platModel.RouterPageRes, err error) {
	// 补充查询关联的路由和路由ID
	routerIds, err = platDB.PermissionRouter{}.FindRouterIds(permissionId)
	if err != nil {
		log.WarnTF(traceID, "GetRouterIdsByPermissionId By %d Fail . Err Is : %v", permissionId, err)
		return
	}
	list := make([]*platDB.Router, 0)
	// 关联查询路由字段
	if withRouter && len(routerIds) > 0 {
		listR, errR := platDB.RouterTable.FindByIds(routerIds)
		if errR != nil {
			err = errR
			log.WarnTF(traceID, "GetRouterByRouterIds By %v Fail . Err Is : %v", routerIds, err)
			return
		}
		list = listR
	}
	routers = platModel.ToRouterPageRes(list)
	return
}

// editPermissionRouters 编辑权限对应的路由
func syncPermissionRouters(permissionId uint64, routerIds []uint64, editFlag bool, traceID string) (err error) {
	if editFlag {
		err = platDB.PermissionRouter{}.DeleteByPermissionId(permissionId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPermissionId By %d Fail . Err Is : %v", permissionId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(routerIds) > 0 {
		permissionRouters := make([]platDB.PermissionRouter, len(routerIds))
		for i, item := range routerIds {
			permissionRouters[i] = platDB.PermissionRouter{
				PermissionId: permissionId,
				RouterId:     item,
			}
		}
		err = platDB.PermissionRouterTable.InsertBatch(&permissionRouters)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPermissionRouter By PermissionId %d Fail . Err Is : %v", permissionId, err)
		}
	}
	return
}
