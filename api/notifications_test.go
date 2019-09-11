package api

import (
	"testing"
)

func TestVK_NotificationsGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.NotificationsGet(map[string]string{
		"count": "30",
	})
	if err != nil {
		t.Errorf("VK.NotificationsGet() err = %v", err)
	}
}

func TestVK_NotificationsMarkAsViewed(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.NotificationsMarkAsViewed(map[string]string{})
	if err != nil {
		t.Errorf("VK.NotificationsMarkAsViewed() err = %v", err)
	}
}
