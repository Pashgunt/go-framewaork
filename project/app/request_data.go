package app

import "project/contract"

type RequestData struct {
	Message string
	Data    contract.DtoInterface
}

func (rd RequestData) GetMessage() string {
	return rd.Message
}

func (rd RequestData) GetData() contract.DtoInterface {
	return rd.Data
}
