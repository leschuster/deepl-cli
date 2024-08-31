package layout

import "github.com/charmbracelet/lipgloss"

type Row struct {
	elements []PositionalElement
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

	for _, el := range r.elements {
		if rendered := el.View(); rendered != "" {
			elementsRendered = append(elementsRendered, rendered)
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		elementsRendered...,
	)
}
