package Templater

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed Templates/*
var f embed.FS

var Templater *template.Template

func Init() error {
	var paths []string

	err := fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	Templater, err = template.ParseFS(f, paths...)
	if err != nil {
		return err
	}

	return nil
}
