package project

import (
	"fmt"
	"testing"

	"github.com/bangwork/ones-ai-api-common/utils/uuid"
	"github.com/yimadai/tools/apitest"
)

func TestAddProject(t *testing.T) {
	projectUUID := apitest.C.UserUUID + uuid.UUID()
	path := fmt.Sprintf("/team/%s/projects/add", apitest.C.TeamUUID)
	body := `
	{
		"project": {
			"uuid": "%s",
			"owner": "%s",
			"name": "test_add_project",
			"status": 1,
			"members": []
		},
		"template_id": "project-t4",
		"members": ["%s"]
	}
	`
	body = fmt.Sprintf(body, projectUUID, apitest.C.UserUUID, apitest.C.UserUUID)
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}

func TestAddComponent(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/components/add", apitest.C.TeamUUID, "JhWrCvGNZEqkMHIZ")
	body := `
	{
		"components": [{
			"uuid": "MLbeCiyg",
			"template_uuid": "com00004",
			"project_uuid": "JhWrCvGNZEqkMHIZ",
			"parent_uuid": "",
			"name": "成员",
			"name_pinyin": "cheng2yuan2",
			"desc": "“成员”组件展示了当前项目中的所有成员及其角色构成。",
			"permissions": [{
				"permission": "view_component",
				"user_domains": [{
					"user_domain_type": "everyone",
					"user_domain_param": ""
				}, {
					"user_domain_type": "project_administrators",
					"user_domain_param": ""
				}]
			}],
			"objects": [],
			"type": 2,
			"views": [],
			"update": 1
		}, {
			"uuid": "EoKNDg7w",
			"template_uuid": "com00009"
		}, {
			"uuid": "ptsn6i8V",
			"template_uuid": "com00007"
		}, {
			"uuid": "gsypKtC2",
			"template_uuid": "com00010"
		}, {
			"uuid": "MxFRgqKv",
			"template_uuid": "com00001"
		}, {
			"uuid": "DmFxfhaP",
			"template_uuid": "com00002"
		}, {
			"uuid": "AV3Nvo9T",
			"template_uuid": "com00003"
		}, {
			"uuid": "7LwPdj2U",
			"template_uuid": "com00008"
		}, {
			"uuid": "3FNW4Vxb",
			"template_uuid": "com00006"
		}, {
			"uuid": "LynpwaJn",
			"template_uuid": "com00005"
		}, {
			"uuid": "F5DfxVAJ",
			"template_uuid": "com00004"
		}]
	}	`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s", ret)
}
