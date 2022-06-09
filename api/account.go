package api

import (
	"errors"
	"net/http"

	db "github.com/Arielcito/simple-bank-go/db/sqlc"
	"github.com/Arielcito/simple-bank-go/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct{
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context){
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account,err := server.store.CreateAccount(ctx,arg)

	if err != nil {
		if pqErr,ok := err.(*pq.Error); ok{
			switch pqErr.Code.Name(){
			case "foreign_key_violation","unique_violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
type getAccountRequest struct{
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) getAccount(ctx *gin.Context){
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account,err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
type listAccountRequest struct{
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
func (server *Server) listAccount(ctx *gin.Context){
	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	account,err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
type updateAccountRequest struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance" binding:"required"`
}
func (server *Server) updateAccount(ctx *gin.Context){
	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID: req.ID,
		Balance: req.Balance,
	}

	account,err := server.store.UpdateAccount(ctx,arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
type deleteAccountRequest struct{
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAccount(ctx *gin.Context){
	var req deleteAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,err)
}