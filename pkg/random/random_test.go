package random

import (
	"encoding/base64"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randm := New(letters)

	t.Run("1. trivial test to check if it works", func(t *testing.T) {
		randStr := randm.Get(4)
		t.Logf("got :> %v\n", randStr)
		if len(randStr) != 4 {
			t.Errorf("got %v of length %v: wanted length == %v\n", randStr, len(randStr), 4)
		}
	})

	var numsAndLetters = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyx!@!#$%&*")
	randm1 := New(numsAndLetters)
	t.Run("2. include numbers too", func(t *testing.T) {
		randStr := randm1.Get(32)
		t.Logf("got : %v of length %v\n", randStr, len(randStr))
		b64 := base64.StdEncoding.EncodeToString([]byte(randStr))
		t.Logf("base64 encoding: %v\n", b64)
		if len(randStr) != 32 {
			t.Errorf("got %v of length %v: wanted length == %v\n", randStr, len(randStr), 32)
		}
	})

}
