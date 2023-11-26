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
//go:generate go run ../../internal/tools/generate/main.go -recv=*ListView -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*ListView -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*ListView -pkg=common/padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*ListView -pkg=common/scrollbar
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
// This will return nil if index is out of bounds.
func (el *ListView) GetColumn(index int) *Column {
	if index < len(el.columns) {
		return el.columns[index]
	}

	return nil
}

// RegisterTag registers a tag to be used by rows.
// See [option.color] for color names. Use color.Default for no color.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *ListView) RegisterTag(name string, foregroundColor, backgroundColor string) {
	tk.Get().Eval("%s tag configure {%s} -foreground {%s} -background {%s}", el.GetID(), name, foregroundColor, backgroundColor)
}

// AddValues adds values to the list view.
func (el *ListView) AddRow(values ...string) *Row {
	row := &Row{}
	row.SetParent(el)

	valStr := strings.Join(values, "} {")
	tk.Get().Eval("%s insert {} end -values [list {%s}]", el.GetID(), valStr)

	rowID := tk.Get().GetStrResult()
	row.SetID(rowID)

	el.rowRef[rowID] = row
	el.rows = append(el.rows, row)

	return row
}

// GetRow gets a row by its index.
// This will return nil if index is out of bounds.
func (el *ListView) GetRow(index int) *Row {
	if index < len(el.rows) {
		return el.rows[index]
	}

	return nil
}

// GetSelectedRow returns the first selected row.
// This will return nil if nothing is selected.
func (el *ListView) GetSelectedRow() *Row {
	rows := el.GetSelectedRows()

	if len(rows) > 0 {
		return rows[0]
	}

	return nil
}

// GetSelectedRows gets all the selected rows as an slice.
func (el *ListView) GetSelectedRows() []*Row {
	tk.Get().Eval("%s selection", el.GetID())
	ids := tk.Get().GetStrSliceResult()

	result := make([]*Row, 0)

	for _, id := range ids {
		if row, ok := el.rowRef[id]; ok {
			result = append(result, row)
		}
	}

	return result
}

// Clear clears the list view.
func (el *ListView) Clear() {
	tk.Get().Eval("%s children {}", el.GetID())
	tk.Get().Eval("%s delete [list %s]", el.GetID(), tk.Get().GetStrResult())

	clear(el.rowRef)
	clear(el.rows)
	clear(el.columns)
}
