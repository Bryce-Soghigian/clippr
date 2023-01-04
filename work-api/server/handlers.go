package server

import (
	"clippr/jobsv1"
	"encoding/json"
	"io"
	"net/http"
)

func (s *ClipprServer) handleIngestWorkItem(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(
			response,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		)
		return
	}
	var workItems []*jobsv1.WorkItem
	err = json.Unmarshal(body, &workItems)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	for _, workitem := range workItems {
		s.workQueue.AddRateLimited(workitem)
	}

}
