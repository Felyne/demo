package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Felyne/demo/model"

	"github.com/gin-gonic/gin"
)

// GetAccount godoc
// @Summary Get a account
// @Description 根据name获取账户Account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param name path string true "Account name" Minlength(1) Maxlength(50) default("")
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{name} [get]
func (c *Controller) GetAccount(ctx *gin.Context) {
	name := strings.TrimSpace(ctx.Param("name"))
	if name == "" {
		NewError(ctx, http.StatusBadRequest, errors.New("name is null"))
		return
	}
	account, err := model.AccountOne(name)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// AddAccount godoc
// @Summary Add a account
// @Description 新增账户
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body model.Account true "Add account"
// @Success 200 {object} model.Account
// @Failure 400 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var addAccount model.Account
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}

	addAccount.Insert()
	ctx.JSON(http.StatusOK, addAccount)
}

// ListAccounts godoc
// @Summary List accounts
// @Description 账户列表
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param status query string false "状态" Format(string)
// @Param offset query int true "偏移" Mininum(0) default(0)
// @Param limit query int true "限制" Maxinum(50) default(10)
// @Success 200 {array} model.Account
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	status := ctx.Request.URL.Query().Get("status")
	//下面两个字段作为演示用
	_ = ctx.Request.URL.Query().Get("offset")
	_ = ctx.Request.URL.Query().Get("limit")
	accounts, err := model.AccountsAll(status)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// DeleteAccount godoc
// @Summary Delete a account
// @Description 删除账户
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param name path string true "账户名" Format(string)
// @Param Authorization header string true "Authentication header"
// @Success 204 {object} model.Account
// @Failure 404 {object} controller.HTTPError
// @Failure 500 {object} controller.HTTPError
// @Router /accounts/{name} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	fmt.Println(auth)
	name := ctx.Param("name")

	err := model.Delete(name)
	if err != nil {
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
