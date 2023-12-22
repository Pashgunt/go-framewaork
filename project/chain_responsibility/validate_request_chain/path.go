package validate_request_chain

import (
	"fmt"
	"project/contract/validate_request_api"
	"project/dto"
)

type Path struct {
	next validate_request_api.DepartmentRequestApi
}

func (p *Path) Execute(dto *dto.ApiRequestDto) {
	if dto.Request().URL.Path != dto.InnerPath() {
		panic(fmt.Sprintf("Invalid Http Request income:%s system:%s", dto.Request().Method, dto.InnerMethod()))
	}

	p.next.Execute(dto)
}

func (p *Path) SetNext(next validate_request_api.DepartmentRequestApi) {
	p.next = next
}
