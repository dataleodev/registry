package uuid

import "testing"

func TestUUIDProviderID(t *testing.T) {
	t.Run("see genarated uuids", func(t *testing.T) {
		up := New()
		str, err := up.ID()

		if err != nil {
			t.Errorf("error occurred %v\n", err)
		}

		t.Logf("uuid: %v\n", str)
	})
}
