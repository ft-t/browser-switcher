package launcher_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ft-t/browser-switcher/pkg/config"
	"github.com/ft-t/browser-switcher/pkg/launcher"
)

func TestLauncher(t *testing.T) {
	l := launcher.New("https://example.com")
	assert.NoError(t, l.Launch(context.TODO(), &config.Browser{}))
}
