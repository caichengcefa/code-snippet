
type UserSignInReq struct {
	g.Meta `path:"/login" method:"get" tags:"User" summary:"Sign in with exist account"`
}

type UserSignInRes struct {
	Token   string   `json:"token"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Color   string   `json:"color"`
	Modules []string `json:"modules"`
}