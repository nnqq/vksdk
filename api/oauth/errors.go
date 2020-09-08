// Package oauth ...
package oauth

// ValidationType ...
type ValidationType string

// Possible values.
const (
	ValidationSMS ValidationType = "2fa_sms"
	ValidationApp ValidationType = "2fa_app"
)

// ErrorType for oauth.
type ErrorType string

// Error types.
//
// See https://tools.ietf.org/html/rfc6749#section-4.2.2.1
const (
	ErrInvalidRequest          ErrorType = "invalid_request"
	ErrUnauthorizedClient      ErrorType = "unauthorized_client"
	ErrUnsupportedResponseType ErrorType = "unsupported_response_type"
	ErrInvalidScope            ErrorType = "invalid_scope"
	ErrServerError             ErrorType = "server_error"
	ErrTemporarilyUnavailable  ErrorType = "temporarily_unavailable"
	ErrAccessDenied            ErrorType = "access_denied"

	ErrInvalidGrant ErrorType = "invalid_grant"

	ErrNeedValidation ErrorType = "need_validation"
	ErrNeedCaptcha    ErrorType = "need_captcha"
)

// Error returns the message of a Error.
func (e ErrorType) Error() string {
	return "oauth: error with type " + string(e)
}

// ErrorReason for oauth.
type ErrorReason string

// Error returns the message of a Error.
func (e ErrorReason) Error() string {
	return "oauth: error with reason " + string(e)
}

// ErrorReason types.
const (
	ErrUserDenied ErrorReason = "user_denied"
)

// Error for oauth.
type Error struct {
	Type        ErrorType   `json:"error"`
	Reason      ErrorReason `json:"error_reason,omitempty"`
	Description string      `json:"error_description"`

	// For auth direct
	CaptchaSID     string         `json:"captcha_sid,omitempty"`
	CaptchaImg     string         `json:"captcha_img,omitempty"`
	RedirectURI    string         `json:"redirect_uri,omitempty"`
	ValidationType ValidationType `json:"validation_type,omitempty"`
	PhoneMask      string         `json:"phone_mask,omitempty"`
}

// Error returns the message of a Error.
func (e Error) Error() string {
	return "oauth: " + e.Description
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e Error) Is(target error) bool {
	if t, ok := target.(*Error); ok {
		return e.Type == t.Type && e.Description == t.Description
	}

	if t, ok := target.(ErrorType); ok {
		return e.Type == t
	}

	if t, ok := target.(ErrorReason); ok {
		return e.Reason == t
	}

	return false
}
