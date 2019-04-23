package bilierrorcode

// 直播 号段 [1000000, 1999999]
// TODO: 检查修复 Description 中的错字错意
// TODO: 补充对应条目的 Description
func getLiveSiteDetail(code ErrorCode) ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}

	switch code {
	// wallet 1000000 - 1001999
	case 1000000:
		result.Message = "CoinNotEnough"
		result.Description = ""
		break
	case 1000001:
		result.Message = "PayFailed"
		result.Description = ""
		break

	// live-test 100 2000 - 100 2999
	case 1002001:
		result.Message = "RoomNotFound"
		result.Description = ""
		break
	case 1002002:
		result.Message = "InvalidParam"
		result.Description = ""
		break
	case 1002003:
		result.Message = "GetAllListRPCError"
		result.Description = ""
		break
	case 1002004:
		result.Message = "GetAllListReturnError"
		result.Description = ""
		break
	case 1002005:
		result.Message = "GetAllListSimpleJSONError"
		result.Description = ""
		break
	case 1002006:
		result.Message = "AttentionRPCError"
		result.Description = ""
		break
	case 1002007:
		result.Message = "AttentionReturnError"
		result.Description = ""
		break
	case 1002008:
		result.Message = "UserTagRPCError"
		result.Description = ""
		break
	case 1002009:
		result.Message = "UserTagReturnError"
		result.Description = ""
		break
	case 1002010:
		result.Message = "UserTagRoomListRPCError"
		result.Description = ""
		break
	case 1002011:
		result.Message = "UserTagRoomListReturnError"
		result.Description = ""
		break
	case 1002012:
		result.Message = "SeaPatrolRPCError"
		result.Description = ""
		break
	case 1002013:
		result.Message = "SeaPatrolReturnError"
		result.Description = ""
		break
	case 1002014:
		result.Message = "ActivityRPCError"
		result.Description = ""
	case 1002015:
		result.Message = "ActivityReturnError"
		result.Description = ""
		break
	case 1002016:
		result.Message = "ChangeGetAllListRPCError"
		result.Description = ""
		break
	case 1002017:
		result.Message = "ChangeGetAllListReturnError"
		result.Description = ""
		break
	case 1002018:
		result.Message = "ChangeGetAllListEmptyError"
		result.Description = ""
		break
	case 1002019:
		result.Message = "SkyHorseError"
		result.Description = ""
		break
	case 1002020:
		result.Message = "ChangeSkyHorseEmptyError"
		result.Description = ""
		break
	case 1002021:
		result.Message = "GetRoomError"
		result.Description = ""
		break
	case 1002024:
		result.Message = "GetRoomEmptyError"
		result.Description = ""
		break
	case 1002022:
		result.Message = "RoomPendantError"
		result.Description = ""
		break
	case 1002023:
		result.Message = "RoomPendantReturnError"
		result.Description = ""
		break
	case 1002100:
		result.Message = "AttentionListRPCError"
		result.Description = ""
		break
	case 1002101:
		result.Message = "RelationFrameWorkCallError"
		result.Description = ""
		break
	case 1002102:
		result.Message = "RelationLiveRPCCodeError"
		result.Description = ""
		break
	case 1002103:
		result.Message = "RelationFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002104:
		result.Message = "RoomGetStatusInfoRPCError"
		result.Description = ""
		break
	case 1002105:
		result.Message = "GetMultipleRPCError"
		result.Description = ""
		break
	case 1002106:
		result.Message = "RoomPendentRPCError"
		result.Description = ""
		break
	case 1002107:
		result.Message = "PKIDRpcError"
		result.Description = ""
		break
	case 1002108:
		result.Message = "UnliveAnchorReqParamsError"
		result.Description = ""
		break
	case 1002109:
		result.Message = "LiveAnchorReqParamsError"
		result.Description = ""
		break
	case 1002110:
		result.Message = "NeedLogIn"
		result.Description = ""
		break
	case 1002111:
		result.Message = "RoomFrameWorkCallError"
		result.Description = ""
		break
	case 1002112:
		result.Message = "RoomLiveRPCCodeError"
		result.Description = ""
		break
	case 1002113:
		result.Message = "RoomFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002114:
		result.Message = "UserFrameWorkCallError"
		result.Description = ""
		break
	case 1002115:
		result.Message = "UserLiveRPCCodeError"
		result.Description = ""
		break
	case 1002116:
		result.Message = "UserFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002117:
		result.Message = "RecordRecordFrameWorkCallError"
		result.Description = ""
		break
	case 1002118:
		result.Message = "RecordLiveRPCCodeError"
		result.Description = ""
		break
	case 1002119:
		result.Message = "RecordFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002120:
		result.Message = "RoomNewsRecordFrameWorkCallError"
		result.Description = ""
		break
	case 1002121:
		result.Message = "RoomNewsLiveRPCCodeError"
		result.Description = ""
		break
	case 1002122:
		result.Message = "RoomNewsFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002123:
		result.Message = "RoomPendentFrameWorkCallError"
		result.Description = ""
		break
	case 1002124:
		result.Message = "RoomPendentLiveRPCCodeError"
		result.Description = ""
		break
	case 1002125:
		result.Message = "RoomPendentFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002126:
		result.Message = "PkIDRecordFrameWorkCallError"
		result.Description = ""
		break
	case 1002127:
		result.Message = "PkIDLiveRPCCodeError"
		result.Description = ""
		break
	case 1002128:
		result.Message = "PkIDFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002129:
		result.Message = "FansMedalFrameWorkCallError"
		result.Description = ""
		break
	case 1002130:
		result.Message = "FansMedalLiveRPCCodeError"
		result.Description = ""
		break
	case 1002131:
		result.Message = "FansMedalFrameWorkGoRoutingError"
		result.Description = ""
		break
	case 1002132:
		result.Message = "GiftFrameWorkCallError"
		result.Description = ""
		break
	case 1002133:
		result.Message = "GiftLiveRPCCodeError"
		result.Description = ""
		break
	case 1002134:
		result.Message = "RoomGetRoomIDCodeRPCError"
		result.Description = ""
		break
	case 1002135:
		result.Message = "LiveAnchorReqParamsNil"
		result.Description = ""
		break
	case 1002136:
		result.Message = "GetGrayRuleError"
		result.Description = ""
		break
	case 1002137:
		result.Message = "AccountGRPCError"
		result.Description = ""
		break
	case 1002138:
		result.Message = "AccountGRPCFrameError"
		result.Description = ""
		break
	case 1002139:
		result.Message = "LiveAnchorReqV2ParamsNil"
		result.Description = ""
		break
	case 1002140:
		result.Message = "LiveAnchorReqV2ParamsError"
		result.Description = ""
		break
	case 1002141:
		result.Message = "UserDHHRPCError"
		result.Description = ""
		break
	case 1002142:
		result.Message = "UserDHHReturnError"
		result.Description = ""
		break
	case 1002143:
		result.Message = "UserDHHDataNil"
		result.Description = ""
		break
	case 1002150:
		result.Message = "XanchorGRPCError"
		result.Description = ""
		break

	// resource 1003000 - 1003199
	case 1003001:
		result.Message = "ResourceParamErr"
		result.Description = "参数传入错误"
		break
	case 1003002:
		result.Message = "TimeForErr"
		result.Description = "时间格式错误"
		break
	case 1003003:
		result.Message = "AddResourceErr"
		result.Description = "添加资源失败"
		break
	case 1003004:
		result.Message = "RepdAddErr"
		result.Description = "添加平台已存在"
		break
	case 1003005:
		result.Message = "SeltResErr"
		result.Description = "资源选择失败"
		break
	case 1003006:
		result.Message = "EditResErr"
		result.Description = "编辑资源失败"
		break
	case 1003007:
		result.Message = "DeviceError"
		result.Description = "参数<device>传入错误"
		break
	case 1003008:
		result.Message = "OfflineResErr"
		result.Description = "下线资源失败"
		break
	case 1003009:
		result.Message = "GetListResErr"
		result.Description = "获取资源列表失败"
		break
	case 1003010:
		result.Message = "GetBannerErr"
		result.Description = "获取Banner配置失败"
		break
	case 1003011:
		result.Message = "GetSplashErr"
		result.Description = "获取闪屏配置失败"
		break
	case 1003012:
		result.Message = "CheckURLErr"
		result.Description = "链接格式错误"
		break
	case 1003101:
		result.Message = "GetConfAdminErr"
		result.Description = "没有获取到配置"
		break
	case 1003102:
		result.Message = "SetConfAdminErr"
		result.Description = "设置配置失败"
		break

	// 弹幕 1003200 - 1003399
	case 1003200:
		result.Message = "DMallUser"
		result.Description = "系统正在维护(全员弹幕禁言)"
		break
	case 1003201:
		result.Message = "DMUserLevel"
		result.Description = "系统正在维护(全员指定等级禁言)"
		break
	case 1003202:
		result.Message = "RealName"
		result.Description = "实名认证才可以发言"
		break
	case 1003203:
		result.Message = "PhoneBind"
		result.Description = "根据国家实名制认证的相关要求，您需要绑定手机号，才能继续进行操作。"
		break
	case 1003204:
		result.Message = "PhoneReal"
		result.Description = "根据国家实名制认证的相关要求，您需要换绑一个非170/171的手机号，才能继续进行操作。"
		break
	case 1003205:
		result.Message = "ShieldUser"
		result.Description = "u (被播主过滤的用户)"
		break
	case 1003206:
		result.Message = "ShieldContent"
		result.Description = "k (被播主过滤的内容)"
		break
	case 1003207:
		result.Message = "BlockUser"
		result.Description = "你在本房间被禁言"
		break
	case 1003208:
		result.Message = "PayLive"
		result.Description = "非常抱歉，本场直播需要购票，即可参与互动(付费直播)"
		break
	case 1003209:
		result.Message = "SecDMLimit"
		result.Description = "msg in 1s(每秒发言限制)"
		break
	case 1003210:
		result.Message = "DMSameMsgLimit"
		result.Description = "msg repeat(消息重复)"
		break
	case 1003211:
		result.Message = "DMLimitPerRoom"
		result.Description = "max limit(单房间每秒限制)"
		break
	case 1003212:
		result.Message = "MsgLengthLimit"
		result.Description = "超出限制长度"
		break
	case 1003213:
		result.Message = "FilterLimit"
		result.Description = "内容非法（敏感词限制）"
		break
	case 1003214:
		result.Message = "CountryLimit"
		result.Description = "你所在的地区暂无法发言(区域限制)"
		break
	case 1003215:
		result.Message = "RoomLeverLimit"
		result.Description = "房间等级限制"
		break
	case 1003216:
		result.Message = "RoomAllLimit"
		result.Description = "房间全员限制"
		break
	case 1003217:
		result.Message = "RoomMedalLimit"
		result.Description = "房间勋章等级限制"
		break
	case 1003218:
		result.Message = "DMServiceERR"
		result.Description = "依赖下游服务失败"
		break

	// user 1004000 - 1004999
	case 1004000:
		result.Message = "UidError"
		result.Description = "Uid错误"
		break
	case 1004001:
		result.Message = "UserNotFound"
		result.Description = "找不到用户信息"
		break

	// dao-anchor  1005000 - 1005999
	case 1005000:
		result.Message = "DaoAnchorCheckAttrSubIdERROR"
		result.Description = "找不到对应的标签子id"
		break

	// user_ex    1006000 - 1006999
	// gift       1007000 - 1007999
	// xuser 1008000 - 1008999
	case 1008001:
		result.Message = "XUserAddUserExpReqBizNotAllow"
		result.Description = ""
		break
	case 1008002:
		result.Message = "XUserAddUserExpTypeNotAllow"
		result.Description = ""
		break
	case 1008003:
		result.Message = "XUserAddUserExpNumNotAllow"
		result.Description = ""
		break
	case 1008004:
		result.Message = "XUserAddUserExpUpdateDBError"
		result.Description = ""
		break
	case 1008005:
		result.Message = "XUserAddUserExpParamsEmptyError"
		result.Description = ""
		break
	case 1008006:
		result.Message = "XUserAddUserExpQueryAfterFail"
		result.Description = ""
		break
	case 1008007:
		result.Message = "XUserAddAnchorUpdateDBError"
		result.Description = ""
		break
	case 1008008:
		result.Message = "XUserAddRExpUpdateDBError"
		result.Description = ""
		break
	case 1008009:
		result.Message = "XUserAddUserRExpQueryAfterFail"
		result.Description = ""
		break
	case 1008020:
		result.Message = "XUserAddRoomAdminOverLimitError"
		result.Description = ""
		break
	case 1008021:
		result.Message = "XUserAddRoomAdminIsAdminError"
		result.Description = ""
		break
	case 1008022:
		result.Message = "XUserAddRoomAdminIsSilentError"
		result.Description = ""
		break
	case 1008023:
		result.Message = "XUserAddRoomAdminNotAdminError"
		result.Description = ""
		break

	case 1008024:
		result.Message = "XUserExpGetExpMcFail"
		result.Description = "获取用户经验缓存失败"
		break
	case 1008025:
		result.Message = "XUserExpGetExpDBFail"
		result.Description = "获取用户经验缓回源db失败"
		break
	case 1008026:
		result.Message = "XUserAddUExpGetParamsNil"
		result.Description = "添加用户经验缓入参为空"
		break

	case 1008027:
		result.Message = "XUserGuardFetchRecentTopListFail"
		result.Description = "获取主播最近总督失败"
		break

	// capusle 1009000 - 1008999
	case 1009001:
		result.Message = "XLotteryCapsuleAreaParamErr"
		result.Description = ""
		break
	case 1009002:
		result.Message = "XLotteryCapsuleSystemErr"
		result.Description = ""
		break
	case 1009003:
		result.Message = "XLotteryCapsuleCoinNotEnough"
		result.Description = ""
		break
	case 1009004:
		result.Message = "XLotteryCapsuleCoinNotChange"
		result.Description = ""
		break
	case 1009005:
		result.Message = "XLotteryCapsulePoolNotChange"
		result.Description = ""
		break
	case 1009006:
		result.Message = "XLotteryCapsulePoolNotOffline"
		result.Description = ""
		break
	case 1009007:
		result.Message = "XLotteryCapsuleOperationFrequent"
		result.Description = ""
		break

	// app-conf 1010000 - 1010100
	case 1010000:
		result.Message = "AppConfKeyErr"
		result.Description = ""
		break

	// 验证码      1990000 - 1990100
	case 1990000:
		result.Message = "VerifyNeed"
		result.Description = ""
		break
	case 1990001:
		result.Message = "VerifyErr"
		result.Description = ""
		break

	// 网关层错误码 1100000 - 1101999
	case 1100000:
		result.Message = "FILTERNOTPASS"
		result.Description = "屏蔽词校验失败"
		break
	case 1100010:
		result.Message = "UploadTokenGenErr"
		result.Description = "Upload token 创建失败"
		break
	case 1100011:
		result.Message = "UploadUploadErr"
		result.Description = "上传失败"
		break
	case 1100012:
		result.Message = "UploadBucketErr"
		result.Description = "上传bucket错误"
		break

	// 调用服务出错 19001000 - 19001999
	case 19001000:
		result.Message = "CallRoomError"
		result.Description = ""
		break
	case 19001001:
		result.Message = "CallUserError"
		result.Description = ""
		break
	case 19001002:
		result.Message = "CallRelationError"
		result.Description = ""
		break
	case 19001003:
		result.Message = "CallFansMedalError"
		result.Description = ""
		break
	case 19001004:
		result.Message = "CallMainMemberError"
		result.Description = ""
		break
	case 19001005:
		result.Message = "CallResourceError"
		result.Description = ""
		break
	case 19001006:
		result.Message = "CallMainFilterError"
		result.Description = ""
		break
	case 19001007:
		result.Message = "CallDaoAnchorError"
		result.Description = ""
		break
	}
	return result
}
