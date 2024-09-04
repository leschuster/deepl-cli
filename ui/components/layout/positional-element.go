package layout

import (
	"github.com/charmbracelet/lipgloss"
)

type Position int

const (
	Left   Position = iota
	Center Position = iota
	Right  Position = iota
)

const (
	empty    = "empty"
	fill     = "fill"
	fillAuto = "fill-auto"
	fixed    = "fixed"
)

type PositionalElement struct {
	elType          string // empty, fill, fixed
	model           *LayoutModel
	position        Position // Left, Center, Right
	widthFixed      int      // only for fixed width
	widthFill       float64  // only for fill; fraction of how much space to take
	calculatedWidth int
	selectable      bool // can the element be selected
}

// Empty element. Fill empty space in Layout.
func Empty() PositionalElement {
	return PositionalElement{
		elType:     empty,
		model:      nil,
		position:   Left,
		selectable: false,
	}
}

// Fill width dynamically.
// position: left | center | right
// width: fraction of how much space to take, e.g. 0.5 for 50%
func Fill(model *LayoutModel, position Position, width float64) PositionalElement {
	return PositionalElement{
		elType:     fill,
		model:      model,
		position:   position,
		widthFill:  width,
		selectable: true,
	}
}

// Fill width dynamically. Split space equally among fill-auto elements
// position: left | center | right
func FillAuto(model *LayoutModel, position Position) PositionalElement {
	return PositionalElement{
		elType:     fillAuto,
		model:      model,
		position:   position,
		selectable: true,
	}
}

// Fixed width.
func Fixed(model *LayoutModel, position Position, width int) PositionalElement {
	return PositionalElement{
		elType:     fixed,
		model:      model,
		position:   position,
		widthFixed: width,
		selectable: true,
	}
}

// Element cannot be selected by user
func (p PositionalElement) NotSelectable() PositionalElement {
	p.selectable = false
	return p
}

func (p *PositionalElement) IsSelectable() bool {
	return p.selectable
}

func (p PositionalElement) getType() string {
	return p.elType
}

func (p PositionalElement) view() string {
	if p.elType == empty {
		return ""
	}

	var pos lipgloss.Position
	switch p.position {
	case Left:
		pos = lipgloss.Left
	case Center:
		pos = lipgloss.Center
	case Right:
		pos = lipgloss.Right
	default:
		pos = lipgloss.Center
	}

	content := ""
	if p.model != nil {
		content = (*p.model).View()
	}

	return lipgloss.PlaceHorizontal(p.calculatedWidth, pos, content)
}

func (p PositionalElement) setActive() PositionalElement {
	if p.model != nil {
		m := (*p.model).SetActive()
		p.model = &m
	}
	return p
}

func (p PositionalElement) unsetActive() PositionalElement {
	if p.model != nil {
		m := (*p.model).UnsetActive()
		p.model = &m
	}
	return p
}

func (p PositionalElement) setCalculatedWidth(width int) PositionalElement {
	p.calculatedWidth = width
	if p.model != nil {
		m := (*p.model).OnAvailWidthChange(width)
		p.model = &m
	}
	return p
}
