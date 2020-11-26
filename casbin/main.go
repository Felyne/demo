package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
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

func CheckPermission(e *casbin.Enforcer, user, domain, obj, op string) {
	enforce(e, user, domain, obj, op)
}
