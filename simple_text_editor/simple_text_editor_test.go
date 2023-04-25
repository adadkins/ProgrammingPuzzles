package simple_text_editor_test

import (
	"math/rand"
	"programmingpuzzles/simple_text_editor"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTextEditorArray(t *testing.T) {
	t.Run("Hacker rank example 1", func(t *testing.T) {
		input := []string{"1 fg", "3 6", "2 5", "4", "3 7", "4", "3 4"}
		result := simple_text_editor.SimpleTextEditorArray("abcde", input)
		expected := []string{"f", "g", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("Hacker rank example 2", func(t *testing.T) {
		input := []string{"1 abc", "3 3", "2 3", "1 xy", "3 2", "4", "4", "3 1"}
		result := simple_text_editor.SimpleTextEditorArray("", input)
		expected := []string{"c", "y", "a"}
		assert.Equal(t, expected, result)
	})
	t.Run("Hacker rank example 3", func(t *testing.T) {
		input := []string{"1 zsfncpxdzl", "3 4", "3 6", "2 1", "3 7", "3 2", "4", "2 4", "2 6", "4", "4", "1 l", "1 hpe", "3 6", "2 7", "4", "3 6", "4", "3 6", "1 zipsqagri", "1 vuqxstnj", "4", "3 13", "4", "3 10", "3 6", "1 uzdpy", "1 bupqp", "1 kn", "2 6", "3 8", "1 iiuvfbn", "4", "2 1", "2 12", "4", "3 7", "4", "2 9", "3 1", "1 axbhx", "1 wovbfyvt", "3 11", "3 7", "3 2", "4", "1 tjmqp", "4", "2 6", "3 4"}
		result := simple_text_editor.SimpleTextEditorArray("", input)
		expected := []string{"n", "p", "x", "s", "p", "p", "p", "i", "l", "p", "d", "x", "z", "b", "x", "s", "n"}
		assert.Equal(t, expected, result)
	})
}

func TestSimpleTextEditorStack(t *testing.T) {
	t.Run("Hacker rank example 1", func(t *testing.T) {
		input := []string{"1 fg", "3 6", "2 5", "4", "3 7", "4", "3 4"}
		result := simple_text_editor.SimpleTextEditorStack("abcde", input)
		expected := []string{"f", "g", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("Hacker rank example 2", func(t *testing.T) {
		input := []string{"1 abc", "3 3", "2 3", "1 xy", "3 2", "4", "4", "3 1"}
		result := simple_text_editor.SimpleTextEditorStack("", input)
		expected := []string{"c", "y", "a"}
		assert.Equal(t, expected, result)
	})
	t.Run("Hacker rank example 3", func(t *testing.T) {
		input := []string{"1 zsfncpxdzl", "3 4", "3 6", "2 1", "3 7", "3 2", "4", "2 4", "2 6", "4", "4", "1 l", "1 hpe", "3 6", "2 7", "4", "3 6", "4", "3 6", "1 zipsqagri", "1 vuqxstnj", "4", "3 13", "4", "3 10", "3 6", "1 uzdpy", "1 bupqp", "1 kn", "2 6", "3 8", "1 iiuvfbn", "4", "2 1", "2 12", "4", "3 7", "4", "2 9", "3 1", "1 axbhx", "1 wovbfyvt", "3 11", "3 7", "3 2", "4", "1 tjmqp", "4", "2 6", "3 4"}
		result := simple_text_editor.SimpleTextEditorStack("", input)
		expected := []string{"n", "p", "x", "s", "p", "p", "p", "i", "l", "p", "d", "x", "z", "b", "x", "s", "n"}
		assert.Equal(t, expected, result)
	})
}

var benchmarkInput = randomString(1000)
var benchmarkCommands = generateRandomCommands(1000)

func BenchmarkSimpleTextEditorArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simple_text_editor.SimpleTextEditorArray(benchmarkInput, benchmarkCommands)
	}
}

func BenchmarkSimpleTextEditorStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simple_text_editor.SimpleTextEditorStack(benchmarkInput, benchmarkCommands)
	}
}
func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateRandomCommands(n int) []string {
	var commands []string
	for i := 0; i < n; i++ {
		commandType := rand.Intn(4) + 1
		switch commandType {
		case 1:
			commands = append(commands, strconv.Itoa(commandType)+" "+randomString(10))
		case 2:
			commands = append(commands, strconv.Itoa(commandType)+" "+strconv.Itoa(rand.Intn(10)+1))
		case 3:
			commands = append(commands, strconv.Itoa(commandType)+" "+strconv.Itoa(rand.Intn(10)+1))
		case 4:
			commands = append(commands, strconv.Itoa(commandType))
		}
	}
	return commands
}
