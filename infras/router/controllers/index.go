package controllers

import "BeGo/infras/router/queries"

type Controllers struct {
	Queries queries.IQueries
}

type IControllers interface {
	IPostControllers
	IUserControllers
}

func NewControllers(q queries.IQueries) IControllers {

	return &Controllers{Queries: q}
}
