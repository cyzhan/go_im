package dao

import (
	"context"
	"database/sql"

	"imws/internal/model/vo"
)

type groupDao struct {
}

var (
	Group = &groupDao{}
)

// todo cache by userId & clear cache when modify
func (g *groupDao) GetGroups(ctx context.Context, userId int64) []*vo.Group {
	s := `
	SELECT gm.group_id, g.name, g.icon_url, g.member_count, g.deleted
	FROM im.group_member AS gm
	INNER JOIN im.group AS g ON gm.group_id = g.id 
	WHERE gm.user_id = ?
	`

	rows, err := maria.QueryContext(ctx, s, userId)
	if err == sql.ErrNoRows {
		return make([]*vo.Group, 0)
	}
	panicIfError(err)

	var list = []*vo.Group{}
	for rows.Next() {
		group := &vo.Group{}
		err := rows.Scan(&group.ID, &group.Name, &group.IconUrl, &group.MemberCount, &group.Deleted)
		panicIfError(err)
		list = append(list, group)
	}
	return list
}
