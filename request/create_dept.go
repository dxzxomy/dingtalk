package request

import (
	"encoding/json"
	"strings"
)

type CreateDept struct {
	//部门名称。长度限制为1~64个字符，不允许包含字符"-"","以及","
	Name string `json:"name" validate:"required,min=1,max=64"`

	//父部门ID，根部门ID为1。
	ParentId uint `json:"parent_id" validate:"required,min=1"`

	//是否隐藏本部门：
	//
	//true：隐藏部门，隐藏后本部门将不会显示在公司通讯录中
	//
	//false（默认值）：显示部门
	HideDept *bool `json:"hide_dept,omitempty"`

	//指定可以查看本部门的其他部门列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	DeptPermits string `json:"dept_ds,omitempty"`

	deptPermits []int

	//指定可以查看本部门的人员userid列表，总数不能超过200
	//
	//当hide_dept为true时，则此值生效
	UserPermits string `json:"user_ds,omitempty"`

	userPermits []string

	//是否限制本部门成员查看通讯录：
	//
	//true：开启限制。开启后本部门成员只能看到限定范围内的通讯录
	//
	//false（默认值）：不限制
	OuterDept *bool `json:"outer_dept,omitempty"`

	//本部门成员是否只能看到所在部门及下级部门通讯录：
	//
	//true：只能看到所在部门及下级部门通讯录
	//
	//false：不能查看所有通讯录，在通讯录中仅能看到自己
	//
	//当outer_dept为true时，此参数生效。
	OuterDeptOnlySelf *bool `json:"outer_dept_only_self,omitempty"`

	//指定本部门成员可查看的通讯录用户userid列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效。
	UserPermitsUsers string `json:"outer_permit_users,omitempty"`

	userPermitsUsers []string

	//指定本部门成员可查看的通讯录部门ID列表，总数不能超过200
	//
	//当outer_dept为true时，此参数生效
	UserPermitsDeptIds string `json:"outer_permit_depts,omitempty"`

	userPermitsDeptIds []int

	//是否创建一个关联此部门的企业群，默认为false即不创建
	CreateDeptGroup *bool `json:"create_dept_group,omitempty"`

	//是否默认同意加入该部门的申请：
	//
	//true：表示加入该部门的申请将默认同意
	//
	//false：表示加入该部门的申请需要有权限的管理员同意
	AutoApproveApply *bool `json:"auto_approve_apply,omitempty"`

	//在父部门中的排序值，order值小的排序靠前
	Order uint `json:"order,omitempty"`

	//部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	SourceIdentifier string `json:"source_identifier,omitempty"`
}

func (cd *CreateDept) String() string {
	str, _ := json.Marshal(cd)
	return string(str)
}

type createDeptBuilder struct {
	cd *CreateDept
}

func NewCreateDept(name string, parentId uint) *createDeptBuilder {
	return &createDeptBuilder{cd: &CreateDept{Name: name, ParentId: parentId}}
}

func (cdb *createDeptBuilder) SetHideDept(hide bool) *createDeptBuilder {
	cdb.cd.HideDept = &hide
	return cdb
}
func (cdb *createDeptBuilder) SetDeptPermits(deptId int, deptIds ...int) *createDeptBuilder {
	ds := cdb.cd.deptPermits
	ds = append(ds, deptId)
	cdb.cd.deptPermits = append(ds, deptIds...)
	return cdb
}

func (cdb *createDeptBuilder) SetUserPermits(userId string, userIds ...string) *createDeptBuilder {
	ds := cdb.cd.userPermits
	ds = append(ds, userId)
	cdb.cd.userPermits = append(ds, userIds...)
	return cdb
}

func (cdb *createDeptBuilder) SetOuterDept(outer bool) *createDeptBuilder {
	cdb.cd.OuterDept = &outer
	return cdb
}

func (cdb *createDeptBuilder) SetOuterDeptOnlySelf(self bool) *createDeptBuilder {
	if *cdb.cd.OuterDept {
		cdb.cd.OuterDeptOnlySelf = &self
	}
	return cdb
}

func (cdb *createDeptBuilder) SetUserPermitsUsers(userId string, userIds ...string) *createDeptBuilder {
	users := cdb.cd.userPermitsUsers
	users = append(users, userId)
	cdb.cd.userPermitsUsers = append(users, userIds...)
	return cdb
}

func (cdb *createDeptBuilder) SetUserPermitsDeptIds(deptId int, deptIds ...int) *createDeptBuilder {
	ids := cdb.cd.userPermitsDeptIds
	ids = append(ids, deptId)
	cdb.cd.userPermitsDeptIds = append(ids, deptIds...)
	return cdb
}

func (cdb *createDeptBuilder) SetCreateDeptGroup(group bool) *createDeptBuilder {
	cdb.cd.CreateDeptGroup = &group
	return cdb
}
func (cdb *createDeptBuilder) SetAutoApproveApply(approve bool) *createDeptBuilder {
	cdb.cd.AutoApproveApply = &approve
	return cdb
}

func (cdb *createDeptBuilder) SetOrder(order uint) *createDeptBuilder {
	cdb.cd.Order = order
	return cdb
}

func (cdb *createDeptBuilder) SetSourceIdentifier(id string) *createDeptBuilder {
	cdb.cd.SourceIdentifier = id
	return cdb
}

func (cdb *createDeptBuilder) Build() *CreateDept {
	cd := cdb.cd
	if cd.HideDept != nil && *cd.HideDept == true {
		cd.DeptPermits = strings.Join(removeIntDuplicatesToString(cd.deptPermits), ",")
		cd.UserPermits = strings.Join(removeStringDuplicates(cd.userPermits), ",")
	}
	if cd.OuterDept != nil && *cd.OuterDept == true {
		cd.UserPermitsDeptIds = strings.Join(removeIntDuplicatesToString(cd.userPermitsDeptIds), ",")
		cd.UserPermitsUsers = strings.Join(removeStringDuplicates(cd.userPermitsUsers), ",")
	}
	return cd
}
