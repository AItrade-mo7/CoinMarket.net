package tmpl

import (
	_ "embed"
)

//go:embed email-template.html
var Email string
