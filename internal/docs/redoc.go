//go:build docs
// +build docs

package docs

import "embed"

//go:embed re-doc.html
var ReDocFS embed.FS

//go:embed main_swagger.json
var SwaggerFS embed.FS

func init() {
	ReDoc = ReDocFS
	Swagger = SwaggerFS
}
