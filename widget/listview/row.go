package listview

import (
	"fmt"
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Row represents a row in the list view.
type Row struct {
	element.Ele
}

// GetValues gets the rows values.
func (el *Row) GetValues() []string {
	tk.Get().Eval("%s item %s -values", el.GetParent().GetID(), el.GetID())
	str := tk.Get().GetStrResult()

	result := make([]string, 0)
	for _, val := range strings.Split(str, " ") {
		fmt.Printf("********* val: %v\n", val)
		if strings.HasPrefix(val, "{") && strings.HasSuffix(val, "}") {
			fmt.Printf("********* found prefix\n")
			result = append(result, val[1:len(val)-1])
		} else {
			result = append(result, val)
		}
	}

	return result
}
