package advent_coin_miner

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func FindFirstValidMd5Hash(input string) int {
	count := 0
	var secret_key string

	for {
		hash := md5.Sum([]byte(input + strconv.Itoa(count)))
		secret_key = hex.EncodeToString(hash[:])

		if md5_hash_has_six_leading_zeroes(secret_key) {
			break
		}
		count++
	}

	return count
}

func md5_hash_has_six_leading_zeroes(secret_key string) bool {
	first_five_chars := secret_key[0:6]

	return first_five_chars == "000000"
}
