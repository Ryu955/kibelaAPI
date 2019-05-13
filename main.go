package main

import (
	"bytes"
	"fmt"
	"github.com/Ryu955/kibela/env"
	"io/ioutil"
	"net/http"
)

func main() {
	//jsonStr :=`{"query":"query {currentUser { account }}"}`
	//query := `query {notes (first: 1) {edges {cursor node {id title content folderName groups {name}}}totalCount}}`
	//jsonStr :=`{"query":"query {notes (first: 1) {edges {cursor node {id title content folderName groups {name}}}totalCount}}","variables":""}`
	query := `mutation($title: String!, $content: String!){createNote(input: { title: $title, content: $content, coediting: false, groupIds: [\"R3JvdXAvMw\"] }) {clientMutationId}}`

	variables := `{ "title": "スクリプトから", "content": "huga" }`
	jsonStr :=`{"query":"`+query +`","variables":`+variables+`}`
	fmt.Println(jsonStr)
	requestUrl := "https://ryu955.kibe.la/api/v1"

	req, err := http.NewRequest(
		"POST",
		requestUrl,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		fmt.Println(err)
	}

	// Content-Type 設定
	req.Header.Set("Authorization","Bearer "+env.APIKey)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}

