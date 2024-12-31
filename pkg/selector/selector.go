package selector

import (
	"context"
	"regexp"

	"github.com/rs/zerolog"

	"github.com/ft-t/browser-switcher/pkg/config"
)

type Selector struct {
	cfg *config.Config
}

func New(
	cfg *config.Config,
) *Selector {
	return &Selector{
		cfg: cfg,
	}
}

func (s *Selector) SelectBrowser(
	ctx context.Context,
	targetURL string,
) *config.Browser {
	for _, browser := range s.cfg.Browsers {
		for _, rule := range browser.Rules {
			lg := zerolog.Ctx(ctx).With().Str("browser_id", browser.ID).
				Str("rule", rule).Logger()

			if matched, err := regexp.MatchString(rule, targetURL); matched {
				lg.Debug().Msg("Matched rule")
				return browser
			} else if err != nil {
				lg.Error().Err(err).Msg("Failed to execute rule")
			}
		}
	}

	return nil
}
