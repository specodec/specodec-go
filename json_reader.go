package specodec

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type JsonReader struct {
	src        string
	pos        int
	firstField []bool
	firstElem  []bool
}

func NewJsonReader(data []byte) *JsonReader {
	return &JsonReader{src: string(data)}
}

func (r *JsonReader) Pos() int { return r.pos }

func (r *JsonReader) ws() {
	for r.pos < len(r.src) {
		c := r.src[r.pos]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			r.pos++
		} else {
			break
		}
	}
}

func (r *JsonReader) peek() byte {
	r.ws()
	if r.pos >= len(r.src) {
		panic("json: unexpected end of input")
	}
	return r.src[r.pos]
}

func (r *JsonReader) read() byte {
	r.ws()
	if r.pos >= len(r.src) {
		panic("json: unexpected end of input")
	}
	b := r.src[r.pos]
	r.pos++
	return b
}

func (r *JsonReader) expect(ch byte) {
	got := r.read()
	if got != ch {
		panic(fmt.Sprintf("json: expected '%c', got '%c' at position %d", ch, got, r.pos-1))
	}
}

func (r *JsonReader) parseString() string {
	r.expect('"')
	var sb strings.Builder
	for r.pos < len(r.src) {
		c := r.src[r.pos]
		if c == '"' {
			r.pos++
			return sb.String()
		}
		if c == '\\' {
			r.pos++
			if r.pos >= len(r.src) {
				panic("json: unexpected end of string escape")
			}
			esc := r.src[r.pos]
			switch esc {
			case '"':
				sb.WriteByte('"')
				r.pos++
			case '\\':
				sb.WriteByte('\\')
				r.pos++
			case '/':
				sb.WriteByte('/')
				r.pos++
			case 'b':
				sb.WriteByte('\b')
				r.pos++
			case 'f':
				sb.WriteByte('\f')
				r.pos++
			case 'n':
				sb.WriteByte('\n')
				r.pos++
			case 'r':
				sb.WriteByte('\r')
				r.pos++
			case 't':
				sb.WriteByte('\t')
				r.pos++
			case 'u':
				r.pos++
				if r.pos+4 > len(r.src) {
					panic("json: incomplete unicode escape")
				}
				hex := r.src[r.pos : r.pos+4]
				cp, err := strconv.ParseInt(hex, 16, 32)
				if err != nil {
					panic(fmt.Sprintf("json: invalid unicode escape \\u%s", hex))
				}
				r.pos += 4
				if cp >= 0xD800 && cp <= 0xDBFF {
					if r.pos+6 <= len(r.src) && r.src[r.pos] == '\\' && r.src[r.pos+1] == 'u' {
						r.pos += 2
						hex2 := r.src[r.pos : r.pos+4]
						low, err := strconv.ParseInt(hex2, 16, 32)
						if err != nil {
							panic(fmt.Sprintf("json: invalid low surrogate \\u%s", hex2))
						}
						r.pos += 4
						if low >= 0xDC00 && low <= 0xDFFF {
							cp = 0x10000 + (cp-0xD800)*0x400 + (low - 0xDC00)
						} else {
							panic("json: expected low surrogate after high surrogate")
						}
					} else {
						panic("json: expected low surrogate after high surrogate")
					}
				}
				sb.WriteRune(rune(cp))
			default:
				panic(fmt.Sprintf("json: invalid escape character '\\%c'", esc))
			}
		} else if c < 0x20 {
			panic(fmt.Sprintf("json: unescaped control character U+%04X", c))
		} else {
			sb.WriteByte(c)
			r.pos++
		}
	}
	panic("json: unterminated string")
}

func (r *JsonReader) parseNumberRaw() string {
	start := r.pos
	if r.pos < len(r.src) && r.src[r.pos] == '-' {
		r.pos++
	}
	if r.pos >= len(r.src) {
		panic("json: unexpected end of number")
	}
	if r.src[r.pos] == '0' {
		r.pos++
	} else if r.src[r.pos] >= '1' && r.src[r.pos] <= '9' {
		r.pos++
		for r.pos < len(r.src) && r.src[r.pos] >= '0' && r.src[r.pos] <= '9' {
			r.pos++
		}
	} else {
		panic("json: invalid number")
	}
	if r.pos < len(r.src) && r.src[r.pos] == '.' {
		r.pos++
		if r.pos >= len(r.src) || r.src[r.pos] < '0' || r.src[r.pos] > '9' {
			panic("json: invalid number fraction")
		}
		for r.pos < len(r.src) && r.src[r.pos] >= '0' && r.src[r.pos] <= '9' {
			r.pos++
		}
	}
	if r.pos < len(r.src) && (r.src[r.pos] == 'e' || r.src[r.pos] == 'E') {
		r.pos++
		if r.pos < len(r.src) && (r.src[r.pos] == '+' || r.src[r.pos] == '-') {
			r.pos++
		}
		if r.pos >= len(r.src) || r.src[r.pos] < '0' || r.src[r.pos] > '9' {
			panic("json: invalid number exponent")
		}
		for r.pos < len(r.src) && r.src[r.pos] >= '0' && r.src[r.pos] <= '9' {
			r.pos++
		}
	}
	return r.src[start:r.pos]
}

func (r *JsonReader) ReadString() string {
	return r.parseString()
}

