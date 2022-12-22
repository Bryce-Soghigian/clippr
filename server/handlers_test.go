package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"clippr/jobsv1"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	workqueue "k8s.io/apimachinery/pkg/util/workqueue"
)

func TestHandleIngestWorkItemBadRequest(t *testing.T) {
	// Set up a mock HTTP request with a malformed body
	request, err := http.NewRequest("POST", "/ingest", strings.NewReader("this is not valid JSON"))
	if err != nil {
		t.Fatal(err)
	}

	// Set up a mock HTTP response recorder
	response := httptest.NewRecorder()

	// Set up a mock ClipprServer
	s := &ClipprServer{
		workQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		logger:    zap.NewNop(),
	}

	// Call the function under test
	s.handleIngestWorkItem(response, request)

	// Assert that the response status code is 400 Bad Request
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected response status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestHandleIngestWorkItem(t *testing.T) {
	// Set up a mock HTTP request with a valid JSON body
	body, err := json.Marshal([]*jobsv1.WorkItem{
		{ID: "work-item-1"},
		{ID: "work-item-2"},
		{ID: "work-item-3"},
	})
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/ingest", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Set up a mock HTTP response recorder
	response := httptest.NewRecorder()

	// Set up a mock ClipprServer
	s := &ClipprServer{
		workQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		logger:    zap.NewNop(),
	}

	// Call the function under test
	s.handleIngestWorkItem(response, request)

	// Assert that the correct work items were added to the work queue
	if s.workQueue.Len() != 3 {
		t.Errorf("Expected 3 work items in the queue, got %d", s.workQueue.Len())
	}
	for i := 0; i < 3; i++ {
		item, _ := s.workQueue.Get()
		if item.(*jobsv1.WorkItem).ID != fmt.Sprintf("work-item-%d", i+1) {
			t.Errorf("Unexpected work item ID: %s", item.(*jobsv1.WorkItem).ID)
		}
		s.workQueue.Done(item)
	}
}
