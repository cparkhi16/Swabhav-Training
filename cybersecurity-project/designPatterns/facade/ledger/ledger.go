package ledger

type Ledger struct {
	entries []string
}

func New(entries []string) *Ledger {
	return &Ledger{
		entries: entries,
	}
}

func (l *Ledger) MakeEntry(operation string) {
	l.entries = append(l.entries, operation)
}

func (l *Ledger) GetDetails() []string {
	return l.entries
}
