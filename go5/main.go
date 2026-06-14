package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) HandleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("Renamed file:", newFilename)

	return nil
}

// sha1
// prefix SC_
// hmac
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename)) // hash is an array of 32 bytes, we need to convert it to a string
	return hex.EncodeToString(hash[:])      // converting the hash to a string
}

func SCPRefixFilename(filename string) string {
	return "SC_" + filename
}

func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {
	s := &Server{
		filenameTransformFunc: prefixFilename("BOB_"), // you can change this to hashFilename to see the difference
	}

	s.HandleRequest("avatar.jpg")
}
