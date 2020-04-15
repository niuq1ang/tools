package apitest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DoPostRequest(path string, body interface{}) (error, string) {
	var b []byte
	var err error
	switch t := body.(type) {
	case string:
		b = []byte(t)
	default:
		if b, err = json.Marshal(t); err != nil {
			return err, ""
		}
	}

	url := fmt.Sprintf("http://%s:%d%s", Host, Port, path)
	fmt.Printf("url: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return err, ""
	}

	return doRequest(req)
}

func DoGetRequest(path string) (error, string) {
	url := fmt.Sprintf("http://%s:%d%s", Host, Port, path)
	fmt.Printf("url: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, ""
	}
	return doRequest(req)
}

func doRequest(req *http.Request) (error, string) {
	req.Header.Set("Ones-User-Id", C.UserUUID)
	req.Header.Set("Ones-Auth-Token", C.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err, ""
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	var out bytes.Buffer
	if err = json.Indent(&out, b, "", "  "); err != nil {
		return err, ""
	}
	return nil, string(out.Bytes())
}

func BuildGraphqlQuery(action string, content string) string {
	var mutation string
	if action == "mutation" {
		mutation = "mutation "
	}

	content = strings.Replace(content, "\n", "\\n", -1)
	content = strings.Replace(content, "\t", "", -1)
	content = strings.Replace(content, "\"", "\\\"", -1)

	body := `
		{
			"query": "%s%s"
		}
		`
	body = fmt.Sprintf(body, mutation, content)
	fmt.Println(fmt.Sprintf("GQL body:%s", body))
	return body
}
