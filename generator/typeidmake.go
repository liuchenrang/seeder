package generator

import (
	"seeder/bootstrap"
	"seeder/generator/idgen"
)

type TypeIDMake struct {
}

func (typeMake TypeIDMake) Make(bizTag string, application *bootstrap.Application) idgen.IDGen {
	return idgen.NewDBGen(bizTag, application)
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}
