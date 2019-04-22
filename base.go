package bilierrorcode

// ErrorCode ErrorCode
type ErrorCode int32

// ErrorCodeDetail Detailed information
type ErrorCodeDetail struct {
	Code        ErrorCode
	Message     string
	Description string
}

// GetRegion Get which part your error code is
func (code ErrorCode) GetRegion() string {
	if code >= 0 && code <= 990000 {
		return "main_or_ep"
	} else if code >= 5000000 && code < 6000000 {
		return "bbq"
	} else if code >= 2000000 && code <= 2099999 {
		return "ticket"
	} else if code >= 2000000 && code <= 2999999 {
		return "open_platform"
	}
	return "unknown"
}

// GetDetail Get detail information about the ErrorCode given
func (code ErrorCode) GetDetail() ErrorCodeDetail {
	result := ErrorCodeDetail{
		Code:        code,
		Message:     "",
		Description: "",
	}
	if code.GetRegion() != "unknown" {
		if getMainSiteDetail(code).Message != "" {
			return getMainSiteDetail(code)
		}
	}
	return result
}
