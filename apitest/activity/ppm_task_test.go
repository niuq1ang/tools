package activity

import (
	"fmt"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestUpdatePPMTask(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		updatePpmTask(
			key: "ppm_task-BEfDpUxy",
			progress: 123,
			name: "new name"
		)
		{
			key
			name
			progress
			startTime
			endTime
		}
	}
	`
	body := apitest.BuildGraphqlQuery("mutation", query)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
