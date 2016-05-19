package semanticui

import (
	//"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

type Semantic struct {
	jquery.JQuery
}

func NewSemantic(i ...interface{}) Semantic {
	return Semantic{jquery.NewJQuery(i)}
}

//Modules

func (s Semantic) Accordion(i ...interface{}) Semantic {
	return NewSemantic(s.Call("accordion", i...))
}

func (s Semantic) Checkbox(i ...interface{}) Semantic {
	return NewSemantic(s.Call("checkbox", i...))
}

func (s Semantic) Dimmer(i ...interface{}) Semantic {
	return NewSemantic(s.Call("dimmer", i...))
}

func (s Semantic) Dropdown(i ...interface{}) Semantic {
	return NewSemantic(s.Call("dropdown", i...))
}

func (s Semantic) Embed(i ...interface{}) Semantic {
	return NewSemantic(s.Call("embed", i...))
}

func (s Semantic) Modal(i ...interface{}) Semantic {
	return NewSemantic(s.Call("modal", i...))
}

func (s Semantic) Popup(i ...interface{}) Semantic {
	return NewSemantic(s.Call("popup", i...))
}

func (s Semantic) Progress(i ...interface{}) Semantic {
	return NewSemantic(s.Call("progress", i...))
}

func (s Semantic) Rating(i ...interface{}) Semantic {
	return NewSemantic(s.Call("rating", i...))
}

func (s Semantic) Search(i ...interface{}) Semantic {
	return NewSemantic(s.Call("search", i...))
}

func (s Semantic) Shape(i ...interface{}) Semantic {
	return NewSemantic(s.Call("shape", i...))
}

func (s Semantic) Sidebar(i ...interface{}) Semantic {
	return NewSemantic(s.Call("sidebar", i...))
}

func (s Semantic) Sticky(i ...interface{}) Semantic {
	return NewSemantic(s.Call("sticky", i...))
}

func (s Semantic) Tab(i ...interface{}) Semantic {
	return NewSemantic(s.Call("tab", i...))
}

func (s Semantic) Transition(i ...interface{}) Semantic {
	return NewSemantic(s.Call("transition", i...))
}
