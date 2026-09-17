package enum

type HttpCode int

const (
	Success HttpCode = 200
	Fail    HttpCode = 0
)

type Reason int

const (
	// RequestError 请求失败
	RequestError Reason = iota
	RequestSuccess
)

//keys

const (
	KeyUser          = "user"
	KeyAuthorization = "Authorization"
	KeyBearer        = "Bearer "
	KeyUserInfo      = "userInfo"
	KeyAdminUser     = "admin"
	KeyDebug         = "debug"
	KeyRelease       = "release"
)

var (
	AllowHeaders = []string{
		"Origin", "Content-Type", "Accept", "Authorization",
		"X-Time", "X-Path", "X-Sign", "X-Key", "X-Body-Sign",
		"X-Request-ID", "X-Data", "Is-Automated",
	}
	ExposeHeaders = []string{"Content-Type", "Authorization", "X-Login"}

	ShowInFrontendConfig = []ConfigKey{
		ConfigKeySiteInfo,
	}
)
