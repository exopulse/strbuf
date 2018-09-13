// Package strbuf contains string buffer type which allows multiple operations on the same buffer while not forcing
// user to check for errors after each operation.
// Each line is separated using new-line string provided during buffer creation.
//
// Once done with the buffer, user can interrogate error state to see if there were errors during string build.
package strbuf

import (
	"fmt"
	"io"
	"strings"
)

// Buffer serves as intermediate storage for string based operations.
// Use NewBuffer() to instantiate new buffer instance.
type Buffer struct {
	newLine string
	builder strings.Builder
	error   error
}

// NewBuffer creates empty buffer. Each line is separated with new-line string.
func NewBuffer(newLine string) *Buffer {
	return &Buffer{newLine: newLine}
}

// Append appends string.
func (b *Buffer) Append(s string) *Buffer {
	return b.append(s)
}

// AppendAll appends all strings.
func (b *Buffer) AppendAll(s []string) *Buffer {
	for _, e := range s {
		b.append(e)
	}

	return b
}

// Appendf appends formatted string.
func (b *Buffer) Appendf(format string, args ...interface{}) *Buffer {
	return b.append(fmt.Sprintf(format, args...))
}

// AppendTitle appends string and underlines it with specified string.
func (b *Buffer) AppendTitle(s string, underline string) *Buffer {
	b.append(s)

	return b.append(strings.Repeat(underline, len(s)))
}

// NewLine appends new line.
func (b *Buffer) NewLine() *Buffer {
	return b.append("")
}

// EnsureEmptyLine ensures there is one empty line at the tail of this buffer.
func (b *Buffer) EnsureEmptyLine() *Buffer {
	if strings.HasSuffix(b.builder.String(), b.newLine) {
		return b
	}

	return b.NewLine()
}

// Error returns error if any.
func (b *Buffer) Error() error {
	return b.error
}

func (b *Buffer) append(s string) *Buffer {
	if b.error != nil {
		return b
	}

	if b.builder.Len() > 0 && len(b.newLine) > 0 {
		if _, b.error = b.builder.WriteString(b.newLine); b.error != nil {
			return b
		}
	}

	_, b.error = b.builder.WriteString(s)

	return b
}

// String returns the accumulated string. Make sure to check for any accumulated errors first.
func (b *Buffer) String() string {
	if b.error != nil {
		return ""
	}

	return b.builder.String()
}

// Bytes return contents of this buffer as bytes. Make sure to check for any accumulated errors first.
func (b *Buffer) Bytes() []byte {
	if b.error != nil {
		return []byte{}
	}

	return []byte(b.builder.String())
}

// Write writes contents of this buffer to writer.
func (b *Buffer) Write(w io.Writer) error {
	if b.error != nil {
		return b.error
	}

	_, err := w.Write(b.Bytes())

	return err
}
