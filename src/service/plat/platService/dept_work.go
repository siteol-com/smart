package platService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"sort"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkDeptDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "dept_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.DeptUniNameNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 递归部门树
func recursionDeptTree(traceID string, treeNode *baseModel.Tree) (err error) {
	// 没有子级了
	if treeNode.Level == "3" {
		return
	}
	// 查询子集
	var deptList platDB.DeptArray
	deptList, err = platDB.DeptTable.GetByObject(&platDB.Dept{Pid: treeNode.Id})
	if err != nil {
		log.WarnTF(traceID, "RecursionDeptTree Fail . PID %d . Err is : %s", treeNode.Id, err)
		return
	}
	if len(deptList) == 0 {
		// 沒有子集推出
		return
	}
	// 数据排序
	sort.Sort(deptList)
	treeNode.Children = make([]*baseModel.Tree, 0)
	// 组装子集
	for _, item := range deptList {
		// 节点对象
		treeChild := &baseModel.Tree{
			Title:    item.Name,
			Key:      fmt.Sprintf("%d", item.Id),
			Children: nil,
			Expand:   item.Name,
			Level:    constant.StatusLock, // 可以移动
			Id:       item.Id,
		}
		// 递归子集
		recursionDeptTree(traceID, treeChild)
		// 加入子集
		treeNode.Children = append(treeNode.Children, treeChild)
	}
	return
}

// getDeptMapByIds 根据ID获取部门MAP
func getDeptMapByIds(ids []uint64) (deptMap map[uint64]string, err error) {
	if len(ids) == 0 {
		return
	}
	deptS, err := platDB.DeptTable.GetByIds(ids)
	if err != nil {
		return
	}
	deptMap = make(map[uint64]string, len(deptS))
	for _, item := range deptS {
		deptMap[item.Id] = item.Name
	}
	return
}
