package graph

import (
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph/model"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/livenotifier"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/notifier"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos        []*model.Todo
	Notifier     *notifier.Notifier
	LiveNotifier *livenotifier.LiveNotifier
}
