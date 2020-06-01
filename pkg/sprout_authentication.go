package pkg

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func LoginToSprout(email string, password string, sproutUrl string, rememberMe string) (*http.Response, error) {

	//Set up request
	var endpoint = "/auth/login_post"
	url := sproutUrl + endpoint
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("email", email)
	_ = writer.WriteField("password", password)
	_ = writer.WriteField("remember_me", rememberMe)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "remember_me=1")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//Return the response body
	return res, err

}
