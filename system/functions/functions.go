package functions

/**
Various helpers
*/

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Log(msg string) {
	log.SetFlags(0)
	log.Print(msg)
}

func BytesToString(bs []byte) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

// SplitString splits a string by whitespaces
func SplitString(stringToSplit string, delimiter string) []string {
	if delimiter == " " {
		return strings.Fields(stringToSplit)
	}
	return strings.Split(stringToSplit, delimiter)
}

func CopyFile(src_path string, dst_path string, executable bool) error {
	sourceFileStat, err := os.Stat(src_path)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src_path)
	}

	source, err := os.Open(src_path)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst_path)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)

	if err != nil {
		return err
	}

	if executable {
		err = os.Chmod(dst_path, 0555)
	}

	return err
}

// FormatStringAsLabel formatting any string to a label form
func FormatStringAsLabel(input string) string {
	output := strings.ReplaceAll(input, " ", "_")

	if len(output) > 63 {
		output = output[0:63]
	}
	return output
}
