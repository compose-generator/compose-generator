/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build !js

package viper

import "github.com/fsnotify/fsnotify"

type watcher = fsnotify.Watcher

func newWatcher() (*watcher, error) {
	return fsnotify.NewWatcher()
}
