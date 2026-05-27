package progressbar

// the progressbar indicates the current status and progress would be desired.
// ref: https://www.pixelstech.net/article/1596946473-A-simple-example-on-implementing-progress-bar-in-GoLang

type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	start   int64  // the init starting value for progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start, total int64) { _ = "STUB: not implemented"; return }

func (bar *Bar) getPercent() int64 { _ = "STUB: not implemented"; return 0 }

func (bar *Bar) Play(cur int64) { _ = "STUB: not implemented"; return }

func (bar *Bar) Finish() { _ = "STUB: not implemented"; return }
