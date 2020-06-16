package livenotifier

import (
	"sync"

	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph/model"
)

// LiveNotifier gets messages and sends them to subscription observers
type LiveNotifier struct {
	notificationMessages chan model.Todo
	observers            map[chan *model.Todo]bool
	done                 <-chan bool
	mu                   sync.Mutex
}

// New creates a new LiveNotifier
func New(done <-chan bool) *LiveNotifier {
	notifier := LiveNotifier{
		notificationMessages: make(chan model.Todo),
		observers:            make(map[chan *model.Todo]bool),
		done:                 done,
	}

	go notifier.dispatch()

	return &notifier
}

func (n *LiveNotifier) dispatch() {
	for {
		select {
		case todo, ok := <-n.notificationMessages:
			if !ok {
				return
			}

			n.mu.Lock()
			for observer := range n.observers {
				observer <- &todo
			}
			n.mu.Unlock()

		case <-n.done:
			close(n.notificationMessages)
			return
		}
	}
}

// SendMessage sends message to subscription observers
func (n *LiveNotifier) SendMessage(todo model.Todo) {
	n.notificationMessages <- todo
}

// RegisterSubscription registers a subscription observer
func (n *LiveNotifier) RegisterSubscription(stop <-chan struct{}) (<-chan *model.Todo, error) {

	observer := make(chan *model.Todo, 1)

	go func() {
		select {
		case <-stop:
		case <-n.done:
		}
		n.mu.Lock()
		delete(n.observers, observer)
		close(observer)
		n.mu.Unlock()
	}()

	n.mu.Lock()
	n.observers[observer] = true
	n.mu.Unlock()

	return observer, nil
}
