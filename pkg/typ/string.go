package typ

type String struct {
	Scalar
	Value string
}

func (p *String) Parse(str string) {
	p.Value = str
}

func (p *String) Serialize() string {
	return p.Value
}

func (p *String) Deserialize(str string) {
	p.Value = str
}
