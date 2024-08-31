package layout

import (
	"fmt"
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
