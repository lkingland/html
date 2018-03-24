package html

import (
	"io/ioutil"
)

// HTML Elements are a curtailed set of HTML5 elements

type Anchor struct{ HTMLElement }

func A() *BASE {
	return &BASE{HTMLElement{key: "a", empty: false}}
}

type BASE struct{ HTMLElement }

func Base() *BASE {
	return &BASE{HTMLElement{key: "base", empty: true}}
}

type BUTTON struct{ HTMLElement }

func Button() *BUTTON {
	return &BUTTON{HTMLElement{key: "button", empty: false}}
}

type BODY struct{ HTMLElement }

func Body() *BODY {
	return &BODY{HTMLElement{key: "body", empty: false}}
}

type DIV struct{ HTMLElement }

func Div() *DIV {
	return &DIV{HTMLElement{key: "div", empty: false}}
}

type EM struct{ HTMLElement }

func Em() *EM {
	return &EM{HTMLElement{key: "em", empty: false, inline: true}}
}

type FORM struct{ HTMLElement }

func Form() *FORM {
	return &FORM{HTMLElement{key: "form", empty: false}}
}

type HEAD struct{ HTMLElement }

func Head() *HEAD {
	return &HEAD{HTMLElement{key: "head", empty: false}}
}

type HTML struct{ HTMLElement }

func Html() *HTML {
	return &HTML{HTMLElement{key: "html", empty: false}}
}

type I struct{ HTMLElement }

func Italic() *I {
	return &I{HTMLElement{key: "i", empty: false, inline: true}}
}

type IMG struct{ HTMLElement }

func Img() *IMG {
	return &IMG{HTMLElement{key: "img", empty: true, inline: true}}
}

type INPUT struct{ HTMLElement }

func Input() *INPUT {
	return &INPUT{HTMLElement{key: "input", empty: false, inline: true}}
}

type LABEL struct{ HTMLElement }

func Label() *LABEL {
	return &LABEL{HTMLElement{key: "label", empty: false, inline: true}}
}

type LINK struct{ HTMLElement }

func Link() *LINK {
	return &LINK{HTMLElement{key: "link", empty: true}}
}

type META struct{ HTMLElement }

func Meta() *META {
	return &META{HTMLElement{key: "meta", empty: true}}
}

type SCRIPT struct{ HTMLElement }

func Script() *SCRIPT {
	return &SCRIPT{HTMLElement{key: "script", empty: false}}
}

type TITLE struct{ HTMLElement }

func Title() *TITLE {
	return &TITLE{HTMLElement{key: "title", empty: false}}
}

type Unordered struct{ HTMLElement }

func UL() *Unordered {
	return &Unordered{HTMLElement{key: "ul", empty: false}}
}

type Ordered struct{ HTMLElement }

func OL() *Ordered {
	return &Ordered{HTMLElement{key: "ol", empty: false}}
}

type Item struct{ HTMLElement }

func LI() *Item {
	return &Item{HTMLElement{key: "li", empty: false}}
}

type Header1 struct{ HTMLElement }

func H1() *Header1 {
	return &Header1{HTMLElement{key: "h1", empty: false}}
}

type Header2 struct{ HTMLElement }

func H2() *Header2 {
	return &Header2{HTMLElement{key: "h2", empty: false}}
}

type Header3 struct{ HTMLElement }

func H3() *Header3 {
	return &Header3{HTMLElement{key: "h3", empty: false}}
}

type Header4 struct{ HTMLElement }

func H4() *Header4 {
	return &Header4{HTMLElement{key: "h4", empty: false}}
}

type Header5 struct{ HTMLElement }

func H5() *Header5 {
	return &Header5{HTMLElement{key: "h5", empty: false}}
}

type Header6 struct{ HTMLElement }

func H6() *Header6 {
	return &Header6{HTMLElement{key: "h6", empty: false}}
}

// ROOT is a special case denoting the ephemeral single
// parent of the two root elements DOCTYPE and HTML.
// It is created as the ROOT of a page, and the DOCTYPE and HTML
// elements added.  The empty key results in no tags or attributes
// being written; only the children elements.
type ROOT struct{ HTMLElement }

func Root() *ROOT {
	return &ROOT{HTMLElement{}}
}

// Special elements that are not of type Element because they do not follow the
// standard HTML tag structure, but are correct HTML5 and are renderable, include the DOCTYPE and CDATA (plain text)

type DOCTYPE string

func Doctype() DOCTYPE {
	return DOCTYPE("<!DOCTYPE html>")
}

func (d DOCTYPE) Render(i int) string {
	return string(d)
}

// C is short for CDATA
type C string

func (c C) Render(i int) string {
	return string(c)
}

// FILE is a renderable structure that is a file substitution for raw
// (unsafe and non-autoformatted) plain text.
type FILE struct {
	Path string
}

func File(path string) FILE {
	return FILE{Path: path}
}

func (f FILE) Render(i int) string {
	// TODO: indent each line by i
	content, err := ioutil.ReadFile(f.Path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

type SPAN struct{ HTMLElement }

func Span() *SPAN {
	return &SPAN{HTMLElement{key: "span", empty: false, inline: true}}
}
