package functions

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	for _, dep := range required {
		valid, err := dep.checksum()
		if !valid {
			t.Fatalf("\ninvalid checksum for dependency: %s", dep.name)
		}
		if err != nil {
			t.Fatalf("\nerror checking dependency for: %s\n\t%s", dep.name, err)
		}
		t.Logf("\nverified %s dependency against known checksum", dep.name)
	}
}
