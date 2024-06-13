package types

type IDatabase interface {
	Connect() bool
	GetDatabase() any
}
