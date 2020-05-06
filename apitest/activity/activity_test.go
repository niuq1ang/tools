package activity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func TestActivityList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activities(
			filter:{
				chartUUID_in:["%s"]
			})
			{
			key
			type
			name 
			parent
			path
			progress
			startTime
			endTime
			description
			createTime
			assign{
				name
			}
			number
			}

	}
	`
	query = fmt.Sprintf(query, activityChartUUID)
	body := apitest.BuildGraphqlQuery("query", query)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestUpateActivity(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
		{
			updateActivity(
				key: "activity-Q1iBfA5o",
				progress: 3300001,
				name: "new name"
			)
			{
				key
				progress
				name
			}
		}
		`
	mutation = strings.Replace(mutation, "\n", "\\n", -1)
	mutation = strings.Replace(mutation, "\t", "", -1)
	mutation = strings.Replace(mutation, "\"", "\\\"", -1)

	body := `
		{
			"query": "mutation %s"
		}
		`
	body = fmt.Sprintf(body, mutation)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
