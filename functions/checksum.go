package functions

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

const (
	__wg_linux_sha256sum         checksum = "71b28d0986cb7e3e496bf0e49735056aa28b16f4800227363c2869d296a0bda3"
	__wg_freebsd_sha256sum       checksum = "774451f56e0f85e147cd4a952a67b7a3852fd2c4e0ccf1bf0e6d1d5a8512231c"
	__wg_darwin_sha256sum        checksum = ""
	__wg_windows_sha256sum       checksum = ""
	__wg_quick_linux_sha256sum   checksum = "cc50912aab62686e60c4844097163567ae7ee1ccf5bd29680a53379810813edc"
	__wg_quick_freebsd_sha256sum checksum = "db33ebd08a6ae9bab620a8c2830822f517d7b8d284d7691c1b10d35b5eeefa1c"
	__wg_quick_darwin_sha256sum  checksum = ""
	__wg_quick_windows_sha256sum checksum = ""
)

type (
	// A hexadecimal encoded checksum.
	checksum   string
	dependency struct {
		name      string
		checksums map[string]checksum
	}
)

var (
	wg dependency = dependency{
		name: "wg",
		checksums: map[string]checksum{
			"linux":   __wg_linux_sha256sum,
			"freebsd": __wg_freebsd_sha256sum,
			"darwin":  __wg_darwin_sha256sum,
			"windows": __wg_windows_sha256sum,
		},
	}
	wg_quick dependency = dependency{
		name: "wg-quick",
		checksums: map[string]checksum{
			"linux":   __wg_quick_linux_sha256sum,
			"freebsd": __wg_quick_freebsd_sha256sum,
			"darwin":  __wg_quick_darwin_sha256sum,
			"windows": __wg_quick_windows_sha256sum,
		},
	}
	required []dependency = []dependency{wg, wg_quick}
)

// Returns the runtime specific checksum for a known dependency.
func (d dependency) knownChecksum() (checksum, error) {
	if cksm, ok := d.checksums[runtime.GOOS]; ok && len(cksm) > 0 {
		return cksm, nil
	}
	return "", errors.New("invalid or non-existent checksum")
}

// Checks if a dependency exists, returning the path and boolean depicting its existence or an empty string.
func (d dependency) exists() (string, bool) {
	if path, err := exec.LookPath(d.name); err == nil {
		return path, true
	}
	return "", false
}

// Returns the hex-decoded checksum for comparison.
func (c checksum) decoded() []byte {
	res := make([]byte, len(c)/2)
	hex.Decode(res, []byte(c))
	return res
}

// Compares the sha256sum of a file against a known checksum.
func (d dependency) checksum() (bool, error) {
	// Ensure dependency exists
	path, exists := d.exists()
	if !exists {
		fmt.Printf("\nmissing required dependency: %s", d.name)
		return false, errors.New("dependency non-existent")
	}
	// Grab a known runtime specific known checksum
	knownChecksum, err := d.knownChecksum()
	if err != nil {
		return false, err
	}

	// take the sha256sum of the dependency
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return false, err
	}
	sum := sha256.Sum256(raw)

	// compare the known checksum with the recently hashed dependency
	return subtle.ConstantTimeCompare(sum[:], knownChecksum.decoded()) == 1, nil
}
