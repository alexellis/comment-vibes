package function

import (
	"testing"
)

func Test_TrimNewLine(t *testing.T) {

	str := `👍
`
	want := `:+1:`

	got := trim(str)

	if got != want {
		t.Errorf("Want: %s\ngot: %s", want, got)
	}
}

func Test_NoTrim(t *testing.T) {

	str := `👍`
	want := `:+1:`

	got := trim(str)

	if got != want {
		t.Errorf("Want: %s\ngot: %s", want, got)
	}
}

func Test_CantFind(t *testing.T) {

	str := `That looks 👀`
	want := `That looks :eyes:`

	got := trim(str)

	if got != want {
		t.Errorf("Want: %s\ngot: %s", want, got)
	}
}

func Test_IsValid_ThumbsUp(t *testing.T) {
	got := isEmoji(trim(`👍`))

	if got != true {
		t.Errorf("Want: %v\ngot: %v", true, got)
	}
}

func Test_NoResult(t *testing.T) {
	str := `no result`
	want := "no result"

	got := trim(str)

	if got != want {
		t.Errorf("Want: %x\ngot: %x", want, got)
	}
}
