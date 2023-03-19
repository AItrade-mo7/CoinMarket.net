package tmpl

import (
	_ "embed"
)

//go:embed ReClear.sh
var ReClear string

//go:embed SysReStart.sh
var SysReStart string
