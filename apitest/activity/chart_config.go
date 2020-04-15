package activity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/CodesInvoker/tools/apitest"
)

func TestActivityChartPersonalConfigList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	query := `
	{
		activityChartPersonalConfigs
			{
			key
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

func TestUpdatePersonalConfigList(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)
	mutation := `
		{
			updateActivityChartPersonalConfig(
				key: "activity_chart_personal_config-SuavaiKZ-DAyGEN19",
				zooming: 60
			)
			{
				key
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
