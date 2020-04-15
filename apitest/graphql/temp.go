package graphql

// &{map[activityDrafts:[map[type:milestone description: key:activity_draft-F63gGRiV path:GR7F5jTL-F63gGRiV startTime:12974400 progress:8000000 endTime:12974400 name:test_activity parent:GR7F5jTL] map[progress:8000000 description: name:test_activity parent: startTime:12974400 endTime:12974400 type:ppm_task_group path:GR7F5jTL key:activity_draft-GR7F5jTL] map[name:test_activity progress:8000000 description: parent: type:ppm_task endTime:13017599 key:activity_draft-T5Nkx9kW path:T5Nkx9kW startTime:11980800]]] []}

// bucketSelectionSet := &gql.SelectionSet{Selections: []*gql.Selection{groupBySelection, taskSelection}}
// updateTimeSL := &gql.Selection{Alias: utname, Name: utname}
// selectionSet

// mustFilter := gql.NewOrderedMap()
// mustFilter.Set("chartUUID", "L3YLH7WC")

// filterGroup := make([]interface{}, 0)
// filterGroup = append(filterGroup, mustFilter)
// selectArg := gql.NewOrderedMap()
// selectArg.Set("filterGroup", filterGroup)

// s2 := &gql.SelectionSet{Selections: []*gql.Selection{&gql.Selection{Alias: "key", Name: "key"}}}
// s1 := &gql.Selection{Alias: "activityDrafts", Name: "activityDrafts", SelectionSet: s2, Args: selectArg}
// selectionSet := &gql.SelectionSet{Selections: []*gql.Selection{s1}}
// q := &gql.Query{Name: "", Kind: "query", SelectionSet: selectionSet}

// data, err := filter.RequestGraphQL(teamUUID, userUUID, q, true)
// fmt.Println("liuyexing_______________data:", data)
// fmt.Println("liuyexing_______________err:", err)

// req := new(item.GraphQLReportRequest)
// reqStr := `
// {
// 	"query": "\n{\nactivityDrafts(\nfilter:{\nchartUUID_in:[\"L3YLH7WC\"]\n})\n{\nkey\ntype\nname \nparent\npath\nprogress\nstartTime\nendTime\ndescription\n}\n}\n"
// }
// `

// reqStr := `
// {
// 	"query": "query ACTIVITYDRAFTS($filter: Filter, $orderBy: OrderBy){activityDrafts(filter: $filter, orderBy: $orderBy){key\ntype\nname\nparent\npath\nprogress\nstartTime\nendTime\ndescription\n}\n}\n",
// 	"variables":{"filter":{"chartUUID_in":["L3YLH7WC"]}}
// }
// `

// {"query":"query PRODUCTS($filter: Filter, $orderBy: OrderBy) {\n    products(filter: $filter, orderBy: $orderBy) {\n      name\n      uuid\n      key\n      owner {\n        uuid\n        name\n      }\n\n      createTime\n      assign {\n        uuid\n        name\n      }\n      \n    productComponents {\n      uuid\n      name\n      parent{\n        uuid\n      }\n      key\n      type\n      contextType\n      contextParam1\n      contextParam2\n      position\n      urlSetting{\n        url\n      }\n      views{\n        key\n        uuid\n        name\n        builtIn\n      }\n    }\n  \n      \n    }\n  }","variables":{"orderBy":{"createTime":"DESC"}}}
// {"query":"query PROJECTS($filter: Filter, $orderBy: OrderBy) {\n    projects(filter: $filter, orderBy: $orderBy) {\n      \n      key\n      uuid\n      name\n      status {\n        uuid\n        name\n        category\n      }\n      isPin\n      statusCategory\n      assign {\n        uuid\n        name\n        avatar\n      }\n      planStartTime\n      planEndTime \n      taskUpdateTime\n      sprintCount\n      taskCount\n      taskCountDone\n      taskCountInProgress\n      taskCountToDo\n      memberCount  \n    }\n  }","variables":{"filter":{"visibleInProject_equal":true},"orderBy":{"isPin":"DESC","namePinyin":"ASC"}}}
// err = json.Unmarshal([]byte(reqStr), req)
// if err != nil {
// 	fmt.Println("liuyexing_______________Unmarshal error:", err)
// }
// fmt.Println("liuyexing_______________req.Query:", req.Query)
// fmt.Println("liuyexing_______________req.Variables:", req.Variables)
