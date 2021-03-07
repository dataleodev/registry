package bcrypt

import (
	"testing"
)

func TestCompare(t *testing.T) {
	bc := New()
	password := "root"

	t.Run("hasher", func(t *testing.T) {
		hash, err := bc.Hash(password)
		if err != nil {
			t.Errorf("could not hash %v\n", err)
		}

		//	err = bc.Compare(hash,password)
		//t.Logf("%v\n",err)

		//t.Logf("password then hash: %v\n",err)
		err = bc.Compare(password, hash)
		t.Logf("%v\n", err)

	})

}
