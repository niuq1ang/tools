package project

import (
	"fmt"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
	"github.com/bangwork/ones-ai-api-common/utils/uuid"
)

func TestAddProject(t *testing.T) {
	projectUUID := apitest.C.UserUUID + uuid.UUID()
	path := fmt.Sprintf("/team/%s/projects/add", apitest.C.TeamUUID)
	body := `
	{
		"project": {
			"uuid": "%s",
			"owner": "%s",
			"name": "test_add_project",
			"status": 1,
			"members": [],
			"type": "waterfall"
		},
		"template_id": "project-t5",
		"members": ["%s"]
	}
	`
	body = fmt.Sprintf(body, projectUUID, apitest.C.UserUUID, apitest.C.UserUUID)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
