package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"imlogic/internal/constant"
	"imlogic/internal/constant/exchange"
	"imlogic/internal/constant/topic"
	"imlogic/internal/dao"
	"imlogic/internal/model/entity"
	"imlogic/internal/model/im"
	"imlogic/internal/rpc/common"
	"imlogic/internal/rpc/user"
	utilsCommon "imlogic/internal/utils/common"
	"imlogic/internal/utils/encrypt"
	"imlogic/internal/utils/imredis"
	"imlogic/internal/utils/mysqlcli"
	"imlogic/internal/utils/rabbit"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewUserService() *userService {
	return &userService{
		WsUrl:     os.Getenv("WS_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}

type userService struct {
	WsUrl     string
	JwtSecret string
}

func (us *userService) Login(ctx context.Context, data *user.LoginData) (*user.LoginResult, error) {
	result := &user.LoginResult{}
	var userID int64
	var status int32

	//Todo add cache and notify mechanism
	langID := dao.UserDao.GetLanguageID(ctx, data.VdID, data.Language)

	u := dao.UserDao.GetUserByName(data.VdID, data.Name)
	log.Printf("%#v\n", u)

	if u == nil {
		log.Printf("%#v not found in db, do new user\n", data)
		userID = dao.UserDao.NewUser(data.VdID, data.Name, encrypt.HashUserPW(data.Password))
		status = 1
	} else if encrypt.HashUserPW(data.Password) != u.Password {
		log.Printf("pw in db: %s, pw from client: %s", u.Password, encrypt.HashUserPW(data.Password))
		log.Printf("%#v password not equal\n", data)
		result.Error = &common.Error{
			Code: 2,
			Msg:  "invalid account or password",
		}
		return result, nil
	} else {
		userID = u.ID
		status = u.Status
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userID,
		"name": data.Name,
		"vdID": data.VdID,
		"ts":   time.Now().UnixMilli(),
	})

	tokenString, err := token.SignedString([]byte(us.JwtSecret))
	signature := strings.Split(tokenString, ".")[2]
	// panicIfError(err)
	if err != nil {
		log.Printf(err.Error())
		result.Error = &common.Error{
			Code: 2,
			Msg:  "invalid account or password",
		}
		return result, nil
	}

	imredis.Client.Set(ctx, signature, signature, 12*time.Hour)

	jsonStr, _ := json.Marshal(&map[string]any{
		"id":     userID,
		"name":   data.Name,
		"vdID":   data.VdID,
		"vip":    data.Vip,
		"status": status,
		"avatar": data.Avatar,
		"langID": langID,
	})

	var uuidStr = uuid.NewString()
	err = imredis.Client.Set(ctx, uuidStr, jsonStr, 180*time.Second).Err()
	if err != nil {
		log.Printf(err.Error())
	}

	result.UserID = userID
	result.Link = us.WsUrl + uuidStr
	return result, nil
}

func (us *userService) GetUserList(ctx context.Context, args *user.GetUserListArgs) (result *user.GetUserListResult, err error) {
	list, totalCount := dao.UserDao.GetUsers(ctx, args)
	result = &user.GetUserListResult{
		TotalCount: totalCount,
		Items:      list,
	}
	return result, nil
}

func (us *userService) PatchUser(ctx context.Context, args *user.PatchUserArgs) (commonErr *common.Error, err error) {
	commonErr = &common.Error{}

	db := mysqlcli.GetClient().Session()
	patchUserReq := &entity.PatchUserArgs{
		Remark:   args.Remark,
		Status:   constant.UserStatusEnum(args.GetStatus()).Check(),
		UserId:   args.Id,
		VendorId: args.VendorID,
	}
	// record date when block
	if constant.UserStatusEnum(args.GetStatus()) == constant.UserStatusEnum_BlackList {
		patchUserReq.DateStr = utilsCommon.UtcMinus4()
	}

	err = dao.UserDao.PatchUser(ctx, db, patchUserReq)
	if err != nil {
		commonErr.Code = 500
		commonErr.Msg = "Error on dao.user.PatchUser. " + err.Error()
		return commonErr, err
	}

	data := user.UserStatusChange{
		UserID: args.Id,
		Status: args.Status,
	}

	any, err := anypb.New(&data)
	if err != nil {
		log.Println(err.Error())
	}
	//Todo move this message type to common
	a := &im.Notify{
		Command: *im.Command_USER_STATUS.Enum(),
		Data:    any,
	}
	rabbit.PbConvertAndSendMQ(a, exchange.AMQ_TOPIC, topic.WS_GENERIC)
	return commonErr, nil
}

func (us *userService) ReportAbuse(ctx context.Context, cond *user.ReportAbuseReq) (*common.Error, error) {
	commonErr := &common.Error{}
	db := mysqlcli.GetClient().Session()

	u, err := dao.UserDao.GetUser(ctx, db, &entity.GetUserArgs{UserId: cond.GetUserID()})
	if err != nil {
		commonErr.Code = 500
		commonErr.Msg = "Error on dao.UserDao.GetUser " + err.Error()
		return commonErr, nil
	}

	tx := db.Begin()
	defer tx.Rollback()

	if err := dao.AbuseRecordDao.InsertAbuseRecord(ctx, tx, &entity.AbuseRecordEntity{
		UserId:      cond.GetUserID(),
		VendorId:    u.VendorID,
		Reason:      cond.GetReason(),
		ReportedBy:  cond.GetReportId(),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}); err != nil {
		commonErr.Code = 500
		commonErr.Msg = "Error on dao.AbuseRecordDao.InsertAbuseRecord " + err.Error()
		return commonErr, nil
	}

	vendorMap, err := dao.VendorDao.GetVdIDMap()
	if err != nil {
		commonErr.Code = 500
		commonErr.Msg = "Error on dao.UserDao.GetUser. " + err.Error()
		return commonErr, nil
	}

	vd, ok := vendorMap[u.VendorID]
	if !ok {
		commonErr.Code = 400
		commonErr.Msg = "vendor no exist"
		return commonErr, nil
	}

	var ban bool
	if u.Report > vd.ReportLock && u.Status == int32(constant.UserStatusEnum_Active) {
		log.Println("lock")
		patchUserReq := &entity.PatchUserArgs{
			Status:   constant.UserStatusEnum_BlackList,
			UserId:   u.ID,
			VendorId: u.VendorID,
			Report:   u.Report + 1,
			DateStr:  utilsCommon.UtcMinus4(),
		}

		if err := dao.UserDao.PatchUser(ctx, tx, patchUserReq); err != nil {
			commonErr.Code = 500
			commonErr.Msg = "Error on dao.UserDao.PatchUser " + err.Error()
			return commonErr, nil
		}

		ban = true
	} else {
		patchUserReq := &entity.PatchUserArgs{
			UserId:   u.ID,
			VendorId: u.VendorID,
			Report:   u.Report + 1,
			DateStr:  utilsCommon.UtcMinus4(),
		}

		if err := dao.UserDao.PatchUser(ctx, tx, patchUserReq); err != nil {
			commonErr.Code = 500
			commonErr.Msg = "Error on dao.UserDao.PatchUser " + err.Error()
			return commonErr, nil
		}
	}

	if err := tx.Commit().Error; err != nil {
		commonErr.Code = 500
		commonErr.Msg = "tx.Commit() " + err.Error()
		return commonErr, nil
	}

	if ban {
		data, _ := anypb.New(&im.UserStatus{
			Name:   u.Name,
			VdID:   u.VendorID,
			Status: int32(constant.UserStatusEnum_BlackList),
		})
		notify := &im.Notify{
			Command: im.Command_USER_STATUS,
			Data:    data,
		}
		rabbit.PbConvertAndSendMQ(notify, exchange.AMQ_TOPIC, topic.WS_GENERIC)
	}

	return &common.Error{}, nil
}
