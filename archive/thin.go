package archive

import (
	"io"
	"os"
)

// Baulk unarchive some files.
// first detect file type
// if gz/xz/zst/br detect and proxy to tar unarchive

// zstd 0x28 0xb5 0x2f 0xfd
// #define ZSTD_MAGICNUMBER            0xFD2FB528    /* valid since v0.8.0 */

// IsZstd Buffer
func IsZstd(b []byte) bool {
	return len(b) > 4 &&
		b[0] == 0x28 && b[1] == 0xb5 && b[2] == 0x2f && b[3] == 0xfd
}

// IsZip buffer is zip
func IsZip(b []byte) bool {
	return len(b) > 3 &&
		b[0] == 0x50 && b[1] == 0x4B &&
		(b[2] == 0x3 || b[2] == 0x5 || b[2] == 0x7) &&
		(b[3] == 0x4 || b[3] == 0x6 || b[3] == 0x8)
}

// IsTar todo
func IsTar(b []byte) bool {
	return len(b) > 261 &&
		b[257] == 0x75 && b[258] == 0x73 &&
		b[259] == 0x74 && b[260] == 0x61 &&
		b[261] == 0x72
}

// IsRar todo
func IsRar(b []byte) bool {
	return len(b) > 6 &&
		b[0] == 0x52 && b[1] == 0x61 && b[2] == 0x72 &&
		b[3] == 0x21 && b[4] == 0x1A && b[5] == 0x7 &&
		(b[6] == 0x0 || b[6] == 0x1)
}

// IsGz todo
func IsGz(b []byte) bool {
	return len(b) > 2 &&
		b[0] == 0x1F && b[1] == 0x8B && b[2] == 0x8
}

// IsBz2 todo
func IsBz2(b []byte) bool {
	return len(b) > 2 &&
		b[0] == 0x42 && b[1] == 0x5A && b[2] == 0x68
}

// Is7Z todo
func Is7Z(b []byte) bool {
	return len(b) > 5 &&
		b[0] == 0x37 && b[1] == 0x7A && b[2] == 0xBC &&
		b[3] == 0xAF && b[4] == 0x27 && b[5] == 0x1C
}

type thinExecutor struct {
	file *os.File
	w    io.WriteCloser /// write
}

func detectFileType(file *os.File) (io.WriteCloser, error) {

	return nil, nil
}

func newThinExecutor(path string) (*thinExecutor, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	w, err := detectFileType(file)
	if err != nil {
		_ = file.Close()
		return nil, err
	}
	return &thinExecutor{file: file, w: w}, nil
}
