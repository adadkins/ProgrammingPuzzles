package philosophers_dining_test

import (
	philosophers_dining "programmingpuzzles/dining_philosophers_tutorial"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiningPhilosophers(t *testing.T) {
	t.Run("Can Philosophers Eat runs successfully", func(t *testing.T) {
		everyoneAte := philosophers_dining.Dine(0)
		assert.Equal(t, true, everyoneAte)
	})

	t.Run("Test that timeouts/deadlocks if everyone picks up right chopstick, so no one can eat. Create a circular dependency", func(t *testing.T) {
		// We set a duration to check the function didnt finish
		timeoutDuration := time.Second * 25

		// Create a channel to signal completion of Dine
		done := make(chan bool)

		// Run in a separate goroutine
		go func() {
			philosophers_dining.Dine(500)
			done <- true
		}()

		message := ""
		// Wait for either Dine to complete or for the timeout to occur
		select {
		case <-done:
			message = "No Timeout"
		case <-time.After(timeoutDuration):
			// Dine timed out
			message = "Time out"
		}
		assert.Equal(t, "Time out", message)
	})
}
