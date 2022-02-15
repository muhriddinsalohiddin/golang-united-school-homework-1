package solution

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	want := GetMessage()
	got := emoji.Sprint("Hello :world_map:!")
	if want != got {
		t.Errorf("Error: Want:%v got %v", want, got)
	}
}
