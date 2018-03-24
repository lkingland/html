// package html defines renderable HTML primitives.
// It's raw.  It's verbose.  It's oft opinionated in what it omits.
// But it endeavors to be strictly correct with what it does.
// For something more friendly, see the ui package.
package html

import "fmt"

// Component is a metastructure that, when rendered, returns
// the root of an arbitrarily large element tree.
type Component interface {
	Render() Element
}

func Render(c Component) string {
	return c.Render().Render(0)
}

// Element is a renderable with other
type Element interface {
	Renderable
	Key() string
	Empty() bool
	Inline() bool
	Set(string, string) Element
	Attributes() []Attribute
	Add(...Renderable) Element
	Children() []Renderable
}

// Element is a representation of an HTML element.
// An element consists of either an empty tag <img />
// or standard tag <div></div>, with attributes <div key="value">
// and potentially children if it is not empty <div><img /></div>
// When rendered, a tag will be indented based on its level of nesting,
// and will have linebreaks in the appropriate place depending on
// it's being either an inline or non-inline element.
type HTMLElement struct {
	key        string
	empty      bool
	inline     bool
	attributes []Attribute
	children   []Renderable
}

func (t *HTMLElement) Key() string {
	return t.key
}

func (t *HTMLElement) Empty() bool {
	return t.empty
}

func (t *HTMLElement) Inline() bool {
	return t.inline
}

func (t *HTMLElement) Attributes() []Attribute {
	return t.attributes
}

func (t *HTMLElement) Children() []Renderable {
	return t.children
}

// Set the value of an attribute on the start tag of the element.
// <tagname key="value"></tagname>
func (t *HTMLElement) Set(k, v string) Element {
	for i, a := range t.attributes {
		if a.Key == k {
			t.attributes[i].Value = v
			break
		}
	}
	a := Attribute{Key: k, Value: v}
	t.attributes = append(t.attributes, a)
	return t
}

// Add something renderable to the element's children
func (t *HTMLElement) Add(r ...Renderable) Element {
	t.children = append(t.children, r...)
	return t
}

// AddText is a convenience method for Add(C(text))
func (t *HTMLElement) AddText(text string) Element {
	return t.Add(C(text))
}

// return a padding string for a given indent level
func padding(indent int) string {
	padFmt := fmt.Sprintf("%%%ds", (indent * 2))
	return fmt.Sprintf(padFmt, "")
}

// Render the element including recursively rendering its children.
func (t *HTMLElement) Render(i int) string {

	// pseudoelement root is denoted by a blank tagname (key)
	// and is rendered by immediately delegating to children
	// without any change to indentation, but with a trailing space
	// for after the closing html tag.  This allows for the entire
	// document to be treated as  a single DAG despite there being
	// multiple root level elements allowd by the spec, eg: doctype
	// and HTML.
	if t.key == "" {
		return t.renderChildren(i) + "\n"
	}

	// Tag is prefixed by a srcbreak and padded to the current level.
	src := "\n" + padding(i)

	// Attributes will be embedded in the open tag
	attributes := t.renderAttributes()

	// Empty tags have a special syntax and are handled and then we immediately return
	if t.empty {
		return src + "<" + t.key + attributes + " />"
	}

	// Regular (nonempty) tags are rendered by first creating the (indented) open tag
	// which includes attribute key value pairs.
	src = src + "<" + t.key + attributes + ">"

	// For each of the child renderables, cast them into their specific type to
	// determine how to indent, and then recursively render their children.
	blockChildren := 0
	for _, child := range t.children {

		// If the child is a Component metaobject, get its root
		// element, and that is the thing
		if component, ok := child.(Component); ok {
			element := component.Render()
			if !element.Inline() {
				blockChildren++
			}
		} else if element, ok := child.(Element); ok {
			if !element.Inline() {
				blockChildren++
			}
		}

		childIndentLevel := i + 1
		src = src + child.Render(childIndentLevel)
	}

	if blockChildren > 0 {
		src = src + "\n" + padding(i)
	}
	src = src + "</" + t.key + ">"
	return src
}

func (t *HTMLElement) renderChildren(i int) string {
	s := ""
	for _, c := range t.children {
		s = s + c.Render(i)
	}
	return s
}

func (t *HTMLElement) renderAttributes() string {
	a := ""
	for _, v := range t.attributes {
		a = a + " " + v.Render()
	}
	return a
}

// Attribute is a tag attribute consiting of a key and a value.
type Attribute struct {
	Key, Value string
}

// Render the attribute, which qquotes the value verbatim.
// TODO: escaping
func (a Attribute) Render() string {
	if a.Value == "" {
		return a.Key
	}
	return a.Key + "=\"" + a.Value + "\""
}
