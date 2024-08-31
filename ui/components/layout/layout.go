package layout

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Layout struct {
	rows                      []Row
	screenWidth, screenHeight int
	colCount, rowCount        int // "Size of Matrix"
	x, y                      int // Active Element
}

func NewLayout(rows ...Row) (*Layout, error) {
	// Check for equal row lenghts
	if len(rows) > 0 {
		length := rows[0].GetLength()

		for _, r := range rows {
			if length != r.GetLength() {
				return nil, fmt.Errorf("rows must all have the same number of elements")
			}
		}
	}

	// Get column count
	colCount := 0
	if len(rows) > 0 {
		colCount = rows[0].GetLength()
	}

	// Get row count
	rowCount := len(rows)

	return &Layout{
		rows:     rows,
		colCount: colCount,
		rowCount: rowCount,
		x:        0, y: 0,
	}, nil
}

func (l *Layout) Init() tea.Cmd {
	//TODO
	return nil
}

func (l *Layout) UpdateActive(msg tea.Msg) tea.Cmd {
	return nil
}

func (l *Layout) UpdateAll(msg tea.Msg) tea.Cmd {
	return nil
}

func (l *Layout) GetActive() *LayoutModel {
	return nil
}

// Set active element
func (l *Layout) SetActive(x, y int) {
	// TODO
}

func (l *Layout) View() string {
	rowsRendered := []string{}

	for _, row := range l.rows {
		if str := row.View(); str != "" {
			rowsRendered = append(rowsRendered, str)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		rowsRendered...,
	)
}

func (l *Layout) Resize(width, height int) {
	l.screenWidth, l.screenHeight = width, height

	for i, row := range l.rows {
		l.rows[i] = row.setWidth(width)
	}
}

func (l *Layout) NavigateUp() {

}

func (l *Layout) NavigateRight() {}

func (l *Layout) NavigateDown() {}

func (l *Layout) NavigateLeft() {}
