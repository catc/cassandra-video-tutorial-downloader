package progress

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const BAR_LENGTH = 30

type ProgressBars struct {
	TotalBars   int // total number of Bar
	LongestName int // longest filename
	NameTabSize int // number of tabs required for name before progress bar
	MinSpaces   int // minimum space between name and progress bar
	WaitGroup   *sync.WaitGroup
}

func New() *ProgressBars {
	clearScreen()

	var wg sync.WaitGroup
	return &ProgressBars{
		MinSpaces: 2,
		WaitGroup: &wg,
	}
}

type Bar struct {
	// file related stuff
	Filename string

	// progress related stuff
	Status    string
	Line      int
	Progress  float32
	BarLength int
}

func (pb *ProgressBars) CreateBar(filename string) *Bar {
	pb.WaitGroup.Add(1)
	pb.TotalBars = pb.TotalBars + 1

	longestName := pb.LongestName
	if longestName < len(filename) {
		pb.LongestName = len(filename)

		// update tab size
		pb.NameTabSize = (len(filename)+pb.MinSpaces)/8 + 1
	}

	bar := &Bar{
		Filename:  filename,
		Line:      pb.TotalBars,
		Status:    "Initializing",
		BarLength: BAR_LENGTH,
	}
	return bar
}

// update bar
func (pb *ProgressBars) SetBarProps(bar *Bar, newProgress float32, status string) {
	if status != "" {
		bar.Status = status
	}
	if newProgress != 0 {
		bar.Progress = newProgress
	}
	pb.updateBar(bar)
}

func (pb *ProgressBars) updateBar(bar *Bar) {
	// go to appropriate line
	goToLine(bar.Line)

	// print out progress
	bar.drawProgress(pb.NameTabSize, pb.MinSpaces)

	// return cursor to the bottom - so it isn't distracting
	goToLine(pb.TotalBars + 1)
}

func (pb *ProgressBars) CompletedBar() {
	pb.WaitGroup.Done()
}

func (pb *ProgressBars) Wait() {
	pb.WaitGroup.Wait()

	goToLine(pb.TotalBars + 1)
	write("Tasks completed.")
}

func (b *Bar) drawProgress(nameTabs int, MinSpaces int) {
	// name
	totalTabs := nameTabs - (len(b.Filename)-MinSpaces)/8
	name := fmt.Sprintf(
		"%s%s",
		b.Filename,
		strings.Repeat("\t", totalTabs))

	// bar
	barLength := float32(b.BarLength)
	completedProgress := int(b.Progress * barLength)
	uncompletedProgress := int(barLength - b.Progress*barLength)
	if uncompletedProgress > 0 {
		uncompletedProgress++
	}
	bar := fmt.Sprintf(
		"[%s%s]",
		strings.Repeat("=", completedProgress),
		strings.Repeat(" ", uncompletedProgress))

	// status
	percent := int(b.Progress * 100)
	status := fmt.Sprintf(" %s \t %s",
		strconv.Itoa(percent)+"%",
		b.Status)

	write(name + bar + status)
}
