package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"imlogic/internal/model/entity"
	"imlogic/internal/utils/esutil"
)

type esMessageDao struct {
	indexName string
	script    *scriptDao
}

func NewESMessageDao() *esMessageDao {
	return &esMessageDao{indexName: "chat-message-*", script: &scriptDao{}}
}

func (dao *esMessageDao) GetChatMsg(ctx context.Context, args *entity.GetChatMsgArgs) ([]*entity.MessageEntity, error) {
	var script string
	if args.Pointer == 0 {
		script = dao.script.LatestChatMsg(args.ChatID)
	} else {
		script = dao.script.OldChatMsg(args.ChatID, args.Pointer)
	}

	resBytes, err := esutil.QueryRequest(script, dao.indexName)
	if err != nil {
		return nil, err
	}

	var k []*entity.ChatMessageListEntity
	obj := &entity.ChatSearchEntity{
		Hits: &entity.ChatSearchHitsEntity{
			Hits: k,
		},
	}

	err = json.Unmarshal(resBytes, obj)
	if err != nil {
		return nil, err
	}

	count := len(obj.Hits.Hits)
	values := make([]*entity.MessageEntity, count)
	for i, item := range obj.Hits.Hits {
		values[i] = item.Source
	}

	return values, nil
}

type scriptDao struct{}

func (sd *scriptDao) LatestChatMsg(chatID int64) string {
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

func (sd *scriptDao) OldChatMsg(chatID int64, pointer int64) string {
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
