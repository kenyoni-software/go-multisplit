package plugin_test

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/stretchr/testify/require"
)

func TestPluginExample(t *testing.T) {
	newPlugin, err := register.GetPlugin("multisplit")
	require.NoError(t, err)

	plugin, err := newPlugin(nil)
	require.NoError(t, err)

	_, err = plugin.BuildAnalyzers()
	require.NoError(t, err)
}
