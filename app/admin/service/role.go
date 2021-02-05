package service

import (
	"encoding/json"
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"strconv"
)

type Role struct {
}

// 多条件查询角色
func (e Role) SelectRoles(p dto.SelectRoleArrayDto, orderData []bo.Order) (roleData bo.SelectRoleArrayBo, err error) {
	role := new(models.SysRole)
	sysRole, err := role.SelectRoles(p, orderData)
	if err != nil {
		return
	}
	if len(sysRole) > 0 {
		for _, value := range sysRole {
			var recordRole bo.RecordRole
			recordRole.CreateBy = value.CreateBy
			recordRole.ID = value.ID
			recordRole.Level = value.Level
			recordRole.UpdateBy = value.UpdateBy
			recordRole.CreateTime = value.CreateTime
			recordRole.DataScope = value.DataScope
			recordRole.Description = value.Description
			recordRole.Name = value.Name
			recordRole.UpdateTime = value.UpdateTime
			if value.IsProtection[0] == 1 {
				recordRole.Protection = true
			} else {
				recordRole.Protection = false
			}
			sysDept, sysMenu, err1 := role.SysDeptAndMenu(value.ID)
			if err1 != nil {
				err = err1
				return
			}
			// Dept
			for _, value := range sysDept {
				var dept bo.Dept
				dept.CreateBy = value.CreateBy
				dept.CreateTime = value.CreateTime
				dept.DeptSort = value.DeptSort
				if value.Enabled[0] == 1 {
					dept.Enabled = true
				} else {
					dept.Enabled = false
				}
				if value.SubCount > 0 {
					dept.HasChildren = true
				} else {
					dept.HasChildren = false
				}
				dept.ID = value.ID
				dept.Name = value.Name
				dept.Pid = value.Pid
				dept.SubCount = value.SubCount
				dept.UpdateTime = value.UpdateTime
				dept.UpdateBy = value.UpdateBy
				recordRole.Depts = append(recordRole.Depts, dept)
			}
			// Menu
			for _, value := range sysMenu {
				var menu bo.Menu
				menu.CreateBy = value.CreateBy
				menu.Icon = value.Icon
				menu.ID = value.ID
				menu.MenuSort = value.MenuSort
				menu.Pid = value.Pid
				menu.SubCount = value.SubCount
				menu.Type = value.Type
				menu.UpdateBy = value.UpdateBy
				menu.Component = value.Component
				// TODO
				//menu.CreateTime = value.CreateTime
				menu.Name = value.Name
				menu.Path = value.Path
				menu.Permission = value.Permission
				menu.Title = value.Title
				//menu.UpdateTime = value.UpdateTime
				if value.Cache[0] == 1 {
					menu.Cache = true
				} else {
					menu.Cache = false
				}
				if value.Hidden[0] == 1 {
					menu.Hidden = true
				} else {
					menu.Hidden = false
				}
				if value.IFrame[0] == 1 {
					menu.Iframe = true
				} else {
					menu.Iframe = false
				}
				recordRole.Menus = append(recordRole.Menus, menu)
			}
			roleData.Records = append(roleData.Records, recordRole)
		}
	}
	roleData.Paging.Current = p.Current
	roleData.Paging.Page = p.Current
	roleData.Paging.SearchCount = true
	roleData.Paging.Size = p.Size
	roleData.Paging.Total = p.Current
	roleData.Paging.HitCount = false
	roleData.Paging.OptimizeCountSql = true
	roleData.Paging.Orders = orderData
	return
}

// 新增角色
func (e Role) InsertRole(p dto.InsertRoleDto) (err error) {
	role := new(models.SysRole)
	role.Level = p.Level
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	depts := []byte(p.Depts)
	deptsData := []int{}
	err = json.Unmarshal(depts, &deptsData)
	if err = role.InsertRole(deptsData); err != nil {
		return
	}
	return
}

// 修改角色
func (e Role) UpdateRole(p dto.UpdateRoleDto) (err error) {
	role := new(models.SysRole)
	role.ID = p.ID
	role.Level = p.Level
	role.CreateBy = p.CreateBy
	role.UpdateBy = p.UpdatedBy
	role.Name = p.Name
	role.DataScope = p.DataScope
	role.Description = p.Description
	//  参数格式化
	updateTime, err := strconv.ParseInt(p.UpdateTime, 10, 64)
	if err != nil {
		return
	}
	role.UpdateTime = updateTime
	if p.Protection == "true" {
		role.IsProtection = append(role.IsProtection, 1)
	} else {
		role.IsProtection = append(role.IsProtection, 0)
	}
	role.IsDeleted = append(role.IsDeleted, 0)
	depts := []byte(p.Depts)
	deptsData := []int{}
	if err = json.Unmarshal(depts, &deptsData); err != nil {
		return
	}
	menus := []byte(p.Menus)
	menusData := []int{}
	if err = json.Unmarshal(menus, &menusData); err != nil {
		return
	}
	if err = role.UpdateRole(deptsData, menusData); err != nil {
		return
	}
	return
}

// 删除角色
func (e Role) DeleteRole(p []int) (err error) {
	role := new(models.SysRole)
	if err = role.DeleteRole(p); err != nil {
		return
	}
	return
}

// 修改角色菜单
func (e Role) UpdateRoleMenu(id int, p []int) (err error) {
	role := new(models.SysRole)
	if err = role.UpdateRoleMenu(id, p); err != nil {
		return
	}
	return
}

// 获取单个角色
func (e Role) SelectRoleOne(id int) (roleone models.SysRole, err error) {
	role := new(models.SysRole)
	role.ID = id
	if roleone, err = role.SelectRoleOne(); err != nil {
		return
	}
	return
}
