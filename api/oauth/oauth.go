// Package oauth ...
package oauth

import (
	"net/url"

	"github.com/SevereCloud/vksdk/v2/api"
)

const (
	scheme             = "https"
	oauthHost          = "oauth.vk.com"
	version            = api.Version
	defaultRedirectURI = "https://oauth.vk.com/blank.html"
)

// Display sets authorization page appearance.
type Display string

// The supported values.
const (
	Page   Display = "page"   // authorization form in a separate window
	Popup  Display = "popup"  // a pop-up window
	Mobile Display = "mobile" //  authorization for mobile devices (uses no Javascript).
)

func parseCode(u *url.URL) (string, error) {
	v := u.Query()

	if errType := v.Get("error"); errType != "" {
		return "", &Error{
			Type:        ErrorType(errType),
			Reason:      ErrorReason(v.Get("error_reason")),
			Description: v.Get("error_description"),
		}
	}

	return v.Get("code"), nil
}
