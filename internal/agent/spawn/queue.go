package spawn

// queuedTask is a pending spawn request ordered by priority.
type queuedTask struct {
	Spec     Spec
	Priority Priority
}

// Queue is a priority queue for spawn requests (high drains first).
type Queue struct {
	high   []queuedTask
	normal []queuedTask
}

// NewQueue creates an empty priority queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds a task to the queue.
func (q *Queue) Enqueue(spec Spec) {
	task := queuedTask{Spec: spec, Priority: spec.Priority}
	if spec.Priority == PriorityHigh {
		q.high = append(q.high, task)
		return
	}
	q.normal = append(q.normal, task)
}

// Dequeue removes and returns the next task (high priority first).
func (q *Queue) Dequeue() (Spec, bool) {
	if len(q.high) > 0 {
		task := q.high[0]
		q.high = q.high[1:]
		return task.Spec, true
	}
	if len(q.normal) > 0 {
		task := q.normal[0]
		q.normal = q.normal[1:]
		return task.Spec, true
	}
	return Spec{}, false
}

// Len returns the total queued task count.
func (q *Queue) Len() int {
	return len(q.high) + len(q.normal)
}
