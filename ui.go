package semanticui

import (
	//"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

type Semantic struct {
	jquery.JQuery
}

func NewSemantic(i ...interface{}) Semantic {
	return Semantic{jquery.NewJQuery(i...)}
}

//Modules

func (s Semantic) Accordion(i ...interface{}) Semantic {
	s.Call("accordion", i...)
	return s
}

func (s Semantic) Checkbox(i ...interface{}) Semantic {
	s.Call("checkbox", i...)
	return s
}

func (s Semantic) Dimmer(i ...interface{}) Semantic {
	s.Call("dimmer", i...)
	return s
}

func (s Semantic) Dropdown(i ...interface{}) Semantic {
	s.Call("dropdown", i...)
	return s
}

func (s Semantic) Embed(i ...interface{}) Semantic {
	s.Call("embed", i...)
	return s
}

func (s Semantic) Modal(i ...interface{}) Semantic {
	s.Call("modal", i...)
	return s
}

func (s Semantic) Popup(i ...interface{}) Semantic {
	s.Call("popup", i...)
	return s
}

func (s Semantic) Progress(i ...interface{}) Semantic {
	s.Call("progress", i...)
	return s
}

func (s Semantic) Rating(i ...interface{}) Semantic {
	s.Call("rating", i...)
	return s
}

func (s Semantic) Search(i ...interface{}) Semantic {
	s.Call("search", i...)
	return s
}

func (s Semantic) Shape(i ...interface{}) Semantic {
	s.Call("shape", i...)
	return s
}

func (s Semantic) Sidebar(i ...interface{}) Semantic {
	s.Call("sidebar", i...)
	return s
}

func (s Semantic) Sticky(i ...interface{}) Semantic {
	s.Call("sticky", i...)
	return s
}

func (s Semantic) Tab(i ...interface{}) Semantic {
	s.Call("tab", i...)
	return s
}

func (s Semantic) Transition(i ...interface{}) Semantic {
	s.Call("transition", i...)
	return s
}
