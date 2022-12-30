package jobsv1

type WorkItem struct {
	action   string
	ID       string
	metadata map[string]string
}
