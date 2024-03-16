package upload

import (
	"bytes"
	"fmt"
	gio "io"
	"os"
	"path/filepath"

	"github.com/lspaccatrosi16/go-cli-tools/credential"
	"github.com/lspaccatrosi16/go-cli-tools/input"
	"github.com/lspaccatrosi16/tools/lib/bupload/io"
	"github.com/lspaccatrosi16/tools/lib/bupload/provider"
)

func Upload(cred credential.Credential) error {
	fmt.Println("Upload File")
	bucket := io.GetBucket()
	provider, err := provider.GetProvider(cred, bucket)

	if err != nil {
		return err
	}

	path := io.GetPath()

	stats, err := os.Stat(path)

	if err != nil {
		return err
	}

	sizeMib := float64(stats.Size()) / (1024 * 1024)
	objKey := filepath.Base(path)

	buf := bytes.NewBuffer(nil)

	fmt.Fprintf(buf, "Object Key: %s\n", objKey)
	fmt.Fprintf(buf, "Size      : %.2fMiB\n", sizeMib)

	fmt.Println(buf.String())

	proceed, err := input.GetConfirmSelection("Proceed with upload")
	if err != nil {
		return err
	}

	if !proceed {
		return fmt.Errorf("cancelled")
	}

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	fileBuf := bytes.NewBuffer(nil)

	gio.Copy(fileBuf, f)

	err = provider.UploadFile(objKey, fileBuf.Bytes())
	return err
}
