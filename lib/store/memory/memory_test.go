package memory

import (
	"testing"

	"github.com/ToastyTheBot/anubis/lib/store/storetest"
)

func TestImpl(t *testing.T) {
	storetest.Common(t, factory{}, nil)
}
