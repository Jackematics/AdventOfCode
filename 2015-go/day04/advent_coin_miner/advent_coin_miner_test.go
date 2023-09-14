package advent_coin_miner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var advent_coin_miner_test_cases = []struct {
	input          string
	valid_md5_hash int
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func TestFindValidMd5Hash(t *testing.T) {
	for _, test_case := range advent_coin_miner_test_cases {
		assert.Equal(t, test_case.valid_md5_hash, FindFirstValidMd5Hash(test_case.input))
	}
}
