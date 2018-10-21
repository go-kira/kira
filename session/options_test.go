package session

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	opt := prepareOptions()

	opt.Path = "/session"

	fmt.Println(opt)

	t.Error("Error")
}
