package command

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/sha3"
)

// baulk command module hash command

type hashOptions struct {
	sub     string
	path    string
	verbose bool
}

func (ho *hashOptions) Invoke(val int, oa, raw string) error {
	switch val {
	case 'v':
	case 'h':
	case 'V':
		ho.verbose = true
	case 'f':
		ho.path = oa
	case 'c':
		ho.sub = oa
	}
	return nil
}

func calculationFileHash(subcmd, path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 1
	}
	defer file.Close()
	var h hash.Hash
	switch subcmd {
	case "md5sum":
		h = md5.New()
	case "sha1sum":
		h = sha1.New()
	case "sha224sum":
		h = sha256.New224()
	case "sha256sum":
		h = sha256.New()
	case "sha384sum":
		h = sha512.New384()
	case "sha512sum":
		h = sha512.New()
	case "sha3-224sum":
		h = sha3.New224()
	case "sha3-256sum":
		h = sha3.New256()
	case "sha3-384sum":
		h = sha3.New384()
	case "sha3-512sum":
		h = sha3.New512()
	default:
		return 1
	}
	if _, err := io.Copy(h, file); err != nil {
		return 1
	}
	hx := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("%s %s\n", hx, filepath.Base(path))
	return 0
}

// HashCalculate impl hash fun
func HashCalculate(args []string) int {
	var ae ArgvEngine
	ae.Add("version", NOARG, 'v')
	ae.Add("help", NOARG, 'h')
	ae.Add("cmd", REQUIRED, 'c')
	var ho hashOptions
	if err := ae.Execute(args, &ho); err != nil {
		fmt.Fprintf(os.Stderr, "ParseArgv: \x1b[31m%s\x1b[0m\n", err.Error())
		os.Exit(1)
	}
	ua := ae.Unresolved()
	if ho.sub == "" && len(ua) > 0 {
		ho.sub = ua[0]
		ua = ua[1:]
	}
	if ho.path == "" && len(ua) > 0 {
		ho.path = ua[0]
	}
	if ho.sub == "" || ho.path == "" {
		fmt.Fprintf(os.Stderr, "unknown hash subcommand: [%v]\n", ae.Unresolved())
		os.Exit(1)
	}
	return calculationFileHash(ho.sub, ho.path)
}
