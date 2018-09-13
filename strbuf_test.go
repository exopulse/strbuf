package strbuf

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"
)

var sample = `Line1

Test
----------------
row1
row2

value=1234`

func TestNewBuffer_String(t *testing.T) {
	b := populateBuffer(false)

	if b.Error() != nil {
		t.Fatal("unexpected error:", b.Error())
	}

	s := b.String()

	if s != sample {
		t.Fatal("s != sample")
	}
}

func TestNewBuffer_Write(t *testing.T) {
	b := populateBuffer(false)

	if b.Error() != nil {
		t.Fatal("unexpected error:", b.Error())
	}

	buf := bytes.NewBufferString("")

	if err := b.Write(buf); err != nil {
		t.Fatal("unexpected error:", b.Error())
	}

	s := buf.String()

	if s != sample {
		t.Fatal("s != sample")
	}
}

func TestNewBuffer_Error(t *testing.T) {
	b := populateBuffer(true)

	if b.Error() == nil {
		t.Fatal("error expected")
	}

	s := b.String()

	if s != "" {
		t.Fatal("unexpected content:", s)
	}
}

func TestNewBuffer_WriteError(t *testing.T) {
	b := populateBuffer(true)

	if b.Error() == nil {
		t.Fatal("error expected")
	}

	buf := bytes.NewBufferString("")

	if err := b.Write(buf); err == nil {
		t.Fatal("error expected")
	}

	s := buf.String()

	if s != "" {
		t.Fatal("unexpected content:", s)
	}
}

func TestNewBuffer_BytesError(t *testing.T) {
	b := populateBuffer(true)

	if b.Error() == nil {
		t.Fatal("error expected")
	}

	bs := b.Bytes()

	if len(bs) > 0 {
		t.Fatal("unexpected content:", string(bs))
	}
}

func populateBuffer(introduceError bool) *Buffer {
	b := NewBuffer("\n")

	if introduceError {
		b.selfInflictError()
	}

	b.Append("Line1")
	b.EnsureEmptyLine()
	b.EnsureEmptyLine()
	b.AppendTitle("Test", "----")
	b.AppendAll([]string{"row1", "row2"})
	b.NewLine()
	b.Appendf("value=%d", 1234)

	return b
}

func (b *Buffer) selfInflictError() {
	b.error = errors.New("self-inflicted error")
}
