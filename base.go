package bilierrorcode

// ErrorCode ErrorCode
type ErrorCode int32

// ErrorCodeDetail Detailed information
type ErrorCodeDetail struct {
	Code        ErrorCode `json:"code"`
	Message     string    `json:"raw"`
	Description string    `json:"message"`
}

// GetRegion Get which part your error code is
func (code ErrorCode) GetRegion() string {
	if code < 0 {
		return "common"
	} else if code == 0 {
		return "ok"
	} else if code > 0 && code <= 990000 {
		return "main_or_ep"
	} else if code >= 1000000 && code <= 1999999 {
		return "live"
	} else if code >= 5000000 && code < 6000000 {
		return "bbq"
	} else if code >= 2000000 && code <= 2099999 {
		return "ticket_or_open_platform"
	} else if code >= 2000000 && code <= 2999999 {
		return "open_platform"
	}
	return "unknown"
}

// GetDetail Get detail information about the ErrorCode given
func (code ErrorCode) GetDetail() ErrorCodeDetail {
	var result ErrorCodeDetail

	switch code.GetRegion() {
	case "ok":
		return ErrorCodeDetail{
			Code:        0,
			Message:     "OK",
			Description: "无错误",
		}
	case "common":
		// 尝试匹配通用错误代码
		result = getCommonCodeDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "main_or_ep":
		// 尝试匹配主站
		result = getMainSiteDetail(code)
		if result.Message != "" {
			return result
		}
		// TODO: ep
		// FIXME: ep 实用性待定 有用再加
		break
	case "live":
		// 尝试匹配直播
		result = getLiveSiteDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "bbq":
		result = getBBQCodeDetail(code)
		if result.Message != "" {
			return result
		}
		break
	case "ticket_or_open_platform":
	case "open_platform":
		// 尝试匹配开放平台
		result = getOpenPlatformSiteDetail(code)
		if result.Message != "" {
			return result
		}
		break
	}

	// 默认情况 返回空
	result.Code = code
	result.Message = ""
	result.Description = ""
	return result
}
