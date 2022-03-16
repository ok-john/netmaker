package functions

const (
	__wg_linux_sha256sum         Checksum = "71b28d0986cb7e3e496bf0e49735056aa28b16f4800227363c2869d296a0bda3"
	__wg_quick_linux_sha256sum   Checksum = "cc50912aab62686e60c4844097163567ae7ee1ccf5bd29680a53379810813edc"
	__wg_freebsd_sha256sum       Checksum = "774451f56e0f85e147cd4a952a67b7a3852fd2c4e0ccf1bf0e6d1d5a8512231c"
	__wg_quick_freebsd_sha256sum Checksum = "db33ebd08a6ae9bab620a8c2830822f517d7b8d284d7691c1b10d35b5eeefa1c"
	// TODO: Add Windows Checksums
)

type (
	Checksum string
)

func (c Checksum) Bytes() []byte {
	return []byte(c)
}
