package web

import (
	"embed"
)

//go:embed all:dist
var embedUI embed.FS

//go:embed dist/index.html
var embedUIIndex embed.FS

// EmbedAdmin is ...
func EmbedUI() embed.FS {
	return embedUI
}

// EmbedAdminIndex is ...
func EmbedUIIndex() embed.FS {
	return embedUIIndex
}
