package helper

import (
	"project/chain_responsibility/validate_request_chain"
	"project/dto"
)

type ValidateStructureRequest struct {
}

func (validator ValidateStructureRequest) Validate(data *dto.ApiRequestDto) {
	path := &validate_request_chain.Path{}
	path.SetNext(&validate_request_chain.Method{})
	path.Execute(data)
}
