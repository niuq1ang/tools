package activity

import (
	"fmt"
	"testing"

	"github.com/yimadai/tools/apitest"
)

func TestGetActivityChartLock(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/activity_chart/%s/locker", apitest.C.TeamUUID, "D2pLSaJ3S1Ji55AX", "TGWFCDHR")
	err, ret := apitest.DoGetRequest(path)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}

func TestLockActivityChart(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/activity_chart/%s/lock", apitest.C.TeamUUID, "D2pLSaJ3S1Ji55AX", "TGWFCDHR")
	err, ret := apitest.DoPostRequest(path, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}

func TestUnLockActivityChart(t *testing.T) {
	path := fmt.Sprintf("/team/%s/project/%s/activity_chart/%s/unlock", apitest.C.TeamUUID, "D2pLSaJ3S1Ji55AX", "TGWFCDHR")
	err, ret := apitest.DoPostRequest(path, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("response: %s\n", ret)
}
