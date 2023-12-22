package validate_request_api

import (
	"project/dto"
)

type DepartmentRequestApi interface {
	Execute(dto *dto.ApiRequestDto)
	SetNext(next DepartmentRequestApi)
}
