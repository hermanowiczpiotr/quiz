package application

type CommandHandler[Q any] interface {
	Handle(q Q) error
}

type QueryHandler[Q any, R any] interface {
	Handle(q Q) (r R, err error)
}
