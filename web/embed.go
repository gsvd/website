package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed static/* src/views/*
var FS embed.FS

func ViewsFS() fs.FS {
	sub, err := fs.Sub(FS, "src/views")
	if err != nil {
		log.Fatalf("failed to get subfs: %v", err.Error())
	}
	return sub
}

func StaticFS() fs.FS {
	sub, err := fs.Sub(FS, "static")
	if err != nil {
		log.Fatalf("failed to get subfs: %v", err.Error())
	}
	return sub
}
