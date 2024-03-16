package download

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	gio "io"

	"github.com/lspaccatrosi16/go-cli-tools/credential"
	"github.com/lspaccatrosi16/tools/lib/bupload/io"
	"github.com/lspaccatrosi16/tools/lib/bupload/provider"
)

func Download(cred credential.Credential) error {
	fmt.Println("Download File")
	bucket := io.GetBucket()

	provider, err := provider.GetProvider(cred, bucket)
	if err != nil {
		return err
	}

	object := io.GetObject()

	file, err := provider.GetFile(object)

	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(file)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fullPath := filepath.Join(wd, object)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer f.Close()

	gio.Copy(f, buf)

	return nil
}
