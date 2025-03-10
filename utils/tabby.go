package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Tabby struct {
	writer *tabwriter.Writer
}

func TabbyNew() *Tabby {
	return &Tabby{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

func NewCustom(writer *tabwriter.Writer) *Tabby {
	return &Tabby{writer: writer}
}

func (t *Tabby) AddLine(args ...interface{}) {
	formatString := t.buildFormatString(args)
	fmt.Fprintf(t.writer, formatString, args...)
}

func (t *Tabby) AddHeader(args ...interface{}) {
	t.AddLine(args...)
	t.addSeparator(args)
}

// Print will write the table to the terminal
func (t *Tabby) Print() {
	t.writer.Flush()
}

func (t *Tabby) addSeparator(args []interface{}) {
	var b bytes.Buffer
	for idx, arg := range args {
		length := len(fmt.Sprintf("%v", arg))
		b.WriteString(strings.Repeat("-", length))
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	b.WriteTo(t.writer)
}

func (t *Tabby) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	return b.String()
}
