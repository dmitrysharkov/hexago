package hexago

type KindOfType int

const (
	KindOfScalar KindOfType = iota
	KindOfObject
)

// Type definition internal structure
type TypeDef struct {
	Kind KindOfType

	PropertiesNum uint
	Property      []PropertyDef
}

type PropertyDef struct {
}
