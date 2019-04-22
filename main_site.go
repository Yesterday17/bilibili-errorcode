package bilierrorcode

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
			result.Description = "稿件不属于该用户"
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

			// Template
			// case 10086:
			// 	result.Message = "ArchiveTypeForbidAdd"
			// 	result.Description = "该类型不支持投稿"
			// 	break
		}
	}

	// not found
	return result
}
