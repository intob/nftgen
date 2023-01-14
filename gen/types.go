package gen

type Variant struct {
	Name string `yaml:"name"`
	Img  string `yaml:"img"`
	X    int    `yaml:"x"`
	Y    int    `yaml:"y"`
}

type Trait map[string]Variant
