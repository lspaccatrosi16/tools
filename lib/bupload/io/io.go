package io

import "github.com/lspaccatrosi16/go-cli-tools/input"

func GetBucket() string {
	return input.GetInput("Bucket Name")
}

func GetObject() string {
	return input.GetInput("Cloud Object ID")
}

func GetPath() string {
	return input.GetInput("Local Path")
}
