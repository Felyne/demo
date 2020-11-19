package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func main() {
	//dsn := "root:kae123456@tcp(10.13.84.186:13306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//db.Logger.LogMode(logger.Info)

	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "root:kae123456@tcp(10.13.84.186:13306)/casbin", true) // Your driver and data source.
	//a, _ := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	if err := e.LoadPolicy(); err != nil {
		log.Fatal(err)
	}

	var (
		role1   = "developer"
		role2   = "viewer"
		user1   = "John"
		domain1 = "team/1"
		//domain2   = "team/2"
		resource1 = "app/1"
		//resource2 = "app/2"
		action1 = "update"
		action2 = "get"
	)
	addPolicy(e, role1, domain1, resource1, action1)
	addPolicy(e, role2, domain1, resource1, action2)

	addRoleForUserInDomain(e, user1, role1, domain1)
	addRoleForUserInDomain(e, user1, role2, domain1)

	enforce(e, user1, domain1, resource1, action1)
	enforce(e, user1, domain1, resource1, action2)

	// Modify the policy.
	// e.AddPolicy(...)
	//e.RemovePolicy("alice", "data1", "read")

	// Save the policy back to DB.
	if err := e.SavePolicy(); err != nil {
		log.Fatal(err)
	}
}

func addRoleForUserInDomain(e *casbin.Enforcer, user, role, domain string) {
	ok, err := e.AddRoleForUserInDomain(user, role, domain)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("add role user mapper ok")
	} else {
		fmt.Println("add role user mapper not ok")
	}
}

func addPolicy(e *casbin.Enforcer, role, domain, resource, action string) {
	ok, err := e.AddPolicy(role, domain, resource, action)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("add policy ok")
	} else {
		fmt.Println("add policy not ok")
	}
}

func enforce(e *casbin.Enforcer, user, domain, resource, action string) {
	ok, err := e.Enforce(user, domain, resource, action)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("check policy ok")
	} else {
		fmt.Println("check policy not ok")
	}
}
