package permission

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func TestAddPermission(t *testing.T) {
	path := fmt.Sprintf("/team/%s/permission_rules/add", apitest.C.TeamUUID)
	body := `
	{"permission_rule":{"context_type":"project","context_param":{"project_uuid":"JhWrCvGNKBUt4KCa"},"permission":"browse_project_schedule","user_domain_type":"single_user","user_domain_param":"%s"}}
	`
	body = fmt.Sprintf(body, apitest.C.UserUUID)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
