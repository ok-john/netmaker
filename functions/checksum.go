package functions

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"io"
	"os"
)

const (
	__wg_linux_sha256sum         checksum = "71b28d0986cb7e3e496bf0e49735056aa28b16f4800227363c2869d296a0bda3"
	__wg_quick_linux_sha256sum   checksum = "cc50912aab62686e60c4844097163567ae7ee1ccf5bd29680a53379810813edc"
	__wg_freebsd_sha256sum       checksum = "774451f56e0f85e147cd4a952a67b7a3852fd2c4e0ccf1bf0e6d1d5a8512231c"
	__wg_quick_freebsd_sha256sum checksum = "db33ebd08a6ae9bab620a8c2830822f517d7b8d284d7691c1b10d35b5eeefa1c"
	// TODO: Add Windows Checksums
)

type (
	// A hexadecimal encoded checksum, for comparison, call Bytes().
	checksum string
)

// Returns the hex-decoded checksum for comparison.
func (c checksum) decoded() []byte {
	res := make([]byte, len(c)/2)
	hex.Decode(res, []byte(c))
	return res
}

// Compares the sha256sum of a file against a known checksum.
func CompareSHA256Sum(of string, matches checksum) (bool, error) {
	f, err := os.Open(of)
	if err != nil {
		return false, err
	}
	defer f.Close()
	raw, err := io.ReadAll(f)
	if err != nil {
		return false, err
	}
	sum := sha256.Sum256(raw)
	return subtle.ConstantTimeCompare(sum[:], matches.decoded()[:]) == 1, nil
}
