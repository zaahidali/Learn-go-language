package graph

import (
	"github.com/zaahidali/Learn-go-language/task_management_api_graphq/graph/model" 
)
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	users []*model.User
	tasks []*model.Task
}
