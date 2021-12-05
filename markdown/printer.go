package markdown

import (
	"fmt"
	"io"
	"sync"
)

// Printer that prints to []string, makes writing and testing console apps easier.
type Printer struct {
	mu      sync.Mutex
	w       io.Writer
	e       io.Writer
	lines   []string
	history []string
	errors  []string
}

// NewPrinter creates a new buffered writer, typically from os.Stdout and os.Stderr
func NewPrinter(w io.Writer, e io.Writer) *Printer {
	p := &Printer{}
	p.w = w
	p.e = e
	return p
}

// NewTestWriter returns a test writer that does not flush to the console
func NewTestWriter() *Printer {
	p := &Printer{}
	p.w = nil
	p.e = nil
	return p
}

// Println prints and appends a line to the internal stdout printer buffer
func (p *Printer) Println(format string, a ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	line := fmt.Sprintf(format+"\n", a...)
	p.lines = append(p.lines, line)
}

// PrintErrln prints and appends a line to the internal std err buffer
func (p *Printer) PrintErrln(format string, a ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	line := fmt.Sprintf(format+"\n", a...)
	p.errors = append(p.errors, line)
}

// GetLines returns all the lines printed.
func (p *Printer) GetLines() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	lines := make([]string, len(p.lines))
	copy(lines, p.lines)
	copy(lines, p.errors)
	return lines
}

// Flush buffered output to stdout and stderr writers
func (p *Printer) Flush() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.history = append(p.history, p.lines...)
	p.history = append(p.history, p.errors...)
	if p.w != nil {
		for _, l := range p.lines {
			fmt.Fprint(p.w, l)
		}
	}
	if p.e != nil {
		for _, l := range p.errors {
			fmt.Fprint(p.e, l)
		}
	}
	p.lines = make([]string, 0)
	p.errors = make([]string, 0)
}
