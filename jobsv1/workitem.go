package jobsv1

type WorkItem struct {
	actionName string
	ID         string
	metadata   map[string]string
}
