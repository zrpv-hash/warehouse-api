package isolation

type IsolationLevel int

const (
	LevelReadUncommitted IsolationLevel = iota
	LevelReadCommitted
	LevelRepeatableRead
	LevelSerializable
)

func (i IsolationLevel) String() string {
	switch i {
	case LevelReadUncommitted:
		return "Read Uncommitted"
	case LevelReadCommitted:
		return "Read Committed"
	case LevelRepeatableRead:
		return "Repeatable Read"
	case LevelSerializable:
		return "Serializable"
	default:
		return "Read Committed"
	}
}
