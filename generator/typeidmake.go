package generator

import (
	"seeder/generator/idgen"
	"seeder/config"
)

type TypeIDMake struct {
}

func (typeMake TypeIDMake) Make(bizTag string, seederConfig config.SeederConfig) idgen.IDGen {
	return idgen.NewDBGen(bizTag, seederConfig)
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}
