package jwt

import (
	"github.com/dataleodev/registry"
	"testing"
)

func TestIssueToken(t *testing.T) {
	t.Run("generate access token", func(t *testing.T) {
		to := NewTokenizer()
		token := registry.NewKey("piusalfred", "access")
		tStr, err := to.Issue(token)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		t.Logf("token : %v\n", tStr)
		tokenRecovered, err := to.Parse(tStr)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		t.Logf("%v\n", tokenRecovered)
	})

}
