package generator

import (
	"seeder/generator/idgen"
)

type TypeIDMake struct {
}

func (typeMake TypeIDMake) Make(bizTag string) idgen.IDGen {
	return idgen.NewDBGen(bizTag)
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}