func (r *JsonReader) ReadBool() bool {
	ch := r.peek()
	if ch == 't' {
		for _, c := range "true" {
			if r.read() != byte(c) {
				panic("json: expected 'true'")
			}
		}
		return true
	}
	if ch == 'f' {
		for _, c := range "false" {
			if r.read() != byte(c) {
				panic("json: expected 'false'")
			}
		}
		return false
	}
	panic(fmt.Sprintf("json: expected bool, got '%c'", ch))
}

func (r *JsonReader) ReadInt32() int32 {
	raw := r.parseNumberRaw()
	v, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("json: invalid int32: %s", raw))
	}
	return int32(v)
}

func (r *JsonReader) ReadInt64() int64 {
	ch := r.peek()
	if ch == '"' {
		s := r.parseString()
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("json: invalid int64: %s", s))
		}
		return v
	}
	raw := r.parseNumberRaw()
	v, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("json: invalid int64: %s", raw))
	}
	return v
}

func (r *JsonReader) ReadUint32() uint32 {
	raw := r.parseNumberRaw()
	v, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("json: invalid uint32: %s", raw))
	}
	return uint32(v)
}

func (r *JsonReader) ReadUint64() uint64 {
	ch := r.peek()
	if ch == '"' {
		s := r.parseString()
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("json: invalid uint64: %s", s))
		}
		return v
	}
	raw := r.parseNumberRaw()
	v, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("json: invalid uint64: %s", raw))
	}
	return v
}

func (r *JsonReader) ReadFloat32() float32 {
	raw := r.parseNumberRaw()
	v, err := strconv.ParseFloat(raw, 32)
	if err != nil {
		panic(fmt.Sprintf("json: invalid float32: %s", raw))
	}
	return float32(v)
}

func (r *JsonReader) ReadFloat64() float64 {
	raw := r.parseNumberRaw()
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		panic(fmt.Sprintf("json: invalid float64: %s", raw))
	}
	return v
}

func (r *JsonReader) ReadNull() {
	for _, c := range "null" {
		if r.read() != byte(c) {
			panic("json: expected 'null'")
		}
	}
}

func (r *JsonReader) ReadBytes() []byte {
	s := r.parseString()
	v, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("json: invalid base64: %v", err))
	}
	return v
}

func (r *JsonReader) ReadEnum() string {
	return r.parseString()
}

func (r *JsonReader) BeginObject() {
	r.expect('{')
	r.firstField = append(r.firstField, true)
}

func (r *JsonReader) HasNextField() bool {
	ch := r.peek()
	if ch == '}' {
		r.firstField = r.firstField[:len(r.firstField)-1]
		return false
	}
	top := len(r.firstField) - 1
	if !r.firstField[top] {
		if ch != ',' {
			panic(fmt.Sprintf("json: expected ',' or '}', got '%c'", ch))
		}
		r.pos++
	} else {
		r.firstField[top] = false
	}
	return true
}

func (r *JsonReader) ReadFieldName() string {
	key := r.parseString()
	r.ws()
	if r.pos < len(r.src) && r.src[r.pos] == ':' {
		r.pos++
	} else {
		panic(fmt.Sprintf("json: expected ':' after field name '%s'", key))
	}
	return key
}

func (r *JsonReader) EndObject() {
	r.expect('}')
}

func (r *JsonReader) BeginArray() {
	r.expect('[')
	r.firstElem = append(r.firstElem, true)
}

func (r *JsonReader) HasNextElement() bool {
	ch := r.peek()
	if ch == ']' {
		r.firstElem = r.firstElem[:len(r.firstElem)-1]
		return false
	}
	top := len(r.firstElem) - 1
	if !r.firstElem[top] {
		if ch != ',' {
			panic(fmt.Sprintf("json: expected ',' or ']', got '%c'", ch))
		}
		r.pos++
	} else {
		r.firstElem[top] = false
	}
	return true
}

func (r *JsonReader) EndArray() {
	r.expect(']')
}

func (r *JsonReader) IsNull() bool {
	ch := r.peek()
	return ch == 'n'
}

func (r *JsonReader) Skip() {
	r.ws()
	if r.pos >= len(r.src) {
		panic("json: unexpected end of input")
	}
	ch := r.src[r.pos]
	switch ch {
	case '"':
		r.pos++
		for r.pos < len(r.src) {
			if r.src[r.pos] == '\\' {
				r.pos += 2
			} else if r.src[r.pos] == '"' {
				r.pos++
				return
			} else {
				r.pos++
			}
		}
		panic("json: unterminated string in skip")
	case '{':
		r.BeginObject()
		for r.HasNextField() {
			r.ReadFieldName()
			r.Skip()
		}
		r.EndObject()
	case '[':
		r.BeginArray()
		for r.HasNextElement() {
			r.Skip()
		}
		r.EndArray()
	case 't':
		for _, c := range "true" {
			if r.read() != byte(c) {
				panic("json: skip expected true")
			}
		}
	case 'f':
		for _, c := range "false" {
			if r.read() != byte(c) {
				panic("json: skip expected false")
			}
		}
	case 'n':
		for _, c := range "null" {
			if r.read() != byte(c) {
				panic("json: skip expected null")
			}
		}
	default:
		if (ch >= '0' && ch <= '9') || ch == '-' {
			r.parseNumberRaw()
		} else {
			panic(fmt.Sprintf("json: unexpected character '%c' in skip", ch))
		}
	}
}
