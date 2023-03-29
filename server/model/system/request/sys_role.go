package request

type RoleParam struct {
	RoleName string `json:"roleName" binding:"required"` // 角色名
	ParentId string `json:"parentId" `                   //父角色ID
	Remark   string `json:"remark"`                      // 备注

}
