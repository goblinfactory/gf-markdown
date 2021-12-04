package printer

import (
	"fmt"
	"io"
	"sync"
)

// Printer that prints to []string, makes writing and testing console apps easier.
type Printer struct {
	mu      sync.Mutex
	w       io.Writer
	lines   []string
	history []string
}

// New creates a new buffered writer
func New(w io.Writer) *Printer {
	p := &Printer{}
	p.w = w
	return p
}

// NewTestWriter returns a test writer that does not flush to the console.
func NewTestWriter() *Printer {
	p := &Printer{}
	p.w = nil
	return p
}

// Println prints and appends a line to the printer.
func (p *Printer) Println(format string, a ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	line := fmt.Sprintf(format, a...)
	p.lines = append(p.lines, line)
}

// GetLines returns all the lines printed.
func (p *Printer) GetLines() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	lines := make([]string, len(p.lines))
	copy(lines, p.lines)
	return lines
}

// Flush buffered output to writer
func (p *Printer) Flush() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.history = append(p.history, p.lines...)
	if p.w != nil {
		for _, l := range p.lines {
			fmt.Fprint(p.w, l)
		}
	}
	p.lines = make([]string, 0)
}
