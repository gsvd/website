package blog

import "embed"

//go:embed articles/*
var FS embed.FS
