package scanner

import (
	"io"
	"strings"
	"unicode/utf8"
)

type FileStats struct {
	BytesCount int64
	CharCount  int64
	LineCount  int64
	WordCount  int64
}

func Scan(r io.Reader) (FileStats, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return FileStats{}, err
	}

	s := string(data)
	lines := strings.Split(s, "\n")
	lineCount := len(lines)
	if len(s) > 0 && s[len(s)-1] == '\n' {
		lineCount--
	}

	words := strings.Fields(s)

	return FileStats{
		BytesCount: int64(len(data)),
		LineCount:  int64(lineCount),
		CharCount:  int64(utf8.RuneCount(data)),
		WordCount:  int64(len(words)),
	}, nil

}
