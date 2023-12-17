package contract

type RequestData interface {
	GetMessage() string
	GetData() DtoInterface
}
