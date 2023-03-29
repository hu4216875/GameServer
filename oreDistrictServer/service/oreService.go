package service

import (
	"github.com/name5566/leaf/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"oreDistrictServer/common"
	"oreDistrictServer/dao"
	"oreDistrictServer/grpc-base/protos"
	"oreDistrictServer/model"
	"oreDistrictServer/template"
	"time"
)

var (
	oreDistircts []*model.OreDistrict
)

type oreService struct {
	protos.UnimplementedOreDistrictServiceServer
}

func (o *oreService) GetOreInfo(ctx context.Context, req *protos.RequestOreInfo) (*protos.ResponseOreInfo, error) {
	ore := getOreDistrict(req.OreId)
	if ore == nil {
		return &protos.ResponseOreInfo{Result: 1}, nil
	}
	return &protos.ResponseOreInfo{OreId: req.OreId, Total: ore.Total, EndTime: ore.EndTime}, nil
}

func (o *oreService) GetOreTotal(ctx context.Context, req *protos.RequestOreTotal) (*protos.ResponseOreTotal, error) {
	ore := getOreDistrict(req.OreId)
	if ore == nil {
		return &protos.ResponseOreTotal{Result: 1}, nil
	}

	total := calcCurOreDistrictTotal(req.OreId)
	return &protos.ResponseOreTotal{OreId: req.OreId, Total: total, EndTime: ore.EndTime}, nil
}

func (o *oreService) AddOrePlayer(ctx context.Context, req *protos.RequestAddOrePlayer) (*protos.ResponseAddOrePlayer, error) {
	ore := getOreDistrict(req.OreId)
	if ore == nil {
		return &protos.ResponseAddOrePlayer{Result: uint32(common.OreErr_ORE_NOT_EXIST)}, nil
	}

	// 没有资源了
	if !canEnterOreDistrict(req.OreId) {
		return &protos.ResponseAddOrePlayer{Result: uint32(common.OreErr_NO_RESOURCE)}, nil
	}

	if existOrePlayer(ore, req.AccountId) {
		return &protos.ResponseAddOrePlayer{Result: uint32(common.OreErr_PLAYER_EXIST)}, nil
	}

	addPlayer(req.OreId, req.AccountId, req.ServerId, req.Speed)
	return &protos.ResponseAddOrePlayer{Result: uint32(common.OreErr_None), AccountId: req.AccountId, EndTime: ore.EndTime, Total: ore.Total}, nil
}

func (o *oreService) SettlePlayer(ctx context.Context, req *protos.RequestSettleOrePlayer) (*protos.ResponseSettleOrePlayer, error) {
	ore := getOreDistrict(req.OreId)
	if ore == nil {
		return &protos.ResponseSettleOrePlayer{Result: uint32(common.OreErr_ORE_NOT_EXIST)}, nil
	}

	if !existOrePlayer(ore, req.AccountId) {
		return &protos.ResponseSettleOrePlayer{Result: uint32(common.OreErr_PLYAER_NOT_EXIST)}, nil
	}
	num := settlePlayer(req.OreId, req.AccountId)
	return &protos.ResponseSettleOrePlayer{
		Result:    uint32(common.OreErr_None),
		AccountId: req.AccountId,
		Total:     ore.Total,
		EndTime:   ore.EndTime,
		Num:       num,
	}, nil
}

func (o *oreService) UpdateOrePlayer(ctx context.Context, req *protos.RequestUpdateOrePlayer) (*protos.ResponseUpdateOrePlayer, error) {
	ore := getOreDistrict(req.OreId)
	if ore == nil {
		return &protos.ResponseUpdateOrePlayer{Result: uint32(common.OreErr_ORE_NOT_EXIST)}, nil
	}

	if !existOrePlayer(ore, req.AccountId) {
		return &protos.ResponseUpdateOrePlayer{Result: uint32(common.OreErr_PLYAER_NOT_EXIST)}, nil
	}

	num := updatePlayer(req.OreId, req.AccountId, req.Speed)
	return &protos.ResponseUpdateOrePlayer{AccountId: req.AccountId, Total: ore.Total, Num: num, EndTime: ore.EndTime}, nil
}

func RegistOreService(server *grpc.Server) {
	protos.RegisterOreDistrictServiceServer(server, new(oreService))
	initOreService()
}

// initOreService 初始化矿洞
func initOreService() {
	oreDistircts = dao.OreDistrictDao.LoadOreDistrict()
	// 获取所有的矿洞配置没有的则需要加上
	ores := template.GetOreTempalte().GetAll()
	for i := 0; i < len(ores); i++ {
		if getOreDistrict(ores[i].Id) == nil {
			data := model.NewOreDistirct(ores[i].Id, ores[i].Total)
			dao.OreDistrictDao.AddOreDistrict(data)
			oreDistircts = append(oreDistircts, data)
		}
	}
	registService(int64(common.MAX_SERVICE_TIME))
}

