package test

import (
	"fmt"
	"strings"

	"github.com/CodesInvoker/tools/apitest"
	. "github.com/bangwork/bang-api/tests/utils"

	"testing"
)

func TestGQLListManhourReport(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/graphql", apitest.C.TeamUUID)

	query := `
	{
		manhourReports(
			orderBy:{
				createTime:DESC
			}
		)
			{
			key
			uuid
			name
			owner{
				uuid
				name
			}
			description
			fold
			builtIn
			config{
				type
				display
				dimensions
				range
				condition
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
	fmt.Printf("body: %s", body)
	err, ret := DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestAddManhourReport(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/add", apitest.C.TeamUUID)
	body := `
	{
		"item": {
			"item_type": "manhour_report",
			"config": {
				"condition": {
					"lock_conditions": [{
						"field_uuid": "field006",
						"operate": {
							"label": "filter.addQueryContent.include",
							"operate_id": "include",
							"predicate": "in"
						},
						"value": [
							"JhWrCvGNyWmcjYnU"
						]
					}],
					"condition_groups": [
						[{
							"field_uuid": "field004",
							"operate": {
								"label": "filter.addQueryContent.include",
								"operate_id": "include",
								"predicate": "in"
							},
							"value": [
								"4LBkqFVC"
							]
						}]
					]
				},
				"dimensions": [{
					"field": "task",
					"order_by": {
						"sprint": "desc"
					}
				}],
				"display": [
					"JhWrCvGNyWmcjYnU"
				],
				"range": {
					"from": "2019-02-14",
					"to": "2019-02-20",
					"quick": "last_7_days"
				},
				"type": "time_series"
			},
			"description": "",
			"fold": false,
			"name": "test"
		}
	}
	`

	err, ret := DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}

func TestUpdateManhourReport(t *testing.T) {
	path := fmt.Sprintf("/team/%s/item/%s/update", apitest.C.TeamUUID, "manhour_report-GpKk1DwU")
	body := `
	{
		"item": {
			"name": "213243",
			"config": {
				"type": "workload",
				"dimensions": [{
					"field": "owner",
					"order_by": {
						"average_workload_rate": "asc"
					}
				}],
				"condition": {
					"condition_groups": [
						[{
							"field_uuid": "field006",
							"operate": {
								"operate_id": "include",
								"predicate": "in",
								"negative": false,
								"label": "filter.addQueryContent.include",
								"filter_query": "in"
							},
							"value": []
						}]
					]
				},
				"display": ["manhour_record", "remaining_manhour", "standard_manhour", "resource_load"],
				"range": {
					"from": "2019-01-01",
					"to": "2019-12-31",
					"quick": "this_year"
				}
			},
			"fold": false,
			"item_type": "manhour_report"
		}
	}
	`

	err, ret := DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}

func TestManhourReportData(t *testing.T) {
	path := fmt.Sprintf("/team/%s/reports/graphql", apitest.C.TeamUUID)
	body := `
	{
		"operationName": "QUERY_MANHOURS",
		"variables": {
			"groupBy": {
				"manhours": {
					"owner": {}
				}
			},
			"orderBy": {
				"aggregateUser": {
					"namePinyin": "ASC"
				}
			},
			"filter": {
				"manhours": {
					"owner_notIn": [null],
					"startTime_range": {
						"unit": "day",
						"quick": "this_month"
					}
				}
			},
			"actualHoursSum": "manhours.hours",
			"actualHoursAvg": "workloadRateSeries",
			"timeSeries": {
				"timeField": "manhours.startTime",
				"valueField": "manhours.hours",
				"unit": "day",
				"quick": "this_month"
			},
			"timeSeries2": {
				"constant": 800000,
				"quick": "this_month",
				"timeField": "manhours.startTime",
				"unit": "day",
				"valueField": "manhours.hours",
				"workdays": [
					"Mon",
					"Tue",
					"Wed",
					"Thu",
					"Fri"
				]
			},
			"columnSource": "owner"
		},
		"query": "%s"}
	  `

	query := `
	query QUERY_MANHOURS($groupBy: GroupBy, $orderBy: OrderBy, $timeSeries: TimeSeriesArgs,  $timeSeries2: TimeSeriesArgs, $actualHoursSum: String, $filter: Filter, $columnSource: Source) {
		buckets(groupBy: $groupBy, orderBy: $orderBy, filter: $filter) {
		  ...ColumnBucketFragment
		  __typename
		}
	  }
	  
	  fragment UserSimple on User {
		key
		uuid
		name
		avatar
		email
		__typename
	  }
	  
	  fragment ColumnBucketFragment on Bucket {
		key
		columnField: aggregateUser(source: $columnSource) {
		  ...UserSimple
		  __typename
		}
		actualHours(sum: $actualHoursSum)
		actualHoursSeries(timeSeries: $timeSeries) {
		  times
		  values
		  __typename
		}
		standardWorkingHoursSeries(timeSeries: $timeSeries2) {
			times
			values
			__typename
		}
		remainingWorkingHoursSeries(timeSeries: $timeSeries) {
			times
			values
			__typename
		}
		workloadRateSeries(timeSeries: $timeSeries) {
			times
			values
			__typename
		}
		averageWorkloadRate
		__typename
	  }
	`

	query = strings.Replace(query, "\n", "\\n", -1)
	query = strings.Replace(query, "\t", "", -1)
	body = fmt.Sprintf(body, query)
	err, ret := DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
