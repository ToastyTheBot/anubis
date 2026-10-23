package challengetest

import (
	"testing"
	"time"

	"github.com/ToastyTheBot/anubis"
	"github.com/ToastyTheBot/anubis/internal"
	"github.com/ToastyTheBot/anubis/lib/challenge"
	"github.com/google/uuid"
)

func New(t *testing.T) *challenge.Challenge {
	t.Helper()

	id := uuid.Must(uuid.NewV7())
	randomData := internal.SHA256sum(time.Now().String())

	return &challenge.Challenge{
		ID:         id.String(),
		RandomData: randomData,
		IssuedAt:   time.Now(),
		Difficulty: anubis.DefaultDifficulty,
	}
}
