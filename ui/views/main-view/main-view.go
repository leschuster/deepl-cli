package mainview

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	languageselect "github.com/leschuster/deepl-cli/ui/components/language-select"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx                 *context.ProgramContext
	inputTextModel      textarea.Model
	outputTextModel     textarea.Model
	sourceLanguageModel languageselect.Model
	targetLanguageModel languageselect.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:                 ctx,
		inputTextModel:      textarea.InitialModel(ctx),
		outputTextModel:     textarea.InitialModel(ctx),
		sourceLanguageModel: languageselect.InitialModel(ctx, "Auto"),
		targetLanguageModel: languageselect.InitialModel(ctx, "-"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	doc := strings.Builder{}

	left := lipgloss.JoinVertical(
		lipgloss.Top,
		m.sourceLanguageModel.View(),
		m.inputTextModel.View(),
	)

	right := lipgloss.JoinVertical(
		lipgloss.Top,
		m.targetLanguageModel.View(),
		m.outputTextModel.View(),
	)

	doc.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		right,
	))

	return doc.String()
}
