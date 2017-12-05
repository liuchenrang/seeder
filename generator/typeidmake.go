package generator

import (
	"seeder/generator/idgen"
	"seeder/bootstrap"
)

type TypeIDMake struct {
}

func (typeMake TypeIDMake) Make(bizTag string,  application *bootstrap.Application) idgen.IDGen {
	return idgen.NewDBGen(bizTag, application )
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}
