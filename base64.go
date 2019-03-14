package EmailWriters

import (
	"encoding/base64"
	"io"
	"strings"
)

func Base64Writer(w io.Writer, text string) (err error) {
	dwr := NewDelimitWriter(w, []byte{0x0d, 0x0a}, 76) // 76 from RFC
	b64Enc := base64.NewEncoder(base64.StdEncoding, dwr)
	reader := strings.NewReader(text)
	_, err = io.Copy(b64Enc, reader)
	if err != nil {
		return err
	}

	return b64Enc.Close()
}
