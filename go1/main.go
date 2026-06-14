package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type HashReader interface {
	io.Reader
	hash() string
}

func main() {
	payload := []byte("Hello world!")
	hashAndBroadcast(NewHashReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println("Hash: ", hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes: ", string(b))

	return nil
}

// ==================================================

// func main() {
// 	payload := []byte("Hello world!")
// 	hashAndBroadcast(bytes.NewReader(payload))
// }

// func hashAndBroadcast(r io.Reader) error {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	hash := sha1.Sum(b)
// 	fmt.Println(hex.EncodeToString(hash[:]))

// 	return broadcast(r)
// }

// func broadcast(r io.Reader) error {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("string of the bytes: ", string(b))

// 	return nil
// }
