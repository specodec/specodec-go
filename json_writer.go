package specodec

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type JsonWriter struct {
	sb        strings.Builder
	firstItem []bool
}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{}
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

func fmtFloat64(value float64) string {
	s := strconv.FormatFloat(value, 'f', -1, 64)
	if strings.HasSuffix(s, ".0") {
		s = s[:len(s)-2]
	}
	return s
}

func (w *JsonWriter) WriteFloat32(value float32) {
	if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
		panic("float32: NaN/Infinity not valid JSON")
	}
	w.sb.WriteString(fmtFloat32(value))
}

func (w *JsonWriter) WriteFloat64(value float64) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		panic("float64: NaN/Infinity not valid JSON")
	}
	w.sb.WriteString(fmtFloat64(value))
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

func (w *JsonWriter) BeginObject(_ int) {
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
	w.sb.WriteString(`":`)
}

func (w *JsonWriter) EndObject() {
	w.firstItem = w.firstItem[:len(w.firstItem)-1]
	w.sb.WriteByte('}')
}

func (w *JsonWriter) BeginArray(_ int) {
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
