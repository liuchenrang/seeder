package generator

import "seeder/generator/idgen"

type TypeIDMake struct{

}

type TypeMake interface {
	factory(makeType string) idgen.IDGen
}

type IdGenFactory func() idgen.IDGen

func (typeMake TypeIDMake ) Make() idgen.IDGen {
	return &idgen.DBGen{}
}
func NewTypeIDMake() TypeIDMake {
	return TypeIDMake{}
}
