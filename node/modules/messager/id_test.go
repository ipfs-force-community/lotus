package messager

import (
	"fmt"
	"testing"
)

func TestNewMId(t *testing.T) {
	for i := 0; i < 10; i++ {
		got, err := NewMId()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(got)
	}
}
