package main

import (
	"bytes"
	"errors"
	"strings"

	perr "github.com/pkg/errors"
)

type Base64 struct {
	Table      [64]byte
	Padding    byte
	paddingCnt int
}

func NewBase64() *Base64 {
	table := [64]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/'}

	return &Base64{
		Table: table,
	}
}

func (b6 *Base64) charAt(index byte) byte {
	return b6.Table[index]
}

func (b6 *Base64) Encode(input []byte) string {
	var buffer bytes.Buffer
	for {
		if len(input) == 0 {
			break
		}
		if len(input) < 3 {
			if len(input) == 1 {
				firstLow2 := input[0] & 0x3
				buffer.WriteByte(input[0] >> 2)
				buffer.WriteByte((firstLow2 << 4) | 0x0)
				b6.paddingCnt = 2
			} else if len(input) == 2 {
				firstLow2 := input[0] & 0x3
				buffer.WriteByte(input[0] >> 2)
				midLow4 := input[1] & 0xf
				buffer.WriteByte((input[1] >> 4) | firstLow2<<4)
				buffer.WriteByte(midLow4 << 2)
				b6.paddingCnt = 1
			}
			break
		}
		// three bytes as a iteration
		firstLow2 := input[0] & 0x3
		one := input[0] >> 2
		buffer.WriteByte(one)
		midLow4 := input[1] & 0xf
		two := (input[1] >> 4) | firstLow2<<4
		buffer.WriteByte(two)
		lastHi2 := input[2] >> 6
		three := (midLow4 << 2) | lastHi2
		buffer.WriteByte(three)
		four := input[2] & 0x3f
		buffer.WriteByte(four)
		// fmt.Printf("input[2]: %b\n", input[2])
		// fmt.Printf("one: %v, two: %v, three: %v, four: %v\n", one, two,
		// three, four)

		input = input[3:]
	}
	var res bytes.Buffer
	for _, b := range buffer.Bytes() {
		res.WriteByte(b6.charAt(b))
	}
	// add padding
	for b6.paddingCnt > 0 {
		res.WriteByte('=')
		b6.paddingCnt--
	}
	return res.String()
}

func (b6 *Base64) indexOf(c byte) (uint8, error) {
	idx := bytes.IndexByte(b6.Table[:], c)
	if idx == -1 {
		return 0, perr.Wrapf(ErrMalformedBase64, "byte: %X", c)
	}
	return uint8(idx), nil
}

var (
	ErrMalformedBase64 = errors.New("malformed base64 character")
)

// A base64 decoder usually works on a window of 4 bytes.
func (bs *Base64) Decode(s string) ([]byte, error) {
	s = strings.TrimSuffix(s, "=") // remove trailing '='
	var buffer bytes.Buffer
	for {
		if len(s) <= 1 {
			break
		}
		var b1, b2, b3, b4 uint8
		var err error
		if len(s) >= 2 {
			b1, err = bs.indexOf(s[0])
			if err != nil {
				return nil, err
			}
			b2, err = bs.indexOf(s[1])
			if err != nil {
				return nil, err
			}
			buffer.WriteByte((b1 << 2) | b2>>4)
			if len(s) == 2 {
				break
			}
		}
		if len(s) >= 3 {
			b2Low4 := b2 & 0xf
			b3, err = bs.indexOf(s[2])
			if err != nil {
				return nil, err
			}
			buffer.WriteByte((b2Low4 << 4) | b3>>2)
			if len(s) == 3 {
				break
			}
		}

		// len(s) >= 4
		b4, err = bs.indexOf(s[3])
		if err != nil {
			return nil, err
		}
		b3Low2 := b3 & 0x3
		buffer.WriteByte((b3Low2 << 6) | b4)

		s = s[4:]
	}

	return buffer.Bytes(), nil
}
