package navigator

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Matrix [][](*NavModal)

// TODO: Support empty rows and cols

type pos struct {
	x, y int
}

type Navigator struct {
	matrix        Matrix
	pos           pos
	width, height int
}

func New(matrix Matrix) Navigator {
	width := 0
	if len(matrix) > 0 {
		width = len(matrix[0])
	}

	return Navigator{
		matrix: matrix,
		pos: pos{
			x: 0,
			y: 0,
		},
		width:  width,
		height: len(matrix),
	}
}

func (n *Navigator) Up() tea.Cmd {
	xOld := n.pos.x
	yOld := n.pos.y

	if yOld == 0 {
		// We cannot go higher
		return nil
	}

	yNew := yOld - 1
	xNew, ok := getBestValue(xOld, n.matrix[yNew])
	if !ok {
		// Nothing there
		return nil
	}

	n.setPos(pos{
		x: xNew,
		y: yNew,
	})

	return nil
}

func (n *Navigator) Right() tea.Cmd {
	xOld := n.pos.x
	yOld := n.pos.y

	if xOld == n.width-1 {
		// We cannot go further
		return nil
	}

	xNew := xOld + 1
	yNew, ok := getBestValue(yOld, n.getCol(xNew))
	if !ok {
		// Nothing there
		return nil
	}

	n.setPos(pos{
		x: xNew,
		y: yNew,
	})

	return nil
}

func (n *Navigator) Down() tea.Cmd {
	xOld := n.pos.x
	yOld := n.pos.y

	if yOld == n.height-1 {
		// We cannot go further down
		return nil
	}

	yNew := yOld + 1
	xNew, ok := getBestValue(xOld, n.matrix[yNew])
	if !ok {
		// Nothing there
		return nil
	}

	n.setPos(pos{
		x: xNew,
		y: yNew,
	})

	return nil
}

func (n *Navigator) Left() tea.Cmd {
	xOld := n.pos.x
	yOld := n.pos.y

	if xOld == 0 {
		// We cannot go further
		return nil
	}

	xNew := xOld - 1
	yNew, ok := getBestValue(yOld, n.getCol(xNew))
	if !ok {
		// Nothing there
		return nil
	}

	n.setPos(pos{
		x: xNew,
		y: yNew,
	})

	return nil
}

func (n *Navigator) UpdateActive(msg tea.Msg) tea.Cmd {
	return n.update(msg, n.pos)
}

func (n *Navigator) UpdateAll(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for y, row := range n.matrix {
		for x, item := range row {
			if item != nil {
				cmd = n.update(msg, pos{x: x, y: y})
				cmds = append(cmds, cmd)
			}
		}
	}

	return tea.Batch(cmds...)
}

func (n *Navigator) update(msg tea.Msg, pos pos) tea.Cmd {
	if model := n.get(pos); model != nil {
		teaModel, cmd := (*model).Update(msg)
		navModel := teaModel.(NavModal)
		n.set(&navModel, pos)

		return cmd
	}
	return nil
}

func (n *Navigator) GetCurr() *NavModal {
	return n.get(n.pos)
}

func (n *Navigator) View() string {
	renderedRows := []string{}

	for _, row := range n.matrix {
		renderedItems := []string{}

		for _, item := range row {
			if item != nil {
				renderedItems = append(renderedItems, (*item).View())
			}
		}

		renderedRows = append(renderedRows, lipgloss.JoinHorizontal(lipgloss.Top, renderedItems...))
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		renderedRows...,
	)
}

func (n *Navigator) get(pos pos) *NavModal {
	return n.matrix[pos.y][pos.x]
}

func (n *Navigator) set(model *NavModal, pos pos) {
	n.matrix[pos.y][pos.x] = model
}

func (n *Navigator) setPos(newPos pos) {
	// TODO: Refactor
	curr := n.GetCurr()
	if curr != nil {
		m := (*curr).UnsetActive()
		n.set(&m, n.pos)
	}

	next := n.get(newPos)
	if next != nil {
		m := (*next).SetActive()
		n.set(&m, newPos)
	}

	n.pos = newPos
}

func (n *Navigator) getCol(idx int) [](*NavModal) {
	col := make([](*NavModal), len(n.matrix))

	for i, row := range n.matrix {
		col[i] = row[idx]
	}

	return col
}

// Helper function to return the value nearest index "best"
func getBestValue(best int, slice [](*NavModal)) (idx int, ok bool) {
	if slice[best] != nil {
		return best, true
	}

	stepsLeft := best
	stepsRight := len(slice) - best - 1

	maxOffset := max(stepsLeft, stepsRight)

	for offset := 1; offset <= maxOffset; offset++ {
		if i := best - offset; i >= 0 && slice[i] != nil {
			return i, true
		}

		if j := best + offset; j < len(slice) && slice[j] != nil {
			return j, true
		}
	}

	return -1, false
}
