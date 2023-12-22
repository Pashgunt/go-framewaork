package helper

import (
	"project/chain_responsibility/validate_request_chain"
	"project/dto"
)

type ValidateStructureRequest struct {
}

func (validator ValidateStructureRequest) Validate(data *dto.ApiRequestDto) {
	path := &validate_request_chain.Path{}
	method := &validate_request_chain.Method{}
	method.SetNext(path)
	path.Execute(data)
}
