package async

//----------------------------------------
// Task

// val: the value returned after task execution.
// err: the error returned during task completion.
// abort: tells Parallel to return, whether or not all tasks have completed.
type Task func(i int) (val interface{}, abort bool, err error)

type TaskResult struct {
	Value interface{}
	Error error
}

type TaskResultCh <-chan TaskResult

type taskResultOK struct {
	TaskResult
	OK bool
}

type TaskResultSet struct {
	chz     []TaskResultCh
	results []taskResultOK
}

func newTaskResultSet(chz []TaskResultCh) *TaskResultSet { _ = "STUB: not implemented"; return nil }

func (trs *TaskResultSet) Channels() []TaskResultCh { _ = "STUB: not implemented"; return nil }

func (trs *TaskResultSet) LatestResult(index int) (TaskResult, bool) {
	_ = "STUB: not implemented"
	return *new(TaskResult), false
}

// NOTE: Not concurrency safe.
// Writes results to trs.results without waiting for all tasks to complete.
func (trs *TaskResultSet) Reap() *TaskResultSet { _ = "STUB: not implemented"; return nil }

// Write result.

// else {
// We already wrote it.
// }

// Do nothing.

// NOTE: Not concurrency safe.
// Like Reap() but waits until all tasks have returned or panic'd.
func (trs *TaskResultSet) Wait() *TaskResultSet { _ = "STUB: not implemented"; return nil }

// Write result.

// else {
// We already wrote it.
// }

// Returns the firstmost (by task index) error as
// discovered by all previous Reap() calls.
func (trs *TaskResultSet) FirstValue() interface{} { _ = "STUB: not implemented"; return nil }

// Returns the firstmost (by task index) error as
// discovered by all previous Reap() calls.
func (trs *TaskResultSet) FirstError() error { _ = "STUB: not implemented"; return nil }

//----------------------------------------
// Parallel

// Run tasks in parallel, with ability to abort early.
// Returns ok=false iff any of the tasks returned abort=true.
// NOTE: Do not implement quit features here.  Instead, provide convenient
// concurrent quit-like primitives, passed implicitly via Task closures. (e.g.
// it's not Parallel's concern how you quit/abort your tasks).
func Parallel(tasks ...Task) (trs *TaskResultSet, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// To return.
// A "wait group" channel, early abort if any true received.
// Keep track of panics to set ok=false later.

// We will set it to false iff any tasks panic'd or returned abort.

// Start all tasks in parallel in separate goroutines.
// When the task is complete, it will appear in the
// respective taskResultCh (associated by task index).

// Capacity for 1 result.

// Recovery

// Send panic to taskResultCh.

// Closing taskResultCh lets trs.Wait() work.

// Decrement waitgroup.

// Run the task.

// Send val/err to taskResultCh.
// NOTE: Below this line, nothing must panic/

// Closing taskResultCh lets trs.Wait() work.

// Decrement waitgroup.

// Wait until all tasks are done, or until abort.
// DONE_LOOP:

// Ok is also false if there were any panics.
// We must do this check here (after DONE_LOOP).
