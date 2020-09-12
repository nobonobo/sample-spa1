package main

import "github.com/nobonobo/spago"

// Top ...
type Top struct {
	spago.Core
}

func (c *Top) Render()spago.HTML {
	return spago.Tag("body", spago.T("Hello World!"))
}

func main(){
	spago.RenderBody(&Top{})
	select{}
}
