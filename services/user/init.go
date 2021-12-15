package user

import (
	"goku.net/framework/network/http"
	"goku.net/services/user/logic"
)

func NewExecutorFactory() http.ModuleExecutorFactory {
	return &logic.UserFactory{}
}
