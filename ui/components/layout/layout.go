package layout

import (
	"fmt"
	"os"

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
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for _, row := range l.rows {
		for el := range row.NotNil() {
			if el.model == nil {
				continue
			}

			cmd = (*el.model).Init()
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

func (l *Layout) UpdateActive(msg tea.Msg) tea.Cmd {
	if el := l.GetActive(); el != nil && el.model != nil {
		return l.update(msg, l.x, l.y)
	}
	return nil
}

func (l *Layout) UpdateAll(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for y := 0; y < l.rowCount; y++ {
		for x := 0; x < l.colCount; x++ {
			cmd = l.update(msg, x, y)
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

func (l *Layout) GetActive() *PositionalElement {
	if l.colCount == 0 || l.rowCount == 0 {
		return nil
	}

	el := l.get(l.x, l.y)

	return &el
}

// Set active element
func (l *Layout) SetActive(x, y int) {
	if curr := l.GetActive(); curr != nil {
		el := curr.unsetActive()
		l.set(l.x, l.y, el)
	}

	if x >= l.colCount || y >= l.rowCount {
		fmt.Fprintf(os.Stderr, "index out of bound: tried accessing element (%d, %d)\n", x, y)
		os.Exit(1)
	}

	next := l.get(x, y)
	el := next.setActive()
	l.set(x, y, el)

	l.x, l.y = x, y
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

func (l *Layout) get(x, y int) PositionalElement {
	return l.rows[y].get(x)
}

func (l *Layout) set(x, y int, el PositionalElement) {
	l.rows[y].set(x, el)
}

func (l *Layout) update(msg tea.Msg, x, y int) tea.Cmd {
	el := l.get(x, y)

	if el.model != nil {
		teaModel, cmd := (*el.model).Update(msg)
		layoutModel := teaModel.(LayoutModel)
		el.model = &layoutModel
		l.set(x, y, el)

		return cmd
	}

	return nil
}
