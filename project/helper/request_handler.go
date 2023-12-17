package helper

import (
	"fmt"
	"net/http"
	"project/app"
	"project/contract"
	"project/enum"
	"regexp"
)

type Handle struct {
	response http.ResponseWriter
	request  http.Request
	method   string
}

func NewHandle(response http.ResponseWriter, request http.Request) *Handle {
	return &Handle{response: response, request: request}
}

func (h Handle) GetRequest() http.Request {
	return h.request
}

func (h Handle) GetResponse() http.ResponseWriter {
	return h.response
}

func (h *Handle) SetMethod(method string) *Handle {
	h.method = method
	return h
}

func (h Handle) Validate(fields map[string]ValidateRule) contract.RequestData {
	h.checkCorrectRequestMethod()
	requestData := make(map[string]string)
	for fieldName, rules := range fields {
		_, err := h.checkValidateRule(fieldName, rules, &requestData)

		if err != nil {
			return app.RequestData{
				Message: err.Error(),
				Data: app.Dto{
					Data:   requestData,
					Status: http.StatusBadRequest,
				},
			}
		}
	}

	return app.RequestData{
		Message: enum.SuccessValidateRequestData,
		Data: app.Dto{
			Data:   requestData,
			Status: http.StatusOK,
		},
	}
}

func (h *Handle) setDefaultMethod() {
	if h.method == "" {
		h.method = http.MethodGet
	}
}
func (h Handle) getFormValue(field string) string {
	request := h.GetRequest()
	request.ParseForm()
	return request.FormValue(field)
}

func (h Handle) checkCorrectRequestMethod() {
	h.setDefaultMethod()
	method := h.GetRequest().Method
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error panic with read request body: %s", err)
			return
		}
	}()

	if method != h.method {
		panic("Error for read request body")
	}
}

func (h Handle) checkValidateRule(
	field string,
	rules ValidateRule,
	requestData *map[string]string,
) (bool, error) {
	val := h.getFormValue(field)

	if rules.Required && (len(val) == 0 || !regexp.MustCompile(rules.Regex).MatchString(val)) {
		return false, fmt.Errorf("%s", enum.ErrorValidateRequestData)
	}

	(*requestData)[field] = val

	return true, nil
}
