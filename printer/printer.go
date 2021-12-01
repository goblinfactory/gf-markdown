package printer

import (
	"fmt"
	"sync"
)

// Printer that prints to []string, makes writing and testing console apps easier.
type Printer struct {
	mu      sync.Mutex
	lines   []string
	history []string
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

// Flush buffered output to console
func (p *Printer) Flush() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.history = append(p.history, p.lines...)
	for _, l := range p.lines {
		fmt.Println(l)
	}
	p.lines = make([]string, 0)
}
