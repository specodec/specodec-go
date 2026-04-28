package specodec

import "math"

type MsgPackWriter struct {
	buf []byte
}

func NewMsgPackWriter() *MsgPackWriter {
	return &MsgPackWriter{buf: make([]byte, 0, 256)}
}

func (w *MsgPackWriter) writeByte(b byte) {
	w.buf = append(w.buf, b)
}

func (w *MsgPackWriter) writeU16(v uint16) {
	w.buf = append(w.buf, byte(v>>8), byte(v))
}

func (w *MsgPackWriter) writeU32(v uint32) {
	w.buf = append(w.buf, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (w *MsgPackWriter) writeU64(v uint64) {
	w.writeU32(uint32(v >> 32))
	w.writeU32(uint32(v))
}

func (w *MsgPackWriter) WriteString(value string) {
	bytes := []byte(value)
	length := len(bytes)
	if length <= 0x1F {
		w.writeByte(0xA0 | byte(length))
	} else if length <= 0xFF {
		w.writeByte(0xD9)
		w.writeByte(byte(length))
	} else if length <= 0xFFFF {
		w.writeByte(0xDA)
		w.writeU16(uint16(length))
	} else {
		w.writeByte(0xDB)
		w.writeU32(uint32(length))
	}
	w.buf = append(w.buf, bytes...)
}

func (w *MsgPackWriter) WriteBool(value bool) {
	if value {
		w.writeByte(0xC3)
	} else {
		w.writeByte(0xC2)
	}
}

func (w *MsgPackWriter) WriteInt32(value int32) {
	if value >= 0 && value <= 0x7F {
		w.writeByte(byte(value))
	} else if value < 0 && value >= -0x20 {
		w.writeByte(byte(value))
	} else if value >= 0 && value <= 0xFF {
		w.writeByte(0xCC)
		w.writeByte(byte(value))
	} else if value >= 0 && value <= 0xFFFF {
		w.writeByte(0xCD)
		w.writeU16(uint16(value))
	} else if value >= 0 {
		w.writeByte(0xCE)
		w.writeU32(uint32(value))
	} else if value >= -0x80 {
		w.writeByte(0xD0)
		w.writeByte(byte(value))
	} else if value >= -0x8000 {
		w.writeByte(0xD1)
		w.writeU16(uint16(value))
	} else {
		w.writeByte(0xD2)
		w.writeU32(uint32(value))
	}
}

func (w *MsgPackWriter) WriteInt64(value int64) {
	if value >= 0 && value <= 0x7F {
		w.writeByte(byte(value))
	} else if value < 0 && value >= -0x20 {
		w.writeByte(byte(value))
	} else if value >= 0 && value <= 0xFF {
		w.writeByte(0xCC)
		w.writeByte(byte(value))
	} else if value >= 0 && value <= 0xFFFF {
		w.writeByte(0xCD)
		w.writeU16(uint16(value))
	} else if value >= 0 && value <= 0xFFFFFFFF {
		w.writeByte(0xCE)
		w.writeU32(uint32(value))
	} else if value >= 0 {
		w.writeByte(0xCF)
		w.writeU64(uint64(value))
	} else if value >= -0x80 {
		w.writeByte(0xD0)
		w.writeByte(byte(value))
	} else if value >= -0x8000 {
		w.writeByte(0xD1)
		w.writeU16(uint16(value))
	} else if value >= -0x80000000 {
		w.writeByte(0xD2)
		w.writeU32(uint32(value))
	} else {
		w.writeByte(0xD3)
		w.writeU64(uint64(value))
	}
}

func (w *MsgPackWriter) WriteUint32(value uint32) {
	if value <= 0x7F {
		w.writeByte(byte(value))
	} else if value <= 0xFF {
		w.writeByte(0xCC)
		w.writeByte(byte(value))
	} else if value <= 0xFFFF {
		w.writeByte(0xCD)
		w.writeU16(uint16(value))
	} else {
		w.writeByte(0xCE)
		w.writeU32(value)
	}
}

func (w *MsgPackWriter) WriteUint64(value uint64) {
	if value <= 0x7F {
		w.writeByte(byte(value))
	} else if value <= 0xFF {
		w.writeByte(0xCC)
		w.writeByte(byte(value))
	} else if value <= 0xFFFF {
		w.writeByte(0xCD)
		w.writeU16(uint16(value))
	} else if value <= 0xFFFFFFFF {
		w.writeByte(0xCE)
		w.writeU32(uint32(value))
	} else {
		w.writeByte(0xCF)
		w.writeU64(value)
	}
}

func (w *MsgPackWriter) WriteFloat32(value float32) {
	w.writeByte(0xCA)
	bits := math.Float32bits(value)
	w.writeU32(bits)
}

func (w *MsgPackWriter) WriteFloat64(value float64) {
	w.writeByte(0xCB)
	bits := math.Float64bits(value)
	w.writeU64(bits)
}

func (w *MsgPackWriter) WriteNull() {
	w.writeByte(0xC0)
}

func (w *MsgPackWriter) WriteBytes(value []byte) {
	length := len(value)
	if length <= 0xFF {
		w.writeByte(0xC4)
		w.writeByte(byte(length))
	} else if length <= 0xFFFF {
		w.writeByte(0xC5)
		w.writeU16(uint16(length))
	} else {
		w.writeByte(0xC6)
		w.writeU32(uint32(length))
	}
	w.buf = append(w.buf, value...)
}

func (w *MsgPackWriter) BeginObject(fieldCount int) {
	if fieldCount <= 0x0F {
		w.writeByte(0x80 | byte(fieldCount))
	} else if fieldCount <= 0xFFFF {
		w.writeByte(0xDE)
		w.writeU16(uint16(fieldCount))
	} else {
		w.writeByte(0xDF)
		w.writeU32(uint32(fieldCount))
	}
}

func (w *MsgPackWriter) WriteField(name string) {
	w.WriteString(name)
}

func (w *MsgPackWriter) EndObject() {
}

func (w *MsgPackWriter) BeginArray(elementCount int) {
	if elementCount <= 0x0F {
		w.writeByte(0x90 | byte(elementCount))
	} else if elementCount <= 0xFFFF {
		w.writeByte(0xDC)
		w.writeU16(uint16(elementCount))
	} else {
		w.writeByte(0xDD)
		w.writeU32(uint32(elementCount))
	}
}

func (w *MsgPackWriter) NextElement() {
}

func (w *MsgPackWriter) EndArray() {
}

func (w *MsgPackWriter) ToBytes() []byte {
	return w.buf
}
