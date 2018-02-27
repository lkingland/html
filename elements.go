package html

import "io/ioutil"

// HTML Elements are a curtailed set of HTML5 elements

type BASE struct{ HTMLElement }

func Base() *BASE {
	return &BASE{HTMLElement{Key: "base", Empty: true}}
}

type BODY struct{ HTMLElement }

func Body() *BODY {
	return &BODY{HTMLElement{Key: "body", Empty: false}}
}

type DIV struct{ HTMLElement }

func Div() *DIV {
	return &DIV{HTMLElement{Key: "div", Empty: false}}
}

type HEAD struct{ HTMLElement }

func Head() *HEAD {
	return &HEAD{HTMLElement{Key: "head", Empty: false}}
}

type HTML struct{ HTMLElement }

func Html() *HTML {
	return &HTML{HTMLElement{Key: "html", Empty: false}}
}

type IMG struct{ HTMLElement }

func Img() *IMG {
	return &IMG{HTMLElement{Key: "img", Empty: true}}
}

type LINK struct{ HTMLElement }

func Link() *LINK {
	return &LINK{HTMLElement{Key: "link", Empty: true}}
}

type META struct{ HTMLElement }

func Meta() *META {
	return &META{HTMLElement{Key: "meta", Empty: true}}
}

type SCRIPT struct{ HTMLElement }

func Script() *SCRIPT {
	return &SCRIPT{HTMLElement{Key: "script", Empty: false}}
}

type TITLE struct{ HTMLElement }

func Title() *TITLE {
	return &TITLE{HTMLElement{Key: "title", Empty: false}}
}

// ROOT is a special case denoting the ephemeral single
// parent of the two root elements DOCTYPE and HTML.
// It is created as the ROOT of a page, and the DOCTYPE and HTML
// elements added.  The empty Key results in no tags or attributes
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
