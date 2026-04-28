package specodec

import (
	"encoding/base64"
	"fmt"
	"math"
	"strings"
)

type JsonWriter struct {
	sb        strings.Builder
	firstItem []bool
}

func (w *JsonWriter) escape(s string) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '"':
			w.sb.WriteString(`\"`)
		case '\\':
			w.sb.WriteString(`\\`)
		case '\b':
			w.sb.WriteString(`\b`)
		case '\f':
			w.sb.WriteString(`\f`)
		case '\n':
			w.sb.WriteString(`\n`)
		case '\r':
			w.sb.WriteString(`\r`)
		case '\t':
			w.sb.WriteString(`\t`)
		default:
			if c < 0x20 {
				w.sb.WriteString(fmt.Sprintf(`\u%04x`, c))
			} else {
				w.sb.WriteByte(c)
			}
		}
	}
}

func (w *JsonWriter) WriteString(value string) {
	w.sb.WriteByte('"')
	w.escape(value)
	w.sb.WriteByte('"')
}

func (w *JsonWriter) WriteBool(value bool) {
	if value {
		w.sb.WriteString("true")
	} else {
		w.sb.WriteString("false")
	}
}

func (w *JsonWriter) WriteInt32(value int32) {
	w.sb.WriteString(fmt.Sprintf("%d", value))
}

func (w *JsonWriter) WriteInt64(value int64) {
	w.sb.WriteByte('"')
	w.sb.WriteString(fmt.Sprintf("%d", value))
	w.sb.WriteByte('"')
}

func (w *JsonWriter) WriteUint32(value uint32) {
	w.sb.WriteString(fmt.Sprintf("%d", value))
}

func (w *JsonWriter) WriteUint64(value uint64) {
	w.sb.WriteByte('"')
	w.sb.WriteString(fmt.Sprintf("%d", value))
	w.sb.WriteByte('"')
}

func (w *JsonWriter) WriteFloat32(value float32) {
	v := float64(value)
	if math.IsNaN(v) || math.IsInf(v, 0) {
		panic("float32: NaN/Infinity not valid JSON")
	}
	w.sb.WriteString(fmt.Sprintf("%g", v))
}

func (w *JsonWriter) WriteFloat64(value float64) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		panic("float64: NaN/Infinity not valid JSON")
	}
	w.sb.WriteString(fmt.Sprintf("%g", value))
}

func (w *JsonWriter) WriteNull() {
	w.sb.WriteString("null")
}

func (w *JsonWriter) WriteBytes(value []byte) {
	w.sb.WriteByte('"')
	encoded := base64.StdEncoding.EncodeToString(value)
	w.sb.WriteString(encoded)
	w.sb.WriteByte('"')
}

func (w *JsonWriter) WriteEnum(value string) {
	w.sb.WriteByte('"')
	w.escape(value)
	w.sb.WriteByte('"')
}

func (w *JsonWriter) BeginObject() {
	w.sb.WriteByte('{')
	w.firstItem = append(w.firstItem, true)
}

func (w *JsonWriter) WriteField(name string) {
	top := len(w.firstItem) - 1
	if !w.firstItem[top] {
		w.sb.WriteByte(',')
	}
	w.firstItem[top] = false
	w.sb.WriteByte('"')
	w.escape(name)
	w.sb.WriteString(`:`)
}

func (w *JsonWriter) EndObject() {
	w.firstItem = w.firstItem[:len(w.firstItem)-1]
	w.sb.WriteByte('}')
}

func (w *JsonWriter) BeginArray() {
	w.sb.WriteByte('[')
	w.firstItem = append(w.firstItem, true)
}

func (w *JsonWriter) NextElement() {
	top := len(w.firstItem) - 1
	if !w.firstItem[top] {
		w.sb.WriteByte(',')
	}
	w.firstItem[top] = false
}

func (w *JsonWriter) EndArray() {
	w.firstItem = w.firstItem[:len(w.firstItem)-1]
	w.sb.WriteByte(']')
}

func (w *JsonWriter) ToBytes() []byte {
	return []byte(w.sb.String())
}
