package helper

import "project/contract"

type Pipeline struct {
	requestData      contract.RequestData
	operationHandler interface{}
	responseHandler  interface{}
}

func (p *Pipeline) SetRequestData(requestData contract.RequestData) *Pipeline {
	p.requestData = requestData

	return p
}

func (p Pipeline) SetOperationHandler() {

}

func (p Pipeline) SetResponseHandler() {

}
