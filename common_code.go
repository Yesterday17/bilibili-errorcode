package bilierrorcode

func getCommonCodeDetail(code ErrorCode) ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}

	switch code {
	case -1:
		result.Message = "AppKeyInvalid"
		result.Description = "应用程序不存在或已被封禁"
		break
	case -2:
		result.Message = "AccessKeyErr"
		result.Description = "Access Key错误"
		break
	case -3:
		result.Message = "SignCheckErr"
		result.Description = "API校验密匙错误"
		break
	case -4:
		result.Message = "MethodNoPermission"
		result.Description = "调用方对该Method没有权限"
		break
	case -101:
		result.Message = "NoLogin"
		result.Description = "账号未登录"
		break
	case -102:
		result.Message = "UserDisabled"
		result.Description = "账号被封停"
		break
	case -103:
		result.Message = "LackOfScores"
		result.Description = "积分不足"
		break
	case -104:
		result.Message = "LackOfCoins"
		result.Description = "硬币不足"
		break
	case -105:
		result.Message = "CaptchaErr"
		result.Description = "验证码错误"
		break
	case -106:
		result.Message = "UserInactive"
		result.Description = "账号未激活"
		break
	case -107:
		result.Message = "UserNoMember"
		result.Description = "账号非正式会员或在适应期"
		break
	case -108:
		result.Message = "AppDenied"
		result.Description = "应用不存在或者被封禁"
		break
	case -110:
		result.Message = "MobileNoVerfiy"
		result.Description = "未绑定手机"
		break
	case -111:
		result.Message = "CsrfNotMatchErr"
		result.Description = "csrf 校验失败"
		break
	case -112:
		result.Message = "ServiceUpdate"
		result.Description = "系统升级中"
		break
	case -113:
		result.Message = "UserIDCheckInvalid"
		result.Description = "账号尚未实名认证"
		break
	case -114:
		result.Message = "UserIDCheckInvalidPhone"
		result.Description = "请先绑定手机"
		break
	case -115:
		result.Message = "UserIDCheckInvalidCard"
		result.Description = "请先完成实名认证"
		break

	case -304:
		result.Message = "NotModified"
		result.Description = "木有改动"
		break
	case -307:
		result.Message = "TemporaryRedirect"
		result.Description = "撞车跳转"
		break
	case -400:
		result.Message = "RequestErr"
		result.Description = "请求错误"
		break
	case -401:
		result.Message = "Unauthorized"
		result.Description = "未认证"
		break
	case -403:
		result.Message = "AccessDenied"
		result.Description = "访问权限不足"
		break
	case -404:
		result.Message = "NothingFound"
		result.Description = "啥都木有"
		break
	case -405:
		result.Message = "MethodNotAllowed"
		result.Description = "不支持该方法"
		break
	case -409:
		result.Message = "Conflict"
		result.Description = "冲突"
		break
	case -500:
		result.Message = "ServerErr"
		result.Description = "服务器错误"
		break
	case -503:
		result.Message = "ServiceUnavailable"
		result.Description = "过载保护,服务暂不可用"
		break
	case -504:
		result.Message = "Deadline"
		result.Description = "服务调用超时"
		break
	case -509:
		result.Message = "LimitExceed"
		result.Description = "超出限制"
		break
	case -616:
		result.Message = "FileNotExists"
		result.Description = "上传文件不存在"
		break
	case -617:
		result.Message = "FileTooLarge"
		result.Description = "上传文件太大"
		break
	case -625:
		result.Message = "FailedTooManyTimes"
		result.Description = "登录失败次数太多"
		break
	case -626:
		result.Message = "UserNotExist"
		result.Description = "用户不存在"
		break
	case -628:
		result.Message = "PasswordTooLeak"
		result.Description = "密码太弱"
		break
	case -629:
		result.Message = "UsernameOrPasswordErr"
		result.Description = "用户名或密码错误"
		break
	case -632:
		result.Message = "TargetNumberLimit"
		result.Description = "操作对象数量限制"
		break
	case -643:
		result.Message = "TargetBlocked"
		result.Description = "被锁定"
		break
	case -650:
		result.Message = "UserLevelLow"
		result.Description = "用户等级太低"
		break
	case -652:
		result.Message = "UserDuplicate"
		result.Description = "重复的用户"
		break
	case -658:
		result.Message = "AccessTokenExpires"
		result.Description = "Token 过期"
		break
	case -662:
		result.Message = "PasswordHashExpires"
		result.Description = "密码时间戳过期"
		break
	case -688:
		result.Message = "AreaLimit"
		result.Description = "地理区域限制"
		break
	case -689:
		result.Message = "CopyrightLimit"
		result.Description = "版权限制"
		break
	case -701:
		result.Message = "FailToAddMoral"
		result.Description = "扣节操失败"
		break

	case -1200:
		result.Message = "Degrade"
		result.Description = "被降级过滤的请求"
		break
	case -1201:
		result.Message = "RPCNoClient"
		result.Description = "rpc服务的client都不可用"
		break
	case -1202:
		result.Message = "RPCNoAuth"
		result.Description = "rpc服务的client没有授权"
		break
	}
	return result
}
