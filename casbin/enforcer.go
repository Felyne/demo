package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func NewEnforcerByMysql() *casbin.Enforcer {
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
	//e, _ := casbin.NewEnforcer("casbin_model/rbac_model.conf", a)

	m, _ := model.NewModelFromString(modelText)

	// Load the policy rules from the .CSV file adapter.
	// Replace it with your adapter to avoid files.
	// Create the enforcer.
	e, _ := casbin.NewEnforcer(m, a)
	return e
}
