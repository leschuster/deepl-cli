// The layout package provides a way of arranging and displaying multiple
// components in a single view with the ability to navigate between them.
// Each component needs to implement the LayoutModel interface, an extension
// of tea.Model. The interface provides methods to select a component as the
// currently active element (thus, changing the style of it).
// Components are arranged in rows and columns. Each column needs to have the same
// number of elements, although you may provide an "empty" element, that is both
// skipped in rendering and navigating.
// You may call NavigateUp, NavigateRight, NavigateDown, NavigateLeft to change
// which element is currently active.

package layout

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// The Layout struct provides a way of arranging and displaying multiple
// components in a single view with the ability to navigate between them.
type Layout struct {
	rows                      []Row
	screenWidth, screenHeight int
	colCount, rowCount        int // "Size of Matrix"
	active                    struct {
		x, y int
	}
}

// Get a new Layout instance
func NewLayout(rows ...Row) *Layout {
	// Check for equal row lenghts
	if len(rows) > 0 {
		length := rows[0].GetLength()

		for _, r := range rows {
			if length != r.GetLength() {
				fmt.Fprintf(os.Stderr, "Error in Layout definition: rows must all have the same number of elements\n")
				os.Exit(1)
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
	}
}

// Call Init() on all components
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

	if l.colCount > 0 && l.rowCount > 0 {
		l.SetActive(l.active.x, l.active.y)
	}

	return tea.Batch(cmds...)
}

// Update only the active element with the provided message
func (l *Layout) UpdateActive(msg tea.Msg) tea.Cmd {
	if el := l.GetActive(); el != nil && el.model != nil {
		return l.update(msg, l.active.x, l.active.y)
	}
	return nil
}

// Update all elements with the provided message
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

// Get a reference to the active element
func (l *Layout) GetActive() *PositionalElement {
	if l.colCount == 0 || l.rowCount == 0 {
		return nil
	}

	el := l.get(l.active.x, l.active.y)

	return &el
}

// Set active element
func (l *Layout) SetActive(x, y int) {
	if curr := l.GetActive(); curr != nil {
		el := curr.unsetActive()
		l.set(l.active.x, l.active.y, el)
	}

	if x >= l.colCount || y >= l.rowCount {
		fmt.Fprintf(os.Stderr, "index out of bound: tried accessing element (%d, %d)\n", x, y)
		os.Exit(1)
	}

	next := l.get(x, y)
	el := next.setActive()
	l.set(x, y, el)

	l.active.x, l.active.y = x, y
}

// Render the layout
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

// Set a new width and height for the Layout to occupy.
func (l *Layout) Resize(width, height int) {
	l.screenWidth, l.screenHeight = width, height

	for i, row := range l.rows {
		l.rows[i] = row.setWidth(width)
	}
}

func (l *Layout) NavigateUp() {
	xOld, yOld := l.active.x, l.active.y
	if yOld <= 0 { // Cannot go higher
		return
	}

	// Go one row higher and search for the element nearest xOld
	// If we do not find anything, go higher...
	yNew := yOld - 1

	for {
		if yNew < 0 {
			// Did not find anything
			// Stay with current element
			return
		}

		xNew, ok := getBestValue(xOld, l.rows[yNew].elements)
		if ok {
			// Yay, found it!
			l.SetActive(xNew, yNew)
			return
		}

		yNew-- // Try one row higher
	}
}

func (l *Layout) NavigateRight() {
	xOld, yOld := l.active.x, l.active.y
	if xOld >= l.colCount-1 { // Cannot go more to the right
		return
	}

	// Go one column to the right and search for the element nearest yOld
	xNew := xOld + 1

	for {
		if xNew >= l.colCount {
			// Did not find anything
			// Stay with current element
			return
		}

		yNew, ok := getBestValue(yOld, l.getCol(xNew))
		if ok {
			// Yay, found it!
			l.SetActive(xNew, yNew)
			return
		}

		xNew++ // Try one row more to the right
	}
}

func (l *Layout) NavigateDown() {
	xOld, yOld := l.active.x, l.active.y
	if yOld >= l.rowCount-1 { // Cannot go lower
		return
	}

	// Go one row lower and search for the element nearest xOld
	// If we do not find anything, go lower...
	yNew := yOld + 1

	for {
		if yNew >= l.rowCount {
			// Did not find anything
			// Stay with current element
			return
		}

		xNew, ok := getBestValue(xOld, l.rows[yNew].elements)
		if ok {
			// Yay, found it!
			l.SetActive(xNew, yNew)
			return
		}

		yNew++ // Try one row lower
	}
}

func (l *Layout) NavigateLeft() {
	xOld, yOld := l.active.x, l.active.y
	if xOld <= 0 { // Cannot go more to the left
		return
	}

	// Go one column to the left and search for the element nearest yOld
	xNew := xOld - 1

	for {
		if xNew < 0 {
			// Did not find anything
			// Stay with current element
			return
		}

		yNew, ok := getBestValue(yOld, l.getCol(xNew))
		if ok {
			// Yay, found it!
			l.SetActive(xNew, yNew)
			return
		}

		xNew-- // Try one row more to the left
	}
}

// Get element at position x, y
func (l *Layout) get(x, y int) PositionalElement {
	return l.rows[y].get(x)
}

func (l *Layout) getCol(x int) []PositionalElement {
	col := make([]PositionalElement, l.rowCount)

	for i, row := range l.rows {
		col[i] = row.get(x)
	}

	return col
}

// Set element at position x, y
func (l *Layout) set(x, y int, el PositionalElement) {
	l.rows[y].set(x, el)
}

// Update element at position x, y with msg
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

// Helper function to return the non-empty, selectable value nearest index "best"
func getBestValue(best int, slice [](PositionalElement)) (idx int, ok bool) {
	if isValidChoice(slice[best]) {
		return best, true
	}

	stepsLeft := best
	stepsRight := len(slice) - best - 1

	maxOffset := max(stepsLeft, stepsRight)

	for offset := 1; offset <= maxOffset; offset++ {
		if i := best - offset; i >= 0 && isValidChoice(slice[i]) {
			return i, true
		}

		if j := best + offset; j < len(slice) && isValidChoice(slice[j]) {
			return j, true
		}
	}

	return -1, false
}

// Checks if the provided element can be the new active element
func isValidChoice(el PositionalElement) bool {
	if el.model == nil {
		return false
	}
	if !el.selectable {
		return false
	}
	if el.elType == empty {
		return false
	}
	return true
}
