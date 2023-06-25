package dao

import (
	"database/sql"

	"imws/internal/model/entity"
	"imws/internal/util/datetime"
)

type userDao struct {
}

func (ud *userDao) GetUser(optTx *sql.Tx, userID int64) *entity.User {
	s := `
	SELECT u.id, u.vendor_id, u.name, u.status, u.report, u.date
    FROM im.user u 
    WHERE u.id = ? FOR UPDATE 
	`
	u := &entity.User{}
	var err error
	if optTx != nil {
		err = optTx.QueryRow(s, userID).Scan(&u.ID, &u.VdID, &u.Name, &u.Status, &u.Report, &u.Date)
	} else {
		err = maria.QueryRow(s, userID).Scan(&u.ID, &u.VdID, &u.Name, &u.Status, &u.Report, &u.Date)
	}

	panicIfError(err)
	return u
}

func (ud *userDao) IncrReportAndBan(optTx *sql.Tx, userID int64) {
	s := `UPDATE im.user SET report = report + 1, status = 2, date = ? WHERE id = ?`
	date := datetime.EestAmerica()
	exec(optTx, s, date, userID)
}

func (ud *userDao) IncrReport(optTx *sql.Tx, userID int64) {
	s := `UPDATE im.user SET report = report + 1 WHERE id = ?`
	exec(optTx, s, userID)
}

func (ud *userDao) InsertAbuseRecord(optTx *sql.Tx, record *entity.AbuseRecord) {
	s := `
	INSERT INTO im.abuse_record
    (user_id, vendor_id, reason, reported_by, created_time, updated_time)
    SELECT a.id, a.vendor_id, ?, ?, current_timestamp(), current_timestamp() 
    FROM im.user AS a
    WHERE a.id = ?
	`
	exec(optTx, s, record.Reason, record.ReportBy, record.UserID)
}
