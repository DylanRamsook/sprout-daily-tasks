package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/DylanRamsook/sprout-daily-tasks/sprout_api/activity_logs"
	"net/http"
	"strings"
)

//Ex: Body
//{"activities":[{"activityUnitId":24988,"quantity":"90","date":"2020-05-28"}],"text":"Pull Day","image":"","shareWithFilterId":"18"}
func CreateSproutActivity(sproutUrl string, date string, quantity string, cookies []*http.Cookie, rememberMe string) (*http.Response, error) {
	url := sproutUrl + "/v1/activity_logs"
	method := "POST"

	//Create a payload
	p := new(activity_logs.CreateActivityRequest)
	var activitySlice = []activity_logs.Activity{
		activity_logs.Activity{24988, date, quantity},
	}
	p.Activities = activitySlice
	p.Image = ""
	p.ShareWithFilterID = "18"
	p.Text = "Daily Sprout Submission"

	q, err := json.Marshal(p)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//payload := strings.NewReader("{\"activities\":[{\"activityUnitId\":24988,\"quantity\":\"90\",\"date\":\"2020-05-28\"}],\"text\":\"Pull Day\",\"image\":\"\",\"shareWithFilterId\":\"18\"}")
	payload := strings.NewReader(string(q))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	for _, s := range cookies {
		req.AddCookie(s)
	}

	//req.Header.Add("Cookie", "sprout_session=qg461esdksqvnejepnfoj817kpfeioqv; remember_me=1; token=7ef8b5187433d77212587b9c4a370c9b6682fbd7")

	res, err := client.Do(req)
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return res, err
}
