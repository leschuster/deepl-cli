package layout

type PositionalElement struct {
	mode       string // empty, fill, fixed
	model      *LayoutModel
	position   string  // left, center, right
	widthFixed int     // only for fixed width
	widthFill  float32 // only for fill; fraction of how much space to take
	selectable bool    // can the element be selected
}

// Empty element. Fill empty space in Layout.
func Empty() PositionalElement {
	return PositionalElement{
		mode:       empty,
		model:      nil,
		selectable: false,
	}
}

// Fill width dynamically.
// position: left | center | right
// width: fraction of how much space to take, e.g. 0.5 for 50%
func Fill(model *LayoutModel, position string, width float32) PositionalElement {
	return PositionalElement{
		mode:       fill,
		model:      model,
		position:   position,
		widthFill:  width,
		selectable: false,
	}
}

// Fill width dynamically. Split space equally among fill-auto elements
// position: left | center | right
func FillAuto(model *LayoutModel, position string) PositionalElement {
	return PositionalElement{
		mode:       fillAuto,
		model:      model,
		position:   position,
		selectable: true,
	}
}

// Fixed width.
func Fixed(model *LayoutModel, width int) PositionalElement {
	return PositionalElement{
		model:      model,
		widthFixed: width,
		selectable: true,
	}
}

// Element cannot be selected by user
func (p PositionalElement) NotSelectable() PositionalElement {
	p.selectable = false
	return p
}

func (p PositionalElement) View() string {
	content := ""
	if p.model != nil {
		content = (*p.model).View()
	}

	// TODO

	return content
}

func (p *PositionalElement) IsSelectable() bool {
	return p.selectable
}
