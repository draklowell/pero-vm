package word

type Integer struct {
	value int64
}

func NewInteger(value int64) *Integer {
	return &Integer{value: value}
}

func (word *Integer) GetType() int {
	return TypeInteger
}

func (word *Integer) GetValue() int64 {
	return word.value
}

type Float struct {
	value float64
}

func NewFloat(value float64) *Float {
	return &Float{value: value}
}

func (word *Float) GetType() int {
	return TypeFloat
}

func (word *Float) GetValue() float64 {
	return word.value
}

func NewBytes(value []byte) (*Bytes, error) {
	if value == nil {
		return nil, ErrNilPointer
	}
	return &Bytes{value: value}, nil
}

type Bytes struct {
	value []byte
}

func (word *Bytes) GetType() int {
	return TypeBytes
}

func (word *Bytes) GetValue() []byte {
	return word.value
}

type NoneType struct{}

func (none *NoneType) GetType() int {
	return TypeNone
}

var None *NoneType = &NoneType{}

type Boolean struct {
	value bool
}

func NewBoolean(value bool) *Boolean {
	if value {
		return True
	} else {
		return False
	}
}

func (word *Boolean) GetType() int {
	return TypeBoolean
}

func (word *Boolean) GetValue() bool {
	return word.value
}

var (
	True  *Boolean = &Boolean{true}
	False *Boolean = &Boolean{false}
)
