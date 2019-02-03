package session

import (
	"fmt"
	"testing"

	"github.com/go-kira/kon"
)

func TestOptions(t *testing.T) {
	opt := prepareOptions(kon.New())

	opt.Path = "/session"

	fmt.Println(opt)

	t.Error("Error")
}
