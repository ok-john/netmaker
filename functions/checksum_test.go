package functions

import "testing"

func TestChecksum(t *testing.T) {
	testFile := "test_data/foo"
	var testCSum checksum = "2f073388335c901c62f4543f60459e29327cb3c157f700eb76f5b77ac73e980c"
	type args struct {
		of      string
		matches checksum
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"test basic",
			args{
				testFile,
				testCSum,
			},
			true,
			false,
		},
		// These tests are only for linux, hmm...
		// what's the best way to switch tests by OS?
		{
			"test wg",
			args{
				"/usr/bin/wg",
				__wg_linux_sha256sum,
			},
			true,
			false,
		},
		{
			"test wg-quick",
			args{
				"/usr/bin/wg-quick",
				__wg_quick_linux_sha256sum,
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareSHA256Sum(tt.args.of, tt.args.matches)
			if (err != nil) != tt.wantErr {
				t.Errorf("Checksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
