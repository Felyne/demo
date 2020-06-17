package model

import (
	"errors"
	"fmt"
)

var (
	ErrNameInvalid = errors.New("name is empty")
	ErrNoRow       = errors.New("no rows in result set")
)

// 参考: https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags
// Account example
type Account struct {
	//备注: 姓名
	Name   string  `form:"name" json:"name" example:"Tom" format:"string" binding:"required,min=1,max=16" minLength:"1" maxLength:"16" validate:"max=10,min=1"` // 用户名
	Age    int     `json:"age" example:"10" binding:"required,gte=1,lte=150" minimum:"1" maximum:"150" default:"10"`                                            // 年龄
	Height float64 `form:"height" json:"height" example:"1.80" binding:"required" minimum:"0.0" maximum:"9.99"`                                                 // 身高,单位米
	Status string  `form:"status" json:"status" enums:"healthy,ill" binding:"required"`                                                                         // 状态
	Email  string  `form:"email" json:"email" binding:"required,email"`                                                                                         //邮箱
}

var accounts = []Account{
	{Name: "account_1", Age: 11, Height: 1.65, Status: "healthy"},
	{Name: "account_2", Age: 12, Height: 1.75, Status: "healthy"},
	{Name: "account_3", Age: 13, Height: 1.85, Status: "ill"},
}

// Validation example
func (a Account) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// AccountOne example
func AccountOne(name string) (Account, error) {
	for _, v := range accounts {
		if name == v.Name {
			return v, nil
		}
	}
	return Account{}, ErrNoRow
}

// Insert example
func (a Account) Insert() {
	accounts = append(accounts, a)
}

// AccountsAll example
func AccountsAll(status string) ([]Account, error) {
	if status == "" {
		return accounts, nil
	}
	var all []Account
	for k, v := range accounts {
		if status == v.Status {
			all = append(all, accounts[k])
		}
	}
	return all, nil
}

// Delete example
func Delete(name string) error {
	for k, v := range accounts {
		if name == v.Name {
			accounts = append(accounts[:k], accounts[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account name=%s is not found", name)
}
