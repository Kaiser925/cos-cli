package output

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Table output rows as table fmt.
// Output will not show border.
func Table(rows [][]any) {
	t := table.NewWriter()
	style := table.StyleDefault
	style.Options.DrawBorder = false
	style.Box.MiddleVertical = ""
	t.SetStyle(style)
	trows := make([]table.Row, 0, len(rows))
	for _, r := range rows {
		trows = append(trows, r)
	}
	t.AppendRows(trows)
	fmt.Println(t.Render())
}

// Fatal prints v to os.Stderr, than call os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...any) {
	fmt.Fprint(os.Stderr, v...)
	os.Exit(1)
}

// Fatalf prints v to os.Stderr, than call os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...any) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

type KVPair struct {
	Key string
	Val any
}

// KeyValues prints key-value pairs with prefix.
func KeyValues(prefix string, kvs []KVPair) {
	maxLength := 0
	for _, v := range kvs {
		if l := len(v.Key); l > maxLength {
			maxLength = l
		}
	}
	for _, v := range kvs {
		s := text.AlignJustify.Apply(fmt.Sprintf("%s : ", v.Key), maxLength+1)
		fmt.Println(fmt.Sprintf("%s%s%v", prefix, s, v.Val))
	}
}
