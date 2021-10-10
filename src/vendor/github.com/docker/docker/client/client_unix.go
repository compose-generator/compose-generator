/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
// +build linux freebsd openbsd netbsd darwin solaris illumos dragonfly

package client // import "github.com/docker/docker/client"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"

const defaultProto = "unix"
const defaultAddr = "/var/run/docker.sock"
