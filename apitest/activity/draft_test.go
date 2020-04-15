package activity

import (
	"fmt"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestActivityDraftList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityDrafts(
			filter:{
				chartUUID_in:["L3YLH7WC"]
			})
			{
			uuid
			key
			chartUUID
			type
			number
			name
			namePinyin
			description
			startTime
			endTime
			progress
			parent
			path
			position
			createTime
			updateTime
			data
			displayStatus
			assign{uuid,name,avatar}
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

func TestAddActivityDraft(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		addActivityDraft(
			name: "test_activity",
			item_type: "activity_draft",
			chart_uuid: "TEJwQRGz",
			type: "ppm_task",
			progress: 8000000,
			start_time: 12000000,
			end_time: 1786793599,
			parent: "LtEFxUiu",
			assign: "%s"
		)
		{
			key
			name
			chartUUID
			type
			parent
			progress
			startTime
			endTime
		}
	}
	`
	query = fmt.Sprintf(query, apitest.C.UserUUID)
	body := apitest.BuildGraphqlQuery("mutation", query)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestUpdateActivityDraft(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		updateActivityDraft(
			key: "activity_draft-GiHCasX6"
			progress:100
			start_time: 100000000,
		)
		{
			key
			name
			chartUUID
			type
			parent
			path
			progress
			startTime
			endTime
			displayStatus
		}
	}
	`

	body := apitest.BuildGraphqlQuery("mutation", query)
	// body := `{"query":"\n        mutation UpdateGanttDraft {\n          updateActivityDraft (key: $key start_time: $start_time end_time: $end_time progress: $progress) {\n            key\n          }\n        }\n      ","variables":{"key":"activity_draft-3gabELgi","start_time":1586361600,"end_time":1586879999,"progress":0}}`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestDeleteActivityDraft(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		deleteActivityDraft(
			key: "activity_draft-RjvuuTza",
		)
		{
			key
			name
			chartUUID
			type
			parent
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
