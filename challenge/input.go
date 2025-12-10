package challenge

import (
	"bufio"
	"io"
	"iter"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"

	"github.com/nlowe/aoc2025/util"
)

// InputFile returns an io.Reader for the file pointed at by the --input flag. If the flag is not specified, it looks
// for a file named "input.txt" in the same package as the caller.
func InputFile() io.Reader {
	path := viper.GetString("input")
	if path == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to determine input path, provide it with -i instead")
		}

		path = filepath.Join(filepath.Dir(f), "input.txt")
	}

	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return r
}

// Raw returns the contents of the provided io.Reader as one giant string
func Raw(r io.Reader) string {
	var buf strings.Builder
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

// Lines returns an iter.Seq[string] over all lines in the provided io.Reader.
func Lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				panic(err)
			}

			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

// SectionsOf returns sections of input from the provided io.Reader delimited by the provided string, exclusive. The
// header up to the first delimiter is returned as the first element in the sequence. Empty sections are omitted.
func SectionsOf(r io.Reader, delim string) iter.Seq[string] {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		// Based off of bufio.ScanLines https://cs.opensource.google/go/go/+/refs/tags/go1.23.3:src/bufio/scan.go;l=355
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), delim); i >= 0 {
			// We have a full delim-terminated section.
			return i + len(delim), data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated section. Return it.
		if atEOF {
			return len(data), data, nil
		}

		// Request more data.
		return 0, nil, nil
	})

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				panic(err)
			}

			tok := scanner.Text()
			if tok == "" {
				continue
			}

			if !yield(tok) {
				return
			}
		}
	}
}

// Sections returns an iter.Seq[string] over all blocks of lines in the provided io.Reader. Blocks are delimited by two
// newlines, and the resulting section has leading and trailing whitespace trimmed by strings.TrimSpace
func Sections(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		for section := range SectionsOf(r, "\n\n") {
			if !yield(strings.TrimSpace(section)) {
				return
			}
		}
	}

}

// Ints returns an iter.Seq[int] over all lines in the provided io.Reader, converting each line to an int. This method
// panics if conversion of any line fails.
func Ints(r io.Reader) iter.Seq[int] {
	scanner := bufio.NewScanner(r)
	return func(yield func(int) bool) {
		for scanner.Scan() {
			err := scanner.Err()
			if err != nil && err != io.EOF {
				panic(err)
			}

			if !yield(util.MustAtoI(scanner.Text())) {
				return
			}
		}
	}
}

func Fields[T any](line string, via func(s string) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, f := range strings.Fields(line) {
			if !yield(via(f)) {
				return
			}
		}
	}
}
