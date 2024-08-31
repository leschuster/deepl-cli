package layout

import (
	"fmt"
	"iter"
	"math"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type Row struct {
	elements []PositionalElement
	width    int
}

func NewRow(elements ...PositionalElement) Row {
	return Row{
		elements: elements,
	}
}

func (r *Row) GetLength() int {
	return len(r.elements)
}

func (r *Row) View() string {
	elementsRendered := []string{}
	widthPerAutoEl := r.calcWithPerAutoEl()

	for _, el := range r.elements {
		if rendered := el.view(widthPerAutoEl, r.width); rendered != "" {
			elementsRendered = append(elementsRendered, rendered)
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		elementsRendered...,
	)
}

func (r *Row) All() iter.Seq[PositionalElement] {
	return func(yield func(PositionalElement) bool) {
		for _, el := range r.elements {
			if !yield(el) {
				return
			}
		}
	}
}

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

func (r *Row) calcWithPerAutoEl() int {
	countAutoEl := 0
	fixedWidth := 0

	for _, el := range r.elements {
		if el.getType() == fillAuto {
			countAutoEl++
		} else {
			fixedWidth += el.getFixedWidth(r.width)
		}
	}

	if fixedWidth > r.width {
		// Thats too much
		fmt.Fprintf(os.Stderr, "Warning: elements in row take up %d in width, but row only is %d characters wide\n", fixedWidth, r.width)
	}

	availWidth := r.width - fixedWidth

	return int(math.Floor(float64(availWidth) / float64(countAutoEl)))
}

func (r Row) setWidth(width int) Row {
	r.width = width
	return r
}

func (r Row) get(i int) PositionalElement {
	return r.elements[i]
}

func (r *Row) set(i int, el PositionalElement) {
	r.elements[i] = el
}
