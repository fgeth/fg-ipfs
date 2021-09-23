// +build !cgo,!noplugin
// +build linux darwin freebsd

package loader

import (
	"errors"

	iplugin "github.com/fgeth/fg-ipfs/plugin"
)

func init() {
	loadPluginFunc = nocgoLoadPlugin
}

func nocgoLoadPlugin(fi string) ([]iplugin.Plugin, error) {
	return nil, errors.New("not built with cgo support")
}