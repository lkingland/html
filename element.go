// package html defines renderable HTML primitives.
// It's raw.  It's verbose.  It's oft opinionated in what it omits.
// But it endeavors to be strictly correct with what it does.
// For something more friendly, see the ui package.
package html

import "fmt"

// Element is a representation of an HTML element.
// An element consists of either an empty tag <img />
// or standard tag <div></div>, with attributes <div key="value">
// and potentially children if it is not empty <div><img /></div>
// When rendered, a tag will be indented based on its level of nesting,
// and will have linebreaks in the appropriate place depending on
// it's being either an inline or non-inline element.
type Element struct {
	Key        string
	Empty      bool
	Inline     bool
	Attributes []Attribute
	Children   []Renderable
}

// Set the value of an attribute on the start tag of the element.
// <tagname key="value"></tagname>
func (t *Element) Set(k, v string) *Element {
	for i, a := range t.Attributes {
		if a.Key == k {
			t.Attributes[i].Value = v
			break
		}
	}
	a := Attribute{Key: k, Value: v}
	t.Attributes = append(t.Attributes, a)
	return t
}

// Add something renderable to the element's children
func (t *Element) Add(r Renderable) *Element {
	t.Children = append(t.Children, r)
	return t
}

// return a padding string for a given indent level
func padding(indent int) string {
	padFmt := fmt.Sprintf("%%%ds", (indent * 2))
	return fmt.Sprintf(padFmt, "")
}

// Render the element including recursively rendering its children.
func (t *Element) Render(i int) string {
	// pseudoelements like root are denoted by a blank tagname (key)
	// They are rendered by immediately delegating to their children
	// without any change to indentation, but with a trailing space
	// for after the closing html tag.
	if t.Key == "" {
		return t.renderChildren(i) + "\n"
	}

	s := "\n" + padding(i)

	a := t.renderAttributes()

	if t.Empty {
		return s + "<" + t.Key + a + " />"
	}

	s = s + "<" + t.Key + a + ">"
	blockChildren := 0
	for _, c := range t.Children {
		t, ok := c.(*Element)
		if ok && !t.Inline {
			blockChildren++
		}
		i2 := i + 1
		s = s + c.Render(i2)
	}

	if blockChildren > 0 {
		s = s + "\n" + padding(i)
	}
	s = s + "</" + t.Key + ">"
	return s
}

func (t *Element) renderChildren(i int) string {
	s := ""
	for _, c := range t.Children {
		s = s + c.Render(i)
	}
	return s
}

func (t *Element) renderAttributes() string {
	a := ""
	for _, v := range t.Attributes {
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
