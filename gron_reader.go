package specodec

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type GronReader struct {
	lines  [][2]string
	cursor int
	ctx    []gronCtx
}

type gronCtx struct {
	prefix  string
	typeStr string
	index   int
}

func NewGronReader(data []byte) *GronReader {
	text := string(data)
	var lines [][2]string
	for _, raw := range strings.Split(text, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		eq := strings.Index(line, " = ")
		if eq < 0 {
			continue
		}
		path := line[:eq]
		val := line[eq+3:]
		val = strings.TrimSuffix(val, ";")
		lines = append(lines, [2]string{path, val})
	}
	return &GronReader{lines: lines}
}

func gronUnescape(s string) (string, error) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return "", fmt.Errorf("gron: expected quoted string")
	}
	var r strings.Builder
	inner := s[1 : len(s)-1]
	for i := 0; i < len(inner); i++ {
		if inner[i] == '\\' && i+1 < len(inner) {
			i++
			switch inner[i] {
			case '"':
				r.WriteByte('"')
			case '\\':
				r.WriteByte('\\')
			case '/':
				r.WriteByte('/')
			case 'b':
				r.WriteByte(0x08)
			case 'f':
				r.WriteByte(0x0C)
			case 'n':
				r.WriteByte('\n')
			case 'r':
				r.WriteByte('\r')
			case 't':
				r.WriteByte('\t')
			case 'u':
				if i+4 < len(inner) {
					v, err := strconv.ParseInt(inner[i+1:i+5], 16, 32)
					if err == nil {
						r.WriteRune(rune(v))
					}
					i += 4
				}
			default:
				r.WriteByte(inner[i])
			}
		} else {
			r.WriteByte(inner[i])
		}
	}
	return r.String(), nil
}

func gronB64decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func (r *GronReader) ReadString() string             { v, _ := gronUnescape(r.lines[r.cursor][1]); r.cursor++; return v }
func (r *GronReader) ReadBool() bool                 { v := r.lines[r.cursor][1] == "true"; r.cursor++; return v }
func (r *GronReader) ReadInt32() int32               { v, _ := strconv.ParseInt(r.lines[r.cursor][1], 10, 32); r.cursor++; return int32(v) }
func (r *GronReader) ReadInt64() int64               { s, _ := gronUnescape(r.lines[r.cursor][1]); v, _ := strconv.ParseInt(s, 10, 64); r.cursor++; return v }
func (r *GronReader) ReadUint32() uint32             { v, _ := strconv.ParseUint(r.lines[r.cursor][1], 10, 32); r.cursor++; return uint32(v) }
func (r *GronReader) ReadUint64() uint64             { s, _ := gronUnescape(r.lines[r.cursor][1]); v, _ := strconv.ParseUint(s, 10, 64); r.cursor++; return v }
func (r *GronReader) ReadNull()                      { r.cursor++ }
func (r *GronReader) ReadBytes() []byte              { s, _ := gronUnescape(r.lines[r.cursor][1]); r.cursor++; b, _ := gronB64decode(s); return b }
func (r *GronReader) ReadEnum() string               { v, _ := gronUnescape(r.lines[r.cursor][1]); r.cursor++; return v }

func (r *GronReader) ReadFloat32() float32 {
	v := r.lines[r.cursor][1]; r.cursor++
	f, _ := strconv.ParseFloat(v, 32); return float32(f)
}

func (r *GronReader) ReadFloat64() float64 {
	v := r.lines[r.cursor][1]; r.cursor++
	f, _ := strconv.ParseFloat(v, 64); return f
}

func (r *GronReader) BeginObject() {
	line := r.lines[r.cursor]; r.cursor++
	r.ctx = append(r.ctx, gronCtx{prefix: line[0], typeStr: "object"})
}

func (r *GronReader) HasNextField() bool {
	if r.cursor >= len(r.lines) {
		return false
	}
	pfx := r.ctx[len(r.ctx)-1].prefix + "."
	p := r.lines[r.cursor][0]
	if !strings.HasPrefix(p, pfx) {
		return false
	}
	rem := p[len(pfx):]
	return !strings.Contains(rem, ".") && !strings.Contains(rem, "[")
}

func (r *GronReader) ReadFieldName() string {
	pfx := r.ctx[len(r.ctx)-1].prefix + "."
	return r.lines[r.cursor][0][len(pfx):]
}

func (r *GronReader) EndObject()          { r.ctx = r.ctx[:len(r.ctx)-1] }

func (r *GronReader) BeginArray() {
	line := r.lines[r.cursor]; r.cursor++
	r.ctx = append(r.ctx, gronCtx{prefix: line[0], typeStr: "array", index: -1})
}

func (r *GronReader) HasNextElement() bool {
	if r.cursor >= len(r.lines) {
		return false
	}
	arr := r.ctx[len(r.ctx)-1]
	ni := arr.index + 1
	exp := fmt.Sprintf("%s[%d]", arr.prefix, ni)
	p := r.lines[r.cursor][0]
	return p == exp || strings.HasPrefix(p, exp+".") || strings.HasPrefix(p, exp+"[")
}

func (r *GronReader) NextElement()          { r.ctx[len(r.ctx)-1].index++ }
func (r *GronReader) EndArray()             { r.ctx = r.ctx[:len(r.ctx)-1] }

func (r *GronReader) IsNull() bool {
	return r.cursor < len(r.lines) && r.lines[r.cursor][1] == "null"
}

func (r *GronReader) Skip() {
	sp := r.lines[r.cursor][0]; r.cursor++
	for r.cursor < len(r.lines) {
		np := r.lines[r.cursor][0]
		if len(np) > len(sp) && (strings.HasPrefix(np, sp+".") || strings.HasPrefix(np, sp+"[")) {
			r.cursor++
		} else {
			break
		}
	}
}
