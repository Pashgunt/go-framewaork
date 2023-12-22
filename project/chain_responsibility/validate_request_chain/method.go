package validate_request_chain

import (
	"fmt"
	"project/contract/validate_request_api"
	"project/dto"
)

type Method struct {
	next validate_request_api.DepartmentRequestApi
}

func (m *Method) Execute(dto *dto.ApiRequestDto) {
	if dto.Request().Method != dto.InnerMethod() {
		panic(fmt.Sprintf("Invalid HTTP Method income:%s system:%s", dto.Request().Method, dto.InnerMethod()))
	}

	m.next.Execute(dto)
}

func (m *Method) SetNext(next validate_request_api.DepartmentRequestApi) {
	m.next = next
}
