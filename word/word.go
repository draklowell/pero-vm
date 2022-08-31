package word

const (
	// None - null pointer, nil
	TypeNone = 0
	// Integer - long, signed int64
	TypeInteger = 1
	// Float - double, float64
	TypeFloat = 2
	// Array - array, slice of words
	TypeArray = 3
	// Container - object, map with string key and word value
	TypeContainer = 4
	// Bytes - array of bytes
	TypeBytes = 5
)

// Word interface ( check TypeX constants )
type Word interface{}

// Get type of word ( check TypeX constants )
func GetType(word Word) int {
	switch word.(type) {
	case int64:
		return TypeInteger
	case float64:
		return TypeFloat
	case []uint8:
		return TypeBytes
	case []Word:
		return TypeArray
	case map[string]Word:
		return TypeContainer
	case nil:
		return TypeNone
	default:
		return -1
	}
}
