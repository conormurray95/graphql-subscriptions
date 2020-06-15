package notifier

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNotifier(t *testing.T) {
	done := make(chan bool)
	defer close(done)

	notifier := New(done)

	stop1 := make(chan struct{})
	stop2 := make(chan struct{})

	sub1, err := notifier.RegisterSubscription(stop1)
	require.Nil(t, err)
	setupNotification1 := <-sub1
	require.Equal(t, "subscription-active", setupNotification1.EventCode)
	require.NotNil(t, setupNotification1.CreatedAt)

	sub2, err := notifier.RegisterSubscription(stop2)
	require.Nil(t, err)
	setupNotification2 := <-sub2
	require.Equal(t, "subscription-active", setupNotification2.EventCode)
	require.NotNil(t, setupNotification2.CreatedAt)

	notifier.SendMessage("foo")

	notification1 := <-sub1
	require.Equal(t, "foo", notification1.EventCode)
	require.NotNil(t, notification1.CreatedAt)

	notification2 := <-sub2
	require.Equal(t, "foo", notification2.EventCode)
	require.NotNil(t, notification2.CreatedAt)

	stop1 <- struct{}{}

	// make sure the subscription 1 is stopped
	time.Sleep(10 * time.Millisecond)

	notifier.SendMessage("bar")

	notification1, ok := <-sub1
	require.False(t, ok)

	notification2, ok = <-sub2
	require.True(t, ok)
	require.Equal(t, "bar", notification2.EventCode)
	require.NotNil(t, notification2.CreatedAt)
}
