package pretty

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(tHead []string, tBody [][]string) {
	t := tablewriter.NewWriter(os.Stdout)
	
	t.SetHeader(tHead)
	for _, row := range tBody {
		t.Append(row)
	}
	t.Render()
}


