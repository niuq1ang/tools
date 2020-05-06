package activity

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func TestActivityReleaseList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityReleases(
			filter:{
				chartUUID_in:["%s"]
			})
			{
				key
				name 
				updateTime
				publishStatus
				updateUser{
					uuid
					name
				}
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
	body := apitest.BuildGraphqlQuery("mutation", mutation)
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
	body := apitest.BuildGraphqlQuery("mutation", mutation)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestSprint(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		sprints{
				key
				name 
				status{
					uuid
					name
					category
				}
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

func TestProject(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		projects{
				key
				name 
				status{
					uuid
					name
					category
				}
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
