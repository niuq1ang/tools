package activity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestActivityReleaseList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityReleases(
			filter:{
				chartUUID_in:[\"Fx47D3aw\"]
			})
			{
				key
				name 
				pushTime
				pushStatus
				pushUser{
					uuid
					name
				}
			}
	}
	`
	query = strings.Replace(query, "\n", "\\n", -1)
	query = strings.Replace(query, "\t", "", -1)
	fmt.Println(query)
	body := `
	{
		"query": "%s"
	}
	`
	body = fmt.Sprintf(body, query)

	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestAddActivityRelease(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
	{
		addActivityRelease(
			name: "test_activity_release",
			chart_uuid: "CL9Zypb7",
			project: "D2pLSaJ3Q3AnLd66"
		)
		{
			key
			name
			createTime
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
	fmt.Println(body)

	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestUpdateActivityRelease(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
	{
		updateActivityRelease(
			key: "activity_release-Rs4iw4AC",
			project: "D2pLSaJ3Q3AnLd61",
			name: "123123"
		)
		{
			key
			name
			createTime
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
	fmt.Println(body)

	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
