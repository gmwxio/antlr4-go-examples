package eval

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/wxio/antlr4-go-examples/ctree"
	. "github.com/wxio/antlr4-go-examples/eval/walker"
)

type precedence struct {
	tree    ctree.Tree
	convert map[reflect.Type](func(Atom) Atom)
}

var MathPrecedences *precedence = &precedence{}
var StringPrecedences *precedence = &precedence{}

func (p *precedence) Tree() ctree.Tree { return p.tree }
func (p *precedence) ConvertUp(from reflect.Type) func(from Atom) Atom {
	return p.convert[from]
}

var _ Precendence = &precedence{}

type Precendence interface {
	Tree() ctree.Tree
	ConvertUp(reflect.Type) func(from Atom) Atom
}

type CUP func(Atom) Atom

func init() {
	MathPrecedences.tree = ctree.BuildTree("", &struct{}{}).
		Add(reflect.TypeOf(&ErrorAtom{})).Down().
		/**/ Add(reflect.TypeOf(&FloatAtom{})).Down().
		/* */ Add(reflect.TypeOf(&IntegerAtom{})).Down().
		/*  */ Add(reflect.TypeOf(&BoolAtom{})).
		/**/ Up().Up().
		/**/ Add(reflect.TypeOf(&StringAtom{})).
		Build()
	MathPrecedences.convert = make(map[reflect.Type](func(Atom) Atom))
	MathPrecedences.convert[reflect.TypeOf(&StringAtom{})] = func(from Atom) Atom {
		return &ErrorAtom{Text: "#Value!"} // TODO return singelton value error
	}
	MathPrecedences.convert[reflect.TypeOf(&BoolAtom{})] =
		func(from Atom) Atom {
			b := from.(*BoolAtom)
			if b.Val {
				return &IntegerAtom{Val: 1}
			}
			return &IntegerAtom{Val: 0}
		}
	MathPrecedences.convert[reflect.TypeOf(&IntegerAtom{})] = func(from Atom) Atom {
		a := from.(*IntegerAtom)
		return &FloatAtom{Val: float64(a.Val)}
	}
	MathPrecedences.convert[reflect.TypeOf(&FloatAtom{})] = func(from Atom) Atom {
		return &ErrorAtom{Text: "#Value!"} // TODO return singelton value error
	}

	StringPrecedences.tree = ctree.BuildTree("", &struct{}{}).
		Add(reflect.TypeOf(&ErrorAtom{})).Down().
		/**/ Add(reflect.TypeOf(&StringAtom{})).Down().
		/* */ Add(reflect.TypeOf(&FloatAtom{})).
		/* */ Add(reflect.TypeOf(&IntegerAtom{})).
		/* */ Add(reflect.TypeOf(&BoolAtom{})).
		Build()
	StringPrecedences.convert = make(map[reflect.Type](func(Atom) Atom))
	StringPrecedences.convert[reflect.TypeOf(&BoolAtom{})] = func(from Atom) Atom {
		b := from.(*BoolAtom)
		if b.Val {
			return &StringAtom{Val: "TRUE"} // return cached values
		}
		return &StringAtom{Val: "FALSE"}
	}
	StringPrecedences.convert[reflect.TypeOf(&IntegerAtom{})] = func(from Atom) Atom {
		a := from.(*IntegerAtom)
		return &StringAtom{Val: strconv.Itoa(a.Val)}
	}
	StringPrecedences.convert[reflect.TypeOf(&FloatAtom{})] = func(from Atom) Atom {
		a := from.(*FloatAtom)
		return &StringAtom{Val: fmt.Sprintf("%v", a.Val)} // TODO Need to apply formating rule to do the convert
	}
	StringPrecedences.convert[reflect.TypeOf(&StringAtom{})] = func(from Atom) Atom {
		return &ErrorAtom{Text: "#Value!"} // TODO return singelton value error
	}

}
