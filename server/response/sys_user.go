package response

import (
	"github.com/gofrs/uuid/v5"
)

type LoginResponse struct {
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
	UserInfo  UserInfo `json:"userInfo"`
	Menus     []Menus  `json:"menus"`
}
type UserInfo struct {
	Id      uuid.UUID `json:"id"`
	Account string    `json:"account"`
	HeadPic string    `json:"headPic"`
}
type Menus struct {
	Path     string  `json:"path"`
	Title    string  `json:"title"`
	Icon     string  `json:"icon"`
	Header   string  `json:"header"`
	IsHeader int32   `json:"isHeader"`
	Children []Menus `json:"children"`
}
