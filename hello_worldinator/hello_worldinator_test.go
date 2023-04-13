package hello_worldinator_test

import (
	"bufio"
	"io"
	"os"
	"programmingpuzzles/hello_worldinator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldinator(t *testing.T) {
	t.Run("Test console output by building a pipe", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatalf("error creating pipe: %v", err)
		}
		// Save the original stdout and replace it with the write end of the pipe
		originalStdout := os.Stdout
		os.Stdout = w

		// Call the function you want to test
		helloWorldFunc := hello_worldinator.HelloWorldinator()
		helloWorldFunc()

		// Close the write end of the pipe and restore the original stdout
		w.Close()
		os.Stdout = originalStdout

		// Read the console output from the read end of the pipe
		var consoleOutput string
		consoleReader := bufio.NewReader(r)
		for {
			line, err := consoleReader.ReadString('\n')
			if err == io.EOF {
				break
			}
			consoleOutput += line
		}

		// Check the console output against the expected value
		assert.Equal(t, "Hello World!\n", consoleOutput)

	})
}
