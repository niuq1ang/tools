package activity

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func TestPPMTaskList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		ppmTasks(
			filter:{
				projectUUID_in:["%s"]
			})
			{
			uuid
			key
			name
			namePinyin
			assign{uuid,name,avatar}
			}
	}
	`
	query = fmt.Sprintf(query, activityChartUUID)
	body := apitest.BuildGraphqlQuery("query", query)
	body = `{"query":"\n    query PPM_TASKS($filterGroup: filterGroup, $orderBy: orderBy) {\n       ppmTasks(filterGroup: $filterGroup, orderBy: $orderBy){\n          key\n          uuid\n          name\n          progress\n          owner {\n            uuid\n          }\n          project {\n            uuid\n          }\n          number\n          assign {\n            uuid\n            name\n            avatar\n          }\n          description\n          startTime\n          endTime\n       }\n    }\n  ","variables":{"filterGroup":[{"project_in":["JhWrCvGNDqWa3JvZ"],"ppmType_in":[1]}],"orderBy":{"createTime":"DESC"}}}`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("response: %s", ret)
}

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
