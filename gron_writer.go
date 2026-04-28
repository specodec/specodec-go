package specodec

type GronWriter struct{}

func NewGronWriter() *GronWriter                                        { return &GronWriter{} }
func (w *GronWriter) WriteString(value string)                         { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteBool(value bool)                             { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteInt32(value int32)                           { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteInt64(value int64)                           { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteUint32(value uint32)                         { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteUint64(value uint64)                         { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteFloat32(value float32)                       { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteFloat64(value float64)                       { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteNull()                                       { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteBytes(value []byte)                          { panic("GronWriter: not implemented") }
func (w *GronWriter) BeginObject(fieldCount int)                       { panic("GronWriter: not implemented") }
func (w *GronWriter) WriteField(name string)                           { panic("GronWriter: not implemented") }
func (w *GronWriter) EndObject()                                       { panic("GronWriter: not implemented") }
func (w *GronWriter) BeginArray(elementCount int)                      { panic("GronWriter: not implemented") }
func (w *GronWriter) NextElement()                                     { panic("GronWriter: not implemented") }
func (w *GronWriter) EndArray()                                        { panic("GronWriter: not implemented") }
func (w *GronWriter) ToBytes() []byte                                  { panic("GronWriter: not implemented") }
