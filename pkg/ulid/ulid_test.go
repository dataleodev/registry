package ulid

import "testing"

func TestUlidIDProvider(t *testing.T) {
	t.Run("observe generated token", func(t *testing.T) {
		up := New()
		id, err := up.ID()
		if err != nil {
			t.Fatalf("error while generating id %v\n", err)
		}

		t.Logf("the id: %v\n", id)
	})
}
