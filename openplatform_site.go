package bilierrorcode

// 开放平台 [2000000, 2999999]
// 票务的 code码 [2000000, 2099999]
func getOpenPlatformSiteDetail(code ErrorCode) ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}

	switch code {
	// 销售-营销
	case 2000000:
		result.Message = "TicketUnKnown"
		result.Description = "未知错误"
		break
	case 2000001:
		result.Message = "TicketParamInvalid"
		result.Description = "参数错误"
		break
	case 2000002:
		result.Message = "TicketRecordDupli"
		result.Description = "重复插入"
		break
	case 2000003:
		result.Message = "TicketRecordLost"
		result.Description = "数据不存在"
		break
	case 2000004:
		result.Message = "TicketPromotionLost"
		result.Description = "活动不存在"
		break
	case 2000005:
		result.Message = "TicketPromotionEnd"
		result.Description = "活动结束"
		break
	case 2000006:
		result.Message = "TicketPromotionRepeatJoin"
		result.Description = "活动重复参加"
		break
	case 2000007:
		result.Message = "TicketPromotionGroupLost"
		result.Description = "拼团不存在"
		break
	case 2000008:
		result.Message = "TicketPromotionGroupFull"
		result.Description = "拼团人数已满"
		break
	case 2000009:
		result.Message = "TicketPromotionGroupNotFull"
		result.Description = "拼团人数未满"
		break
	case 2000010:
		result.Message = "TicketPromotionOrderLost"
		result.Description = "拼团订单不存在"
		break
	case 2000011:
		result.Message = "TicketPromoExistSameTime"
		result.Description = "同时间段存在已上架拼团活动"
		break
	case 2000012:
		result.Message = "TicketAddPromoOrderFail"
		result.Description = "添加活动订单失败"
		break
	case 2000013:
		result.Message = "TicketAddPromoGroupFail"
		result.Description = "添加拼团 团订单失败"
		break
	case 2000014:
		result.Message = "TicketPromoGroupEnd"
		result.Description = "拼团 团订单已失效"
		break
	case 2000015:
		result.Message = "TicketUpdatePromoOrderFail"
		result.Description = "更新拼团订单失败"
		break
	case 2000016:
		result.Message = "TicketUpdatePromoGroupFail"
		result.Description = "更新拼团 团订单失败"
		break
	case 2000017:
		result.Message = "IllegalPromoOperate"
		result.Description = "拼团 不支持的操作类型"
		break
	case 2000018:
		result.Message = "PromoStatusChanged"
		result.Description = "拼团状态无法变更"
		break
	case 2000019:
		result.Message = "TicketPromoGroupStatusErr"
		result.Description = "拼团状态不对"
		break
	case 2000020:
		result.Message = "TicketPromoOrderTypeErr"
		result.Description = "订单类型不对"
		break
	case 2000021:
		result.Message = "PromoEditNotALlowed"
		result.Description = "不可编辑"
		break
	case 2000022:
		result.Message = "PromoEditFieldNotALlowed"
		result.Description = "不可编辑部分字段"
		break
	case 2000023:
		result.Message = "PromoExists"
		result.Description = "拼团已存在"
		break

		//销售-交易
	case 2000101:
		result.Message = "TicketGetOidFail"
		result.Description = "获取订单号失败"
		break
	case 2000102:
		result.Message = "TicketExceedLimit"
		result.Description = "超过购买限制"
		break
	case 2000103:
		result.Message = "TicketMissData"
		result.Description = "信息不完整"
		break
	case 2000104:
		result.Message = "TicketSaleNotStart"
		result.Description = "没开售"
		break
	case 2000105:
		result.Message = "TicketSaleEnd"
		result.Description = "已结束"
		break
	case 2000106:
		result.Message = "TicketNoPriv"
		result.Description = "无权操作"
		break
	case 2000107:
		result.Message = "TicketInvalidUser"
		result.Description = "无效用户"
		break
	case 2000108:
		result.Message = "TicketPriceChanged"
		result.Description = "价格变化"
		break

	//销售-库存
	case 2000201:
		result.Message = "TicketStockLack"
		result.Description = "库存不足"
		break
	case 2000202:
		result.Message = "TicketStockLogNotFound"
		result.Description = "没有库存操作记录"
		break
	case 2000203:
		result.Message = "TicketStockUpdateFail"
		result.Description = "库存更新失败"
		break

		//番剧推荐
	case 2002000:
		result.Message = "SugEsSearchErr"
		result.Description = "es搜索错误"
		break
	case 2002001:
		result.Message = "SugSearchTypeErr"
		result.Description = "搜索类型错误"
		break
	case 2002002:
		result.Message = "SugOpTypeErr"
		result.Description = "操作类型错误"
		break
	case 2002003:
		result.Message = "SugOpErr"
		result.Description = "add or del match fail"
		break
	case 2002004:
		result.Message = "SugItemNone"
		result.Description = "商品不存在"
		break
	case 2002005:
		result.Message = "SugSeasonNone"
		result.Description = "番剧不存在"
		break

	// 防刷工具
	case 2001000:
		result.Message = "ParamInvalid"
		result.Description = "参数错误"
		break
	case 2001002:
		result.Message = "UpdateError"
		result.Description = "更新失败"
		break
	case 2001003:
		result.Message = "QusbNotFound"
		result.Description = "找不到题库"
		break
	case 2001005:
		result.Message = "QusIDInvalid"
		result.Description = "题目id错误"
		break
	case 2001007:
		result.Message = "BankUsing"
		result.Description = "题目正在使用"
		break
	case 2001009:
		result.Message = "BindBankNotFound"
		result.Description = "未找到题库绑定关系"
		break
	case 2001010:
		result.Message = "AnswerError"
		result.Description = "答案错误"
		break
	case 2001011:
		result.Message = "GetQusBankInfoCache"
		result.Description = "获取题库缓存失败"
		break
	case 2001012:
		result.Message = "GetComponentTimesErr"
		result.Description = "获取组件缓存失败"
		break
	case 2001013:
		result.Message = "SetComponentTimesErr"
		result.Description = "设置答题次数缓存失败"
		break
	case 2001014:
		result.Message = "SetComponentIDErr"
		result.Description = "设置组件缓存失败"
		break
	case 2001015:
		result.Message = "GetComponentIDErr"
		result.Description = "获取组件ID缓存失败"
		break
	case 2001016:
		result.Message = "SameCompentErr"
		result.Description = "相同组件"
		break
	case 2001017:
		result.Message = "GetQusIDsErr"
		result.Description = "获取题目失败"
		break
	case 2001018:
		result.Message = "AnswerPoiError"
		result.Description = "答案错误"
		break
	case 2001019:
		result.Message = "NotEnoughQuestion"
		result.Description = "部分题库不足3题，无法绑定，请绑定别的题库，或者修改题库"
		break
	case 2001020:
		result.Message = "AntiSalesTimeErr"
		result.Description = "售卖时间有错"
		break
	case 2001021:
		result.Message = "AntiIPChangeLimit"
		result.Description = "用户IP变更"
		break
	case 2001022:
		result.Message = "AntiLimitNumUpper"
		result.Description = "次数达到上限"
		break
	case 2001023:
		result.Message = "AntiCheckVoucherErr"
		result.Description = "用户凭证验证失败"
		break
	case 2001024:
		result.Message = "AntiValidateFailed"
		result.Description = "验证失败"
		break
	case 2001025:
		result.Message = "AntiGeetestCountUpper"
		result.Description = "极验总数达到上线"
		break
	case 2001026:
		result.Message = "AntiCustomerErr"
		result.Description = "业务方错误"
		break
	case 2001027:
		result.Message = "AntiBlackErr"
		result.Description = "黑名单用户"
		break

	// 项目
	case 2004000:
		result.Message = "TicketCannotDelTk"
		result.Description = "无法删除票价"
		break
	case 2004001:
		result.Message = "TicketDelTkFailed"
		result.Description = "删除票价失败"
		break
	case 2004002:
		result.Message = "TicketLkTkNotFound"
		result.Description = "关联票种不存在"
		break
	case 2004003:
		result.Message = "TicketLkTkTypeNotFound"
		result.Description = "关联票种类型不存在"
		break
	case 2004004:
		result.Message = "TicketLkScNotFound"
		result.Description = "关联场次不存在"
		break
	case 2004005:
		result.Message = "TicketCannotDelSc"
		result.Description = "无法删除场次"
		break
	case 2004006:
		result.Message = "TicketLkScTimeNotFound"
		result.Description = "关联的场次时间不存在"
		break
	case 2004007:
		result.Message = "TicketPidIsEmpty"
		result.Description = "项目id为空"
		break
	case 2004008:
		result.Message = "TicketMainInfoTooLarge"
		result.Description = "项目版本详情信息量过大"
		break
	case 2004009:
		result.Message = "TicketDelTkExFailed"
		result.Description = "删除票价额外信息失败"
		break
	case 2004010:
		result.Message = "TicketAddVersionFailed"
		result.Description = "添加版本信息失败"
		break
	case 2004011:
		result.Message = "TicketAddVerExtFailed"
		result.Description = "添加版本详情失败"
		break
	case 2004012:
		result.Message = "TicketBannerIDEmpty"
		result.Description = "BannerID为空"
		break
	case 2004013:
		result.Message = "TicketVerCannotEdit"
		result.Description = "版本不可编辑"
		break
	case 2004014:
		result.Message = "TicketVerCannotReview"
		result.Description = "无法审核 非待审核版本"
		break
	case 2004015:
		result.Message = "TicketAddTagFailed"
		result.Description = "添加项目标签失败"
		break
	}
	return result
}
