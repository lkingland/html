package html

// Renderable indicates that with a passed indentation level, the implementing
// instance will return a string suitable for rendering.
type Renderable interface {
	Render(i int) string
}
