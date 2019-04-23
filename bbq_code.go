package bilierrorcode

func getBBQCodeDetail(code ErrorCode) ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}

	switch code {
	case 3001016:
		result.Message = "CheckInviteCodeErr"
		result.Description = "检查邀请码失败(特殊处理请勿修改) todo"
		break

	// Common [5000000,5001000)
	case 5000000:
		result.Message = "Common"
		result.Description = ""
		break
	case 5000001:
		result.Message = "TypeDismatch"
		result.Description = "类型不匹配"
		break
	case 5000002:
		result.Message = "ExternalErr"
		result.Description = "外部错误"
		break
	case 5000003:
		result.Message = "ReqParamErr"
		result.Description = "参数错误"
		break
	case 5000004:
		result.Message = "BBQSystemErr"
		result.Description = "用于一些异常请求"
		break
	case 5000005:
		result.Message = "BBQNoBindPhone"
		result.Description = "未绑定手机号"
		break
	case 5000006:
		result.Message = "BBQUserBanned"
		result.Description = "已被封禁，无法进行相关操作，如有疑问可在“设置-吐槽”中进行反馈"
		break
	case 5000007:
		result.Message = "ArchiveDatabusNilErr"
		result.Description = "预发环境不配置稿件databus"
		break

	// Search [5001000,5002000)
	case 5001000:
		result.Message = "SearchCreateIndexErr"
		result.Description = "创建索引失败"
		break
	case 5001001:
		result.Message = "SearchVideoDataErr"
		result.Description = "获取视频信息失败"
		break

	// web [5002000,5003000)
	case 5002001:
		result.Message = "CommentClosed"
		result.Description = "评论已关闭"
		break
	case 5002003:
		result.Message = "VideoUnExists"
		result.Description = "视频不存在"
		break
	case 5002004:
		result.Message = "VideoUnReachable"
		result.Description = "视频不存在，由于状态原因不可访问"
		break
	case 5002005:
		result.Message = "VideoInAudit"
		result.Description = "视频审核中"
		break
	case 5002014:
		result.Message = "InviteCodeInvalid"
		result.Description = "无效邀请码"
		break
	case 5002015:
		result.Message = "InviteCodeUsed"
		result.Description = "邀请码已使用"
		break
	case 5002021:
		result.Message = "CommentForbidden"
		result.Description = "禁止评论"
		break
	case 5002023:
		result.Message = "CommentTooShort"
		result.Description = "评论过短"
		break
	case 5002024:
		result.Message = "CommentTooLong"
		result.Description = "评论过长"
		break
	case 5002025:
		result.Message = "SvNotReachable"
		result.Description = ""
		break
	case 5002026:
		result.Message = "NoticeTypeErr"
		result.Description = "通知类型错误"
		break
	case 5002027:
		result.Message = "CommentForbidLike"
		result.Description = "禁止赞或踩"
		break
	case 5002028:
		result.Message = "CommentLengthIllegal"
		result.Description = "评论长度不合法"
		break

	// video-service [5003000,5004000)
	case 5003000:
		result.Message = "UnKnownBPS"
		result.Description = "未知码率"
		break
	case 5003001:
		result.Message = "SyncBVCFail"
		result.Description = "同步bvc转码失败"
		break
	case 5003002:
		result.Message = "VideoDelFail"
		result.Description = "视频删除失败，不能删除别人的视频"
		break

	// UserLike [5005000, 5005100]
	case 5005000:
		result.Message = "UserLike"
		result.Description = "UserLike [5005000, 5005100]"
		break
	case 5005001:
		result.Message = "AddUserLikeErr"
		result.Description = "点赞失败"
		break
	case 5005002:
		result.Message = "CancelUserLikeErr"
		result.Description = "取消点赞失败"
		break

	// UserInfo [5005100, 5005200]
	case 5005100:
		result.Message = "UserInfo"
		result.Description = "UserInfo"
		break
	case 5005101:
		result.Message = "BatchUserTooLong"
		result.Description = "用户批量请求太多"
		break
	case 5005102:
		result.Message = "UPMIDNotExists"
		result.Description = "up主不存在"
		break
	case 5005103:
		result.Message = "GetUserBaseErr"
		result.Description = "获取用户信息失败"
		break
	case 5005104:
		result.Message = "EditUserBaseErr"
		result.Description = "更新用户基础信息失败"
		break
	case 5005105:
		result.Message = "UserUnameSpecial"
		result.Description = "昵称包含特殊字符"
		break
	case 5005106:
		result.Message = "UserUnameLength"
		result.Description = "昵称长度不符合"
		break
	case 5005107:
		result.Message = "UserUnameExisted"
		result.Description = "昵称已被占用"
		break
	case 5005108:
		result.Message = "UserUnameFilterErr"
		result.Description = "昵称包含敏感词"
		break
	case 5005109:
		result.Message = "UserUnamePrefixErr"
		result.Description = "该昵称无法注册"
		break

	// UserRelation [5005200, 5005300]
	case 5005200:
		result.Message = "UserRelation"
		result.Description = ""
		break
	case 5005201:
		result.Message = "AddUserFollowErr"
		result.Description = "关注失败，请稍后重试"
		break
	case 5005202:
		result.Message = "CancelUserFollowErr"
		result.Description = "取消关注失败，请稍后重试"
		break
	case 5005203:
		result.Message = "UserFollowLimitErr"
		result.Description = "关注失败，关注已达上限"
		break
	case 5005204:
		result.Message = "FollowMyselfErr"
		result.Description = "不能关注自己"
		break
	case 5005205:
		result.Message = "UserAlreadyBlackFollowErr"
		result.Description = "关注失败，请将用户移出黑名单后重试"
		break
	case 5005206:
		result.Message = "UserBlackLimitErr"
		result.Description = "拉黑失败，黑名单已达上限"
		break
	case 5005207:
		result.Message = "UserBlackErr"
		result.Description = "黑名单请求系统错误"
		break
	case 5005208:
		result.Message = "UserBlackSelfErr"
		result.Description = "拉黑失败，不能拉黑自己"
		break

	// Danmu [5005300, 5005400]
	case 5005300:
		result.Message = "Danmu"
		result.Description = ""
		break
	case 5005301:
		result.Message = "FilterErr"
		result.Description = "弹幕包含敏感词"
		break
	case 5005302:
		result.Message = "DanmuGetErr"
		result.Description = "弹幕获取失败"
		break
	case 5005303:
		result.Message = "DanmuPostErr"
		result.Description = "弹幕发送失败"
		break
	case 5005304:
		result.Message = "DanmuLimitErr"
		result.Description = "该视频暂时无法发送弹幕"
		break

		// Comment [5005400, 5005500]
	case 5005400:
		result.Message = "Comment"
		result.Description = ""
		break
	case 5005401:
		result.Message = "CommentFilterErr"
		result.Description = "评论包含敏感词"
		break
	case 5005402:
		result.Message = "CommentMissErr"
		result.Description = "评论不见了"
		break
	case 5005403:
		result.Message = "CommentLengthErr"
		result.Description = "评论需要2-96字"
		break
	case 5005404:
		result.Message = "CommentOptLimitErr"
		result.Description = "操作太快了，休息一下"
		break
	case 5005405:
		result.Message = "CommentLimithErr"
		result.Description = "该视频暂时无法发送评论"
		break

		// report [5005500, 5005599]
	case 5005501:
		result.Message = "ReportDanmuError"
		result.Description = "弹幕举报失败"
		break

		//Upload [5005600, 5005700]
	case 5005600:
		result.Message = "Upload"
		result.Description = ""
		break
	case 5005601:
		result.Message = "UploadFailed"
		result.Description = "上传失败"
		break

		// Topic [5005700, 5005800]
	case 5005700:
		result.Message = "Topic"
		result.Description = ""
		break
	case 5005701:
		result.Message = "TopicReqParamErr"
		result.Description = "参数错误"
		break
	case 5005702:
		result.Message = "TopicNumTooManyErr"
		result.Description = "一次性插入db的话题数量太大"
		break
	case 5005703:
		result.Message = "TopicNameLenErr"
		result.Description = "话题长度太长"
		break
	case 5005704:
		result.Message = "TopicIDErr"
		result.Description = "话题ID错误"
		break
	case 5005705:
		result.Message = "TopicIDNotFound"
		result.Description = "话题ID没找到"
		break
	case 5005706:
		result.Message = "TopicStateErr"
		result.Description = "话题为下架状态"
		break
	case 5005707:
		result.Message = "TopicTooManyInOneVideo"
		result.Description = "一个视频的话题数量太多了"
		break
	case 5005708:
		result.Message = "TopicDescLenErr"
		result.Description = "话题描述长度太长"
		break
	case 5005709:
		result.Message = "TopicInsertErr"
		result.Description = "话题插入失败"
		break
	case 5005721:
		result.Message = "TopicVideoStateErr"
		result.Description = "话题视频状态错误"
		break
	}
	return result
}
