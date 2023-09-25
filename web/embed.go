package web

import (
	"embed"
)

//go:embed all:dist
var embedAdmin embed.FS

//go:embed dist/index.html
var embedAdminIndex embed.FS

// EmbedAdmin is ...
func EmbedAdmin() embed.FS {
	return embedAdmin
}

// EmbedAdminIndex is ...
func EmbedAdminIndex() embed.FS {
	return embedAdminIndex
}
