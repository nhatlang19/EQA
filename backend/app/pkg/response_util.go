package pkg

import (
	"EQA/backend/app/constant"
	"EQA/backend/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse(responseStatus constant.ResponseStatus, data interface{}) dto.ApiResponse {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_(status string, message string, data interface{}) dto.ApiResponse {
	return dto.ApiResponse{
		Key:     status,
		Message: message,
		Data:    data,
	}
}
