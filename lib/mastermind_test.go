package go_mastermind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueDigits(t *testing.T) {
	for i := 0; i < 100; i++ {
		number := Generate()
		assert.Regexp(t, `^\d{4}$`, number, "Does not match 4 digits, apparently")
		for pos, char := range number {
			for j := pos + 1; j < 4; j++ {
				assert.NotEqual(t, string(char), string(number[j]))
			}
		}
	}
}