// addPlayer 添加玩家
func addPlayer(oreId uint32, accountId int64, serverId, speed uint32) {
	player := model.NewOreDistirctPlayer(accountId, serverId, speed, uint32(time.Now().Unix()))
	oreDistrict := addOreDistrictPlayer(oreId, player)
	dao.OreDistrictDao.AddOreDistrictPlayer(oreDistrict, player)
}

func existOrePlayer(data *model.OreDistrict, accountId int64) bool {
	for i := 0; i < len(data.Players); i++ {
		if data.Players[i].AccountId == accountId {
			return true
		}
	}
	return false
}

// settlePlayer 结算玩家(离线，离开矿区)
func settlePlayer(oreId uint32, accountId int64) uint32 {
	ore := getOreDistrict(oreId)
	if ore == nil {
		return 0
	}

	var player *model.OreDistrictPlayer = nil
	for i := 0; i < len(ore.Players); i++ {
		if ore.Players[i].AccountId == accountId {
			player = ore.Players[i]
			ore.Players = append(ore.Players[:i], ore.Players[i+1:]...)
			break
		}
	}

	if player == nil {
		return 0
	}

	curTime := uint32(time.Now().Unix())
	num := player.Speed * (curTime - player.StartTime)
	if num > ore.Total {
		num = ore.Total
		ore.Total = 0
	} else {
		ore.EndTime = calcOreEndTime(ore)
		ore.Total -= num
	}

	dao.OreDistrictDao.RemoveOreDistrictPlayer(ore, player)
	if num > 0 {
		dao.OreDistrictDao.UpdateOreRecord(ore.OreDistId, player.AccountId, uint32(num))
	}
	return num
}

// updatePlayer 修改玩家
func updatePlayer(oreId uint32, accountId int64, newSpeed uint32) uint32 {
	ore := getOreDistrict(oreId)
	if ore == nil {
		log.Error("changePlayerSpeed oreId:%v, accountId:%v new Speed:%v ore null", oreId, accountId, newSpeed)
		return 0
	}

	var player *model.OreDistrictPlayer = nil
	for i := 0; i < len(ore.Players); i++ {
		if ore.Players[i].AccountId == accountId {
			player = ore.Players[i]
			break
		}
	}

	if player == nil {
		return 0
	}

	curTime := uint32(time.Now().Unix())
	num := player.Speed * (curTime - player.StartTime)
	if num > ore.Total {
		num = ore.Total
		ore.Total = 0
	} else {
		ore.Total -= num
	}
	ore.EndTime = calcOreEndTime(ore)
	player.Speed = newSpeed
	player.StartTime = uint32(time.Now().Unix())
	dao.OreDistrictDao.UpdateOreDistrictPlayer(ore, player)
	dao.OreDistrictDao.UpdateOreRecord(ore.OreDistId, player.AccountId, uint32(num))
	return num
}

func getOreDistrict(oreId uint32) *model.OreDistrict {
	for i := 0; i < len(oreDistircts); i++ {
		if oreDistircts[i].OreDistId == oreId {
			return oreDistircts[i]
		}
	}
	return nil
}

func getOreDistrictTotal(oreId uint32) uint32 {
	data := getOreDistrict(oreId)
	if data == nil {
		return 0
	}
	return data.Total
}

func calcCurOreDistrictTotal(oreId uint32) uint32 {
	data := getOreDistrict(oreId)
	if data == nil {
		return 0
	}
	curTime := uint32(time.Now().Unix())

	var temp uint32 = 0
	for i := 0; i < len(data.Players); i++ {
		temp += (curTime - data.Players[i].StartTime) * data.Players[i].Speed
	}
	if data.Total >= temp {
		data.Total = data.Total - temp
		return data.Total
	}
	return 0
}

func canEnterOreDistrict(oreId uint32) bool {
	if calcCurOreDistrictTotal(oreId) <= 0 {
		return false
	}
	return true
}

// addOreDistrictPlayer 增加挖矿的人
func addOreDistrictPlayer(oreId uint32, player *model.OreDistrictPlayer) *model.OreDistrict {
	ore := getOreDistrict(oreId)
	if ore == nil {
		return nil
	}
	ore.Players = append(ore.Players, player)
	// 重新计算总时间

	var total int64
	var totalSpeed int64 = 0

	total = int64(ore.Total)
	for i := 0; i < len(ore.Players); i++ {
		total += int64(ore.Players[i].Speed) * int64(ore.Players[i].StartTime)
		totalSpeed += int64(ore.Players[i].Speed)
	}
	ore.EndTime = uint32(total / totalSpeed)
	return ore
}

func calcOreEndTime(ore *model.OreDistrict) uint32 {
	var total int64
	var totalSpeed int64
	total = int64(ore.Total)
	for i := 0; i < len(ore.Players); i++ {
		total += int64(ore.Players[i].Speed) * int64(ore.Players[i].StartTime)
		totalSpeed += int64(ore.Players[i].Speed)
	}

	var endTime uint32
	if totalSpeed > 0 {
		endTime = uint32(total / totalSpeed)
	} else {
		endTime = uint32(common.UINT32_MAX)
	}
	return endTime
}
