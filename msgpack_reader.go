package specodec

import (
	"encoding/binary"
	"fmt"
	"math"
)

type MsgPackReader struct {
	data           []byte
	pos            int
	containerCount []int
}

func NewMsgPackReader(data []byte) *MsgPackReader {
	return &MsgPackReader{data: data}
}

func (r *MsgPackReader) Pos() int { return r.pos }

func (r *MsgPackReader) readByte() byte {
	if r.pos >= len(r.data) {
		panic("msgpack: unexpected end of data")
	}
	b := r.data[r.pos]
	r.pos++
	return b
}

func (r *MsgPackReader) readU16() uint16 {
	v := binary.BigEndian.Uint16(r.data[r.pos:])
	r.pos += 2
	return v
}

func (r *MsgPackReader) readU32() uint32 {
	v := binary.BigEndian.Uint32(r.data[r.pos:])
	r.pos += 4
	return v
}

func (r *MsgPackReader) readI16() int16 { return int16(r.readU16()) }
func (r *MsgPackReader) readI32() int32 { return int32(r.readU32()) }

func (r *MsgPackReader) ReadMapHeader() int {
	b := r.readByte()
	if b&0xF0 == 0x80 {
		return int(b & 0x0F)
	}
	if b == 0xDE {
		return int(r.readU16())
	}
	if b == 0xDF {
		return int(r.readU32())
	}
	panic(fmt.Sprintf("msgpack: expected map, got 0x%02X", b))
}

func (r *MsgPackReader) ReadArrayHeader() int {
	b := r.readByte()
	if b&0xF0 == 0x90 {
		return int(b & 0x0F)
	}
	if b == 0xDC {
		return int(r.readU16())
	}
	if b == 0xDD {
		return int(r.readU32())
	}
	panic(fmt.Sprintf("msgpack: expected array, got 0x%02X", b))
}

func (r *MsgPackReader) ReadString() string {
	b := r.readByte()
	var length int
	if b&0xE0 == 0xA0 {
		length = int(b & 0x1F)
	} else if b == 0xD9 {
		length = int(r.readByte())
	} else if b == 0xDA {
		length = int(r.readU16())
	} else if b == 0xDB {
		length = int(r.readU32())
	} else {
		panic(fmt.Sprintf("msgpack: expected string, got 0x%02X", b))
	}
	s := string(r.data[r.pos : r.pos+length])
	r.pos += length
	return s
}

func (r *MsgPackReader) ReadInt() int64 {
	b := r.readByte()
	if b <= 0x7F {
		return int64(b)
	}
	if b >= 0xE0 {
		return int64(int8(b))
	}
	switch b {
	case 0xCC:
		return int64(r.readByte())
	case 0xCD:
		return int64(r.readU16())
	case 0xCE:
		return int64(r.readU32())
	case 0xD0:
		return int64(int8(r.readByte()))
	case 0xD1:
		return int64(r.readI16())
	case 0xD2:
		return int64(r.readI32())
	default:
		panic(fmt.Sprintf("msgpack: expected int, got 0x%02X", b))
	}
}

func (r *MsgPackReader) ReadFloat() float64 {
	b := r.readByte()
	if b == 0xCA {
		bits := binary.BigEndian.Uint32(r.data[r.pos:])
		r.pos += 4
		return float64(math.Float32frombits(bits))
	}
	if b == 0xCB {
		bits := binary.BigEndian.Uint64(r.data[r.pos:])
		r.pos += 8
		return math.Float64frombits(bits)
	}
	panic(fmt.Sprintf("msgpack: expected float, got 0x%02X", b))
}

func (r *MsgPackReader) ReadBool() bool {
	b := r.readByte()
	if b == 0xC3 {
		return true
	}
	if b == 0xC2 {
		return false
	}
	panic(fmt.Sprintf("msgpack: expected bool, got 0x%02X", b))
}

func (r *MsgPackReader) ReadNull() {
	b := r.readByte()
	if b != 0xC0 {
		panic(fmt.Sprintf("msgpack: expected null, got 0x%02X", b))
	}
}

func (r *MsgPackReader) IsNull() bool {
	if r.pos >= len(r.data) {
		return false
	}
	return r.data[r.pos] == 0xC0
}

func (r *MsgPackReader) Skip() {
	b := r.readByte()
	if b <= 0x7F || b >= 0xE0 {
		return
	}
	if b&0xF0 == 0x80 {
		n := int(b & 0x0F)
		for i := 0; i < n; i++ {
			r.Skip()
			r.Skip()
		}
		return
	}
	if b&0xF0 == 0x90 {
		n := int(b & 0x0F)
		for i := 0; i < n; i++ {
			r.Skip()
		}
		return
	}
	if b&0xE0 == 0xA0 {
		r.pos += int(b & 0x1F)
		return
	}
	switch b {
	case 0xC0, 0xC2, 0xC3:
	case 0xCC, 0xD0:
		r.pos++
	case 0xCD, 0xD1:
		r.pos += 2
	case 0xCE, 0xD2, 0xCA:
		r.pos += 4
	case 0xCF, 0xD3, 0xCB:
		r.pos += 8
	case 0xD9:
		r.pos += int(r.readByte())
	case 0xDA:
		r.pos += int(r.readU16())
	case 0xDB:
		r.pos += int(r.readU32())
	case 0xC4:
		r.pos += int(r.readByte())
	case 0xC5:
		r.pos += int(r.readU16())
	case 0xC6:
		r.pos += int(r.readU32())
	case 0xD4:
		r.pos += 2
	case 0xD5:
		r.pos += 3
	case 0xD6:
		r.pos += 5
	case 0xD7:
		r.pos += 9
	case 0xD8:
		r.pos += 17
	case 0xC7:
		r.pos += 1 + int(r.readByte())
	case 0xC8:
		r.pos += 1 + int(r.readU16())
	case 0xC9:
		r.pos += 1 + int(r.readU32())
	case 0xDC:
		n := int(r.readU16())
		for i := 0; i < n; i++ {
			r.Skip()
		}
	case 0xDD:
		n := int(r.readU32())
		for i := 0; i < n; i++ {
			r.Skip()
		}
	case 0xDE:
		n := int(r.readU16())
		for i := 0; i < n; i++ {
			r.Skip()
			r.Skip()
		}
	case 0xDF:
		n := int(r.readU32())
		for i := 0; i < n; i++ {
			r.Skip()
			r.Skip()
		}
	default:
		panic(fmt.Sprintf("msgpack: unknown format 0x%02X", b))
	}
}

func (r *MsgPackReader) BeginObject() {
	n := r.ReadMapHeader()
	r.containerCount = append(r.containerCount, n)
}

func (r *MsgPackReader) HasNextField() bool {
	top := len(r.containerCount) - 1
	if r.containerCount[top] > 0 {
		r.containerCount[top]--
		return true
	}
	r.containerCount = r.containerCount[:top]
	return false
}

func (r *MsgPackReader) ReadFieldName() string {
	return r.ReadString()
}

func (r *MsgPackReader) EndObject() {
}

func (r *MsgPackReader) BeginArray() {
	n := r.ReadArrayHeader()
	r.containerCount = append(r.containerCount, n)
}

func (r *MsgPackReader) HasNextElement() bool {
	top := len(r.containerCount) - 1
	if r.containerCount[top] > 0 {
		r.containerCount[top]--
		return true
	}
	r.containerCount = r.containerCount[:top]
	return false
}

func (r *MsgPackReader) EndArray() {
}
