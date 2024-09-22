package handler

import (
	"swaggertest/models"
	"swaggertest/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func ConfigGetAccount(params operations.GetAccountAllParams) middleware.Responder {
	var result []*models.AccountEntry
	result = make([]*models.AccountEntry, 0)
	result = append(result, &models.AccountEntry{
		UserID:   "1",
		Password: "password",
		Email:    "test@test.io",
	})
	return operations.NewGetAccountAllOK().WithPayload(&operations.GetAccountAllOKBody{Attr: result})
}

func ConfigDeleteAccount(params operations.DeleteAccountUserIDParams) middleware.Responder {
	// delete account "user_id "1"
	if params.UserID == "1" {
		// success delete
		return operations.NewDeleteAccountUserIDNoContent()
	}

	// fail delete
	return operations.NewDeleteAccountUserIDBadRequest().WithPayload(&models.Error{
		Code:    400,
		Message: "Invalid user_id provided",
	})

}

func ConfigPostAccount(params operations.PostAccountParams) middleware.Responder {
	// create account
	newAccount := params.Attr

	// check if data is not empty
	if newAccount.UserID == "" || newAccount.Password == "" || newAccount.Email == "" {
		return operations.NewPostAccountBadRequest().WithPayload(&models.Error{
			Code:    400,
			Message: "missing required fields",
		})
	}

	// success create, return 201 Created (but, we use 200 OK here)
	return operations.NewPostAccountOK().WithPayload(&models.PostSuccess{
		Code:    200,
		Message: "Account created successfully",
	})
}
