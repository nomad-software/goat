package directorydialog

import "github.com/nomad-software/goat/internal/tk"

func init() {
	tk.Get().Eval("catch {tk_getOpenFile foo bar}")
	tk.Get().Eval("set ::tk::dialog::file::showHiddenVar 0")
	tk.Get().Eval("set ::tk::dialog::file::showHiddenBtn 1")
}
