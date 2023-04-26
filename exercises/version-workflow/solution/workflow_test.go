package loanprocess

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/worker"
)

func TestReplayWorkflowHistoryFromFile(t *testing.T) {
	replayer := worker.NewWorkflowReplayer()

	replayer.RegisterWorkflow(LoanProcessingWorkflow)

	// NOTE: Your path will be that of the file you downloaded, such as
	// /Users/twheeler/Downloads/02c502fe-846c-4493-abfd-b6909935693c_events.json
	// instead of the one you see here (which was changed so that you can 
	// run the test).
	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "history_for_original_execution.json")

	assert.NoError(t, err)
}
