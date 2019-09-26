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

type thinExecutor struct {
	file *os.File
	w    io.WriteCloser /// write
}
