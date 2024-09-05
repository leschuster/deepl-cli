package layout

import (
	"fmt"
	"iter"
	"math"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// Provides one Row in the Layout
// Make sure that all rows have the same number of elements.
// You can use an "empty" element to fill the gaps.
type Row struct {
	elements []PositionalElement
	width    int
}

// Create a new row
func NewRow(elements ...PositionalElement) Row {
	return Row{
		elements: elements,
	}
}

// Get count of element in row
func (r *Row) GetLength() int {
	return len(r.elements)
}

// Render row
func (r *Row) View() string {
	elementsRendered := []string{}

	for _, el := range r.elements {
		if rendered := el.view(); rendered != "" {
			elementsRendered = append(elementsRendered, rendered)
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		elementsRendered...,
	)
}

// Iterate over all elements
func (r *Row) All() iter.Seq[PositionalElement] {
	return func(yield func(PositionalElement) bool) {
		for _, el := range r.elements {
			if !yield(el) {
				return
			}
		}
	}
}

// Iterate over all non-nil elements
func (r *Row) NotNil() iter.Seq[PositionalElement] {
	return func(yield func(PositionalElement) bool) {
		for _, el := range r.elements {
			if el.model == nil {
				continue
			}
			if !yield(el) {
				return
			}
		}
	}
}

// set actual width
func (r Row) setWidth(width int) Row {
	r.width = width

	widthPerAutoEl := r.calcWithPerAutoEl()
	r.applyCalculatedWidths(widthPerAutoEl)

	return r
}

// Calculate the width that remains per fill-auto element
// It is calculated by starting with the row's width,
// substracting all fixed widths and fill widths,
// then dividing the remainder by the number of fill-auto elements
func (r *Row) calcWithPerAutoEl() int {
	countAutoEl := 0
	fixedWidth := 0

	for _, el := range r.elements {
		switch el.getType() {
		case empty:
			continue
		case fixed:
			fixedWidth += el.widthFixed
		case fill:
			fixedWidth += int(math.Floor(el.widthFill * float64(r.width)))
		case fillAuto:
			countAutoEl++
		}
	}

	if fixedWidth > r.width {
		// Thats too much
		fmt.Fprintf(os.Stderr, "Warning: elements in row take up %d in width, but row only is %d characters wide\n", fixedWidth, r.width)
	}

	availWidth := r.width - fixedWidth

	return int(math.Floor(float64(availWidth) / float64(countAutoEl)))
}

// Send calculated widths to elements
func (r *Row) applyCalculatedWidths(widthPerAutoEl int) {
	for i, el := range r.elements {
		var res PositionalElement

		switch el.getType() {
		case empty:
			res = el.setCalculatedWidth(0)
		case fixed:
			res = el.setCalculatedWidth(el.widthFixed)
		case fill:
			res = el.setCalculatedWidth(int(math.Floor(el.widthFill * float64(r.width))))
		case fillAuto:
			res = el.setCalculatedWidth(widthPerAutoEl)
		}

		r.set(i, res)
	}
}

// Get element at position i
func (r Row) get(i int) PositionalElement {
	return r.elements[i]
}

// Set element at position i
func (r *Row) set(i int, el PositionalElement) {
	r.elements[i] = el
}
