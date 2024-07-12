package system

import "meet_directly/service"

type ApiGroup struct {
	OperationRecordApi
}

var (
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
