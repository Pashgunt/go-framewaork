package app

import "net/http"

type Dto struct {
	Data   map[string]string
	Status int16
}

func (d Dto) GetByField(field string) string {
	return d.Data[field]
}

func (d Dto) GetStatus() int16 {
	return d.Status
}

func (d Dto) GetStatusResult() bool {
	if d.GetStatus() >= http.StatusBadRequest {
		return false
	}

	return true
}
