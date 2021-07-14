// go:generate stringer -type=Pill
package painingkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	lbuprofen
)
