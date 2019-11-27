package render

import (
	"sort"
	"time"

	"vbom.ml/util/sortorder"
)

// Fields represents a collection of row fields.
type Fields []string

// Row represents a colllection of columns.
type Row struct {
	ID     string
	Fields Fields
}

// Rows represents a collection of rows.
type Rows []Row

// Header represent a table header
type Header struct {
	Name  string
	Align int
}

// HeaderRow represents a table header.
type HeaderRow []Header

// RowSorter sorts rows.
type RowSorter struct {
	Rows  Rows
	Index int
	Asc   bool
}

// Delete removes an element by id.
func (rr Rows) Delete(id string) Rows {
	idx, ok := rr.Find(id)
	if !ok {
		return rr
	}

	if idx == 0 {
		return rr[1:]
	}
	if idx+1 == len(rr) {
		return rr[:len(rr)-1]
	}

	return append(rr[:idx], rr[idx+1:]...)
}

// NewRow returns a new row with initialized fields.
func NewRow(cols int) Row {
	return Row{Fields: make([]string, cols)}
}

// Find locates a row by id. Retturns false is not found.
func (rr Rows) Find(id string) (int, bool) {
	for i, r := range rr {
		if r.ID == id {
			return i, true
		}
	}

	return 0, false
}

// Sort rows based on column index and order.
func (rr Rows) Sort(col int, asc bool) {
	t := RowSorter{Rows: rr, Index: col, Asc: asc}
	sort.Sort(t)
}

func (s RowSorter) Len() int {
	return len(s.Rows)
}

func (s RowSorter) Swap(i, j int) {
	s.Rows[i], s.Rows[j] = s.Rows[j], s.Rows[i]
}

func (s RowSorter) Less(i, j int) bool {
	return Less(s.Asc, s.Rows[i].Fields[s.Index], s.Rows[j].Fields[s.Index])
}

func Less(asc bool, c1, c2 string) bool {
	if o, ok := isDurationSort(asc, c1, c2); ok {
		return o
	}

	b := sortorder.NaturalLess(c1, c2)
	if asc {
		return b
	}
	return !b
}

func isDurationSort(asc bool, s1, s2 string) (bool, bool) {
	d1, ok1 := isDuration(s1)
	d2, ok2 := isDuration(s2)
	if !ok1 || !ok2 {
		return false, false
	}

	if asc {
		return d1 <= d2, true
	}
	return d1 >= d2, true
}

func isDuration(s string) (time.Duration, bool) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return d, false
	}
	return d, true
}
