package enum

// go install github.com/alvaroloes/enumer@latest

//go:generate enumer -type=HttpCode -json
//go:generate enumer -type=Reason -json

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
	ShowInFrontendConfig = []ConfigKey{
		ConfigKeySiteInfo,
	}
)
