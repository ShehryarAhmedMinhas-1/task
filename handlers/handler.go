package handlers

import (
	"github.com/go-openapi/loads"
	runtime "wancloudsV2"
	"wancloudsV2/gen/restapi/operations"
)

type taskAPIInstence *operations.TaskAPI

func NewHandler(rt *runtime.Runtime, specs *loads.Document) taskAPIInstence {
	handler := operations.NewTaskAPI(specs)

	handler.UsersRegisterUserHandler = NewRegisterUserHandler(*rt)
	handler.UsersLoginUserHandler = NewLoginUserHandler(*rt)
	handler.UsersUpdateUserHandler = NewUpdateUserHandler(*rt)
	handler.UsersLogoutUserHandler = NewLogoutUserHandler(*rt)

	return handler
}
