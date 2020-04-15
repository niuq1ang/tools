package activity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestActivityList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activities(
			filter:{
				chartUUID_in:["TEJwQRGz"]
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
				key: "activity-T5Nkx9kW",
				progress: 3300000,
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
