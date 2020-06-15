package graph

import (
	"github.com/conormurraypuppet/gqlbackend/graph/model"
	"github.com/conormurraypuppet/gqlbackend/notifier"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos    []*model.Todo
	Notifier *notifier.Notifier
}
