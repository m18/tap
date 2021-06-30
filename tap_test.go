package tap_test

import (
	"fmt"
	"testing"

	"github.com/m18/tap"
)

func TestNew(t *testing.T) {
	closedTap := tap.New()
	if closedTap == nil {
		t.Fatalf("expected tap to not be nil but it was")
	}
	select {
	case <-closedTap.Stream():
		t.Fatalf("unexpected stream flow from closed tap")
	default:
	}
}

func TestNewOpen(t *testing.T) {
	openTap := tap.NewOpen()
	if openTap == nil {
		t.Fatalf("expected tap to not be nil but it was")
	}
	select {
	case <-openTap.Stream():
	default:
		t.Fatalf("no stream flow from open tap")
	}
}

func TestOpenClose(t *testing.T) {
	toggleTap := tap.New()
	count := 0
	run := func() {
		select {
		case <-toggleTap.Stream():
			count++
		default:
		}
	}
	steps := []struct {
		open          bool
		runs          int
		expectedCount int
	}{
		{
			open:          false,
			runs:          1,
			expectedCount: 0,
		},
		{
			open:          false,
			runs:          2,
			expectedCount: 0,
		},
		{
			open:          true,
			runs:          1,
			expectedCount: 1,
		},
		{
			open:          false,
			runs:          1,
			expectedCount: 1,
		},
		{
			open:          true,
			runs:          2,
			expectedCount: 3,
		},
	}
	for _, step := range steps {
		if step.open {
			toggleTap.Open()
		} else {
			toggleTap.Close()
		}
		for i := 0; i < step.runs; i++ {
			fmt.Println(step)
			run()
		}
		if count != step.expectedCount {
			t.Fatalf("expected count to be %b but it was %d", step.expectedCount, count)
		}
	}
}
