package specodec

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type GronWriter struct {
	lines    []string
	segments []string
	nesting  []nestInfo
}

type nestInfo struct {
	depth      int
	arrayIndex int
}

func NewGronWriter() *GronWriter {
	return &GronWriter{segments: []string{"json"}}
}

func (w *GronWriter) buildPath() string {
	r := w.segments[0]
	for i := 1; i < len(w.segments); i++ {
		if strings.HasPrefix(w.segments[i], "[") {
			r += w.segments[i]
		} else {
			r += "." + w.segments[i]
		}
	}
	return r
}

func gronEscape(s string) string {
	var r strings.Builder
	for _, c := range s {
		switch c {
		case '"':
			r.WriteString(`\"`)
		case '\\':
			r.WriteString(`\\`)
		case '\b':
			r.WriteString(`\b`)
		case '\f':
			r.WriteString(`\f`)
		case '\n':
			r.WriteString(`\n`)
		case '\r':
			r.WriteString(`\r`)
		case '\t':
			r.WriteString(`\t`)
		default:
			if c < 0x20 {
				r.WriteString(fmt.Sprintf(`\u%04x`, c))
			} else {
				r.WriteRune(c)
			}
		}
	}
	return r.String()
}

func gronB64(data []byte) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var s strings.Builder
	for i := 0; i < len(data); i += 3 {
		b0 := int(data[i])
		b1, b2 := 0, 0
		if i+1 < len(data) {
			b1 = int(data[i+1])
		}
		if i+2 < len(data) {
			b2 = int(data[i+2])
		}
		s.WriteByte(chars[b0>>2])
		s.WriteByte(chars[((b0&3)<<4)|(b1>>4)])
		if i+1 < len(data) {
			s.WriteByte(chars[((b1&0xF)<<2)|(b2>>6)])
		} else {
			s.WriteByte('=')
		}
		if i+2 < len(data) {
			s.WriteByte(chars[b2&0x3F])
		} else {
			s.WriteByte('=')
		}
	}
	return s.String()
}

func (w *GronWriter) emit(raw string) {
	w.lines = append(w.lines, fmt.Sprintf("%s = %s;", w.buildPath(), raw))
}

func (w *GronWriter) WriteString(value string)         { w.emit(`"` + gronEscape(value) + `"`) }
func (w *GronWriter) WriteBool(value bool)             { w.emit(strconv.FormatBool(value)) }
func (w *GronWriter) WriteInt32(value int32)           { w.emit(strconv.FormatInt(int64(value), 10)) }
func (w *GronWriter) WriteInt64(value int64)           { w.emit(`"` + strconv.FormatInt(value, 10) + `"`) }
func (w *GronWriter) WriteUint32(value uint32)         { w.emit(strconv.FormatUint(uint64(value), 10)) }
func (w *GronWriter) WriteUint64(value uint64)         { w.emit(`"` + strconv.FormatUint(value, 10) + `"`) }
func (w *GronWriter) WriteNull()                       { w.emit("null") }
func (w *GronWriter) WriteBytes(value []byte)          { w.emit(`"` + gronB64(value) + `"`) }

func (w *GronWriter) WriteFloat32(value float32) {
	if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
		panic("float32: NaN/Infinity not valid")
	}
	w.emit(formatFloat32(value))
}

func (w *GronWriter) WriteFloat64(value float64) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		panic("float64: NaN/Infinity not valid")
	}
	w.emit(formatFloat64(value))
}

func (w *GronWriter) BeginObject(fieldCount int) {
	w.lines = append(w.lines, fmt.Sprintf("%s = {};", w.buildPath()))
	w.nesting = append(w.nesting, nestInfo{depth: len(w.segments)})
}

func (w *GronWriter) WriteField(name string) {
	top := &w.nesting[len(w.nesting)-1]
	if len(w.segments) > top.depth {
		w.segments[len(w.segments)-1] = name
	} else {
		w.segments = append(w.segments, name)
	}
}

func (w *GronWriter) EndObject() {
	info := w.nesting[len(w.nesting)-1]
	w.nesting = w.nesting[:len(w.nesting)-1]
	w.segments = w.segments[:info.depth]
}

func (w *GronWriter) BeginArray(elementCount int) {
	w.lines = append(w.lines, fmt.Sprintf("%s = [];", w.buildPath()))
	w.nesting = append(w.nesting, nestInfo{depth: len(w.segments), arrayIndex: -1})
}

func (w *GronWriter) NextElement() {
	info := &w.nesting[len(w.nesting)-1]
	info.arrayIndex++
	seg := fmt.Sprintf("[%d]", info.arrayIndex)
	if len(w.segments) > info.depth {
		w.segments[len(w.segments)-1] = seg
	} else {
		w.segments = append(w.segments, seg)
	}
}

func (w *GronWriter) EndArray() {
	info := w.nesting[len(w.nesting)-1]
	w.nesting = w.nesting[:len(w.nesting)-1]
	w.segments = w.segments[:info.depth]
}

func (w *GronWriter) WriteEnum(value string) { w.emit(`"` + gronEscape(value) + `"`) }

func (w *GronWriter) ToBytes() []byte {
	return []byte(strings.Join(w.lines, "\n") + "\n")
}
