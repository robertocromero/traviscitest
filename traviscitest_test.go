package traviscitest

import (
	"github.com/bmizerany/assert"
	"testing"
)


func TestIsHelloWorldOK(t *testing.T) {
	word := "hello"
	boolTrue, err := IsHelloWorld(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, boolTrue)
}

func TestIsHelloWorldFail(t *testing.T) {
	word := "notHello"
	boolFalse, err := IsHelloWorld(word)
	if err == nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, boolFalse)
	
}

