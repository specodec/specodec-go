package specodec

type GronReader struct{}

func NewGronReader(data []byte) *GronReader                          { return &GronReader{} }
func (r *GronReader) ReadString() string                             { panic("GronReader: not implemented") }
func (r *GronReader) ReadBool() bool                                 { panic("GronReader: not implemented") }
func (r *GronReader) ReadInt32() int32                                { panic("GronReader: not implemented") }
func (r *GronReader) ReadInt64() int64                                { panic("GronReader: not implemented") }
func (r *GronReader) ReadUint32() uint32                              { panic("GronReader: not implemented") }
func (r *GronReader) ReadUint64() uint64                              { panic("GronReader: not implemented") }
func (r *GronReader) ReadFloat32() float32                            { panic("GronReader: not implemented") }
func (r *GronReader) ReadFloat64() float64                            { panic("GronReader: not implemented") }
func (r *GronReader) ReadNull()                                       { panic("GronReader: not implemented") }
func (r *GronReader) ReadBytes() []byte                               { panic("GronReader: not implemented") }
func (r *GronReader) BeginObject()                                    { panic("GronReader: not implemented") }
func (r *GronReader) HasNextField() bool                              { panic("GronReader: not implemented") }
func (r *GronReader) ReadFieldName() string                           { panic("GronReader: not implemented") }
func (r *GronReader) NextFieldSeparator()                             { panic("GronReader: not implemented") }
func (r *GronReader) EndObject()                                      { panic("GronReader: not implemented") }
func (r *GronReader) BeginArray()                                     { panic("GronReader: not implemented") }
func (r *GronReader) HasNextElement() bool                            { panic("GronReader: not implemented") }
func (r *GronReader) NextElementSeparator()                           { panic("GronReader: not implemented") }
func (r *GronReader) EndArray()                                       { panic("GronReader: not implemented") }
func (r *GronReader) IsNull() bool                                    { panic("GronReader: not implemented") }
func (r *GronReader) Skip()                                           { panic("GronReader: not implemented") }
