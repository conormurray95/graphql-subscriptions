package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/conormurraypuppet/gqlbackend/graph/generated"
	"github.com/conormurraypuppet/gqlbackend/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	r.Notifier.SendMessage("todo-added")
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *subscriptionResolver) Notifications(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Notifier.RegisterSubscription(ctx.Done())
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
