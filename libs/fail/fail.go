package fail

func envSet() int { _ = "STUB: not implemented"; return 0 }

// Fail when FAIL_TEST_INDEX == callIndex
var callIndex int // indexes Fail calls

func Fail() { _ = "STUB: not implemented"; return }

func Exit() { _ = "STUB: not implemented"; return }

//	proc, _ := os.FindProcess(os.Getpid())
//	proc.Signal(os.Interrupt)
//	panic(fmt.Sprintf("*** fail-test %d ***", callIndex))
