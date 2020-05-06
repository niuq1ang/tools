package temp_test

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func Test001(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/JhWrCvGNkANipdTc/components/add", apitest.C.TeamUUID)
	body := `
	{"components":[{"uuid":"7wSB65kV","template_uuid":"com00012","project_uuid":"JhWrCvGNkANipdTc","parent_uuid":"","name":"自定义链接","name_pinyin":"zi4ding4yi4lian4jie1","desc":"“自定义链接”组件可以通过配置URL，点击后跳转到对应网页。","permissions":[{"permission":"view_component","user_domains":[{"user_domain_type":"everyone","user_domain_param":""},{"user_domain_type":"project_administrators","user_domain_param":""}]}],"objects":[],"type":5,"views":[],"url_setting":{"url":"http://www.baidu.com"},"update":1},{"uuid":"4qZV6FpK","template_uuid":"com00009"},{"uuid":"4bBiDt8k","template_uuid":"com00014"},{"uuid":"TH9cB4cN","template_uuid":"com00015"},{"uuid":"3TjVhHrL","template_uuid":"com00006"},{"uuid":"TeDgkC8q","template_uuid":"com00005"},{"uuid":"M1d7XhLc","template_uuid":"com00004"}]}	`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
func Test002(t *testing.T) {
	path := fmt.Sprintf("/team/%s/items/add", apitest.C.TeamUUID)
	body := `
{"item":{"name":"test3","item_type":"gantt_data","gantt_data_type":"task","gantt_chart_uuid":"7JJfRn2A","plan_start_time":1589904600,"plan_end_time":1590076799,"parent":"","assign":"GxBCYd6q"}}
`
	err, ret := apitest.DoPostRequest(path, body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
