package proofofwork

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/ToastyTheBot/anubis/internal"
	"github.com/ToastyTheBot/anubis/lib/challenge"
	"github.com/ToastyTheBot/anubis/lib/config"
	"github.com/ToastyTheBot/anubis/lib/policy"
)

func TestInvalidProofDoesNotRevealExpectedDigest(t *testing.T) {
	for _, algorithm := range []string{"fast", "slow"} {
		t.Run(algorithm, func(t *testing.T) {
			data := "oracle-regression"
			expected := internal.SHA256sum(data + "0")
			err := (&Impl{Algorithm: algorithm}).Validate(mkRequest(t, map[string]string{"nonce": "0", "elapsedTime": "1", "response": strings.Repeat("f", 64)}), slog.Default(), &challenge.ValidateInput{Challenge: &challenge.Challenge{RandomData: data}, Rule: &policy.Bot{Challenge: &config.ChallengeRules{Difficulty: 1}}})
			if err == nil {
				t.Fatal("invalid proof accepted")
			}
			if strings.Contains(err.Error(), expected) {
				t.Error("error reveals expected digest")
			}
		})
	}
}
