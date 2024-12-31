package bubbles

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ft-t/browser-switcher/pkg/config"
)

type Bubbles struct {
	cfg      *config.Config
	launcher Launcher
}

func NewBubbles(
	cfg *config.Config,
	launcher Launcher,
) *Bubbles {
	return &Bubbles{
		cfg:      cfg,
		launcher: launcher,
	}
}

func (b *Bubbles) ShowManualSelect(
	ctx context.Context,
) error {
	var items []list.Item
	for _, browser := range b.cfg.Browsers {
		items = append(items, item{
			title: browser.Name,
			desc:  fmt.Sprintf("Rules: %v", len(browser.Rules)),
			launch: func() error {
				return b.launcher.Launch(ctx, browser)
			},
		})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Browsers"

	_, err := tea.NewProgram(m, tea.WithAltScreen()).Run()

	return err
}
