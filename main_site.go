package bilierrorcode

// TODO: 检查修复 Description 中的错字错意
func getMainSiteDetail(code ErrorCode) ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}
	if code >= 0 && code <= 990000 {
		switch code {
		// archive
		case 10001:
			result.Message = "ArchiveTypeForbidAdd"
			result.Description = "该类型不支持投稿"
			break
		case 10002:
			result.Message = "ArchiveExist"
			result.Description = "已经存在该稿件了"
			break
		case 10003:
			result.Message = "ArchiveNotExist"
			result.Description = "不存在该稿件"
			break
		case 10004:
			result.Message = "ArchiveAlreadyDel"
			result.Description = "稿件已经被删除"
			break
		case 10005:
			result.Message = "ArchiveDelayTimeErr"
			result.Description = "延迟变更的时间早于当前"
			break
		case 10006:
			result.Message = "ArchiveAlreadyDelay"
			result.Description = "延迟变更的时间已存在"
			break
		case 10007:
			result.Message = "ArchiveOwnerErr"
			result.Description = "稿件不属于这个人"
			break
		case 10008:
			result.Message = "VideoshotNotExist"
			result.Description = "稿件的缩略图不存在"
			break
		case 10009:
			result.Message = "VideoAbnormalSubmit"
			result.Description = "异常视频提交"
			break
		case 10010:
			result.Message = "ArchiveBlocked"
			result.Description = "当前稿件已锁定，可能处于以下状态之一(审核锁定，网警锁定,用户删除)"
			break

		// appeal
		case 10101:
			result.Message = "AppealNotExist"
			result.Description = "不存在该申诉"
			break
		case 10102:
			result.Message = "AppealAlreadyClose"
			result.Description = "申诉工单已经被关闭"
			break
		case 10103:
			result.Message = "AppealInterval"
			result.Description = "申诉间隔时间内，不能再发起申诉"
			break
		case 10104:
			result.Message = "AppealOpen"
			result.Description = "该稿件已处于申诉中"
			break
		case 10105:
			result.Message = "AppealOwner"
			result.Description = "只能申诉自己的稿件"
			break
		case 10106:
			result.Message = "AppealHasStar"
			result.Description = "该申诉已评分"
			break
		case 10107:
			result.Message = "AppealLimit"
			result.Description = "仅支持打回和锁定的稿件申诉"
			break

		// favorite
		case 11001:
			result.Message = "FavNameTooLong"
			result.Description = "收藏夹名称过长"
			break
		case 11002:
			result.Message = "FavMaxFolderCount"
			result.Description = "达到最大收藏夹数"
			break
		case 11003:
			result.Message = "FavCanNotDelDefault"
			result.Description = "不能删除默认收藏夹"
			break
		case 11004:
			result.Message = "FavFolderNoPublic"
			result.Description = "收藏夹目录未公开"
			break
		case 11005:
			result.Message = "FavMaxVideoCount"
			result.Description = "视频数达到目录最大收藏数"
			break
		case 11006:
			result.Message = "FavFolderExist"
			result.Description = "已经存在该收藏夹"
			break
		case 11007:
			result.Message = "FavVideoExist"
			result.Description = "已经存在该视频了"
			break
		case 11008:
			result.Message = "FavOnlyPublic"
			result.Description = "仅仅只能设置公开0 或 非公开1"
			break
		case 11009:
			result.Message = "FavDefaultFolder"
			result.Description = "默认收藏夹"
			break
		case 11010:
			result.Message = "FavFolderNotExist"
			result.Description = "没有该收藏夹"
			break
		case 11011:
			result.Message = "FavSearchReqErr"
			result.Description = "请求搜索出错"
			break
		case 11012:
			result.Message = "FavFloderAlreadyDel"
			result.Description = "已经删除该收藏夹了"
			break
		case 11013:
			result.Message = "FavVideoAlreadyDel"
			result.Description = "已经取消收藏该视频了"
			break
		case 11014:
			result.Message = "FavMaxOperNum"
			result.Description = "超出允许的最大操作数75"
			break
		case 11015:
			result.Message = "FavFolderSame"
			result.Description = "一样的收藏夹"
			break
		case 11016:
			result.Message = "FavFolderMoveFailed"
			result.Description = "收藏夹视频移动失败"
			break
		case 11017:
			result.Message = "FavFolderSortErr"
			result.Description = "收藏夹列表信息错误"
			break
		case 11018:
			result.Message = "FavSearchWordIllegal"
			result.Description = "收藏夹视频搜索关键词非法"
			break
		case 11019:
			result.Message = "FavMaintenance"
			result.Description = "收藏服务维护中"
			break
		case 11020:
			result.Message = "FavFolderBanned"
			result.Description = "收藏夹名称包含敏感词"
			break
		case 11021:
			result.Message = "FavCleaneInProgress"
			result.Description = "正在删除失效视频…请过段时间再来访问"
			break
		case 11022:
			result.Message = "FavCleanedLocked"
			result.Description = "清除操作锁定中"
			break
		case 11201:
			result.Message = "FavResourceExist"
			result.Description = "已经收藏过了"
			break
		case 11202:
			result.Message = "FavResourceAlreadyDel"
			result.Description = "已经取消收藏了"
			break
		case 11203:
			result.Message = "FavResourceOverflow"
			result.Description = "达到收藏上限"
			break
		case 11204:
			result.Message = "FavDescTooLang"
			result.Description = "收藏夹描述过长"
			break
		case 11205:
			result.Message = "FavRetryLater"
			result.Description = "请稍后重试"
			break
		case 11206:
			result.Message = "FavHitSensitive"
			result.Description = "收藏夹命中敏感词"
			break

		// reply
		case 12001:
			result.Message = "ReplySubjectExist"
			result.Description = "已经存在评论主题"
			break
		case 12002:
			result.Message = "ReplyForbidReply"
			result.Description = "禁止评论"
			break
		case 12003:
			result.Message = "ReplyForbidRootReply"
			result.Description = "禁止回复"
			break
		case 12004:
			result.Message = "ReplyForbidAction"
			result.Description = "禁止操作： 赞或踩"
			break
		case 12005:
			result.Message = "ReplyForbidReport"
			result.Description = "禁止举报"
			break
		case 12006:
			result.Message = "ReplyNotExist"
			result.Description = "没有该评论"
			break
		case 12007:
			result.Message = "ReplyActioned"
			result.Description = "已经操作过了"
			break
		case 12008:
			result.Message = "ReplyReported"
			result.Description = "已经举报过了"
			break
		case 12009:
			result.Message = "ReplyIllegalSubType"
			result.Description = "评论主体的type不合法"
			break
		case 12010:
			result.Message = "ReplyIllegalRoot"
			result.Description = "不合法的父评论"
			break
		case 12011:
			result.Message = "ReplyIllegalAction"
			result.Description = "不合法的赞或踩"
			break
		case 12012:
			result.Message = "ReplyIllegalReport"
			result.Description = "不合法的举报"
			break
		case 12013:
			result.Message = "ReplyTooManyAts"
			result.Description = "at人数过多"
			break
		case 12014:
			result.Message = "ReplyDeniedAsCD"
			result.Description = "cd时间未到不能评论"
			break
		case 12015:
			result.Message = "ReplyDeniedAsCaptcha"
			result.Description = "验证码错误未到不能评论"
			break
		case 12016:
			result.Message = "ReplyDeniedByFilter"
			result.Description = "评论内容被 word filter 禁止"
			break
		case 12017:
			result.Message = "ReplyMosaicByFilter"
			result.Description = "评论内容被过滤掉敏感词打***"
			break
		case 12018:
			result.Message = "ReplyUpDeniedAsState"
			result.Description = "目前状态不能执行up主操作"
			break
		case 12019:
			result.Message = "ReplyReportDeniedAsCD"
			result.Description = "cd时间未到不能举报"
			break
		case 12020:
			result.Message = "ReplyReportNotExist"
			result.Description = "举报不存在"
			break
		case 12021:
			result.Message = "ReplyReportDealed"
			result.Description = "已经处理该举报了"
			break
		case 12022:
			result.Message = "ReplyDeleted"
			result.Description = "已经被删除了"
			break
		case 12023:
			result.Message = "ReplyRepeat"
			result.Description = "短时间内发重复评论"
			break
		case 12024:
			result.Message = "ReplyNoNeedCaptcha"
			result.Description = "用户不需要验证码"
			break
		case 12025:
			result.Message = "ReplyContentOver"
			result.Description = "评论内容过少或过多"
			break
		case 12026:
			result.Message = "ReplyDeniedByLink"
			result.Description = "一级用户评论里含有链接"
			break
		case 12027:
			result.Message = "ReplyPending"
			result.Description = "评论进入待审"
			break
		case 12028:
			result.Message = "ReplyDeniedAsGarbage"
			result.Description = "大数据垃圾评论"
			break
		case 12029:
			result.Message = "ReplyHaveTop"
			result.Description = "已经有置顶评论"
			break
		case 12030:
			result.Message = "ReplyNotRootReply"
			result.Description = "不能置顶非一级评论"
			break
		case 12031:
			result.Message = "ReplyIllegalSubState"
			result.Description = "sub状态不合法"
			break
		case 12032:
			result.Message = "ReplyEmojiOverMax"
			result.Description = "vip表情超过上限"
			break
		case 12033:
			result.Message = "ReplyVipStatusIllegal"
			result.Description = "vip状态异常无法使用评论特权"
			break
		case 12034:
			result.Message = "ReplyDelTopForbidden"
			result.Description = "置顶评论不允许删除"
			break
		case 12035:
			result.Message = "ReplyBlacklistFilter"
			result.Description = "黑名单屏蔽不予评论"
			break
		case 12036:
			result.Message = "ReplyUpgrading"
			result.Description = "系统升级中"
			break
		case 12037:
			result.Message = "ReplyUserTelVerified"
			result.Description = "请绑定手机号后重试"
			break
		case 12038:
			result.Message = "ReplyForbidList"
			result.Description = "评论已关闭"
			break
		case 12039:
			result.Message = "ReplyHitBlacklist"
			result.Description = "评论含有违规内容"
			break
		case 12040:
			result.Message = "ReplyOverRateLimit"
			result.Description = "一段时间内相似内容出现过多，请稍后再试"
			break
		case 12041:
			result.Message = "ReplyContestNotPassed"
			result.Description = "该账号处于封禁中,点击申请答题"
			break
		case 12042:
			result.Message = "ReplyNoticeConflict"
			result.Description = "时间冲突，发布公告失败"
			break
		case 12043:
			result.Message = "ReplySubjectFrozen"
			result.Description = "该评论区已冻结"
			break
		case 12044:
			result.Message = "ReplyEmojiExits"
			result.Description = "表情已存在"
			break
		case 12045:
			result.Message = "ReplyForbidReplyNotPay"
			result.Description = "禁止评论，未付费"
			break
		case 12048:
			result.Message = "ReplyForbidFolded"
			result.Description = "该评论禁止被折叠"
			break
		case 12049:
			result.Message = "ReplyFolded"
			result.Description = "该评论被折叠"
			break

		// elec
		case 13001:
			result.Message = "ElecUserForbid"
			result.Description = "用户被禁止参与充电计划"
			break
		case 13002:
			result.Message = "ElecUserAudit"
			result.Description = "用户的是否充电正在审核中"
			break
		case 13003:
			result.Message = "ElecNotUpper"
			result.Description = "不是up主"
			break
		case 13004:
			result.Message = "ElecArchiveForbid"
			result.Description = "稿件被屏蔽"
			break

		// stat
		case 14001:
			result.Message = "ClickAesDecryptErr"
			result.Description = "aes解密失败"
			break
		case 14002:
			result.Message = "ClickQueryFormatErr"
			result.Description = "解密粗来的参数格式有问题"
			break
		case 14003:
			result.Message = "ClickServerTimeout"
			result.Description = "服务端时间超时"
			break
		case 14004:
			result.Message = "ClickQuerySignErr"
			result.Description = "参数中的sign错误"
			break
		case 14005:
			result.Message = "ClickHmacSignErr"
			result.Description = "hmac算的签名错误"
			break

		// topic
		case 15001:
			result.Message = "TopicNotExist"
			result.Description = "不存在该话题"
			break
		case 15002:
			result.Message = "FavTopicExist"
			result.Description = "已经存在该话题了"
			break
		case 15003:
			result.Message = "FavTopicAlreadyDel"
			result.Description = "已经取消收藏该话题了"
			break

		// tag
		case 16000:
			result.Message = "TagOperateFail"
			result.Description = "操作失败"
			break
		case 16001:
			result.Message = "TagNotExist"
			result.Description = "tag不存在"
			break
		case 16002:
			result.Message = "TagAidNotExist"
			result.Description = "视频tag关联不存在"
			break
		case 16003:
			result.Message = "TagArcTagMaxNum"
			result.Description = "超过视频最大tag数"
			break
		case 16004:
			result.Message = "TagMaxSub"
			result.Description = "超过个人最大订阅数"
			break
		case 16005:
			result.Message = "TagOpDenied"
			result.Description = "tag操作权限不够"
			break
		case 16006:
			result.Message = "TagAidNoTags"
			result.Description = "该视频不存在tag"
			break
		case 16007:
			result.Message = "TagRidNotExist"
			result.Description = "rid分区不存在"
			break
		case 16008:
			result.Message = "TagAvNotExist"
			result.Description = "该视频不存在"
			break
		case 16009:
			result.Message = "TagArcTagExist"
			result.Description = "视频存在此tag"
			break
		case 16010:
			result.Message = "TagArcTagAddMaxFre"
			result.Description = "添加太多标签啦,休息休息~"
			break
		case 16011:
			result.Message = "TagArcTagDelMaxFre"
			result.Description = "删除太多标签啦,休息休息~"
			break
		case 16012:
			result.Message = "TagArcTagLikeMaxFre"
			result.Description = "顶/踩太多次啦,休息休息~"
			break
		case 16013:
			result.Message = "TagArcTagRptMaxFre"
			result.Description = "举报太多次啦,休息休息~"
			break
		case 16014:
			result.Message = "TagArcCannotAddTag"
			result.Description = "管理员暂时不让添加标签哦~"
			break
		case 16015:
			result.Message = "TagArcCannotDelTag"
			result.Description = "管理员暂时不让删除标签哦~"
			break
		case 16016:
			result.Message = "TagArcAddLevelLower"
			result.Description = "升级到Lv3就能添加标签啦~"
			break
		case 16017:
			result.Message = "TagUpAddNotAllowed"
			result.Description = "然而up主并不希望你添加~"
			break
		case 16018:
			result.Message = "TagLikeNotDelAllowed"
			result.Description = "被大家顶起来的标签,不能随便删除哦~"
			break
		case 16020:
			result.Message = "TagArcTagNotExist"
			result.Description = "视频不存在此tag"
			break
		case 16021:
			result.Message = "TagArcTagLiked"
			result.Description = "已经顶过啦~"
			break
		case 16022:
			result.Message = "TagArcTagHated"
			result.Description = "已经踩过啦~"
			break
		case 16023:
			result.Message = "TagArcTagisLocked"
			result.Description = "已经是锁定状态"
			break
		case 16024:
			result.Message = "TagArcTagnoLocked"
			result.Description = "已经是非锁定状态"
			break
		case 16025:
			result.Message = "TagIsSealing"
			result.Description = "tag已经被封印了~"
			break
		case 16026:
			result.Message = "TagArcDelLevelLower"
			result.Description = "升级到Lv4就能删除标签啦~"
			break
		case 16027:
			result.Message = "TagArcRptLevelLower"
			result.Description = "升级到Lv1就能举报标签啦~"
			break
		case 16028:
			result.Message = "TagArcLikeLevelLower"
			result.Description = "升级到Lv1就能喜欢标签啦~"
			break
		case 16029:
			result.Message = "TagArcHateLevelLower"
			result.Description = "升级到Lv1就能讨厌标签啦~"
			break
		case 16030:
			result.Message = "TagTagIsSubed"
			result.Description = "tag 已经订阅啦~"
			break
		case 16031:
			result.Message = "TagRankingTagNotExist"
			result.Description = "排行榜热门tag不存在"
			break
		case 16032:
			result.Message = "TagRankingBsNotExist"
			result.Description = "排行榜番剧不存在"
			break
		case 16033:
			result.Message = "TagRankingPsNotExist"
			result.Description = "排行榜分区不存在"
			break
		case 16034:
			result.Message = "TagRankingSimNotExist"
			result.Description = "排行榜分区不存在"
			break
		case 16035:
			result.Message = "TagNotSub"
			result.Description = "未订阅"
			break
		case 16036:
			result.Message = "TagArcTagRpted"
			result.Description = "已经举报啦~"
			break
		case 16037:
			result.Message = "TagArcTagDelMore"
			result.Description = "超过单个视频日删除次数上线~"
			break
		case 16038:
			result.Message = "TagRptNotRptPassed"
			result.Description = "已经被审核通过的,不能再继续举报哦~"
			break
		case 16039:
			result.Message = "TagArcTagLogNotExsit"
			result.Description = "视频tag操作日志不存在"
			break
		case 16040:
			result.Message = "TagDelNotRptPassed"
			result.Description = "该标签已经审核为正常标签,不能删除哦~"
			break
		case 16041:
			result.Message = "TagUpTagCannotDel"
			result.Description = "不能删除UP主添加的tag哦~"
			break
		case 16042:
			result.Message = "TagAddNotRptPassed"
			result.Description = "该标签已经认定为恶意标签,不能添加哦~"
			break
		case 16043:
			result.Message = "TagAdminOpCanNotTpt"
			result.Description = "管理员操作的标签不能举报~"
			break
		case 16044:
			result.Message = "TagArcAccountBlocked"
			result.Description = "你被封禁，禁止操作~"
			break
		case 16045:
			result.Message = "TagIsOfficailTag"
			result.Description = "此tag是官方活动tag,不可以操作~"
			break
		case 16046:
			result.Message = "TagAlreadyExist"
			result.Description = "Tag已存在，不能重复添加~"
			break
		case 16047:
			result.Message = "TagResTagMaxNum"
			result.Description = "资源绑定tag超过最大数量~"
			break
		case 16048:
			result.Message = "TagResTagExist"
			result.Description = "资源绑定tag关系已存在~"
			break
		case 16049:
			result.Message = "TagResTagNotExist"
			result.Description = "资源绑定tag关系不存在~"
			break
		case 16050:
			result.Message = "TagServiceUpdate"
			result.Description = "Tag服务升级中~"
			break
		case 16051:
			result.Message = "TagUpdateLimitUserFail"
			result.Description = "Tag服务更新白名单用户失败"
			break
		case 16052:
			result.Message = "TagLimitUserExist"
			result.Description = "Tag服务白名单用户已存在"
			break
		case 16053:
			result.Message = "TagLimitUserFail"
			result.Description = "Tag服务获取白名单用户信息失败"
			break
		case 16054:
			result.Message = "TagResTagFail"
			result.Description = "资源绑定tag失败"
			break
		case 16055:
			result.Message = "TagReportNotExist"
			result.Description = "该举报不存在"
			break
		case 16056:
			result.Message = "TagReportIsDealed"
			result.Description = "该举报已经处理无法移动"
			break
		case 16057:
			result.Message = "TagPunishMsgFailed"
			result.Description = "Tag惩罚通知用户失败"
			break
		case 16058:
			result.Message = "TagReportLogAddFailed"
			result.Description = "Tag举报日志增加失败"
			break
		case 16059:
			result.Message = "TagReportUpdateFailed"
			result.Description = "Tag举报更新失败"
			break
		case 16060:
			result.Message = "TagHotUpdateFailed"
			result.Description = "热门Tag更新失败"
			break
		case 16061:
			result.Message = "TagAlreadyExamined"
			result.Description = "tag已经审核过了"
			break
		case 16062:
			result.Message = "TagAlreadyShield"
			result.Description = "tag已经被屏蔽"
			break
		case 16063:
			result.Message = "TagArcTagLockFail"
			result.Description = "tag稿件锁定失败"
			break
		case 16064:
			result.Message = "TagUpdateFail"
			result.Description = "tag更新失败"
			break
		case 16065:
			result.Message = "TagAddFail"
			result.Description = "tag增加失败"
			break
		case 16066:
			result.Message = "TagChangeStateFail"
			result.Description = "tag更改状态失败"
			break
		case 16067:
			result.Message = "TagArcTagUnLockFail"
			result.Description = "tag稿件解除锁定失败"
			break
		case 16068:
			result.Message = "TagReportLogNotExist"
			result.Description = "举报日志不存在"
			break
		case 16069:
			result.Message = "TagAlreadyDelete"
			result.Description = "tag已经被删除"
			break
		case 16070:
			result.Message = "TagOnlyUpAdd"
			result.Description = "只有UP主可以添加频道"
			break
		case 16071:
			result.Message = "TagOnlyUpDel"
			result.Description = "只有UP主可以删除频道"
			break
		case 16101:
			result.Message = "ChannelAlreadyExist"
			result.Description = "频道已存在"
			break
		case 16102:
			result.Message = "ChannelNotExist"
			result.Description = "频道不存在"
			break
		case 16103:
			result.Message = "ChanRuleException"
			result.Description = "频道规则有异常"
			break
		case 16104:
			result.Message = "ChanTypeNotExist"
			result.Description = "频道类型不存在"
			break
		case 16105:
			result.Message = "ChanRuleCanotUse"
			result.Description = "频道规则不符合标准"
			break
		case 16106:
			result.Message = "ChannelNoLike"
			result.Description = "频道不能点赞"
			break
		case 16107:
			result.Message = "ChannelNoHate"
			result.Description = "频道不能点踩"
			break
		case 16108:
			result.Message = "ChannelNotAllowDel"
			result.Description = "这个不能删除哦~"
			break
		case 16109:
			result.Message = "ChannelAINoData"
			result.Description = "该频道下没数据了哦"
			break
		case 16110:
			result.Message = "ChannelAITimeout"
			result.Description = "AI超时了，技能冷却中"
			break
		case 16111:
			result.Message = "ChannelAleadyMigrated"
			result.Description = "频道已经被迁移过了"
			break
		case 16112:
			result.Message = "ChannelCanNotDel"
			result.Description = "频道分类下有数据，禁止操作"
			break
		case 16113:
			result.Message = "ChannelRecommended"
			result.Description = "频道已经被推荐了，无须再推荐"
			break
		case 16114:
			result.Message = "ChannelTypeExist"
			result.Description = "频道分类已存在"
			break
		case 16115:
			result.Message = "ChannelTypeDeled"
			result.Description = "频道分类已删除"
			break
		case 16116:
			result.Message = "ChannelNoChange"
			result.Description = "频道没变化"
			break
		case 16117:
			result.Message = "CategoryNoChange"
			result.Description = "频道分类没变化"
			break

		// zlimit
		case 17001:
			result.Message = "ZlimitAllow"
			result.Description = "稿件无限制"
			break
		case 17002:
			result.Message = "ZlimitForbidden"
			result.Description = "稿件禁止查看"
			break
		case 17003:
			result.Message = "ZlimitFormal"
			result.Description = "正式会员"
			break
		case 17004:
			result.Message = "ZlimitPay"
			result.Description = "付费会员"
			break
		case 17005:
			result.Message = "ZlimitShared"
			result.Description = "共享地址错误提示"
			break

		// Feedback
		case 18001:
			result.Message = "FeedbackSsnNotExist"
			result.Description = "session不存在"
			break
		case 18002:
			result.Message = "FeedbackBodyTooLarge"
			result.Description = "上传的文件太大"
			break
		case 18003:
			result.Message = "FeedbackBodyNotExist"
			result.Description = "上传的文件无内容"
			break
		case 18004:
			result.Message = "FeedbackContentOver"
			result.Description = "创建会话过多"
			break

		// short utl
		case 30001:
			result.Message = "ShortURLAlreadyExist"
			result.Description = "短链已存在"
			break
		case 30002:
			result.Message = "ShortURLNotFound"
			result.Description = "短链不存在"
			break
		case 30003:
			result.Message = "ShortURLIllegalSrc"
			result.Description = "生成短链的长链来源不合法"
			break

		// captcha
		case 33001:
			result.Message = "CaptchaSignErr"
			result.Description = "签名错误"
			break
		case 33002:
			result.Message = "CaptchaTSOverTolerant"
			result.Description = "时间戳超过规定时间范围"
			break
		case 33003:
			result.Message = "CaptchaBusinessNotAllowed"
			result.Description = "业务id不存在"
			break
		case 33004:
			result.Message = "CaptchaCodeNotFound"
			result.Description = "code不存在，一般是token已过期"
			break
		case 33005:
			result.Message = "CaptchaTokenErr"
			result.Description = "验证时token错误"
			break
		case 33006:
			result.Message = "CaptchaNotCreate"
			result.Description = "验证码未创建"
			break
		case 33007:
			result.Message = "CaptchaTokenNotExist"
			result.Description = "验证码Token不存在"
			break
		case 33008:
			result.Message = "CaptchaTokenExpired"
			result.Description = "验证码已过期"
			break

		// Coin
		case 34001:
			result.Message = "CoinRatingErr"
			result.Description = "评级参数错误"
			break
		case 34002:
			result.Message = "CoinCannotAddToSelf"
			result.Description = "up主不能自己投币"
			break
		case 34003:
			result.Message = "CoinIllegaMultiply"
			result.Description = "非法的投币数量"
			break
		case 34004:
			result.Message = "CoinDeniedAsCD"
			result.Description = "投币间隔太短"
			break
		case 34005:
			result.Message = "CoinOverMax"
			result.Description = "超过单个视频投币上限"
			break

		// account
		case 35001:
			result.Message = "AccountOverdue"
			result.Description = "token过期 -658"
			break
		case 35002:
			result.Message = "AccountNotLogin"
			result.Description = "未登录 -400"
			break
		case 35003:
			result.Message = "AccountInexistence"
			result.Description = "用户不存在 -626"
			break
		case 35004:
			result.Message = "AccountAKNotFound"
			result.Description = "accessKey不存在 -2"
			break

		// player-interface
		case 19001:
			result.Message = "PLayerPolicyNotExist"
			result.Description = "策略不存在"
			break
		case 19002:
			result.Message = "PLayerPolicyNotStart"
			result.Description = "策略未开始"
			break
		case 19003:
			result.Message = "PLayerPolicyEnded"
			result.Description = "策略未开始"
			break

		// creative
		case 20001:
			result.Message = "CreativeArchiveAPIErr"
			result.Description = "投稿API错误"
			break
		case 20002:
			result.Message = "CreativeSearchErr"
			result.Description = "搜索API错误"
			break
		case 20003:
			result.Message = "CreativeDataErr"
			result.Description = "大数据API错误"
			break
		case 20004:
			result.Message = "CreativeElecErr"
			result.Description = "充电API错误"
			break
		case 20005:
			result.Message = "CreativeActivityErr"
			result.Description = "活动API错误"
			break
		case 20006:
			result.Message = "CreativeArcServiceErr"
			result.Description = "稿件服务错误"
			break
		case 20007:
			result.Message = "CreativeAccServiceErr"
			result.Description = "账号服务错误"
			break
		case 20008:
			result.Message = "CreativeOrderAPIErr"
			result.Description = "商单API错误"
			break
		case 20009:
			result.Message = "CreativeGeetestErr"
			result.Description = "验证码错误"
			break
		case 20010:
			result.Message = "CreativeFeedbackErr"
			result.Description = "用户反馈服务错误"
			break
		case 20011:
			result.Message = "CreativeRecommendOverMax"
			result.Description = "推荐关联的稿件已超过上限(最多8个)"
			break
		case 20012:
			result.Message = "CreativeRecommendForbid"
			result.Description = "当前稿件尚未开放浏览，不允许被关联推荐"
			break
		case 20013:
			result.Message = "CreativeTemplateOverMax"
			result.Description = "添加的稿件模板数量已超过上限(最多5个)"
			break
		case 20014:
			result.Message = "CreativeAccBanned"
			result.Description = "用户被禁言"
			break
		case 20015:
			result.Message = "CreativeAccNotActivated"
			result.Description = "用户未激活"
			break
		case 20016:
			result.Message = "CreativeTemplateDeleted"
			result.Description = "稿件模板已删除"
			break
		case 20017:
			result.Message = "CreativeArticleRPCErr"
			result.Description = "文章RPC接口错误"
			break
		case 20018:
			result.Message = "CreativeArticleParamErr"
			result.Description = "文章投稿参数错误"
			break
		case 20019:
			result.Message = "CreativeArticleCanNotRepeat"
			result.Description = "新增文章同一个标题短时间内不能重复提交"
			break
		case 20020:
			result.Message = "CreativeArticleNotExist"
			result.Description = "修改的文章不存在"
			break
		case 20021:
			result.Message = "CreativeArticleOwnerErr"
			result.Description = "不是你的文章"
			break
		case 20022:
			result.Message = "CreativeArticleCategoryErr"
			result.Description = "文章分类不合法"
			break
		case 20023:
			result.Message = "CreativeArticleReprintErr"
			result.Description = "文章转载类型不合法"
			break
		case 20024:
			result.Message = "CreativeArticleTagErr"
			result.Description = "文章tag不合法"
			break
		case 20025:
			result.Message = "CreativeArticleCoverErr"
			result.Description = "文章封面不合法"
			break
		case 20026:
			result.Message = "CreativeArticleImageURLsErr"
			result.Description = "文章小图片不合法"
			break
		case 20027:
			result.Message = "CreativeArticleTitleErr"
			result.Description = "文章标题不合法"
			break
		case 20028:
			result.Message = "CreativeArticleContentErr"
			result.Description = "文章内容不合法"
			break
		case 20029:
			result.Message = "CreativeArticleUploadErr"
			result.Description = "文件上传失败"
			break
		case 20030:
			result.Message = "CreativeArticleTIDErr"
			result.Description = "文章模板类型不合法"
			break
		case 20031:
			result.Message = "CreativeArticleImageTypeErr"
			result.Description = "图片类型不合法"
			break
		case 20032:
			result.Message = "CreativeFansMedalErr"
			result.Description = "粉丝勋章服务错误"
			break
		case 20033:
			result.Message = "CreativeWaterMarkWrongState"
			result.Description = "水印状态参数错误"
			break
		case 20034:
			result.Message = "CreativeWaterMarkWrongType"
			result.Description = "水印类型参数错误"
			break
		case 20035:
			result.Message = "CreativeWaterMarkWrongPosition"
			result.Description = "水印位置参数错误"
			break
		case 20036:
			result.Message = "CreativeWaterMarkCreateFailed"
			result.Description = "水印生成错误"
			break
		case 20037:
			result.Message = "CreativeLiveErr"
			result.Description = "直播间服务错误"
			break
		case 20038:
			result.Message = "CreativeLiveNotOpenErr"
			result.Description = "没有开通直播间"
			break
		case 20039:
			result.Message = "CreativePhoneNotBind"
			result.Description = "请先绑定手机"
			break
		case 20040:
			result.Message = "CreativeAssistErr"
			result.Description = "协管服务错误"
			break
		case 20041:
			result.Message = "CreativeAssistLogAlreadyRevoc"
			result.Description = "协管日志已被撤销，不允许重复撤销"
			break
		case 20042:
			result.Message = "CreativeTagErr"
			result.Description = "查询tag出错"
			break
		case 20043:
			result.Message = "CreativeDanmuErr"
			result.Description = "弹幕服务错误"
			break
		case 20044:
			result.Message = "CreativeReplyAPIErr"
			result.Description = "评论API错误"
			break
		case 20045:
			result.Message = "CreativeCoinAPIErr"
			result.Description = "硬币API错误"
			break
		case 20046:
			result.Message = "CreativeUpNoMedalRight"
			result.Description = "Up主没有粉丝勋章资格"
			break
		case 20047:
			result.Message = "CreativeMissionDeltimeInvalid"
			result.Description = "禁止删除，当前稿件已过所参加活动的可删除截止时间"
			break
		case 20048:
			result.Message = "CreativeDanmuFilterParamError"
			result.Description = "弹幕过滤参数错误"
			break
		case 20049:
			result.Message = "CreativeNotLogin"
			result.Description = "未登录"
			break
		case 20050:
			result.Message = "CreativeDynamicIntroErr"
			result.Description = "专栏动态推荐语不合法"
			break
		case 20051:
			result.Message = "CreativeGameOpenAPIErr"
			result.Description = "游戏开放平台API错误"
			break
		case 20052:
			result.Message = "CreativePorderForbidShowFront"
			result.Description = "当前稿件的私单已被禁止在前端展现"
			break
		case 20053:
			result.Message = "CreativeMusicErr"
			result.Description = "音乐API错误"
			break
		case 20054:
			result.Message = "CreativeGeetestAPIErr"
			result.Description = "投稿极验服务不可用"
			break
		case 20055:
			result.Message = "CreativeAcademyOIDExistErr"
			result.Description = "已存在该稿件，不能重复添加"
			break
		case 20056:
			result.Message = "CreativeMessageAPIErr"
			result.Description = "消息服务API错误"
			break
		case 20057:
			result.Message = "CreativeLotteryAPIErr"
			result.Description = "抽奖服务API错误"
			break
		case 20058:
			result.Message = "CreativeSubtitleAPIErr"
			result.Description = "字幕服务API错误"
			break
		case 20059:
			result.Message = "CreativeAcademyH5RecommendErr"
			result.Description = "创作学院h5推荐服务错误"
			break
		case 20060:
			result.Message = "CreativePayForbidDeleteAfterOpen"
			result.Description = "第一次开放之后禁止修改稿件信息，包括付费和分P，请再等60天"
			break
		case 20061:
			result.Message = "CreativePayAPIErr"
			result.Description = "UGC付费查询API错误"
			break
		case 20062:
			result.Message = "CreativeNewcomerMallAPIErr"
			result.Description = "会员购服务错误"
			break
		case 20063:
			result.Message = "CreativeNewcomerBCoinAPIErr"
			result.Description = "B币券服务错误"
			break
		case 20064:
			result.Message = "CreativeNewcomerPendantAPIErr"
			result.Description = "挂件服务错误"
			break
		case 20065:
			result.Message = "CreativeNewcomerRepeatRewardErr"
			result.Description = "重复领取奖励"
			break
		case 20066:
			result.Message = "CreativeNewcomerNotCompleteErr"
			result.Description = "任务未完成不可领取奖励"
			break
		case 20067:
			result.Message = "CreativeNewcomerReBindTaskErr"
			result.Description = "重复绑定任务"
			break
		case 20068:
			result.Message = "CreativeNotUper"
			result.Description = "用户非up主"
			break
		case 20069:
			result.Message = "CreativeNewcomerBigMemberErr"
			result.Description = "大会员服务错误"
			break
		case 20070:
			result.Message = "CreativeNewcomerReceiveExpireErr"
			result.Description = "奖励已过期"
			break
		case 20071:
			result.Message = "CreativeNewcomerNoTask"
			result.Description = "任务体系正在配置任务，请稍候"
			break
		case 20072:
			result.Message = "CreativeNewcomerGroupIDErr"
			result.Description = "任务组ID不存在"
			break
		case 20073:
			result.Message = "CreativeNewcomerDuplicateGiftRewardIDErr"
			result.Description = "任务礼包包含重复任务分类和奖励ID"
			break
		case 20075:
			result.Message = "CreativeAcademyDuplicateSoftIDErr"
			result.Description = "稿件已绑定该软件"
			break

		// videoup
		case 21001:
			result.Message = "VideoupParamErr"
			result.Description = "参数错误"
			break
		case 21002:
			result.Message = "VideoupTypeidErr"
			result.Description = "该分区不存在"
			break
		case 21003:
			result.Message = "VideoupCopyrightErr"
			result.Description = "该稿件类型不合法"
			break
		case 21004:
			result.Message = "VideoupMissionErr"
			result.Description = "该活动不存在"
			break
		case 21005:
			result.Message = "VideoupTagErr"
			result.Description = "Tag参数不合法"
			break
		case 21006:
			result.Message = "VideoupDelayTimeErr"
			result.Description = "定时发布设置错误"
			break
		case 21007:
			result.Message = "VideoupSubmitErr"
			result.Description = "当前时间限制提交"
			break
		case 21008:
			result.Message = "VideoupCoverErr"
			result.Description = "该封面地址不合法"
			break
		case 21009:
			result.Message = "VideoupTitleErr"
			result.Description = "标题不合法"
			break
		case 21010:
			result.Message = "VideoupDescErr"
			result.Description = "描述信息不合法"
			break
		case 21011:
			result.Message = "VideoupZeroVideos"
			result.Description = "新增稿件分P不能为空"
			break
		case 21012:
			result.Message = "VideoupCanotRepeat"
			result.Description = "新增稿件同一个标题短时间内不能重复提交"
			break
		case 21013:
			result.Message = "VideoupVideoTitleErr"
			result.Description = "分P标题有违禁符号"
			break
		case 21014:
			result.Message = "VideoupVideoDescErr"
			result.Description = "分P描述有违禁符号"
			break
		case 21015:
			result.Message = "VideoupVideoFilenameErr"
			result.Description = "分P的filename不能为空"
			break
		case 21016:
			result.Message = "VideoupBodyFormatErr"
			result.Description = "代码模式格式错误"
			break
		case 21017:
			result.Message = "VideoupBodyCidErr"
			result.Description = "代码模式cid不为数字"
			break
		case 21018:
			result.Message = "VideoupBodyTitleEmpty"
			result.Description = "代码模式title不能为空（如果多于1P"
			break
		case 21019:
			result.Message = "VideoupBodyForbidCid"
			result.Description = "代码模式存在不可用cid"
			break
		case 21020:
			result.Message = "VideoupBodyForbidAllCid"
			result.Description = "代码模式的cid全部不可用"
			break
		case 21021:
			result.Message = "VideoupSourceErr"
			result.Description = "稿件转载来源不能为空"
			break
		case 21022:
			result.Message = "VideoupOrderAPIErr"
			result.Description = "商单API错误"
			break
		case 21023:
			result.Message = "VideoupUperIDNotAllow"
			result.Description = "当前用户尚未加入商单计划"
			break
		case 21024:
			result.Message = "VideoupOrderIDNotAllow"
			result.Description = "当前商单ID不合法"
			break
		case 21025:
			result.Message = "VideoupLaunchTimeIllegal"
			result.Description = "当前商单的发布时间不合法"
			break
		case 21026:
			result.Message = "VideoupForbidNoreprint"
			result.Description = "已允许转载的稿件不允许改为禁止转载"
			break
		case 21027:
			result.Message = "VideoupFilenameCanotRepeat"
			result.Description = "同个稿件内不能重复提交相同的视频"
			break
		case 21028:
			result.Message = "VideoupMissionEtimeInvalid"
			result.Description = "该非法活动禁止参加,原因(结束时间不合法)"
			break
		case 21029:
			result.Message = "VideoupPvodAderTooMany"
			result.Description = "稿件私单广告主过多"
			break
		case 21030:
			result.Message = "VideoupPvodAderTooLong"
			result.Description = "稿件私单单个广告主长度已超出限制"
			break
		case 21031:
			result.Message = "VideoupPvodWithoutCommercial"
			result.Description = "稿件私单参数错误，商业广告标记为否"
			break
		case 21032:
			result.Message = "VideoupPvodEpyFlowRemark"
			result.Description = "稿件私单参数错误，商业广告类型为空"
			break
		case 21033:
			result.Message = "VideoupPvodForbidFlowID"
			result.Description = "稿件私单参数错误，商业广告类型ID错误"
			break
		case 21034:
			result.Message = "VideoupPvodForbidFlowRemark"
			result.Description = "稿件私单参数错误，商业广告类型名称错误"
			break
		case 21035:
			result.Message = "VideoupPvodForbidOrderAlready"
			result.Description = "稿件私单参数错误，当前用户已选择有效的商单ID，与私单互斥"
			break
		case 21036:
			result.Message = "VideoupFilenameExpired"
			result.Description = "第(%d)P的文件已超过了限定的时间范围(48小时)，请在删除此分P后再提交"
			break
		case 21037:
			result.Message = "VideoupForbidMultiVideoForTypes"
			result.Description = "[电影]分区的特殊二级分区[欧美电影,日本电影,国产电影,其他国家]不允许提交多P稿件"
			break
		case 21038:
			result.Message = "VideoupDynamicErr"
			result.Description = "稿件动态文案不合法"
			break
		case 21051:
			result.Message = "VideoupFmDesZero"
			result.Description = "稿件描述长度为零"
			break
		case 21052:
			result.Message = "VideoupFmDesLenOverLimit"
			result.Description = "稿件描述长度太长，已超过限制"
			break
		case 21053:
			result.Message = "VideoupFmDesIDForbid"
			result.Description = "稿件描述类型不存在或者不匹配"
			break
		case 21054:
			result.Message = "VideoupFmDesIDNoMatchTypeID"
			result.Description = "稿件描述类型和对应的分区类型不匹配"
			break
		case 21055:
			result.Message = "VideoupFmDesIDNoMatchCopyright"
			result.Description = "稿件描述类型和对应的创作类型不匹配"
			break
		case 21057:
			result.Message = "VideoupTagForbid"
			result.Description = "第(%d)个Tag已被封印"
			break
		case 21058:
			result.Message = "VideoupVideosMaxLimit"
			result.Description = "当前单次操作添加的视频数已经超过限制"
			break
		case 21059:
			result.Message = "VideoupAdBrandIDErr"
			result.Description = "稿件商业推广信息(官方品牌ID)有误，请修正"
			break
		case 21060:
			result.Message = "VideoupAdIndustryIDErr"
			result.Description = "稿件商业推广信息(推广行业)有误，请修正"
			break
		case 21061:
			result.Message = "VideoupAdShowTypeErr"
			result.Description = "稿件商业推广信息(推广形式)有误，请修正"
			break
		case 21062:
			result.Message = "VideoupAdOfficialIndustryIDErr"
			result.Description = "稿件商业推广信息(官方行业)有误，请修正"
			break
		case 21063:
			result.Message = "VideoupAppMustUploadCover"
			result.Description = "APP投稿封面不能为空，请添加稿件封面后再继续"
			break
		case 21064:
			result.Message = "VideoupFilterServiceErr"
			result.Description = "投稿过滤服务暂不可用"
			break
		case 21065:
			result.Message = "VideoupSingleFilterForbid"
			result.Description = "当前输入有敏感信息,请修正"
			break
		case 21066:
			result.Message = "VideoupFieldFilterForbid"
			result.Description = "%s中包含敏感词"
			break
		case 21067:
			result.Message = "VideoupTagForbidNotJoinMission"
			result.Description = "自定义标签包含不可选的活动tag，请修改后重新提交"
			break
		case 21068:
			result.Message = "VideoupMissionNoMatch"
			result.Description = "当前活动和分区不匹配，请重新选择活动或者分区"
			break
		case 21069:
			result.Message = "VideoupChannelReviewNotTrigger"
			result.Description = "当前稿件所属频道不需要回查"
			break
		case 21070:
			result.Message = "VideoupAddLimitHalfMin"
			result.Description = "您投稿的频率过快，请稍等30秒"
			break
		case 21071:
			result.Message = "VideoupCopyForbidJoinMission"
			result.Description = "转载类型稿件不支持活动参加哦~"
			break
		case 21072:
			result.Message = "VideoupEditLocked"
			result.Description = "稿件后台处理中,请10秒后再尝试"
			break
		case 21073:
			result.Message = "VideoupMaxAllVsCntLimit"
			result.Description = "当前总提交视频个数已经超过上限:(%d)"
			break
		case 21074:
			result.Message = "VideoupPayCopyrightErr"
			result.Description = "非自制稿件禁止参与UGC内容付费"
			break
		case 21075:
			result.Message = "VideoupPayProtocolLimit"
			result.Description = "UGC付费修改之前必须接受当前最新的投稿协议"
			break
		case 21076:
			result.Message = "VideoupPayAPIErr"
			result.Description = "UGC付费服务API错误"
			break
		case 21077:
			result.Message = "VideoupPayPriceErr"
			result.Description = "当前稿件定价错误"
			break
		case 21078:
			result.Message = "VideoupPayUserNotAllow"
			result.Description = "当前用户暂无参与UGC内容付费的资格"
			break
		case 21079:
			result.Message = "VideoupPayCommericalLimit"
			result.Description = "禁止付费投稿时同时参与其他商业性投稿活动"
			break
		case 21080:
			result.Message = "VideoupPayForbidEditAfterOpen"
			result.Description = "第一次开放之后禁止修改付费信息"
			break
		case 21085:
			result.Message = "VideoupStaffBlocked"
			result.Description = "因参与者%s隐私设置，暂无法提交"
			break
		case 21086:
			result.Message = "VideoupStaffCopyright"
			result.Description = "非自制稿件不可提交多人合作稿件"
			break
		case 21087:
			result.Message = "VideoupStaffChangeCopyright"
			result.Description = "多人合作稿件不可更改稿件类型"
			break
		case 21088:
			result.Message = "VideoupStaffTypeNotExists"
			result.Description = "该分区暂未开放多人合作稿件"
			break
		case 21089:
			result.Message = "VideoupStaffAuth"
			result.Description = "该up主暂未在多人合作稿件公测范围内"
			break
		case 21090:
			result.Message = "VideoupStaffCountLimit"
			result.Description = "超出团队成员上限，暂不可提交"
			break
		case 21091:
			result.Message = "VideoupStaffMidInvalid"
			result.Description = "无效参与者"
			break
		case 21092:
			result.Message = "VideoupStaffTitleLength"
			result.Description = "参与者类型长度超出上限，暂不可提交"
			break
		case 21093:
			result.Message = "VideoupStaffTitleFilter"
			result.Description = "职能已经被封印了"
			break
		case 21094:
			result.Message = "VideoupStaffMidRepeat"
			result.Description = "已经添加过啦，暂不可重复添加"
			break
		case 21096:
			result.Message = "VideoupStaffTitleChar"
			result.Description = "参与者类型暂时仅支持中文、英文、数字哦"
			break
		case 21097:
			result.Message = "VideoupStaffTitleShort"
			result.Description = "参与者类型长度不足，暂不可提交"
			break
		case 21098:
			result.Message = "VideoupStaffUpSilence"
			result.Description = "很抱歉参与者%s的账号目前为封禁状态"
			break
		case 21099:
			result.Message = "VideoupStaffChangeTypeCopyright"
			result.Description = "暂不支持修改分区和转载类型"
			break
		case 21081:
			result.Message = "VideoupStaffApply404"
			result.Description = "申请单不存在"
			break
		case 21082:
			result.Message = "VideoupStaffApplyTypeChange"
			result.Description = "申请单性质改变"
			break
		case 21083:
			result.Message = "VideoupStaffApplyStateNotMatch"
			result.Description = "申请单操作不允许"
			break
		case 21084:
			result.Message = "VideoupStaffApplyMidNotIn"
			result.Description = "申请单staff不存在"
			break
		case 21101:
			result.Message = "VideoupStaffApplyUpMidBlack"
			result.Description = "staff被UP拉黑"
			break
		case 21102:
			result.Message = "VideoupStaffMidSilence"
			result.Description = "Staff 被封禁"
			break

		// relation
		case 22001:
			result.Message = "RelFollowSelfBanned"
			result.Description = "不能关注自己"
			break
		case 22002:
			result.Message = "RelFollowBlacked"
			result.Description = "被用户拉黑，无法关注"
			break
		case 22003:
			result.Message = "RelFollowAlreadyBlack"
			result.Description = "已经拉黑用户，无法关注"
			break
		case 22004:
			result.Message = "RelFollowAttrAlreadySet"
			result.Description = "已经设置该属性了"
			break
		case 22005:
			result.Message = "RelFollowAttrNotSet"
			result.Description = "未设置该属性，不能取消"
			break
		case 22006:
			result.Message = "RelFollowReachTelLimit"
			result.Description = "关注已达上限，答题成为正式会员或者绑定手机号才能继续关注"
			break
		case 22007:
			result.Message = "RelFollowingGuestLimit"
			result.Description = "访客只限制访问前五页"
			break
		case 22008:
			result.Message = "RelBlackReachMaxLimit"
			result.Description = "黑名单达到上限"
			break
		case 22009:
			result.Message = "RelFollowReachMaxLimit"
			result.Description = "关注失败，已达关注上限"
			break
		case 22010:
			result.Message = "RelBatchFollowAlreadyBlack"
			result.Description = "部分拉黑用户未成功关注"
			break
		case 22101:
			result.Message = "RelTagExistNotAllowedWords"
			result.Description = "分组名称存在不允许的字符"
			break
		case 22102:
			result.Message = "RelTagNumLimit"
			result.Description = "分组数量超过限制"
			break
		case 22103:
			result.Message = "RelTagLenLimit"
			result.Description = "分组名称长度超过限制"
			break
		case 22104:
			result.Message = "RelTagNotExist"
			result.Description = "分组不存在"
			break
		case 22105:
			result.Message = "RelTagAddFollowingFirst"
			result.Description = "请先公开关注后再添加分组"
			break
		case 22106:
			result.Message = "RelTagExisted"
			result.Description = "分组已存在"
			break
		case 22107:
			result.Message = "RelAwardPhoneRequired"
			result.Description = "该账号未通过手机认证"
			break
		case 22108:
			result.Message = "RelAwardIsBlocked"
			result.Description = "该账号处于封禁状态"
			break
		case 22109:
			result.Message = "RelAwardInsufficientFollower"
			result.Description = "该账号粉丝数不符合满足10000的标准"
			break
		case 22110:
			result.Message = "RelAwardGetFailed"
			result.Description = "获取粉丝成就奖励失败"
			break
		case 22111:
			result.Message = "RelAwardInfoFailed"
			result.Description = "获取成就信息失败"
			break
		case 22112:
			result.Message = "RelAwardInsufficientRank"
			result.Description = "很抱歉您的账号为非转正会员"
			break

		// kvo
		case 23001:
			result.Message = "KvoTimestampErr"
			result.Description = "时间戳不合法"
			break
		case 23002:
			result.Message = "KvoCheckSumErr"
			result.Description = "checksum不合法"
			break
		case 23003:
			result.Message = "KvoDataOverLimit"
			result.Description = "数据超过限制"
			break
		case 23004:
			result.Message = "KvoNotModified"
			result.Description = "数据没有修改"
			break

		// credit
		case 25001:
			result.Message = "CreditLevelLimit"
			result.Description = "风纪委申请等级限制(会员等级≥3)"
			break
		case 25002:
			result.Message = "CreditIsVerify"
			result.Description = "风纪委申请没有实名认证"
			break
		case 25003:
			result.Message = "CreditIsBlock"
			result.Description = "风纪委申请90天内有封禁记录"
			break
		case 25004:
			result.Message = "CreditNoJoinCase"
			result.Description = "风纪委没有参与案件"
			break
		case 25005:
			result.Message = "CreditNotJury"
			result.Description = "风纪委不是风纪委成员"
			break
		case 25006:
			result.Message = "CreditJuryExpired"
			result.Description = "风纪委资格过期"
			break
		case 25007:
			result.Message = "CreditUnderVoteRate"
			result.Description = "风纪委低于投准率"
			break
		case 25008:
			result.Message = "CreditNoCase"
			result.Description = "风纪委没有案件"
			break
		case 25009:
			result.Message = "CreditCaseNotExist"
			result.Description = "风纪委案件不存在"
			break
		case 25010:
			result.Message = "CreditCaseLimit"
			result.Description = "风纪委没有权限查看案件"
			break
		case 25011:
			result.Message = "CreditVoteNotExist"
			result.Description = "风纪委投票类型错误"
			break
		case 25012:
			result.Message = "CreditNovote"
			result.Description = "风纪委不能投票(可能重复投票或case不存在或投票过期vote被job设置为"
			break
		case 25013:
			result.Message = "CreditNoApply"
			result.Description = "风纪委不能重复申请风纪委资格"
			break
		case 25014:
			result.Message = "CreditCaseMax"
			result.Description = "风纪委每日获取案件到达上限"
			break
		case 25015:
			result.Message = "CreditBlack"
			result.Description = "风纪委在黑名单内用户申请风纪委资格"
			break
		case 25016:
			result.Message = "CreditNoJuryNum"
			result.Description = "风纪委当日风纪委员名额已发完"
			break
		case 25101:
			result.Message = "CreditNoblock"
			result.Description = "劳改未被封禁 没有答题资格"
			break
		case 25102:
			result.Message = "CreditQsNumError"
			result.Description = "劳改问题总数不正确"
			break
		case 25103:
			result.Message = "CreditAnsNumError"
			result.Description = "劳改答案总数不正确"
			break
		case 25104:
			result.Message = "CreditForeverBlock"
			result.Description = "劳改 永久封禁 不能答题"
			break
		case 25105:
			result.Message = "CreditBlockNotExist"
			result.Description = "封禁信息不存在"
			break
		case 25106:
			result.Message = "CreditBlockExpired"
			result.Description = "封禁信息已过7天有效期"
			break
		case 25107:
			result.Message = "CreditAppealExisted"
			result.Description = "申诉已存在"
			break

		// dm
		case 36001:
			result.Message = "DMFilterIllegalType"
			result.Description = "不支持该屏蔽类型"
			break
		case 36002:
			result.Message = "DMFilterTooLong"
			result.Description = "屏蔽词超过上限啦（关键词50字，正则200字）"
			break
		case 36003:
			result.Message = "DMFilterOverMax"
			result.Description = "屏蔽规则超过条数限制"
			break
		case 36004:
			result.Message = "DMFitlerIllegalRegex"
			result.Description = "屏蔽规则正则格式不对"
			break
		case 36005:
			result.Message = "DMFilterExist"
			result.Description = "屏蔽规则已经存在"
			break
		case 36006:
			result.Message = "DMFilterIsEmpty"
			result.Description = "屏蔽规则不允许为空"
			break
		case 36007:
			result.Message = "DMAdvNotAllow"
			result.Description = "不允许购买"
			break
		case 36009:
			result.Message = "DMAdvConfirm"
			result.Description = "正在确认中"
			break
		case 36010:
			result.Message = "DMAdvBought"
			result.Description = "已购买"
			break
		case 36011:
			result.Message = "DMAdvNoFound"
			result.Description = "高级弹幕购买记录不存在"
			break
		case 36101:
			result.Message = "DMPADMNotOwner"
			result.Description = "别人的弹幕不可以申请弹幕保护~"
			break
		case 36102:
			result.Message = "DMPAUserLevel"
			result.Description = "目前仅限 lv4及以上的用户可以直接申请哦~"
			break
		case 36103:
			result.Message = "DMPAUserLimit"
			result.Description = "一个人一天最多只能申请保护100条哦"
			break
		case 36104:
			result.Message = "DMPADMLimit"
			result.Description = "该弹幕已经申请保护~"
			break
		case 36105:
			result.Message = "DMPADMProtected"
			result.Description = "本弹幕已经被保护了~"
			break
		case 36106:
			result.Message = "DMNotFound"
			result.Description = "该弹幕已被删除"
			break
		case 36107:
			result.Message = "DMPAFailed"
			result.Description = "申请失败"
			break
		case 36108:
			result.Message = "DMPoolLimit"
			result.Description = "弹幕池超过大小"
			break
		case 36201:
			result.Message = "DMReportNotExist"
			result.Description = "举报弹幕不存在"
			break
		case 36202:
			result.Message = "DMReportReasonTooLong"
			result.Description = "举报原因过长"
			break
		case 36203:
			result.Message = "DMReportReasonError"
			result.Description = "举报原因类型错误"
			break
		case 36204:
			result.Message = "DMReportExist"
			result.Description = "已举报"
			break
		case 36205:
			result.Message = "DMReportLimit"
			result.Description = "操作过于频繁，请稍后再试"
			break
		case 36301:
			result.Message = "DMRecallTimeout"
			result.Description = "撤回失败，弹幕发送已过2分钟"
			break
		case 36302:
			result.Message = "DMRecallDeleted"
			result.Description = "撤回失败，弹幕已经被删除或撤回"
			break
		case 36303:
			result.Message = "DMRecallLimit"
			result.Description = "撤回失败，今天撤回的机会已经用完"
			break
		case 36304:
			result.Message = "DMRecallError"
			result.Description = "撤回失败，服务器出错"
			break
		case 36401:
			result.Message = "DMAssistNo"
			result.Description = "不是协管"
			break
		case 36402:
			result.Message = "DMAssistLimit"
			result.Description = "操作次数不足"
			break
		case 36501:
			result.Message = "DMTransferSame"
			result.Description = "弹幕转移源cid和目标cid相等"
			break
		case 36502:
			result.Message = "DMTransferNotFound"
			result.Description = "弹幕转移cid不存在"
			break
		case 36503:
			result.Message = "DMTransferNotBelong"
			result.Description = "弹幕转移cid不属于该用户"
			break
		case 36504:
			result.Message = "DMTransferRepet"
			result.Description = "弹幕转移任务重复"
			break
		case 36601:
			result.Message = "DMActSilence"
			result.Description = "弹幕点赞被禁言"
			break
		case 36700:
			result.Message = "DMUpgrading"
			result.Description = "系统升级中"
			break
		case 36701:
			result.Message = "DMMsgIlleagel"
			result.Description = "弹幕包含被禁止的内容"
			break
		case 36702:
			result.Message = "DMMsgTooLong"
			result.Description = "您的弹幕长度大于100"
			break
		case 36703:
			result.Message = "DMMsgPubTooFast"
			result.Description = "您发送弹幕的频率过快"
			break
		case 36704:
			result.Message = "DMArchiveIlleagel"
			result.Description = "禁止向未审核的视频发送弹幕"
			break
		case 36705:
			result.Message = "DMMsgNoPubPerm"
			result.Description = "您的等级不足，不能发送弹幕"
			break
		case 36706:
			result.Message = "DMMsgNoPubTopPerm"
			result.Description = "您的等级不足，不能发送顶端弹幕"
			break
		case 36707:
			result.Message = "DMMsgNoPubBottomPerm"
			result.Description = "您的等级不足，不能发送底端弹幕"
			break
		case 36708:
			result.Message = "DMMsgNoColorPerm"
			result.Description = "您的等级不足，不能发送彩色弹幕"
			break
		case 36709:
			result.Message = "DMMsgNoPubAdvancePerm"
			result.Description = "您的等级不足，不能发送高级弹幕"
			break
		case 36710:
			result.Message = "DMMsgNoPubStylePerm"
			result.Description = "您的权限不足，不能发送这种样式的弹幕"
			break
		case 36711:
			result.Message = "DMForbidPost"
			result.Description = "该视频禁止发送弹幕"
			break
		case 36712:
			result.Message = "DMMsgTooLongLevel1"
			result.Description = "level 1用户发送弹幕的最大长度为20"
			break
		case 36713:
			result.Message = "DMNotpayForPost"
			result.Description = "稿件未付费，不能发送弹幕"
			break
		case 36714:
			result.Message = "DMProgressTooBig"
			result.Description = "弹幕发送时间不合法"
			break
		case 36715:
			result.Message = "DMAssistOpToMuch"
			result.Description = "当日操作数量超过上限，请明天再试"
			break
		case 36800:
			result.Message = "DMTaskRegexTooLong"
			result.Description = "任务正则过长"
			break
		case 36801:
			result.Message = "DMTaskRegexIllegal"
			result.Description = "任务正则不合法"
			break

		// article
		case 37001:
			result.Message = "ArtLikeDupErr"
			result.Description = "重复点赞"
			break
		case 37002:
			result.Message = "ArtCancelLikeErr"
			result.Description = "取消点赞失败 用户未点赞"
			break
		case 37003:
			result.Message = "ArtDislikeDupErr"
			result.Description = "重复不喜欢"
			break
		case 37004:
			result.Message = "ArtCancelDislikeErr"
			result.Description = "取消不喜欢失败 用户未不喜欢"
			break
		case 37101:
			result.Message = "ArtCreationNoPrivilege"
			result.Description = "创作中心:用户没有权限发文章"
			break
		case 37102:
			result.Message = "ArtCreationStateErr"
			result.Description = "创作中心:文章状态错误"
			break
		case 37103:
			result.Message = "ArtCreationIDErr"
			result.Description = "创作中心:文章ID错误"
			break
		case 37104:
			result.Message = "ArtCreationMIDErr"
			result.Description = "创作中心:非文章作者"
			break
		case 37105:
			result.Message = "ArtCreationDelPendingErr"
			result.Description = "创作中心:审核中的文章不能删除"
			break
		case 37106:
			result.Message = "ArtCreationDraftFull"
			result.Description = "创作中心:草稿数已达最大上限"
			break
		case 37107:
			result.Message = "ArtCreationTplErr"
			result.Description = "创作中心:模板和图片数量不匹配"
			break
		case 37108:
			result.Message = "ArtCreationDraftDeleted"
			result.Description = "创作中心:草稿已被删除，不可编辑"
			break
		case 37109:
			result.Message = "ArtCreationArticleFull"
			result.Description = "创作中心：当日投稿数量已到达上限"
			break
		case 37200:
			result.Message = "ArtUserDisabled"
			result.Description = "用户被封禁 无法操作"
			break
		case 37300:
			result.Message = "ArtNoCategory"
			result.Description = "文章分区错误"
			break
		case 37400:
			result.Message = "ArtApplyPass"
			result.Description = "申请已通过"
			break
		case 37401:
			result.Message = "ArtApplyReject"
			result.Description = "已经申请处于冷冻期"
			break
		case 37402:
			result.Message = "ArtApplySubmit"
			result.Description = "已经申请待审"
			break
		case 37403:
			result.Message = "ArtApplyClose"
			result.Description = "关闭申请"
			break
		case 37404:
			result.Message = "ArtApplyFull"
			result.Description = "今日申请名额已满"
			break
		case 37405:
			result.Message = "ArtApplyVerify"
			result.Description = "用户未实名认证"
			break
		case 37406:
			result.Message = "ArtApplyForbid"
			result.Description = "用户已被封禁"
			break
		case 37407:
			result.Message = "ArtApplyPhone"
			result.Description = "用户未绑定手机"
			break
		case 37408:
			result.Message = "ArtApplyPhoneVirtual"
			result.Description = "绑定的手机号是虚拟号码"
			break
		case 37409:
			result.Message = "ArtCannotEditErr"
			result.Description = "文章不能被编辑"
			break
		case 37410:
			result.Message = "ArtAuthorReject"
			result.Description = "申请被拒绝"
			break
		case 37411:
			result.Message = "ArtNoActivity"
			result.Description = "活动未开始"
			break
		case 37412:
			result.Message = "ArtMaxListErr"
			result.Description = "达到文集上限 无法再增加新文集"
			break
		case 37413:
			result.Message = "ArtListNameErr"
			result.Description = "文集标题不合法"
			break
		case 37414:
			result.Message = "ArtArtAddListErr"
			result.Description = "文章已存在于其他文集或者文章不存在"
			break
		case 37415:
			result.Message = "ArtAddListLimitErr"
			result.Description = "达到文集文章数量上限 无法再增加新文章"
			break
		case 37416:
			result.Message = "ArtPermClosedErr"
			result.Description = "操作失败，你的专栏权限已被关闭"
			break
		case 37417:
			result.Message = "ArtLevelFailedErr"
			result.Description = "等级未达到要求"
			break
		case 37418:
			result.Message = "ArtMediaExistedErr"
			result.Description = "已经存在长评了"
			break
		case 37419:
			result.Message = "ArtUpdateFullErr"
			result.Description = "重复编辑次数已用完"
			break

		// Assist
		case 24001:
			result.Message = "AssistAlreadyExist"
			result.Description = "当前协管关系已存在,不能重复添加一个用户为另一个用户的协管"
			break
		case 24002:
			result.Message = "AssistNotExist"
			result.Description = "不存在当前的协管关系"
			break
		case 24003:
			result.Message = "AssistUserNotExist"
			result.Description = "用户不存在"
			break
		case 24004:
			result.Message = "AssistUserNotFollowUp"
			result.Description = "没有关注UP"
			break
		case 24005:
			result.Message = "AssistUserNotRealAuth"
			result.Description = "协管尚未实名认证"
			break
		case 24006:
			result.Message = "AssistOverMaxLimit"
			result.Description = "禁止再添加协管,当前用户的协管数已经超出最大值"
			break
		case 24007:
			result.Message = "AssistAlreadyCancel"
			result.Description = "已经撤销过协管操作日志"
			break
		case 24008:
			result.Message = "AssistNotFollowUp"
			result.Description = "当前用户尚未是up主的粉丝，所以不能添加为协管"
			break
		case 24009:
			result.Message = "AssistForbidType"
			result.Description = "禁止添加不允许的业务类型"
			break
		case 24010:
			result.Message = "AssistForbidAction"
			result.Description = "禁止添加不允许的操作类型"
			break
		case 24011:
			result.Message = "AssistLogNotExist"
			result.Description = "不存在的操作记录"
			break
		case 24012:
			result.Message = "AssistLogOverMaxLimit"
			result.Description = "禁止再添加操作日志,当前协管操作数已超过当天阈值"
			break
		case 24013:
			result.Message = "AssistOverMaxLimitDailyAddAll"
			result.Description = "禁止再添加骑士,当天已经添加所有用户为骑士超过最大次数"
			break
		case 24014:
			result.Message = "AssistOverMaxLimitDailyAddSame"
			result.Description = "禁止再添加骑士,当天已经添加同一个用户为骑士超过最大次数"
			break

		// Member
		case 40001:
			result.Message = "UpdateBirthdayFormat"
			result.Description = "出生日期格式不正确"
			break
		case 40002:
			result.Message = "UpdateUnameSensitive"
			result.Description = "昵称包含敏感词"
			break
		case 40003:
			result.Message = "UpdateSexError"
			result.Description = "请选择正常的性别"
			break
		case 40004:
			result.Message = "UpdateUnameFormat"
			result.Description = "昵称不可包含除-和_以外的特殊字符"
			break
		case 40005:
			result.Message = "UpdateUnameTooLong"
			result.Description = "昵称过长，不能修改"
			break
		case 40006:
			result.Message = "UpdateUnameTooShort"
			result.Description = "昵称过短，不能修改"
			break
		case 40007:
			result.Message = "UpdateUnameMoneyIsNot"
			result.Description = "硬币不足,改昵称需要6个硬币"
			break
		case 40008:
			result.Message = "UpdateUnameHadVerified"
			result.Description = "已过实名验证，不能修改"
			break
		case 40009:
			result.Message = "UpdateUnameHadLocked"
			result.Description = "昵称已锁定不能修改"
			break
		case 40010:
			result.Message = "UpdateUnameHadOfficial"
			result.Description = "认证账号不得随意修改昵称，如有需要请联系客服娘~"
			break

		case 40012:
			result.Message = "UpdateFaceFormat"
			result.Description = "头像格式错误，允许：png/jpg/jpeg/jp2"
			break
		case 40013:
			result.Message = "UpdateFaceSize"
			result.Description = "头像超过限制的大小，允许2M"
			break
		case 40014:
			result.Message = "UpdateUnameRepeated"
			result.Description = "昵称已存在"
			break
		case 40015:
			result.Message = "MemberSignSensitive"
			result.Description = "签名包含敏感词"
			break
		case 40016:
			result.Message = "MemberPhoneRequired"
			result.Description = "根据国家实名制认证的相关要求，需要绑定手机号"
			break
		case 40017:
			result.Message = "MemberRealPhoneRequired"
			result.Description = "根据国家实名制认证的相关要求，需要绑定非虚拟手机号"
			break
		case 40021:
			result.Message = "MemberSignHasEmoji"
			result.Description = "签名不能包含表情图片"
			break
		case 40022:
			result.Message = "MemberSignOverLimit"
			result.Description = "签名最多支持70个字"
			break
		case 40043:
			result.Message = "BirthdayInfoIsNull"
			result.Description = "该用户没有生日信息 // 答题系统使用"
			break
		case 40050:
			result.Message = "MemberUpdate"
			result.Description = "系统维护中"
			break
		case 40051:
			result.Message = "MemberBlocked"
			result.Description = "用户被封禁"
			break
		case 40052:
			result.Message = "MemberNameFormatErr"
			result.Description = "用户名不合法"
			break
		case 40053:
			result.Message = "MemberNameOverLimit"
			result.Description = "用户名长度超过限制"
			break
		case 40054:
			result.Message = "MemberNameUnmodify"
			result.Description = "用户名未修改"
			break
		case 40055:
			result.Message = "MemberNameHasEmoji"
			result.Description = "用户名包含emoji"
			break
		case 40056:
			result.Message = "MemberNameCoinErr"
			result.Description = "扣除硬币失败"
			break
		case 40058:
			result.Message = "MemberUnRealName"
			result.Description = "用户名未实名"
			break
		case 40059:
			result.Message = "MemberCerted"
			result.Description = "用户名包含敏感词"
			break
		case 40060:
			result.Message = "MemberOverLimit"
			result.Description = "批量请求超过限制"
			break
		case 40061:
			result.Message = "MemberNotExist"
			result.Description = "用户不存在"
			break
		case 40071:
			result.Message = "MemberUpdateBirthdayFaild"
			result.Description = "修改生日失败"
			break
		case 40072:
			result.Message = "MemberBirthdayNotAllow"
			result.Description = "生日信息不合法"
			break
		case 40073:
			result.Message = "MemberBirthdayInfoIsNull"
			result.Description = "该用户没有生日信息"
			break
		case 40080:
			result.Message = "MemberTagsOverLen"
			result.Description = "用户 Tag 不合法"
			break
		case 40081:
			result.Message = "MemberTagsOverCount"
			result.Description = "用户 Tag 不合法"
			break
		case 40083:
			result.Message = "SubmitOfficialDocFailed"
			result.Description = "提交官方认证请求失败"
			break
		case 40084:
			result.Message = "NoOfficialDoc"
			result.Description = "未提交过官方认证请求"
			break
		case 40085:
			result.Message = "SearchMidOverLimit"
			result.Description = "Mid查询数量过大"
			break

		// Answer
		case 41001:
			result.Message = "QuestionStrNotAllow"
			result.Description = "分院帽题目不合法"
			break
		case 41002:
			result.Message = "QuestionAnsNotAllow"
			result.Description = "分院帽题目答案不合法"
			break
		case 41003:
			result.Message = "QuestionTipsNotAllow"
			result.Description = "分院帽题目提示不合法"
			break
		case 41004:
			result.Message = "QuestionTypeNotAllow"
			result.Description = "分院帽题目类型不合法"
			break
		case 41010:
			result.Message = "AnswerDenied"
			result.Description = "用户答题非法访问"
			break
		case 41011:
			result.Message = "AnswerTimeExpire"
			result.Description = "用户答题时间已超时"
			break
		case 41012:
			result.Message = "AnswerIdsErr"
			result.Description = "用户答题提交题目id不合法"
			break
		case 41013:
			result.Message = "AnswerQsNumErr"
			result.Description = "用户答题提交题目数量不合法"
			break
		case 41014:
			result.Message = "AnswerBlock"
			result.Description = "用户自选题提交过快（2分钟内）被封禁12小时"
			break
		case 41016:
			result.Message = "AnswerSorceZero"
			result.Description = "该用户答题分数为0"
			break
		case 41017:
			result.Message = "AnswerGeetestErr"
			result.Description = "答题验证码错误"
			break
		case 41018:
			result.Message = "AnswerFormalFailed"
			result.Description = "答题转正失败"
			break
		case 41020:
			result.Message = "AnswerBasePassed"
			result.Description = "用户基础题已通过"
			break
		case 41021:
			result.Message = "AnswerBaseNotPassed"
			result.Description = "用户基础题未通过"
			break
		case 41023:
			result.Message = "AnswerHistoryNotFound"
			result.Description = "用户答题记录不存在"
			break
		case 41024:
			result.Message = "AnswerMidCacheQidsErr"
			result.Description = "获取用户题目id缓存异常"
			break
		case 41025:
			result.Message = "AnswerQidDiffRequestErr"
			result.Description = "用户答题提交题目ID和实际用户的答题id不一致"
			break
		case 41026:
			result.Message = "AnswerMidDBQueErr"
			result.Description = "获取用户DB题目信息异常"
			break
		case 41027:
			result.Message = "AnswerCheckFaild"
			result.Description = "基础题检查不通过"
			break
		case 41031:
			result.Message = "AnswerProNoPass"
			result.Description = "自选题未通过"
			break
		case 41050:
			result.Message = "AnswerCaptchaPassed"
			result.Description = "用户答题验证码已通过"
			break
		case 41051:
			result.Message = "AnswerCaptchaNoPassed"
			result.Description = "用户答题验证码未通过"
			break
		case 41052:
			result.Message = "AnswerTypeIDsErr"
			result.Description = "用户题目类型不合法"
			break
		case 41053:
			result.Message = "AnswerGeetestVaErr"
			result.Description = "极验验证异常"
			break
		case 41054:
			result.Message = "AnswerExtraHadPass"
			result.Description = "基础附加题已通过"
			break
		case 41055:
			result.Message = "AnswerExtraNoPass"
			result.Description = "基础附加题未通过"
			break
		case 41090:
			result.Message = "AnswerAccCallErr"
			result.Description = "调用账号异常"
			break
		case 41091:
			result.Message = "AnswerNeedBindTel"
			result.Description = "答题需要绑定手机"
			break

		// bfs upload
		case 42001:
			result.Message = "BfsUploadCodeErr"
			result.Description = "bfs响应code错误"
			break
		case 42002:
			result.Message = "BfsUploadStatusErr"
			result.Description = "返回状态错误（非常规捕捉）"
			break
		case 42400:
			result.Message = "BfsRequestErr"
			result.Description = "bfs参数错误"
			break
		case 42401:
			result.Message = "BfsUploadAuthErr"
			result.Description = "上传验证错误"
			break
		case 42404:
			result.Message = "BfsUplaodBucketNotExist"
			result.Description = "bucket不存在"
			break
		case 42503:
			result.Message = "BfsUploadServiceUnavailable"
			result.Description = "服务不可用"
			break
		case 42601:
			result.Message = "BfsUploadFileTooLarge"
			result.Description = "上传的文件太大"
			break
		case 42602:
			result.Message = "BfsUploadFilePixelError"
			result.Description = "不能获取图片的像素信息"
			break
		case 42603:
			result.Message = "BfsUploadFilePixelWidthIllegal"
			result.Description = "宽像素不合法"
			break
		case 42604:
			result.Message = "BfsUploadFilePixelHeightIllegal"
			result.Description = "高像素不合法"
			break
		case 42605:
			result.Message = "BfsUploadFilePixelAspectRatioIllegal"
			result.Description = "像素宽高比不合法"
			break
		case 42606:
			result.Message = "BfsUploadFileContentTypeIllegal"
			result.Description = "文件类型不合法"
			break

		// remote login
		case 43001:
			result.Message = "RemoteLoginStatusQueryError"
			result.Description = "查询失败"
			break
		case 43002:
			result.Message = "RemoteLoginFeedBackError"
			result.Description = "反馈失败"
			break
		case 43003:
			result.Message = "RemoteLoginWarnCloseError"
			result.Description = "关闭失败"
			break

		// Spy
		case 50001:
			result.Message = "SpyEventNotExist"
			result.Description = "反作弊事件类型不存在"
			break
		case 50002:
			result.Message = "SpyServiceNotExist"
			result.Description = "反作弊服务不存在"
			break
		case 50003:
			result.Message = "SpyFactorNotExist"
			result.Description = "反作弊因子不存在"
			break
		case 50004:
			result.Message = "SpySettingUnknown"
			result.Description = "反作弊配置类型不存在"
			break
		case 50005:
			result.Message = "SpySettingValTypeError"
			result.Description = "反作弊配置值类型错误"
			break
		case 50006:
			result.Message = "SpySettingValueOutOfRange"
			result.Description = "反作弊配置值超出范围"
			break
		case 50007:
			result.Message = "SpyRuleNotExist"
			result.Description = "反作弊规则不存在"
			break
		case 50008:
			result.Message = "SpyRulesNotMatch"
			result.Description = "反作弊规则不匹配"
			break

		// filter-service and filter-job
		case 52001:
			result.Message = "FilterHitLimitBlack"
			result.Description = "命中黑名单"
			break
		case 52002:
			result.Message = "FilterHitRubLimit"
			result.Description = "超过发送次数"
			break
		case 52003:
			result.Message = "FilterLimitTypeNotExist"
			result.Description = "限制类型不存"
			break
		case 52004:
			result.Message = "FilterLimitContentNotExist"
			result.Description = "限制关键词不存在"
			break
		case 52005:
			result.Message = "FilterHitStrictLimit"
			result.Description = "命中严格限制"
			break
		case 52006:
			result.Message = "FilterIllegalRegexp"
			result.Description = "非法正则"
			break
		case 52007:
			result.Message = "FilterIllegalArea"
			result.Description = "业务不存在"
			break
		case 52010:
			result.Message = "FilterWhiteSampleHit"
			result.Description = "敏感词可能误杀较大"
			break
		case 52011:
			result.Message = "FilterBlackSampleHit"
			result.Description = "敏感词导致高危内容失效"
			break
		case 52012:
			result.Message = "FilterDuplicateContent"
			result.Description = "已存在相同内容敏感词/白名单"
			break
		case 52013:
			result.Message = "FilterRegexpError1"
			result.Description = "含有.*容易引起误伤，请换用.{0,10}"
			break
		case 52014:
			result.Message = "FilterRegexpError2"
			result.Description = "含有||容易引起误伤"
			break
		case 52020:
			result.Message = "FilterInvalidAreaGroupName"
			result.Description = "不合法的业务组命名"
			break
		case 52021:
			result.Message = "FilterDuplicateAreaGroup"
			result.Description = "业务组重复"
			break
		case 52022:
			result.Message = "FilterAreaGroupNotFound"
			result.Description = "业务组不存在"
			break
		case 52023:
			result.Message = "FilterInvalidAreaShowName"
			result.Description = "不合法的业务模块命名"
			break
		case 52024:
			result.Message = "FilterInvalidAreaName"
			result.Description = "不合法的业务id命名"
			break
		case 52025:
			result.Message = "FilterDuplicateArea"
			result.Description = "业务重复"
			break
		case 52026:
			result.Message = "FilterInvalidArea"
			result.Description = "业务不存在"
			break
		case 52027:
			result.Message = "FilterInvalidAIWhiteUID"
			result.Description = "AI过滤白名单重复添加"
			break

		// space channel
		case 53001:
			result.Message = "ChNameToLong"
			result.Description = "频道名字数超过限制啦"
			break
		case 53002:
			result.Message = "ChIntroToLong"
			result.Description = "频道简介字数超过限制啦"
			break
		case 53003:
			result.Message = "ChMaxArcCount"
			result.Description = "本频道里的视频已经满啦"
			break
		case 53004:
			result.Message = "ChMaxCount"
			result.Description = "你创建的频道已经满额了哦"
			break
		case 53005:
			result.Message = "ChFakeAid"
			result.Description = "频道内有失效视频了哦"
			break
		case 53006:
			result.Message = "ChAidsExist"
			result.Description = "你提交的视频已失效或者频道里已经有了哦"
			break
		case 53007:
			result.Message = "ChNameExist"
			result.Description = "频道名称已经存在了哦"
			break
		case 53008:
			result.Message = "ChNoArcs"
			result.Description = "频道内没有视频"
			break
		case 53009:
			result.Message = "ChNoArc"
			result.Description = "频道内没有该视频"
			break
		case 53010:
			result.Message = "ChNameBanned"
			result.Description = "频道名称有敏感词，请重新编写"
			break
		case 53011:
			result.Message = "ChIntroBanned"
			result.Description = "频道简介有敏感词，请重新编写"
			break
		case 53012:
			result.Message = "SpaceNoShop"
			result.Description = "非营业中商户号"
			break
		case 53013:
			result.Message = "SpaceNoPrivacy"
			result.Description = "用户隐私设置未公开"
			break
		case 53014:
			result.Message = "SpaceFakeAid"
			result.Description = "该稿件已失效"
			break
		case 53015:
			result.Message = "TopReasonLong"
			result.Description = "置顶理由字数超过限制啦"
			break
		case 53016:
			result.Message = "SpaceNoTopArc"
			result.Description = "没有置顶视频"
			break
		case 53017:
			result.Message = "SpaceNotAuthor"
			result.Description = "只能操作自己的稿件"
			break
		case 53018:
			result.Message = "SpaceTextBanned"
			result.Description = "提交文本有敏感词"
			break
		case 53019:
			result.Message = "SpaceMpMaxCount"
			result.Description = "代表作已达上限"
			break
		case 53020:
			result.Message = "SpaceMpExist"
			result.Description = "代表作内已有该视频"
			break
		case 53021:
			result.Message = "SpaceMpNoArc"
			result.Description = "代表作内没有该视频"
			break

		// search
		case 54001:
			result.Message = "SearchArchiveCheckFailed"
			result.Description = "搜索稿件管理失败"
			break
		case 54002:
			result.Message = "SearchArticleDataFailed"
			result.Description = "搜索专栏数据失败"
			break
		case 54003:
			result.Message = "SearchReplyRecordFailed"
			result.Description = "搜索个人中心评论记录数据失败"
			break
		case 54010:
			result.Message = "SearchBlockedListFailed"
			result.Description = "搜索风纪委封禁列表失败"
			break
		case 54011:
			result.Message = "SearchBlockedPublishFailed"
			result.Description = "搜索公告列表失败"
			break
		case 54012:
			result.Message = "SearchBlockedCaseFailed"
			result.Description = "搜索案件列表失败"
			break
		case 54013:
			result.Message = "SearchBlockedJuryFailed"
			result.Description = "搜索委员列表失败"
			break
		case 54014:
			result.Message = "SearchBlockedOpinionFailed"
			result.Description = "搜索观点列表失败"
			break
		case 54015:
			result.Message = "SearchWorkflowGroupFailed"
			result.Description = "工作流获取反馈列表失败"
			break
		case 54016:
			result.Message = "SearchWorkflowChaFailed"
			result.Description = "工作流获取用户工单失败"
			break
		case 54017:
			result.Message = "SearchWorkflowTagFailed"
			result.Description = "工作流获取Tag列表失败"
			break
		case 54018:
			result.Message = "SearchWorkflowLogFailed"
			result.Description = "工作流获取日志列表失败"
			break
		case 54019:
			result.Message = "SearchWorkflowCommonFaild"
			result.Description = "工作流举报列表获取失败"
			break
		case 54900:
			result.Message = "SearchUpdateIndexFailed"
			result.Description = "更新索引失败"
			break
		case 54020:
			result.Message = "SearchDmFailed"
			result.Description = "弹幕列表获取失败"
			break
		case 54021:
			result.Message = "SearchVideoFailed"
			result.Description = "弹幕列表获取失败"
			break
		case 54022:
			result.Message = "SearchMusicSongsFailed"
			result.Description = "音乐审核列表获取失败"
			break
		case 54023:
			result.Message = "SearchKpiPointFailed"
			result.Description = "搜索kpi评分列表失败"
			break
		case 54024:
			result.Message = "SearchWorkflowFeedbackFailed"
			result.Description = "工作流获取feedback列表失败"
			break
		case 54025:
			result.Message = "SearchUNameFailed"
			result.Description = "用户昵称查询失败"
			break
		case 54026:
			result.Message = "SearchDmmonitorFailed"
			result.Description = "弹幕监控查询失败"
			break
		case 54027:
			result.Message = "SearchPgcMediaFailed"
			result.Description = "pgc影视查询失败"
			break
		case 54028:
			result.Message = "SearchFeedbackFailed"
			result.Description = "用户反馈查询失败"
			break
		case 54029:
			result.Message = "SearchFeedbackReplyFailed"
			result.Description = "用户反馈报告查询失败"
			break
		case 54030:
			result.Message = "SearchLogAuditFailed"
			result.Description = "审核日志查询失败"
			break
		case 54031:
			result.Message = "SearchLogAuditOidFailed"
			result.Description = "根据oid查询审核日志查询失败"
			break
		case 54032:
			result.Message = "SearchLogUserActionFailed"
			result.Description = "用户操作日志查询失败"
			break
		case 54033:
			result.Message = "SearchUserApplyReviewsFailed"
			result.Description = "用户头像挂件查询失败"
			break

		case 54901:
			result.Message = "SearchBusinessFailed"
			result.Description = "后台接口Business报错"
			break
		case 54902:
			result.Message = "SearchAppidFailed"
			result.Description = "后台接口Appid报错"
			break
		case 54903:
			result.Message = "SearchBusinessExistErr"
			result.Description = "该业务已经存在"
			break
		case 54904:
			result.Message = "SearchAssetExistErr"
			result.Description = "该数据源已经存在"
			break
		case 54905:
			result.Message = "SearchAppExistErr"
			result.Description = "该应用已经存在"
			break

		// figure 信用分服务
		case 55001:
			result.Message = "FigureNotFound"
			result.Description = "未找到用户信用分"
			break

		// workflow 工作流
		case 56001:
			result.Message = "WkfGroupNotFound"
			result.Description = "未找到工单"
			break
		case 56002:
			result.Message = "WkfChallNotFound"
			result.Description = "未找到工单详情"
			break
		case 56201:
			result.Message = "WkfAppealNotUserOwned"
			result.Description = "只能查看或操作自己的申诉"
			break
		case 56202:
			result.Message = "WkfAppealNotAllowDegree"
			result.Description = "当前申诉状态不允许提交满意度"
			break
		case 56401:
			result.Message = "WkfBusinessNotFound"
			result.Description = "未找到业务id"
			break
		case 56402:
			result.Message = "WkfBusinessNotConsistent"
			result.Description = "工单业务id不一致"
			break
		case 56403:
			result.Message = "WkfTagNotFound"
			result.Description = "未找到tid数据"
			break
		case 56404:
			result.Message = "WkfBusinessCallbackConfigNotFound"
			result.Description = "未找到业务回调配置"
			break
		case 56405:
			result.Message = "WkfBanNotSupportBatchOperate"
			result.Description = "不支持批量封禁账号"
			break
		case 56406:
			result.Message = "WkfBidNotSupportPublicReferee"
			result.Description = "业务不支持移交众裁"
			break
		case 56407:
			result.Message = "WkfBidNotSupportQuerySource"
			result.Description = "业务不支持查询来源"
			break
		case 56408:
			result.Message = "WkfBidNotSupportQueryContentState"
			result.Description = "业务不支持查询内容状态"
			break
		case 56501:
			result.Message = "WkfSearchGroupFailed"
			result.Description = "es search工单失败"
			break
		case 56502:
			result.Message = "WkfSearchChallFailed"
			result.Description = "es search工单详情失败"
			break
		case 56503:
			result.Message = "WkfGetBlockInfoFailed"
			result.Description = "获取封禁信息失败"
			break
		case 56504:
			result.Message = "WkfSetPublicRefereeFailed"
			result.Description = "提交众裁失败"
			break

		// account common
		case 61000:
			result.Message = "UserLoginInvalid"
			result.Description = "使用登录状态访问了，并且登录状态无效，客服端可以／需要删除登录状态"
			break
		case 61001:
			result.Message = "UserCheckNoPhone"
			result.Description = "根据国家实名制认证的相关要求，您需要绑定手机号，才能继续进行操作"
			break
		case 61002:
			result.Message = "UserCheckInvalidPhone"
			result.Description = "根据国家实名制认证的相关要求，您需要换绑一个非170/171的手机号，才能继续进行操作"
			break

		// web-interface
		case 62001:
			result.Message = "ElecDenied"
			result.Description = "不需要展示充电信息"
			break
		case 62002:
			result.Message = "ArchiveDenied"
			result.Description = "稿件不可见"
			break
		case 62003:
			result.Message = "ArchivePass"
			result.Description = "稿件已审核通过，等待发布中"
			break
		case 62004:
			result.Message = "ArchiveChecking"
			result.Description = "视频正在审核中，请耐心等待～"
			break
		case 62005:
			result.Message = "ArchiveNotLogin"
			result.Description = "视频不见了？你可以试试登录！"
			break
		case 62006:
			result.Message = "HelpListError"
			result.Description = "智齿列表结果错误"
			break
		case 62007:
			result.Message = "HelpDetailError"
			result.Description = "智齿详情结果错误"
			break
		case 62008:
			result.Message = "HelpSearchError"
			result.Description = "智齿搜索结果错误"
			break
		case 62009:
			result.Message = "ArcAppealLimit"
			result.Description = "短时间内请勿重复投诉相同稿件"
			break

		// playlist
		case 63001:
			result.Message = "PlNameTooLong"
			result.Description = "播单标题超出最大限制"
			break
		case 63002:
			result.Message = "PlDescTooLong"
			result.Description = "播单简介超出最大限制"
			break
		case 63003:
			result.Message = "PlMaxCount"
			result.Description = "播单个数超出最大限制"
			break
		case 63004:
			result.Message = "PlCanNotDelDefault"
			result.Description = "不能删除默认播单"
			break
		case 63005:
			result.Message = "PlExist"
			result.Description = "已经存在相同标题的播单"
			break
		case 63006:
			result.Message = "PlNotExist"
			result.Description = "播单无法访问"
			break
		case 63007:
			result.Message = "PlAlreadyDel"
			result.Description = "播单已经删除"
			break
		case 63008:
			result.Message = "PlDenied"
			result.Description = "播单暂未开放"
			break
		case 63009:
			result.Message = "PlVideoOverflow"
			result.Description = "播单内视频个数超出最大限制"
			break
		case 63010:
			result.Message = "PlVideoExist"
			result.Description = "视频已经添加进此播单"
			break
		case 63011:
			result.Message = "PlVideoAlreadyDel"
			result.Description = "视频已经不属于此播单"
			break
		case 63012:
			result.Message = "PlSortOverflow"
			result.Description = "播单内视频排序不生效"
			break
		case 63013:
			result.Message = "PlFavExist"
			result.Description = "播单已经收藏"
			break
		case 63014:
			result.Message = "PlFavAlreadyDel"
			result.Description = "播单未收藏"
			break
		case 63015:
			result.Message = "PlNotUser"
			result.Description = "非创建者不能修改此播单"
			break

		// usersuit
		case 64001:
			result.Message = "UsersuitInviteLevelLow"
			result.Description = "你还不满足购买激活码的条件哦，升级到Lv5再来吧~"
			break
		case 64002:
			result.Message = "UsersuitInviteReachCurrentMonthLimit"
			result.Description = "当月邀请码申请数达到上限"
			break
		case 64003:
			result.Message = "UsersuitInviteAlreadyFormal"
			result.Description = "已经转正不能使用邀请码"
			break
		case 64004:
			result.Message = "UsersuitInviteCodeNotExists"
			result.Description = "邀请码不存在"
			break
		case 64005:
			result.Message = "UsersuitInviteCodeUsed"
			result.Description = "邀请码已使用"
			break
		case 64006:
			result.Message = "UsersuitInviteCodeExpired"
			result.Description = "邀请码已过期"
			break
		case 64007:
			result.Message = "UsersuitInviteReachMaxGeneLimit"
			result.Description = "超过批量生成邀请码上限（最多1000个）"
			break

		// pendant 挂件相关
		case 64101:
			result.Message = "PendantNotFound"
			result.Description = "挂件不存在"
			break
		case 64102:
			result.Message = "PendantCanNotBuy"
			result.Description = "大会员挂件不能购买"
			break
		case 64103:
			result.Message = "PendantAlreadyGet"
			result.Description = "大会员挂件已领取过"
			break
		case 64104:
			result.Message = "PendantGetVIPErr"
			result.Description = "获取大会员信息错误"
			break
		case 64105:
			result.Message = "PendantPayErr"
			result.Description = "订单接口错误"
			break
		case 64106:
			result.Message = "PendantOrderNotFound"
			result.Description = "订单不存在"
			break
		case 64107:
			result.Message = "PendantPackageNotFound"
			result.Description = "背包里无此挂件"
			break
		case 64108:
			result.Message = "PendantPayTypeErr"
			result.Description = "该挂件无此种支付方式"
			break
		case 64109:
			result.Message = "PendantVIPOverdue"
			result.Description = "大会员过期"
			break
		case 64110:
			result.Message = "PendantAboveSendMaxLimit"
			result.Description = "超过批量发放挂件上限 (最多1000个)"
			break

		// usersuit medal 勋章
		case 64201:
			result.Message = "MedalNotFound"
			result.Description = "勋章不存在"
			break
		case 64202:
			result.Message = "MedalNotGet"
			result.Description = "不拥有该勋章"
			break
		case 64203:
			result.Message = "MedalHasGet"
			result.Description = "已拥有该勋章"
			break

		// thumbup
		case 65001:
			result.Message = "ThumbupBusinessBlankErr"
			result.Description = "业务id不存在"
			break
		case 65002:
			result.Message = "ThumbupOriginErr"
			result.Description = "origin id 错误"
			break
		case 65003:
			result.Message = "ThumbupBusinessErr"
			result.Description = "未开通此业务"
			break
		case 65004:
			result.Message = "ThumbupCancelLikeErr"
			result.Description = "取消点赞失败 用户未点赞"
			break
		case 65005:
			result.Message = "ThumbupCancelDislikeErr"
			result.Description = "取消不喜欢失败 用户未不喜欢"
			break
		case 65006:
			result.Message = "ThumbupDupLikeErr"
			result.Description = "重复点赞"
			break
		case 65007:
			result.Message = "ThumbupDupDislikeErr"
			result.Description = "重复点踩"
			break

		// sms
		case 66001:
			result.Message = "SmsTemplateNotExist"
			result.Description = "模版不存在"
			break
		case 66002:
			result.Message = "SmsTemplateParamNotEnough"
			result.Description = "模版参数不足"
			break
		case 66003:
			result.Message = "SmsTemplateCodeExist"
			result.Description = "模版code已存在"
			break
		case 66004:
			result.Message = "SmsTemplateParamIllegal"
			result.Description = "模版参数值不合法"
			break
		case 66005:
			result.Message = "SmsTemplateModifyForbind"
			result.Description = "修改已审核的模版必须提供approver3"
			break
		case 66006:
			result.Message = "SmsTemplateNotAct"
			result.Description = "模版不是营销短信"
			break
		case 66023:
			result.Message = "SmsSendBatchOverLimit"
			result.Description = "批量发送超出限制"
			break
		case 66024:
			result.Message = "SmsSendBothMidAndMobile"
			result.Description = "mid,mobile只能传一个参数"
			break
		case 66031:
			result.Message = "SmsMobilePatternErr"
			result.Description = "手机号码格式不正确"
			break

		// growup admin and interface and job
		case 68001:
			result.Message = "GrowupDisabled"
			result.Description = "up主在黑名单"
			break
		case 68002:
			result.Message = "GrowupTagForbit"
			result.Description = "不允许操作该标签"
			break
		case 68003:
			result.Message = "GrowupNotExist"
			result.Description = "有不存在的UP主账号"
			break
		case 68004:
			result.Message = "GrowupAuthorityUserNotFound"
			result.Description = "用户在权限系统中不存在"
			break
		case 68005:
			result.Message = "GrowupTagAddForbit"
			result.Description = "该标签在本业务分区下已存在"
			break
		case 68006:
			result.Message = "GrowupAuthorityExist"
			result.Description = "用户名/任务组名/角色组名/权限点名已存在"
			break
		case 68007:
			result.Message = "GrowupBodyTooLarge"
			result.Description = "上传的文件太大"
			break
		case 68008:
			result.Message = "GrowupBodyNotExist"
			result.Description = "上传的文件无内容"
			break
		case 68009:
			result.Message = "GrowupGetTypeError"
			result.Description = "获取视频全部分区失败"
			break
		case 68010:
			result.Message = "GrowupGetActivityError"
			result.Description = "根据活动ID获取稿件失败"
			break
		case 68020:
			result.Message = "GrowupPriceErr"
			result.Description = "购买大会员价格错误"
			break
		case 68021:
			result.Message = "GrowupPriceNotEnough"
			result.Description = "购买大会员余额不足"
			break
		case 68022:
			result.Message = "GrowupBuyErr"
			result.Description = "激励兑换购买失败"
			break
		case 68023:
			result.Message = "GrowupGoodsNotExist"
			result.Description = "激励兑换商品不存在"
			break
		case 68024:
			result.Message = "GrowupGoodsTimeErr"
			result.Description = "不在兑换时间内"
			break
		case 68101:
			result.Message = "GrowupArchiveNotYours"
			result.Description = "稿件不属于此人"
			break
		case 68102:
			result.Message = "GrowupSubTidNotExist"
			result.Description = "此二级分类不存在"
			break
		case 68103:
			result.Message = "GrowupActivityCountNotEnough"
			result.Description = "活动列表不存在"
			break
		case 68104:
			result.Message = "GrowupRecommendUpNotExist"
			result.Description = "推荐UP主列表不存在"
			break
		case 68105:
			result.Message = "GrowupUpInfoNotExist"
			result.Description = "UP主信息不存在"
			break
		case 68106:
			result.Message = "GrowupRecommendUpInfoNotExist"
			result.Description = "推荐UP主信息不存在"
			break
		case 68107:
			result.Message = "GrowupTidNotExist"
			result.Description = "此一级分类不存在"
			break
		case 68201:
			result.Message = "GrowupSpecialAwardJoined"
			result.Description = "已参加过专项奖"
			break
		case 68202:
			result.Message = "GrowupSpecialAwardUnqualified"
			result.Description = "没有资格参加专项奖"
			break

		// Vip
		case 69001:
			result.Message = "VipBatchIDErr"
			result.Description = "大会员资源批次不存在"
			break
		case 69002:
			result.Message = "VipPoolIDErr"
			result.Description = "大会员资源池不存在"
			break
		case 69003:
			result.Message = "VipBusinessErr"
			result.Description = "大会员资源与业务方不匹配"
			break
		case 69004:
			result.Message = "VipAppkeyExitErr"
			result.Description = "appkey已存在"
			break
		case 69005:
			result.Message = "VipChangeHistoryErr"
			result.Description = "该订单已消费"
			break
		case 69006:
			result.Message = "VipBatchNotEnoughErr"
			result.Description = "资源池数量不足"
			break
		case 69007:
			result.Message = "VipBatchTTLErr"
			result.Description = "资源池有效时间无效"
			break
		case 69008:
			result.Message = "VipMidErr"
			result.Description = "mid无效"
			break
		case 69009:
			result.Message = "VipRemarkErr"
			result.Description = "请填写备注"
			break
		case 69010:
			result.Message = "VipRelationIDErr"
			result.Description = "请填写关联ID"
			break
		case 69011:
			result.Message = "VipDaysErr"
			result.Description = "大会员天数不能少于一天"
			break
		case 69012:
			result.Message = "VipBusinessNameExitErr"
			result.Description = "业务方名称已存在"
			break
		case 69013:
			result.Message = "VipConfigNotExitErr"
			result.Description = "配置不存在"
			break
		case 69014:
			result.Message = "VipUpdateErr"
			result.Description = "大会员更新失败"
			break
		case 69015:
			result.Message = "VipPoolTTLErr"
			result.Description = "资源池过期"
			break
		case 69016:
			result.Message = "VipBusinessStatusErr"
			result.Description = "业务方无效"
			break
		case 69017:
			result.Message = "VipOrderNoErr"
			result.Description = "大会员订单号有误"
			break
		case 69018:
			result.Message = "VipOpenErr"
			result.Description = "大会员资源池开通失败"
			break
		case 69020:
			result.Message = "VipPoolNameErr"
			result.Description = "资源池名称不能为空"
			break
		case 69021:
			result.Message = "VipPoolReasonErr"
			result.Description = "资源池使用事由不能为空"
			break
		case 69022:
			result.Message = "VipPoolStartTimeErr"
			result.Description = "资源池开始时间无效"
			break
		case 69023:
			result.Message = "VipPoolEndTimeErr"
			result.Description = "资源池结束时间无效"
			break
		case 69025:
			result.Message = "VipPoolNameExitErr"
			result.Description = "资源池名称已存在"
			break
		case 69026:
			result.Message = "VipBusinessNotExitErr"
			result.Description = "业务方不存在"
			break
		case 69027:
			result.Message = "VipPoolValidityTimeErr"
			result.Description = "资源池有效期数据有误"
			break
		case 69028:
			result.Message = "VipBatchUnitErr"
			result.Description = "资源批次单位不能小于等于零且单位小于3660"
			break
		case 69029:
			result.Message = "VipBatchCountErr"
			result.Description = "资源批次总量出错"
			break
		case 69030:
			result.Message = "VipBatchPlusResouceErr"
			result.Description = "资源批次增量不能小于零"
			break
		case 69031:
			result.Message = "VipPayWayNotSupport"
			result.Description = "vip支付方式不支持"
			break
		case 69032:
			result.Message = "VipUserInfoNotExit"
			result.Description = "不存在该会员信息"
			break
		case 69033:
			result.Message = "VipCaptchaVerifyCheckErr"
			result.Description = "参数 captcha 验证失败"
			break
		case 69050:
			result.Message = "VipMonthErr"
			result.Description = "月份输入错误"
			break
		case 69060:
			result.Message = "VipMonthPriceErr"
			result.Description = "月份价格输入错误"
			break
		case 69061:
			result.Message = "VipMonthsNotFoundErr"
			result.Description = "月份找不到"
			break
		case 69062:
			result.Message = "VipBatchPriceErr"
			result.Description = "资源批次价格异常"
			break
		case 69063:
			result.Message = "VipCodeIDErr"
			result.Description = "激活码不存在"
			break
		case 69064:
			result.Message = "VipPointExchangeErr"
			result.Description = "积分兑换失败"
			break
		case 69065:
			result.Message = "VipBatchCodeCountErr"
			result.Description = "资源激活码批次数量最多20W"
			break
		case 69066:
			result.Message = "VipBatchCodeNameErr"
			result.Description = "资源批次激活码的名称已存在"
			break
		case 69067:
			result.Message = "VipOpenCodeCountErr"
			result.Description = "开通查询达到上限20次"
			break
		case 69068:
			result.Message = "VipCodeNotExitErr"
			result.Description = "激活码不存在"
			break
		case 69069:
			result.Message = "VipOpenCodeErr"
			result.Description = "开通失败"
			break
		case 69070:
			result.Message = "VipCodeUsedErr"
			result.Description = "激活码已被使用"
			break
		case 69071:
			result.Message = "VipCodeFrozenErr"
			result.Description = "激活码被冻结"
			break
		case 69072:
			result.Message = "VipCodeTTLErr"
			result.Description = "激活码已过期"
			break
		case 69073:
			result.Message = "VipCodeErr"
			result.Description = "无效激活码"
			break
		case 69074:
			result.Message = "VipCodeNotStartErr"
			result.Description = "激活码未开始"
			break
		case 69075:
			result.Message = "VipOrderNotFoundErr"
			result.Description = "订单号不存在"
			break
		case 69076:
			result.Message = "VipOrderHadExistErr"
			result.Description = "订单号已存在"
			break
		case 69077:
			result.Message = "VipPanelPlatNotExitErr"
			result.Description = "价格面板的业务不存在"
			break
		case 69078:
			result.Message = "VipPanelValidOPriceErr"
			result.Description = "价格面板的原价不能小于等于0"
			break
		case 69079:
			result.Message = "VipPanelConfNotUQErr"
			result.Description = "同类型的价格面板已存在"
			break
		case 69080:
			result.Message = "VipPanelConfNotExistErr"
			result.Description = "基础价格面版不存在"
			break
		case 69081:
			result.Message = "VipPanelValidDPriceErr"
			result.Description = "折扣价不能大于原价"
			break
		case 69082:
			result.Message = "VipPanelProductIDNotNilErr"
			result.Description = "IOS的价格面板的产品ID不能为空"
			break
		case 69083:
			result.Message = "VipPanelValidTimeErr"
			result.Description = "面板的有效时间跟已有的配置冲突"
			break
		case 69084:
			result.Message = "VipPanelSTGeqETErr"
			result.Description = "面板的结束时间必须大于开始时间"
			break
		case 69085:
			result.Message = "VipPanelSuperscriptTooLongErr"
			result.Description = "价格面板角标内容过长"
			break
		case 69086:
			result.Message = "VipPanelFirstPriceNotSupportErr"
			result.Description = "非自动续费不支持录入首月价格"
			break

		case 69100:
			result.Message = "VipTipTooLoogErr"
			result.Description = "小黄条提示至多允许填入28字"
			break
		case 69101:
			result.Message = "VipTipNotFoundErr"
			result.Description = "小黄条记录没有找到"
			break
		case 69102:
			result.Message = "VipTipStartTimeCatNotModifyErr"
			result.Description = "小黄条生效时间已不允许更新"
			break
		case 69103:
			result.Message = "VipTipEndTimeCatNotModifyErr"
			result.Description = "小黄条失效时间已不允许更新"
			break
		case 69104:
			result.Message = "VipTipCatNotDeleteErr"
			result.Description = "小黄条只允许删除未生效记录"
			break
		case 69105:
			result.Message = "VipTipCatNotExpireErr"
			result.Description = "小黄条只允许下线生效中的记录"
			break
		case 69106:
			result.Message = "VipTipTimeErr"
			result.Description = "小黄条生效时间有误"
			break
		case 69201:
			result.Message = "VipOpenOnlyNotVipErr"
			result.Description = "仅非大会员使用"
			break
		case 69202:
			result.Message = "VipCodeLimitErr"
			result.Description = "激活码查询数量到达上限"
			break
		case 69250:
			result.Message = "VipSuitPirceNotFound"
			result.Description = "大会员套餐价格信息未找到"
			break
		case 69251:
			result.Message = "VipOrderNoEmptyErr"
			result.Description = "订单号不能为空"
			break
		case 69252:
			result.Message = "VipOrderInfoNotFoundErr"
			result.Description = "未找到订单信息"
			break
		case 69253:
			result.Message = "VipOrderStatusPayingErr"
			result.Description = "订单状态必需是未支付"
			break
		case 69254:
			result.Message = "VipOrderCancelFaildErr"
			result.Description = "订单取消失败或已支付"
			break
		case 69298:
			result.Message = "VipGetPirceErr"
			result.Description = "获取价格信息失败"
			break
		case 69299:
			result.Message = "VipOrderPirceErr"
			result.Description = "订单价格有误"
			break
		case 69301:
			result.Message = "VipPushGroupLenErr"
			result.Description = "group长度过长"
			break
		case 69302:
			result.Message = "VipPushTitleLenErr"
			result.Description = "title长度过长"
			break
		case 69303:
			result.Message = "VipPushContentLenErr"
			result.Description = "内容长度过长"
			break
		case 69304:
			result.Message = "VipPushLinkTypeErr"
			result.Description = "链接类型错误"
			break
		case 69305:
			result.Message = "VipPushFmtTimeErr"
			result.Description = "推送时间格式错误"
			break
		case 69306:
			result.Message = "VipPushPlatformErr"
			result.Description = "平台及版本错误"
			break
		case 69307:
			result.Message = "VipPushEffectTimeErr"
			result.Description = "生效时间错误"
			break
		case 69308:
			result.Message = "VipPushDataNotExitErr"
			result.Description = "推送数据不存在"
			break
		case 69309:
			result.Message = "VipPushDataUpdateErr"
			result.Description = "不允许修改"
			break
		case 69310:
			result.Message = "VipPushDataDelErr"
			result.Description = "不允许删除"
			break
		case 69311:
			result.Message = "VipPushDataDisableErr"
			result.Description = "不允许下架"
			break
		case 69312:
			result.Message = "VipOrderAlreadyHandlerErr"
			result.Description = "订单已处理"
			break
		case 69313:
			result.Message = "VipRenewTypeErr"
			result.Description = "请先签订契约吧！"
			break
		case 69314:
			result.Message = "VipPayChannelNotExitErr"
			result.Description = "不存在该渠道"
			break
		case 69315:
			result.Message = "VipRescissionErr"
			result.Description = "解约失败"
			break
		case 69316:
			result.Message = "VipBisTypeErr"
			result.Description = "业务方类型错误"
			break
		case 69317:
			result.Message = "VipWhiteIPListErr"
			result.Description = "不存在ip白名单中"
			break
		case 69318:
			result.Message = "VipOrderTypeErr"
			result.Description = "只能是非续费订单"
			break
		case 69319:
			result.Message = "VipOldOrderErr"
			result.Description = "请选择新支付平台订单"
			break
		case 69320:
			result.Message = "VipOrderMoneyErr"
			result.Description = "退款金额大于购买总金额"
			break
		case 69321:
			result.Message = "VipOrderStatusErr"
			result.Description = "订单状态错误"
			break
		case 69322:
			result.Message = "VipOrderToMidErr"
			result.Description = "赠送大会员不允许退款"
			break
		case 69323:
			result.Message = "VipRefundErr"
			result.Description = "退款失败"
			break
		case 69324:
			result.Message = "VipBatchMaxCountErr"
			result.Description = "已达到最大开通次数"
			break
		case 69325:
			result.Message = "VipBatchLimitDayErr"
			result.Description = "仅限新用户可用"
			break
		case 69330:
			result.Message = "VipDialogTimeErr"
			result.Description = "大会员支付弹框时间错误"
			break
		case 69331:
			result.Message = "VipDialogConflictErr"
			result.Description = "大会员支付弹框时间冲突"
			break
		case 69332:
			result.Message = "VipPlatformConfDelErr"
			result.Description = "平台下有价格列表或弹框不允许被删除"
			break
		case 69333:
			result.Message = "VipEleOrderCanNotReFundErr"
			result.Description = "饿了么联合会员不允许退款"
			break

		case 69400:
			result.Message = "VipStartTimeErr"
			result.Description = "结束时间必须大于开始时间"
			break
		case 69401:
			result.Message = "VipTitleTooLongErr"
			result.Description = "标题过长"
			break
		case 69402:
			result.Message = "VipContentTooLongErr"
			result.Description = "内容过长"
			break

		case 69350:
			result.Message = "VipFileUploadFaildErr"
			result.Description = "权益图片上传失败"
			break
		case 69351:
			result.Message = "VipFileImgEmptyErr"
			result.Description = "权益图片不能为空"
			break
		case 69352:
			result.Message = "VipFileTypeErr"
			result.Description = "权益图片格式不合法"
			break
		case 69353:
			result.Message = "VipPrivilegeNameTooLongErr"
			result.Description = "权益名称过长"
			break
		case 69354:
			result.Message = "VipPrivilegeTitleTooLongErr"
			result.Description = "权益标题过长"
			break
		case 69355:
			result.Message = "VipPrivilegeExplainTooLongErr"
			result.Description = "权益内容过长"
			break

		case 69450:
			result.Message = "VipAssociateBindPurchasedErr"
			result.Description = "该用户已消费，不可换绑"
			break
		case 69452:
			result.Message = "VipAssociateOpenIDNotExsitErr"
			result.Description = "openID 不存在"
			break
		case 69460:
			result.Message = "VipAssociateGrantOutTradeNoExsitErr"
			result.Description = "第三方订单号已存在"
			break
		case 69461:
			result.Message = "VipAssociateGrantDurationErr"
			result.Description = "联合权益发放时长不合法"
			break
		case 69462:
			result.Message = "VipAssociateGrantLimitErr"
			result.Description = "用户联合权益超过限制"
			break
		case 69470:
			result.Message = "VipAssociatePrizeKeyErr"
			result.Description = "奖品key不合法"
			break
		case 69471:
			result.Message = "VipAssociateGrantDayErr"
			result.Description = "联合权益发放天数不合法"
			break
		case 69472:
			result.Message = "VipAssociateWhiteIPListErr"
			result.Description = "IP限制"
			break
		case 69480:
			result.Message = "VipCouponUniqueNoErr"
			result.Description = "奖品发放受限,重复的unique_no"
			break

		case 69500:
			result.Message = "VipJavaAPIErr"
			result.Description = "vip java接口调用失败"
			break
		case 69550:
			result.Message = "VipNotAutoRenewUserErr"
			result.Description = "非自动续费用户"
			break
		case 69551:
			result.Message = "VipRenewPayErr"
			result.Description = "续约扣费失败"
			break
		case 69552:
			result.Message = "VipRetryRenewPayNotExpireErr"
			result.Description = "无法重放未过期用户续约扣款"
			break
		case 69600:
			result.Message = "VipUserUnFrozenErr"
			result.Description = "vip 未冻结"
			break
		case 69700:
			result.Message = "VipWelfareCodeRunOut"
			result.Description = "福利已经被领完了"
			break
		case 69701:
			result.Message = "VipWelfareOffLine"
			result.Description = "该福利已下线，逛逛别的吧~"
			break
		case 69702:
			result.Message = "VipWelfareNotStart"
			result.Description = "该福利还未上线"
			break
		case 69703:
			result.Message = "VipWelfareAlreadyReceived"
			result.Description = "你已领取过该福利"
			break
		case 69704:
			result.Message = "VipWelfareVipOnly"
			result.Description = "只有会员才可领取福利"
			break
		case 69705:
			result.Message = "VipWelfareYearVipOnly"
			result.Description = "只有年度大会员才可领取该福利"
			break
		case 69706:
			result.Message = "VipWelfareCodeUpdConflict"
			result.Description = "多人同时获取同一优惠码冲突"
			break
		case 69707:
			result.Message = "VipWelfareNotExist"
			result.Description = "你访问的福利不存在"
			break
		case 69708:
			result.Message = "VipWelfareUploadMaxErr"
			result.Description = "上传文件超过2w行"
			break
		case 69709:
			result.Message = "VipWelfareRequestErr"
			result.Description = "跳转url不能为空!"
			break
		case 69710:
			result.Message = "VipWelfareUrlUnvalid"
			result.Description = "上传图片url不合法"
			break

		case 69900:
			result.Message = "VipActivityNotStart"
			result.Description = "活动未开始,请稍后再来"
			break
		case 69901:
			result.Message = "VipActivityHadEnd"
			result.Description = "活动已结束"
			break
		case 69902:
			result.Message = "VipActivityDeviceNotSupport"
			result.Description = "请将本页面分享至浏览器中购买"
			break
		case 69903:
			result.Message = "VipActivityAccountNotSupport"
			result.Description = "账号未绑定"
			break
		case 69904:
			result.Message = "VipActivityProductLimit"
			result.Description = "超过限额，试试其他商品"
			break
		case 69905:
			result.Message = "VipEleUnionReceivePrizesErr"
			result.Description = "饿了么奖品发放失败"
			break
		case 69906:
			result.Message = "VipEleUnionBuyCanPurchaseErr"
			result.Description = "饿了么联合会员限制购买"
			break
		case 69907:
			result.Message = "VipEleUnionReqErr"
			result.Description = "饿了么服务请求失败"
			break
		case 69908:
			result.Message = "VipEleUnionRespCodeErr"
			result.Description = "饿了么服务请求响应有误"
			break
		case 69909:
			result.Message = "VipEleUnionBuyCanProductErr"
			result.Description = "不支持购买的商品类型"
			break
		case 69910:
			result.Message = "VipEleUnionBuyGiveFirendErr"
			result.Description = "联合会员不支持好友赠送"
			break
		case 69911:
			result.Message = "VipPlatformByIDNotFoundErr"
			result.Description = "未查到平台数据"
			break
		case 69912:
			result.Message = "VipOrderStatusNotSuccessErr"
			result.Description = "订单未支付成功"
			break
		case 69913:
			result.Message = "VipOrderEleVipGrantFaildErr"
			result.Description = "饿了么会员发放失败"
			break
		case 69914:
			result.Message = "VipPriceInfoNotFoundErr"
			result.Description = "价格信息未找到"
			break
		case 69915:
			result.Message = "VipPriceNotEleProductErr"
			result.Description = "非饿了么商品"
			break
		case 69916:
			result.Message = "VipMailReqErr"
			result.Description = "会员购服务请求失败"
			break
		case 69917:
			result.Message = "VipMailRespCodeErr"
			result.Description = "会员购服务请求响应有误"
			break
		case 69918:
			result.Message = "VipActivityOrderNotFoundErr"
			result.Description = "未找到活动订单信息"
			break
		case 69919:
			result.Message = "VipAssociateThirdBindErr"
			result.Description = "饿了么绑定不成功"
			break

		case 70001:
			result.Message = "SvenRepeat"
			result.Description = "sven 数据重复"
			break
		case 70002:
			result.Message = "CanalAddrFmtErr"
			result.Description = "canal地址格式错误{host:port}!"
			break
		case 70003:
			result.Message = "CanalAddrExist"
			result.Description = "canal地址已存在"
			break
		case 70004:
			result.Message = "CanalAddrNotFound"
			result.Description = "canal地址不存在"
			break
		case 70005:
			result.Message = "CanalApplyUpdateErr"
			result.Description = "canal申请信息更新失败"
			break
		case 70006:
			result.Message = "CanalApplyErr"
			result.Description = "canal申请失败"
			break
		case 70007:
			result.Message = "DatabasesUnmarshalErr"
			result.Description = "Databases解析失败"
			break
		case 70008:
			result.Message = "GetConfigByNameErr"
			result.Description = "根据name获取config信息失败"
			break
		case 70009:
			result.Message = "DatabusGroupErr"
			result.Description = "databus group信息获取失败"
			break
		case 70010:
			result.Message = "DatabusAppErr"
			result.Description = "databus app信息获取失败"
			break
		case 70011:
			result.Message = "DatabusActionErr"
			result.Description = "databus action信息获取失败"
			break
		case 70012:
			result.Message = "ConfigCreateErr"
			result.Description = "配置中心生成配置文件失败"
			break
		case 70013:
			result.Message = "ConfigUpdateErr"
			result.Description = "配置中心更新配置文件失败"
			break
		case 70014:
			result.Message = "SetConfigIDErr"
			result.Description = "canal更新configId失败"
			break
		case 70015:
			result.Message = "CheckMasterErr"
			result.Description = "canalcheckMaster验证失败"
			break
		case 70016:
			result.Message = "ConfigParseErr"
			result.Description = "canal config信息解析失败"
			break
		case 70017:
			result.Message = "NeedInfoErr"
			result.Description = "需求不存在"
			break
		case 70018:
			result.Message = "NeedEditErr"
			result.Description = "需求不符合编辑要求"
			break
		case 70019:
			result.Message = "NeedVerifyErr"
			result.Description = "需求已审核"
			break
		case 70020:
			result.Message = "ConfigNotNow"
			result.Description = "配置中心配置源文件非最新"
			break
		case 70021:
			result.Message = "DatabusDuplErr"
			result.Description = "databus group重复"
			break
		// share
		case 71000:
			result.Message = "ShareAlreadyAdd"
			result.Description = "已经分享过了"
			break

		// push
		case 72001:
			result.Message = "PushSensitiveWordsErr"
			result.Description = "推送信息中有敏感词"
			break
		case 72002:
			result.Message = "PushUUIDErr"
			result.Description = "调用添加推送任务接口请求重放了"
			break
		case 72003:
			result.Message = "PushBizAuthErr"
			result.Description = "业务方调用接口时token校验未通过"
			break
		case 72004:
			result.Message = "PushSilenceErr"
			result.Description = "业务方处于免打扰时间段，不允许推送"
			break
		case 72005:
			result.Message = "PushBizForbiddenErr"
			result.Description = "业务方被禁止推送"
			break
		case 72006:
			result.Message = "PushUploadInvalidErr"
			result.Description = "上传的文件内容不符合规范"
			break
		case 72007:
			result.Message = "PushRecordRepeatErr"
			result.Description = "该记录已经存在"
			break
		case 72008:
			result.Message = "PushServiceBusyErr"
			result.Description = "系统繁忙，请稍后再试"
			break
		case 72009:
			result.Message = "PushServiceFileSizeErr"
			result.Description = "图片大小超过限制"
			break
		case 72010:
			result.Message = "PushServiceFileExtErr"
			result.Description = "图片格式不支持"
			break
		case 72101:
			result.Message = "PushAdminDPNoDataErr"
			result.Description = "数据平台数据未准备好"
			break

		// coupon
		case 73001:
			result.Message = "CouPonUseTooFrequently"
			result.Description = "观影劵使用劵操作太频繁"
			break
		case 73002:
			result.Message = "CouPonNotEnoughErr"
			result.Description = "观影劵数量不够"
			break
		case 73003:
			result.Message = "CouPonConsumeFaildErr"
			result.Description = "劵消费失败"
			break
		case 73004:
			result.Message = "CouPonTypeNotExistErr"
			result.Description = "劵类型不存在"
			break
		case 73005:
			result.Message = "CouPonBatchNotExistErr"
			result.Description = "劵批次号不存在"
			break
		case 73006:
			result.Message = "CouPonBatchNotEnoughErr"
			result.Description = "劵批次可发放劵数不足"
			break
		case 73007:
			result.Message = "CouPonBatchTimeErr"
			result.Description = "批次开始时间不能大于等于过期时间"
			break
		case 73008:
			result.Message = "CouPonBatchLimitErr"
			result.Description = "超过单人最大可领取劵数"
			break
		case 73009:
			result.Message = "CouPonGrantTooFrequently"
			result.Description = "劵领取操作太频繁"
			break
		case 73010:
			result.Message = "CouPonGrantMaxCountErr"
			result.Description = "该批次劵已发放完"
			break
		case 73011:
			result.Message = "CouPonHadUseErr"
			result.Description = "代金券已被使用"
			break
		case 73012:
			result.Message = "CouPonTokenNotFoundErr"
			result.Description = "券不存在"
			break
		case 73013:
			result.Message = "CouPonHadBlockErr"
			result.Description = "券已被冻结"
			break
		case 73014:
			result.Message = "CouPonHadExpireErr"
			result.Description = "券不在有效期内"
			break
		case 73015:
			result.Message = "CouPonNotFullPriceErr"
			result.Description = "券未满足可使用条件"
			break
		case 73016:
			result.Message = "CouPonOrderHadUseErr"
			result.Description = "订单已经使用过代金券"
			break
		case 73017:
			result.Message = "CouPonStateCanNotCancelErr"
			result.Description = "不能解锁非使用中的劵"
			break
		case 73018:
			result.Message = "CouPonUseFaildErr"
			result.Description = "劵使用失败"
			break
		case 73019:
			result.Message = "CouPonNotifyStateErr"
			result.Description = "发货状态未知"
			break
		case 73020:
			result.Message = "CouponAmountErr"
			result.Description = "代金券的满额金额必须小于等于券金额"
			break
		case 73021:
			result.Message = "CouponUpdateFileErr"
			result.Description = "批量发放文件格式有误"
			break
		case 73022:
			result.Message = "CouponBatchSalaryLimitErr"
			result.Description = "批量发放数量超过上限"
			break
		case 73023:
			result.Message = "CouponBatchSalaryCountZeroErr"
			result.Description = "批量发放数量不能为0"
			break
		case 73024:
			result.Message = "CouPonPlatformNotSupportErr"
			result.Description = "平台限制使用"
			break
		case 73025:
			result.Message = "CouponExpireErr"
			result.Description = "请填写过期时间"
			break
		case 73026:
			result.Message = "CouponReceiveErr"
			result.Description = "券接收失败"
			break
		case 73027:
			result.Message = "CouponBatchBlockErr"
			result.Description = "批次信息被冻结"
			break
		case 73028:
			result.Message = "CouponBatchExpireTimeErr"
			result.Description = "批次已过期"
			break
		case 73029:
			result.Message = "CouPonProductNotSupportErr"
			result.Description = "商品限制使用"
			break
		case 73100:
			result.Message = "CouponInfoStateBlockErr"
			result.Description = "当前观影券状态不允许冻结"
			break
		case 73101:
			result.Message = "CouponInfoStateUnblockErr"
			result.Description = "当前观影券状态不允许解冻"
			break
		case 73102:
			result.Message = "CouponSalaryCountErr"
			result.Description = "批量发放数量不匹配"
			break
		case 73103:
			result.Message = "CouponSalaryHadRunErr"
			result.Description = "批量发放任务已开始"
			break
		case 73104:
			result.Message = "CouponTypeNotSupportErr"
			result.Description = "批量发放类型不支持"
			break
		case 73200:
			result.Message = "CouponNewYearNotStartErr"
			result.Description = "元旦活动未开始"
			break
		case 73201:
			result.Message = "CouponNewYearIsEndErr"
			result.Description = "元旦活动已结束"
			break
		case 73202:
			result.Message = "CouponNewYearIsOpenErr"
			result.Description = "元旦活动卡牌已翻开"
			break
		case 73203:
			result.Message = "CouponNewYearGrantErr"
			result.Description = "元旦活动卡牌发放失败"
			break

		// coupon code
		case 73300:
			result.Message = "CouponCodeVerifyFaildErr"
			result.Description = "验证码错误,请重新输入"
			break
		case 73301:
			result.Message = "CouponCodeNotFoundErr"
			result.Description = "兑换码不存在,请重新输入"
			break
		case 73302:
			result.Message = "CouponCodeUsedErr"
			result.Description = "兑换码已使用,请重新输入"
			break
		case 73303:
			result.Message = "CouponCodeBlockErr"
			result.Description = "兑换码已冻结,请重新输入"
			break
		case 73304:
			result.Message = "CouponCodeLimitByMidErr"
			result.Description = "单个用户兑换数量超过上限,请重新输入"
			break
		case 73305:
			result.Message = "CouponCodeMaxLimitByMidErr"
			result.Description = "批次兑换数量超过上限,请重新输入"
			break
		case 73306:
			result.Message = "CouponCodeCanNotUseErr"
			result.Description = "兑换码非使用中状态"
			break
		case 73307:
			result.Message = "CouponCodeMaxLimitErr"
			result.Description = "未设置批次码最大上线，或超过最大限制5W"
			break

		// realname 实名认证
		case 74001:
			result.Message = "RealnameCaptureErr"
			result.Description = "实名验证码输入错误"
			break
		case 74002:
			result.Message = "RealnameCaptureSendTooMany"
			result.Description = "实名验证码发送次数过于频繁"
			break
		case 74003:
			result.Message = "RealnameCaptureInvalid"
			result.Description = "实名验证码未发送或已失效"
			break
		case 74004:
			result.Message = "RealnameCaptureErrTooMany"
			result.Description = "实名验证码错误次数过多"
			break
		case 74005:
			result.Message = "RealnameApplyAlready"
			result.Description = "实名认证已提交申请"
			break
		case 74006:
			result.Message = "RealnameInvalidName"
			result.Description = "实名姓名错误"
			break
		case 74007:
			result.Message = "RealnameImageIDErr"
			result.Description = "实名照片错误"
			break
		case 74008:
			result.Message = "RealnameImageExpired"
			result.Description = "实名认证照片过期"
			break
		case 74009:
			result.Message = "RealnameCardNumErr"
			result.Description = "实名证件号码错误"
			break
		case 74010:
			result.Message = "RealnameCardBindAlready"
			result.Description = "实名证件已绑定"
			break
		case 74011:
			result.Message = "RealnameAlipayAntispam"
			result.Description = "实名触发芝麻认证防刷"
			break
		case 74012:
			result.Message = "RealnameAlipayErr"
			result.Description = "实名芝麻认证服务错误"
			break
		case 74013:
			result.Message = "RealnameAlipayApplyInvalid"
			result.Description = "实名芝麻认证申请不合法"
			break

		// activity
		case 75001:
			result.Message = "ActivityNotExist"
			result.Description = "活动不存在"
			break
		case 75002:
			result.Message = "ActivityNotStart"
			result.Description = "活动没有开始"
			break
		case 75003:
			result.Message = "ActivityOverEnd"
			result.Description = "活动已结束"
			break
		case 75004:
			result.Message = "ActivityHaveGuess"
			result.Description = "活动已竞猜"
			break
		case 75005:
			result.Message = "ActivityNotEnoughCoin"
			result.Description = "硬币不足"
			break
		case 75006:
			result.Message = "ActivityOverCoin"
			result.Description = "超额投币"
			break
		case 75007:
			result.Message = "ActivityServerTimeout"
			result.Description = "服务超时"
			break
		case 75008:
			result.Message = "ActivityKeyNotExists"
			result.Description = "key不存在"
			break
		case 75009:
			result.Message = "ActivityKeyBindAlready"
			result.Description = "key已绑定"
			break
		case 75010:
			result.Message = "ActivityMidBindAlready"
			result.Description = "用户已绑定"
			break
		case 75011:
			result.Message = "ActivityNotBind"
			result.Description = "用户未绑定"
			break
		case 75012:
			result.Message = "ActivityIDNotExists"
			result.Description = "ID不存在"
			break
		case 75013:
			result.Message = "ActivityAwardAlready"
			result.Description = "奖品已兑换"
			break
		case 75014:
			result.Message = "ActivityNoAward"
			result.Description = "没有奖品"
			break
		case 75015:
			result.Message = "ActivityNotOwner"
			result.Description = "不是该point点Owner"
			break
		case 75016:
			result.Message = "ActivityHasUnlock"
			result.Description = "该point点已解锁"
			break
		case 75017:
			result.Message = "ActivityGameResult"
			result.Description = "请选择游戏结果"
			break
		case 75018:
			result.Message = "ActivityUserAchieveFail"
			result.Description = "用户成就记录获取失败"
			break
		case 75019:
			result.Message = "ActivityUserPointFail"
			result.Description = "用户point记录获取失败"
			break
		case 75020:
			result.Message = "ActivityAchieveFail"
			result.Description = "成就列表获取失败"
			break
		case 75021:
			result.Message = "ActivityPointFail"
			result.Description = "point列表获取失败"
			break
		case 75022:
			result.Message = "ActivityNoAchieve"
			result.Description = "没有成就"
			break
		case 75023:
			result.Message = "ActivityKeyFail"
			result.Description = "key获取失败"
			break
		case 75024:
			result.Message = "ActivityMidFail"
			result.Description = "mid获取失败"
			break
		case 75025:
			result.Message = "ActivityNotAdmin"
			result.Description = "非管理员登录"
			break
		case 75026:
			result.Message = "ActivityAddAchieveFail"
			result.Description = "获得成就失败"
			break
		case 75027:
			result.Message = "ActivityLackHp"
			result.Description = "您的HP不足"
			break
		case 75028:
			result.Message = "ActivityMaxHp"
			result.Description = "您的HP已满"
			break
		case 75029:
			result.Message = "ActivityNotAwardAdmin"
			result.Description = "您不是奖品兑换管理员"
			break
		case 75030:
			result.Message = "ActivityNotLotteryAdmin"
			result.Description = "您不是抽奖管理员"
			break
		case 75031:
			result.Message = "ActivityLotteryFail"
			result.Description = "抽奖失败"
			break
		case 75032:
			result.Message = "ActivityUnlockFail"
			result.Description = "解锁失败"
			break
		case 75033:
			result.Message = "ActivityNotLotteryAchieve"
			result.Description = "该成就不支持抽奖"
			break
		case 75034:
			result.Message = "ActivityLikeIPFrequence"
			result.Description = "点赞活动ip访问过于频繁-访问过快"
			break
		case 75035:
			result.Message = "ActivityLikeScoreLower"
			result.Description = "账户score过低不支持点赞-异常账号!"
			break
		case 75036:
			result.Message = "ActivityLikeLevelLimit"
			result.Description = "-用户等级不够!"
			break
		case 75037:
			result.Message = "ActivityLikeMemberLimit"
			result.Description = "-用户注册不足7天！"
			break
		case 75038:
			result.Message = "ActivityLikeNotStart"
			result.Description = "-评分未开始!"
			break
		case 75039:
			result.Message = "ActivityLikeHasEnd"
			result.Description = "-评分已结束!"
			break
		case 75040:
			result.Message = "ActivityLikeHasLike"
			result.Description = "-已点赞过!"
			break
		case 75041:
			result.Message = "ActivityLikeHasGrade"
			result.Description = "-已评过分!"
			break
		case 75042:
			result.Message = "ActivityLikeRegisterLimit"
			result.Description = "-晚于活动限制注册时间!"
			break
		case 75043:
			result.Message = "ActivityLikeHasVote"
			result.Description = "-已投过票!"
			break
		case 75044:
			result.Message = "ActivityHasOffLine"
			result.Description = "-活动已经下线"
			break
		case 75045:
			result.Message = "ActivityLikeHasOffLine"
			result.Description = "-活动稿件已经下线"
			break
		case 75046:
			result.Message = "ActivityLikeBeforeRegister"
			result.Description = "-早于活动限制注册时间!"
			break
		case 75047:
			result.Message = "ActivityHasMissionGroup"
			result.Description = "-已发起过活动!"
			break
		case 75048:
			result.Message = "ActivityMGNotYourself"
			result.Description = "-不支持给自己助力哦!"
			break
		case 75049:
			result.Message = "ActivityMissionNotStart"
			result.Description = "-助攻未开始!"
			break
		case 75050:
			result.Message = "ActivityMissionHasEnd"
			result.Description = "-助攻已结束!"
			break
		case 75051:
			result.Message = "ActivityHasMission"
			result.Description = "-已助攻过!"
			break
		case 75052:
			result.Message = "ActivityOverMissionLimit"
			result.Description = "-超出可助攻上限!"
			break
		case 75053:
			result.Message = "ActivityHasAward"
			result.Description = "-重复领取"
			break
		case 75054:
			result.Message = "ActivityNotAward"
			result.Description = "-非法领取"
			break
		case 75055:
			result.Message = "ActivityTelValid"
			result.Description = "-未绑定有效手机号码"
			break
		case 75056:
			result.Message = "ActivityOverDailyScore"
			result.Description = "-超过单日投票上线"
			break
		case 75057:
			result.Message = "ActivityBnjTimeCancel"
			result.Description = "倒计时活动取消"
			break
		case 75058:
			result.Message = "ActivityBnjResetCD"
			result.Description = "倒计时重置CD"
			break
		case 75059:
			result.Message = "ActivityBnjTimeFinish"
			result.Description = "倒计时已完成"
			break
		case 75060:
			result.Message = "ActivityBnjNotSub"
			result.Description = "未预约拜年祭活动"
			break
		case 75061:
			result.Message = "ActivityBnjSubLow"
			result.Description = "该宝箱预约人数未达成"
			break
		case 75062:
			result.Message = "ActivityBnjHasReward"
			result.Description = "该宝箱奖励已被领取"
			break
		case 75063:
			result.Message = "ActivityBnjRewardFail"
			result.Description = "宝箱奖励领取失败"
			break
		case 75064:
			result.Message = "ActivityRewardConfErr"
			result.Description = "宝箱奖励配置错误"
			break
		case 75065:
			result.Message = "ActivityMemberBlocked"
			result.Description = "-封禁用户无法操作"
			break
		case 75066:
			result.Message = "ActivityKfcHasUsed"
			result.Description = "-code已经被使用"
			break
		case 75067:
			result.Message = "ActivityKfcNotExist"
			result.Description = "-code不存在"
			break
		case 75068:
			result.Message = "ActivityKfcNotGiveOut"
			result.Description = "-code未发放"
			break
		case 75069:
			result.Message = "ActivityKfcSqlError"
			result.Description = "-发生未知错误"
			break

		// TV
		case 76001:
			result.Message = "TvUpperNotInList"
			result.Description = "up主不存在"
			break
		case 76002:
			result.Message = "TvDangbeiWrongType"
			result.Description = "当贝页面-内容类型错误"
			break
		case 76003:
			result.Message = "TvDangbeiPageNotExist"
			result.Description = "当贝页面不存在"
			break
		case 76004:
			result.Message = "TvSearchOrderWrong"
			result.Description = "搜索所提供的order字段有误"
			break
		case 76005:
			result.Message = "TvVideoNotFound"
			result.Description = "视频找不到，文案：啊叻？视频不见啦！看看别的吧～"
			break
		case 76006:
			result.Message = "TvLabelExist"
			result.Description = "所添加的索引标签已存在"
			break
		case 76007:
			result.Message = "TvSyncErr"
			result.Description = "同步牌照错误"
			break
		case 76008:
			result.Message = "TvPGCRankEmpty"
			result.Description = "pgc排行榜结果为空"
			break
		case 76009:
			result.Message = "TvPGCRankNewEPNil"
			result.Description = "pgc排行榜结果中newEP为空"
			break
		case 76011:
			result.Message = "TvAllDataAuditing"
			result.Description = "tv稿件或者season下所有的数据为审核中，即详情页无可展示数据时"
			break
		case 76101:
			result.Message = "TvPriceTimeConflict"
			result.Description = "折扣价时间冲突，请检查已有折扣时间段"
			break
		case 76102:
			result.Message = "TvVipProductExit"
			result.Description = "该产品id已存在"
			break
		case 76103:
			result.Message = "TvVipProdSyncErr"
			result.Description = "同步产品包到云视听错误,请尝试重试"
			break
		case 76104:
			result.Message = "TVVipSuitTypeConflict"
			result.Description = "请先下线已存在的升级产品"
			break

		// manager
		case 77001:
			result.Message = "ManagerTagDelErr"
			result.Description = "manager tag不能删除"
			break
		case 77002:
			result.Message = "ManagerUIDNOTExist"
			result.Description = "manager管理员不存在"
			break
		case 77003:
			result.Message = "ManagerFlowForbidden"
			result.Description = "manager 工作流被禁用"
			break
		case 77004:
			result.Message = "ManagerTagTypeDelErr"
			result.Description = "manager tag类型不能删除"
			break
		// APP
		case 78000:
			result.Message = "AppNotData"
			result.Description = "app not data"
			break

		// subtitle
		case 79001:
			result.Message = "SubtitleDrfatExist"
			result.Description = "当前已存在未过审字幕"
			break
		case 79002:
			result.Message = "SubtitleDrfatNotExist"
			result.Description = "当前字幕草稿不存在"
			break
		case 79003:
			result.Message = "SubtitleDrfatUnSubmit"
			result.Description = "当前字幕未提交"
			break
		case 79004:
			result.Message = "SubtitleDelUnExist"
			result.Description = "删除不存在的字幕"
			break
		case 79005:
			result.Message = "SubtitlePermissionDenied"
			result.Description = "字幕操作权限不足"
			break
		case 79006:
			result.Message = "SubtitleDenied"
			result.Description = "视频禁止投稿"
			break
		case 79007:
			result.Message = "SubtileLanLocked"
			result.Description = "当前语言版本锁定"
			break
		case 79008:
			result.Message = "SubtitleUnValid"
			result.Description = "字幕不合法"
			break
		case 79009:
			result.Message = "SubtitleWaveFormFailed"
			result.Description = "波形图调用失败"
			break
		case 79010:
			result.Message = "SubtitlePubNotExist"
			result.Description = "发布的字幕不存在"
			break
		case 79011:
			result.Message = "SubtitleIllegalLanguage"
			result.Description = "不合法的语言"
			break
		case 79012:
			result.Message = "SubtitleNotPublish"
			result.Description = "当前字幕未发布"
			break
		case 79013:
			result.Message = "SubtitleTimeUnValid"
			result.Description = "字幕时间不合法"
			break
		case 79014:
			result.Message = "SubtitleSizeLimit"
			result.Description = "字幕超过限制"
			break
		case 79015:
			result.Message = "SubtitleOriginUnValid"
			result.Description = "原字幕不合法"
			break
		case 79016:
			result.Message = "SubtitleLocationUnValid"
			result.Description = "字幕位置不合法"
			break
		case 79017:
			result.Message = "SubtitleUserBalcked"
			result.Description = "账号黑名单"
			break
		case 79018:
			result.Message = "SubtitleStatusUnValid"
			result.Description = "字幕状态不合法"
			break
		case 79019:
			result.Message = "SubtitleAlreadyHasDraft"
			result.Description = "前存在草稿或者待审核状态的字幕"
			break
		case 79020:
			result.Message = "SubtitleVideoDurationOverFlow"
			result.Description = "字幕时间点超过视频时间长度"
			break
		case 79021:
			result.Message = "SubtitleDuarionMustThanZero"
			result.Description = "字幕的持续时间必须大于0"
			break

		// MCN
		// --82000~82499 前台错误
		case 82001:
			result.Message = "MCNNotAllowed"
			result.Description = "没有权限操作"
			break
		case 82002:
			result.Message = "MCNUpCannotBind"
			result.Description = "无法绑定Up主，已被绑定"
			break
		case 82003:
			result.Message = "MCNUpBindUpAlreadyInProgress"
			result.Description = "已存在与up的绑定在审核中"
			break
		case 82004:
			result.Message = "MCNUpBindUpDuplicatePhone"
			result.Description = "该电话号码已绑定"
			break
		case 82005:
			result.Message = "MCNUpBindUpDuplicateIDCard"
			result.Description = "该身份证号码已绑定"
			break
		case 82006:
			result.Message = "MCNUpBindUpDuplicateCompanyLicenseID"
			result.Description = "该营业执照号码已绑定"
			break
		case 82007:
			result.Message = "MCNUpBindUpDuplicateCompanyName"
			result.Description = "该公司名称已绑定"
			break
		case 82008:
			result.Message = "MCNUpBindUpSTimeLtETime"
			result.Description = "up主绑定的开始时间必须小于结束时间"
			break
		case 82009:
			result.Message = "MCNUpBindUpIsBlocked"
			result.Description = "up主已被封禁"
			break
		case 82010:
			result.Message = "MCNUpBindUpDateError"
			result.Description = "日期错误，请检查"
			break
		case 82011:
			result.Message = "MCNStateInvalid"
			result.Description = "MCN状态异常"
			break
		case 82012:
			result.Message = "MCNUpBindInvalid"
			result.Description = "该绑定请求已失效"
			break
		case 82013:
			result.Message = "MCNUpBindInvalidURL"
			result.Description = "该绑定的站外up主链接错误，请检查"
			break
		case 82014:
			result.Message = "MCNUpOutSiteIsNotQualified"
			result.Description = "站外Up主需要满足（1.粉丝数≤100  或  2. 投稿数＜2及90天内未投稿）"
			break
		case 82015:
			result.Message = "MCNUpBindUpIsBlueUser"
			result.Description = "该up主为蓝V用户，不能被绑定"
			break
		case 82016:
			result.Message = "MCNUpSignStateInvalid"
			result.Description = "该Up主签约状态异常"
			break
		case 82020:
			result.Message = "MCNChangePermissionAlreadyInProgress"
			result.Description = "等待UP主授权或运营审核"
			break
		case 82021:
			result.Message = "MCNChangePermissionLackPermission"
			result.Description = "您缺少要申请的权限"
			break
		case 82022:
			result.Message = "MCNChangePermissionSamePermission"
			result.Description = "权限没有任何变更"
			break
		case 82030:
			result.Message = "MCNPublicationFailTimeLimit"
			result.Description = "刊例价每月只能更改一次"
			break
		case 82040:
			result.Message = "MCNPermissionInsufficient"
			result.Description = "权限不足"
			break

		// --82500~82999 后台错误
		case 82501:
			result.Message = "MCNSignCycleNotUQErr"
			result.Description = "mcn签约周期冲突"
			break
		case 82502:
			result.Message = "MCNUnknownFileTypeErr"
			result.Description = "只允许jpg、png、pdf、word格式的文件上传"
			break
		case 82503:
			result.Message = "MCNSignUnknownReviewErr"
			result.Description = "mcn签约未知审核方式"
			break
		case 82504:
			result.Message = "MCNSignOnlyReviewOpErr"
			result.Description = "只有待审核中的mcn才能操作"
			break
		case 82505:
			result.Message = "MCNUpUnknownReviewErr"
			result.Description = "mcn和up主绑定未知审核方式"
			break
		case 82506:
			result.Message = "MCNUpOnlyReviewOpErr"
			result.Description = "有待审核中的mcn和up主绑定才能操作"
			break
		case 82507:
			result.Message = "MCNContractFileSize"
			result.Description = "mcn合同上传的大小，允许20M"
			break
		case 82508:
			result.Message = "MCNCSignUnknownInfoErr"
			result.Description = "未知mcn信息"
			break
		case 82509:
			result.Message = "MCNCUpUnknownInfoErr"
			result.Description = "未知mcn-up主绑定信息"
			break
		case 82510:
			result.Message = "MCNUnknownFileExt"
			result.Description = "未知上传文件后缀名"
			break
		case 82511:
			result.Message = "MCNSignNoOkState"
			result.Description = "处于封禁中、审核中、未申请中、驳回中、签约中、待开启中的状态不可录入"
			break
		case 82512:
			result.Message = "MCNUpPassOnEffectSign"
			result.Description = "up主通过必须是有效的签约状态"
			break
		case 82513:
			result.Message = "MCNSignIsBlocked"
			result.Description = "mcn管理用户已被封禁"
			break
		case 82514:
			result.Message = "MCNSignEtimeNLEQNowTime"
			result.Description = "mcn签约结束时间不能小于等于当前时间"
			break
		case 82515:
			result.Message = "MCNSignStateFlowErr"
			result.Description = "mcn状态流转错误,请联系开发人员"
			break
		case 82516:
			result.Message = "MCNRecommendUpStateFlowErr"
			result.Description = "推荐的up状态流转错误,请联系开发人员"
			break
		case 82517:
			result.Message = "MCNRecommendUpInPool"
			result.Description = "该up主已经被推荐"
			break
		case 82518:
			result.Message = "MCNRecommendUpMidsIsEmpty"
			result.Description = "推荐池被操作的up主不能为空"
			break
		case 82519:
			result.Message = "MCNUpPermitUnknownReviewErr"
			result.Description = "未知mcn-up的权限变更审核方式"
			break
		case 82520:
			result.Message = "MCNUpAbnormalDataErr"
			result.Description = "异常数据"
			break
		case 82521:
			result.Message = "MCNUpPermitStateFlowErr"
			result.Description = "mcn-up的权限变更的状态流转错误,请联系开发人员"
			break

		case 82601:
			result.Message = "MCNRenewalNotInRangeErr"
			result.Description = "MCN续约不在范围内"
			break
		case 82602:
			result.Message = "MCNRenewalAlreadyErr"
			result.Description = "MCN已续过约"
			break
		case 82603:
			result.Message = "MCNRenewalDateErr"
			result.Description = "MCN续约周期错误"
			break

		// esports
		case 83001:
			result.Message = "EsportsContestNotExist"
			result.Description = "你所订阅的赛程不存在"
			break
		case 83002:
			result.Message = "EsportsContestMaxCount"
			result.Description = "你订阅赛程数已达上限"
			break
		case 83003:
			result.Message = "EsportsContestFavDel"
			result.Description = "该赛程未订阅，不能取消哦~"
			break
		case 83004:
			result.Message = "EsportsContestFavExist"
			result.Description = "该赛程已订阅，不能重复订阅哦~"
			break
		case 83005:
			result.Message = "EsportsContestNotDay"
			result.Description = "仅可订阅15天内的赛事哦~"
			break
		case 83006:
			result.Message = "EsportsContestStart"
			result.Description = "你订阅的赛程已经开始啦~快来直播间观看吧~"
			break
		case 83007:
			result.Message = "EsportsContestEnd"
			result.Description = "你订阅的赛程已经结束啦~可以点击回放和集锦进行观看哦~"
			break
		case 83008:
			result.Message = "EsportsContestFavNot"
			result.Description = "该赛程不可订阅哦~"
			break
		case 83009:
			result.Message = "EsportsActNotExist"
			result.Description = "赛事活动不存在"
			break
		case 83010:
			result.Message = "EsportsActVideoNotExist"
			result.Description = "赛事活动视频不存在"
			break
		case 83011:
			result.Message = "EsportsActPointNotExist"
			result.Description = "比赛数据不存在"
			break
		case 83012:
			result.Message = "EsportsActKnockNotExist"
			result.Description = "淘汰赛数据不存在"
			break
		case 83050:
			result.Message = "EsportsModNameErr"
			result.Description = "模块名称重复~"
			break
		case 83051:
			result.Message = "EsportsActModNot"
			result.Description = "模块不属于该赛事活动~"
			break
		case 83052:
			result.Message = "EsportsActModErr"
			result.Description = "模块信息不正确~"
			break
		case 83053:
			result.Message = "EsportsModArcErr"
			result.Description = "模块稿件不正确~"
			break
		case 83054:
			result.Message = "EsportsTreeNodeErr"
			result.Description = "节点不属于该赛事详情~"
			break
		case 83055:
			result.Message = "EsportsTreeDetailErr"
			result.Description = "赛事活动详情不存在~"
			break
		case 83056:
			result.Message = "EsportsTreeEmptyErr"
			result.Description = "当前没有任何记录，请编辑后提交~"
			break
		case 83057:
			result.Message = "EsportsMultiEdit"
			result.Description = "节点不能多人同时保存~"
			break
		case 83058:
			result.Message = "EsportsArcServerErr"
			result.Description = "稿件服务出错~"
			break
		case 83059:
			result.Message = "EsportsContestDataErr"
			result.Description = "比赛数据不正确~"
			break

		case 85000:
			result.Message = "CardNotEffectiveErr"
			result.Description = "非有效卡片"
			break
		case 85001:
			result.Message = "CardEquipNotVipErr"
			result.Description = "非大会员不可装备"
			break
		case 85003:
			result.Message = "CardIDNotFoundErr"
			result.Description = "不存在的卡片ID"
			break
		case 85004:
			result.Message = "GroupIDNotFoundErr"
			result.Description = "不存在的分类ID"
			break
		case 85005:
			result.Message = "CardFileUploadFaildErr"
			result.Description = "卡片图片文件上传失败"
			break
		case 85006:
			result.Message = "CardNameTooLongErr"
			result.Description = "卡片名称过长"
			break
		case 85007:
			result.Message = "CardImageEmptyErr"
			result.Description = "卡片图片不能为空"
			break
		case 85008:
			result.Message = "CardNameExistErr"
			result.Description = "卡片名称已存在"
			break
		case 85009:
			result.Message = "CardGroupNameExistErr"
			result.Description = "卡片组名称已存在"
			break

			//passport
			//passport-login 86000~86299
			//passport-user  86300~86599
			//passport-sns	 86600~
		case 86600:
			result.Message = "PassportSnsMidAlreadyBindQQ"
			result.Description = "mid已绑定QQ号"
			break
		case 86601:
			result.Message = "PassportSnsMidAlreadyBindWEIBO"
			result.Description = "mid已绑定微博"
			break
		case 86610:
			result.Message = "PassportSnsQQAlreadyBind"
			result.Description = "QQ号已被绑定"
			break
		case 86611:
			result.Message = "PassportSnsWEIBOAlreadyBind"
			result.Description = "微博已经绑定"
			break
		case 86660:
			result.Message = "PassportSnsRequestErr"
			result.Description = "请求第三方失败"
			break

			//playurl
		case 87000:
			result.Message = "PlayURLNotLogin"
			result.Description = "用户未登录"
			break
		case 87001:
			result.Message = "PlayURLNotPay"
			result.Description = "稿件未付费"
			break

		// upcpay
		case 88001:
			result.Message = "UGCPayAssetInvalid"
			result.Description = "ugcpay 内容无效"
			break
		case 88002:
			result.Message = "UGCPayAssetCantBuy"
			result.Description = "ugcpay 内容不可购买"
			break
		case 88003:
			result.Message = "UGCPayOrderInvalid"
			result.Description = "ugcpay 订单无效"
			break
		case 88004:
			result.Message = "UGCPayOrderNotPay"
			result.Description = "ugcpay 订单未付款"
			break
		case 88005:
			result.Message = "UGCPayAssetPaid"
			result.Description = "ugcpay 内容已支付"
			break
		case 88101:
			result.Message = "UGCPayDependArchiveErr"
			result.Description = "ugcpay archive依赖错误"
			break
		case 88102:
			result.Message = "UGCPayDependPayPlatformErr"
			result.Description = "ugcpay payplatform依赖错误"
			break

		case 90000:
			result.Message = "ToViewPayUGC"
			result.Description = "付费稿件"
			break
		case 90001:
			result.Message = "ToViewOverMax"
			result.Description = "塞满啦！先看看库存吧~"
			break

		// uprating
		case 91000:
			result.Message = "UpRatingNoPermission"
			result.Description = "用户无权限"
			break
		case 91001:
			result.Message = "UpRatingNoData"
			result.Description = "用户无数据"
			break
		case 91002:
			result.Message = "UpRatingScoreLimit"
			result.Description = "用户分数未达标"
			break

			//aegis
		case 92001:
			result.Message = "AegisUniqueAlreadyExist"
			result.Description = "当前流程%s %s 已存在"
			break
		case 92002:
			result.Message = "AegisTokenNotAssign"
			result.Description = "当前绑定令牌%s不是赋值语句"
			break
		case 92003:
			result.Message = "AegisTokenNotFound"
			result.Description = "当前令牌不存在"
			break
		case 92004:
			result.Message = "AegisFlowNoToken"
			result.Description = "当前节点没绑定令牌"
			break
		case 92005:
			result.Message = "AegisFlowNotFound"
			result.Description = "当前节点不存在"
			break
		case 92006:
			result.Message = "AegisFlowDisabled"
			result.Description = "当前节点已被禁用"
			break
		case 92007:
			result.Message = "AegisFlowNoEnableTran"
			result.Description = "当前节点所有下游变化不可用"
			break
		case 92008:
			result.Message = "AegisFlowNoFromDir"
			result.Description = "当前节点上游没有有向线"
			break
		case 92009:
			result.Message = "AegisFlowBinded"
			result.Description = "当前流程节点已被关联，请取消后再禁用"
			break
		case 92010:
			result.Message = "AegisTranBinded"
			result.Description = "当前流程变化已被关联，请取消后再禁用"
			break
		case 92011:
			result.Message = "AegisTranNotFound"
			result.Description = "当前变化不存在"
			break
		case 92012:
			result.Message = "AegisTranDisabled"
			result.Description = "当前变化已被禁用"
			break
		case 92013:
			result.Message = "AegisTranNoFlow"
			result.Description = "当前变化没有可用输出节点"
			break
		case 92014:
			result.Message = "AegisRunInDiffNet"
			result.Description = "当前资源在不同网内运行"
			break
		case 92015:
			result.Message = "AegisNotRunInFlow"
			result.Description = "当前资源不在当前节点上运行"
			break
		case 92016:
			result.Message = "AegisFlowDiffNet"
			result.Description = "节点不在同一网上"
			break
		case 92017:
			result.Message = "AegisNotRunInRange"
			result.Description = "当前资源不在该业务或网内运行"
			break
		case 92018:
			result.Message = "AegisNotTriggerFlow"
			result.Description = "当前资源没有触发当前节点上的变化"
			break
		case 92019:
			result.Message = "AegisDirOrderConflict"
			result.Description = "当前有向线顺序(%s)冲突:%s"
			break
		case 92020:
			result.Message = "AegisNetErr"
			result.Description = "当前网配置错误"
			break

		case 92021:
			result.Message = "AegisTaskErr"
			result.Description = "任务操作错误"
			break
		case 92022:
			result.Message = "AegisResourceErr"
			result.Description = "资源操作失败"
			break
		case 92023:
			result.Message = "AegisBusinessCfgErr"
			result.Description = "业务配置错误"
			break
		case 92024:
			result.Message = "AegisSearchErr"
			result.Description = "搜索接口错误"
			break
		case 92025:
			result.Message = "AegisDuplicateErr"
			result.Description = "资源重复"
			break
		case 92026:
			result.Message = "AegisReservedCfgErr"
			result.Description = "保留字段配置错误"
			break
		case 92027:
			result.Message = "AegisBusinessSyncErr"
			result.Description = "业务回调错误"
			break
		case 92028:
			result.Message = "AegisTaskBusy"
			result.Description = "任务调度繁忙，请重试"
			break
		case 92029:
			result.Message = "AegisTaskFinish"
			result.Description = "任务已被处理，无需再操作"
			break

		// TV VIP
		case 93000:
			result.Message = "TVIPTokenErr"
			result.Description = "查询令牌无效"
			break
		case 93001:
			result.Message = "TVIPTokenExpire"
			result.Description = "查询令牌过期"
			break
		case 93002:
			result.Message = "TVIPPanelNotFound"
			result.Description = "套餐不存在"
			break
		case 93003:
			result.Message = "TVIPPanelNotSuitalbe"
			result.Description = "套餐不用"
			break
		case 93004:
			result.Message = "TVIPNotContracted"
			result.Description = "用户未签约连续包月"
			break
		case 93005:
			result.Message = "TVIPRenewTooEarly"
			result.Description = "续费过早"
			break
		case 93006:
			result.Message = "TVIPRenewTooLate"
			result.Description = "续费过晚"
			break
		case 93007:
			result.Message = "TVIPNotVip"
			result.Description = "非VIP用户"
			break
		case 93008:
			result.Message = "TVIPPayOrderExpired"
			result.Description = "订单超时适"
			break
		case 93009:
			result.Message = "TVIPPayOrderNotFound"
			result.Description = " 订单不存在"
			break
		case 93010:
			result.Message = "TVIPSignErr"
			result.Description = "签名错误"
			break
		case 93011:
			result.Message = "TVIPBuyNumExceeded"
			result.Description = "购买数量超标"
			break
		case 93012:
			result.Message = "TVIPYstSystemErr"
			result.Description = "系统异常"
			break
		case 93013:
			result.Message = "TVIPYstRequestErr"
			result.Description = "云视听请求错误"
			break
		case 93014:
			result.Message = "TVIPYstUnknownErr"
			result.Description = "云视听未知异常"
			break
		case 93015:
			result.Message = "TVIPYstUnknownTradeState"
			result.Description = "云视听交易状态异常"
			break
		case 93016:
			result.Message = "TVIPDupOrderNo"
			result.Description = "重复的订单号"
			break
		case 93017:
			result.Message = "TVIPQrDevideUnsupported"
			result.Description = "扫码设备不支持"
			break
		case 93018:
			result.Message = "TVIPGiveMVipFailed"
			result.Description = "主站大会员赠送失败"
			break
		case 93019:
			result.Message = "TVIPBatchIdNotFound"
			result.Description = "主站大会员批次id不存在"
			break
		case 93020:
			result.Message = "TVIPBuyRateExceeded"
			result.Description = "购买过于频繁"
			break
		case 93021:
			result.Message = "TVIPMVipRateExceeded"
			result.Description = "购买大会员升级套餐频繁"
			break
		}
	}

	// not found
	return result
}
