package word

const (
	TypeNone      = 0
	TypeInteger   = 1
	TypeFloat     = 2
	TypeArray     = 3
	TypeContainer = 4
	TypeBytes     = 5
	TypeBoolean   = 6
)

type Word interface {
	GetType() int
}
