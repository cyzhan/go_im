package esutil

import (
	"fmt"
	"strings"
)

var script = &queryScript{}

type queryScript struct {
}

func (q *queryScript) LatestChatMsg(chatID int64) string {
	s := `
	{
		"sort":[{"msgID":"asc"}],
		"query": {
			"bool":{
				"must":[
					{
						"term":{
							"chatID":%d
						}
					}
				]
			}
		},
		"from":0,
		"size":100
	}`
	return fmt.Sprintf(s, chatID)
}

func (q *queryScript) OldChatMsg(chatID int64, pointer int64) string {
	s := `
	{
		"sort":[{"msgID":"asc"}],
		"query": {
			"bool":{
				"must":[
					{
						"term":{
							"chatID":%d
						}
					}
				],
				"filter":{
					"range":{
						"msgID":{
							"lt":%d
						}
					}
				}
			}
		},
		"from":0,
		"size":100
	}`
	return fmt.Sprintf(s, chatID, pointer)
}

func (q *queryScript) fetchTextingMsg(msgId []int64) string {
	s := `
	{
		"sort":[{"msgID":"asc"}],
		"query": {
			"bool":{
				"must":[
					{
						"terms":{
							"msgID":%s
						}
					}
				]
			}
		},
		"from":0,
		"size":100
	}`

	idStr := strings.Replace(fmt.Sprint(msgId), " ", ",", -1)
	fmt.Println(fmt.Sprintf(s, idStr))
	return fmt.Sprintf(s, idStr)
}
