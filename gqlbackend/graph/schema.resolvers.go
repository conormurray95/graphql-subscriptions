package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph/generated"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
	}
	r.todos = append(r.todos, todo)
	r.Notifier.SendMessage("todo-added")
	r.LiveNotifier.SendMessage(*todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	for index, todo := range r.todos {
		if todo.ID == input.ID {
			r.todos[index].Done = input.Done
			r.todos[index].Text = input.Text
			r.Notifier.SendMessage("todo-updated")
			return r.todos[index], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *subscriptionResolver) Notifications(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Notifier.RegisterSubscription(ctx.Done())
}

func (r *subscriptionResolver) TodoAdded(ctx context.Context) (<-chan *model.Todo, error) {
	return r.LiveNotifier.RegisterSubscription(ctx.Done())
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
