// +build noplugin

package loader

import (
	"errors"

	iplugin "github.com/fgeth/fg-ipfs/plugin"
)

func init() {
	loadPluginFunc = nopluginLoadPlugin
}

func nopluginLoadPlugin(string) ([]iplugin.Plugin, error) {
	return nil, errors.New("not built with plugin support")
}
