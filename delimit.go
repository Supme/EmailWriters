package EmailWriters

import "io"

type delimitWriter struct {
	n      int
	cnt    int
	dr     []byte
	writer io.Writer
}

func NewDelimitWriter(writer io.Writer, dr []byte, cnt int) *delimitWriter {
	return &delimitWriter{n: 0, cnt: cnt, dr: dr, writer: writer}
}

func (w *delimitWriter) Write(p []byte) (n int, err error) {
	for i := range p {
		_, err = w.writer.Write(p[i : i+1])
		if err != nil {
			break
		}
		if w.n++; w.n%w.cnt == 0 {
			_, err = w.writer.Write(w.dr)
			if err != nil {
				break
			}
		}
	}
	return w.n, err
}
