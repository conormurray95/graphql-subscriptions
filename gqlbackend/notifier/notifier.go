package notifier

import (
	"sync"
	"time"

	"github.com/conormurraypuppet/gqlbackend/graph/model"
)

// Notifier gets messages and sends them to subscription observers
type Notifier struct {
	notificationMessages chan string
	observers            map[chan *model.Notification]bool
	done                 <-chan bool
	mu                   sync.Mutex
}

// New creates a new Notifier
func New(done <-chan bool) *Notifier {
	notifier := Notifier{
		notificationMessages: make(chan string),
		observers:            make(map[chan *model.Notification]bool),
		done:                 done,
	}

	go notifier.dispatch()

	return &notifier
}

func (n *Notifier) dispatch() {
	for {
		select {
		case msg, ok := <-n.notificationMessages:
			if !ok {
				return
			}

			notification := model.Notification{
				EventCode: msg,
				CreatedAt: time.Now(),
			}

			n.mu.Lock()
			for observer := range n.observers {
				observer <- &notification
			}
			n.mu.Unlock()

		case <-n.done:
			close(n.notificationMessages)
			return
		}
	}
}

// SendMessage sends message to subscription observers
func (n *Notifier) SendMessage(msg string) {
	n.notificationMessages <- msg
}

// RegisterSubscription registers a subscription observer
func (n *Notifier) RegisterSubscription(stop <-chan struct{}) (<-chan *model.Notification, error) {

	observer := make(chan *model.Notification, 1)

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
	notification := model.Notification{
		EventCode: "subscription-active",
		CreatedAt: time.Now(),
	}
	observer <- &notification
	n.mu.Unlock()

	return observer, nil
}
