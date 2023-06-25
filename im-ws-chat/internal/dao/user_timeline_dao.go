package dao

import (
	"context"
	"database/sql"

	"imws/internal/model/entity"
	"imws/internal/model/im"
)

type userTimelineDao struct{}

func (dao *userTimelineDao) GetSelf(ctx context.Context, cond *entity.GetSelfTimelineCond) []*entity.UserTimeline {
	if cond.UserId == 0 {
		return nil
	}

	sqlStr := "SELECT msg_id, user_id, group_id FROM im.user_timeline WHERE user_id = ? "
	args := []interface{}{cond.UserId}

	if cond.GroupId != 0 {
		sqlStr += " AND group_id =? "
		args = append(args, cond.GroupId)
	}

	if cond.Pointer != 0 {
		sqlStr += " AND msg_id > ? "
		args = append(args, cond.Pointer)
		sqlStr += " ORDER BY msg_id DESC"
	} else {
		sqlStr += "ORDER BY msg_id DESC LIMIT 1000"
	}

	rows, err := maria.QueryContext(ctx, sqlStr, args...)
	if err == sql.ErrNoRows {
		return make([]*entity.UserTimeline, 0)
	}
	panicIfError(err)

	defer rows.Close()
	var list []*entity.UserTimeline
	for rows.Next() {
		timeline := &entity.UserTimeline{}
		if err := rows.Scan(&timeline.MsgId, &timeline.UserId, &timeline.GroupId); err != nil {
			panicIfError(err)
		}
		list = append(list, timeline)
	}

	return list
}

func (dao *userTimelineDao) GetUserGroupPointer(ctx context.Context, userID int64) []*im.ReadPointer {
	script := `SELECT ugp.group_id, IFNULL(ugp.current,0) AS current
	FROM im.user_group_pointer ugp 
	WHERE ugp.user_id = ?`

	rows, err := maria.QueryContext(ctx, script, userID)
	if err == sql.ErrNoRows {
		return make([]*im.ReadPointer, 0)
	}
	panicIfError(err)

	var list = []*im.ReadPointer{}
	for rows.Next() {
		rp := &im.ReadPointer{}
		err := rows.Scan(&rp.GroupID, &rp.Pointer)
		panicIfError(err)
		list = append(list, rp)
	}
	return list
}
