//go:build !cgo && !noplugin && (linux || darwin || freebsd)
// +build !cgo
// +build !noplugin
// +build linux darwin freebsd

package loader

import (
	"errors"

	iplugin "github.com/uss2022sayahi/kubo/plugin"
)

func init() {
	loadPluginFunc = nocgoLoadPlugin
}

func nocgoLoadPlugin(fi string) ([]iplugin.Plugin, error) {
	return nil, errors.New("not built with cgo support")
}
