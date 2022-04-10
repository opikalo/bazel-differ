package internal

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type BazelSourceFileTarget interface {
	GetName() *string
	GetDigest() []byte
}

type bazelSourceFileTarget struct {
	name   *string
	digest []byte
}

func NewBazelSourceFileTarget(name string, digest []byte, filesystem fs.FS,
	workingDirectory string) (BazelSourceFileTarget,
	error) {
	finalDigest := bytes.NewBuffer([]byte{})
	if workingDirectory != "" && strings.HasPrefix(name, "//") {
		filenameSubstring := name[2:]
		filenamePath := strings.Replace(filenameSubstring, ":", "/", 1)
		sourceFile := path.Join(workingDirectory, filenamePath)
		if _, err := os.Stat(sourceFile); !errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does not exist
			contents, err := ioutil.ReadFile(sourceFile)
			if err != nil {
				return nil, err
			}
			finalDigest.Write(contents)
		}
	}
	finalDigest.Write(digest)
	finalDigest.Write([]byte(name))
	checksum := sha256.Sum256(finalDigest.Bytes())
	return &bazelSourceFileTarget{
		name:   &name,
		digest: checksum[:],
	}, nil
}

func (b *bazelSourceFileTarget) GetName() *string {
	return b.name
}

func (b *bazelSourceFileTarget) GetDigest() []byte {
	return b.digest
}
