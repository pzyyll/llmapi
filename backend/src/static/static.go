package static

import (
	"embed"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const BasePrefix = "/dashboard"

//go:embed all:dist
var StaticDistFS embed.FS
var Fs static.ServeFileSystem

//go:embed dist/index.html
var IndexHTML []byte

func ServeSPA() gin.HandlerFunc {
	return static.Serve(BasePrefix, Fs)
}

func ServeSSR() gin.HandlerFunc {
	return static.Serve(BasePrefix, Fs)
}

func init() {
	// This is a placeholder for any initialization logic if needed in the future.
	// Currently, it does nothing but can be expanded later.
	// For example, you could log that the package has been initialized or perform some setup.
	// fmt.Println("Static package initialized")
	var err error
	Fs, err = static.EmbedFolder(StaticDistFS, "dist")
	if err != nil {
		panic(err)
	}
}
