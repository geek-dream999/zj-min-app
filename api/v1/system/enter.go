package system

import "meet_directly/service"

type ApiGroup struct {
	OperationRecordApi
	BaseApi
}

var (
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	jwtService             = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
)
