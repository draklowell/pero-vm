package word

const (
	TypeNone      = 0
	TypeInteger   = 1
	TypeFloat     = 2
	TypeArray     = 3
	TypeContainer = 4
	TypeBytes     = 5
)

type Word interface{}

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
