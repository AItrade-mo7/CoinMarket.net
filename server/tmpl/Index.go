package tmpl

import (
	_ "embed"
)

//go:embed email-sys.html
var SysEmail string

type SysParam struct {
	Message      string
	SysTime      string
	SecurityCode string
}

//go:embed ReClear.sh
var ReClear string

//go:embed SysReStart.sh
var SysReStart string
