package activity

import (
	"fmt"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestActivityChartList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityCharts(
			filter:{
				project_in:["D2pLSaJ36RKexu9p"]
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
				key: "activity_chart-TEJwQRGz"
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
				key: "activity_chart-TEJwQRGz"
				drafting: false
			)
			{
				key
				drafting
				draftName
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

func TestChartPublish(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/activity_chart/%s/publish", apitest.C.TeamUUID, "JhWrCvGNK4JGVsOq", "TEJwQRGz")
	err, ret := apitest.DoPostRequest(path, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
