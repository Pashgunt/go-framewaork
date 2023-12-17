package contract

type DtoInterface interface {
	GetByField(field string) string
	GetStatus() int16
	GetStatusResult() bool
}
