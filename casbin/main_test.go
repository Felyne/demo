package main

import "testing"

var (
	// role对应用户组
	role1 = "developer"
	//role2   = "viewer"
	user1 = "23160589"
	// domain对应云服务
	domain1 = "kae:app"

	pObj = "team/*/app/*"
	rObj = "team/1/app/1"
	//resource2 = "team/"
	op1 = "UpdateBuildInfo"
	//op2 = "Get"
)

func init() {
	e := NewEnforcerByMysql()
	addPolicy(e, role1, domain1, pObj, op1)
	//addPolicy(e, role1, domain1, pObj, op2)

	addRoleForUserInDomain(e, user1, role1, domain1)
	//addRoleForUserInDomain(e, user1, role2, domain1)

	enforce(e, user1, domain1, rObj, op1)
	//enforce(e, user1, domain1, rObj, op2)
}

var tests = []struct {
	in       string // input
	expected string // expected result
}{
	{"???", "???"},
	{"filename.go", "filename.go"},
	{"hello/filename.go", "filename.go"},
	{"main/hello/filename.go", "filename.go"},
}

func TestCheckPermission(t *testing.T) {
}
