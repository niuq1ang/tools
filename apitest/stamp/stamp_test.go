package stamp_test

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

/*
	name = "team"
	name = "team_member"
	name = "department"
	name = "group"
	name = "project"
	name = "sprint"
	name = "issue_type"
	name = "issue_type_config"
	name = "field"
	name = "field_config"
	name = "task_status"
	name = "task_status_config"
	name = "transition"
	name = "permission_rule"
	name = "evaluated_permission"
	name = "task_stats"
*/
func TestListStampDatas(t *testing.T) {
	path := fmt.Sprintf("/team/%s/stamps/data", apitest.C.TeamUUID)
	body := `
	{
		"component_template":123
	}
	`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
