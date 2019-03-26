package tinterface

type Router interface {
	Handle(request Request)
}
