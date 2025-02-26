package confluence

type (
	Config struct {
		Name   string
		Fields []Field // ([]Field == nil) == (IsLeaf == true)
	}

	Field struct {
		Name     *string // embeddings have no names
		Tag      *Tag
		TypeName string
	}

	Tag struct {
		ZeroValue string
		Env       string
	}
)
