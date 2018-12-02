package vksdk

import "encoding/json"

// UsersGetResponse users.get response
type UsersGetResponse []User

// UsersGet returns detailed information on users
// users.get Returns detailed information on users.
func (vk *VK) UsersGet(params map[string]string) (UsersGetResponse, error) {
	var response UsersGetResponse

	rawResponse, err := vk.Request("users.get", params)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// UsersGetFollowersResponse struct
type UsersGetFollowersResponse struct{}

// TODO users.getFollowers Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.

// UsersGetSubscriptionsResponse struct
type UsersGetSubscriptionsResponse struct{}

// TODO users.getSubscriptions Returns a list of IDs of users and public pages followed by the user.

// UsersIsAppUserResponse struct
type UsersIsAppUserResponse struct{}

// TODO users.isAppUser Returns information whether a user installed the application.

// UsersReportResponse struct
type UsersReportResponse struct{}

// TODO users.report Reports (submits a complain about) a user.

// UsersSearchResponse struct
type UsersSearchResponse struct{}

// TODO users.search Returns a list of users matching the search criteria.