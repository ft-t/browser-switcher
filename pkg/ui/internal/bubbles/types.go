package bubbles

import (
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title  string
	desc   string
	launch func() error
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list   list.Model
	logger zerolog.Logger
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			if err := m.list.SelectedItem().(item).launch(); err != nil {
				m.logger.Panic().Err(err).Msg("failed to launch browser")
			}

			os.Exit(0)
		}
		if parsed, _ := strconv.ParseInt(msg.String(), 10, 64); parsed > 0 {
			if err := m.list.Items()[parsed-1].(item).launch(); err != nil {
				m.logger.Panic().Err(err).Msg("failed to launch browser")
			}

			os.Exit(0)
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
