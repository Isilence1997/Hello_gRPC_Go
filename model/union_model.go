package model

// SocietyUserInfoUnion2071 社区用户2071表信息
type SocietyUserInfoUnion2071 struct {
	UserHead          string `union:"user_head"`            //用户头像
	UserNick          string `union:"user_nick"`            //用户昵称
	UserDesc          string `union:"user_desc"`            //用户认证信息
	CpAuthLogo        string `union:"cp_auth_logo"`         //用户认证头像
}
