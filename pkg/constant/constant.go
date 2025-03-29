package constant

const (
	// Common headers
	HeaderContentType    = "Content-Type"
	HeaderAuthorization  = "Authorization"
	HeaderXRequestID     = "X-Request-ID"
	HeaderXCorrelationID = "X-Correlation-ID"

	// Auth related headers
	HeaderXAccessToken  = "X-Access-Token"
	HeaderXRefreshToken = "X-Refresh-Token"

	// Rate limiting headers
	HeaderRateLimit     = "X-RateLimit-Limit"
	HeaderRateRemaining = "X-RateLimit-Remaining"
	HeaderRateReset     = "X-RateLimit-Reset"
)

// Content types
const (
	ContentTypeJSON      = "application/json"
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeText      = "text/plain"
	ContentTypeHTML      = "text/html"
)

// Service Names
const (
	ServiceAuth         = "auth-service"
	ServiceTransaction  = "transaction-service"
	ServiceNotification = "notification-service"
)

// Circuit Breaker settings
const (
	CircuitBreakerTimeout     = 30
	CircuitBreakerMaxRequests = 5
	CircuitBreakerInterval    = 60
	CircuitBreakerThreshold   = 3
)

// Rate Limiting settings
const (
	RateLimitPerMinute = 60
	RateLimitBurst     = 10
)

// JWT settings
const (
	JWTAccessTokenExpiry  = 15
	JWTRefreshTokenExpiry = 7
)

// Error codes
const (
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	ErrCodeRateLimitExceeded  = "RATE_LIMIT_EXCEEDED"
)
