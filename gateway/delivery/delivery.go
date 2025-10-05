package delivery

type Handler interface {
	RegisterRoutes() error
}
