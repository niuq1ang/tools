package activity

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

var (
	waterfallProjectUUID = "JhWrCvGNDqWa3JvZ"
	activityChartUUID    = "4FLuyMeH"
)

func TestActivityChartList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityCharts(
			filter:{
				project_in:["%s"]
			})
			{
			key
			name 
			drafting
			draftName
			createTime
			owner{
				name
			}
			}
	}
	`
	query = fmt.Sprintf(query, waterfallProjectUUID)
	body := apitest.BuildGraphqlQuery("query", query)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestUpdateChartOpenDrafting(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
		{
			updateActivityChart(
				key: "activity_chart-%s"
				drafting: true
				draft_name: "lalalalal125"
			)
			{
				key
				drafting
				draftName
			}
		}
		`
	mutation = fmt.Sprintf(mutation, activityChartUUID)
	body := apitest.BuildGraphqlQuery("mutation", mutation)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestUpdateChartCloseDrafting(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
		{
			updateActivityChart(
				key: "activity_chart-%s"
				drafting: false
			)
			{
				key
				drafting
				draftName
			}
		}
		`
	mutation = fmt.Sprintf(mutation, activityChartUUID)
	body := apitest.BuildGraphqlQuery("mutation", mutation)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestChartPublish(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/activity_chart/%s/publish", apitest.C.TeamUUID, waterfallProjectUUID, activityChartUUID)
	err, ret := apitest.DoPostRequest(path, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
