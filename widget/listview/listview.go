package listview

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "listview"
)

// The listview wiget is a secondary mode of the treeview wiget.
//
// There are two varieties of columns. The first is the main tree view column
// that is present all the time. The second are data columns that can be added
// when needed. This widget only uses the data columns.
//
// Each tree item has a list of tags, which can be used to associate event
// bindings and control their appearance. Treeview widgets support horizontal
// and vertical scrolling with the standard scroll commands.
//
// Virtual events that can also be bound to.
// <<TreeviewSelect>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_treeview.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=height
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=scrollbar
type ListView struct {
	widget.Widget

	rowRef  map[string]*Row
	rows    []*Row
	columns []*Column
}

// New creates a new tree view.
// See [option.selectionmode] for mode values.
func New(parent element.Element, numberOfColumns int, selectionMode string) *ListView {
	list := &ListView{
		rowRef:  make(map[string]*Row),
		columns: make([]*Column, 0),
	}

	list.SetParent(parent)
	list.SetType(Type)

	ids := make([]string, 0)

	for i := 0; i < numberOfColumns; i++ {
		col := &Column{}
		col.SetType("column")

		// Override the id, not using the parent.
		col.SetID(col.GetID())
		col.SetParent(list)

		list.columns = append(list.columns, col)

		ids = append(ids, col.GetID())
	}

	idStr := strings.Join(ids, " ")

	tk.Get().Eval("ttk::treeview %s -show {headings} -columns [list %s] -selectmode {%s}", list.GetID(), idStr, selectionMode)

	return list
}

// EnableHeadings controls showing the headings.
func (el *ListView) EnableHeadings(enable bool) {
	if enable {
		tk.Get().Eval("%s configure -show {headings}", el.GetID())
	} else {
		tk.Get().Eval("%s configure -show {}", el.GetID())
	}
}

// GetColumn gets a column by its index.
func (el *ListView) GetColumn(index int) *Column {
	return el.columns[index]
}

// AddValues adds values to the list view.
func (el *ListView) AddRow(values ...string) *Row {
	row := &Row{}
	row.SetParent(el)

	valStr := strings.Join(values, "\" \"")
	tk.Get().Eval("%s insert {} end -values [list \"%s\"]", el.GetID(), valStr)

	rowID := tk.Get().GetStrResult()
	row.SetID(rowID)

	el.rowRef[rowID] = row
	el.rows = append(el.rows, row)

	return row
}

// GetRow gets a row by its index.
func (el *ListView) GetRow(index int) *Row {
	return el.rows[index]
}
