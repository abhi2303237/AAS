package backend

type IDatabase interface {
	Connect() bool
}