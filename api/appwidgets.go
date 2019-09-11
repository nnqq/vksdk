package api // import "github.com/SevereCloud/vksdk/api"

// AppWidgetsGetAppImageUploadServerResponse struct
type AppWidgetsGetAppImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetAppImageUploadServer returns a URL for uploading a
// photo to the app collection for community app widgets.
//
// https://vk.com/dev/appWidgets.getAppImageUploadServer
func (vk *VK) AppWidgetsGetAppImageUploadServer(params map[string]string) (response AppWidgetsGetAppImageUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getAppImageUploadServer", params, &response)
	return
}

// AppWidgetsGetAppImagesResponse struct
type AppWidgetsGetAppImagesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Images struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
}

// AppWidgetsGetAppImages returns an app collection of images for community app widgets.
//
// https://vk.com/dev/appWidgets.getAppImages
func (vk *VK) AppWidgetsGetAppImages(params map[string]string) (response AppWidgetsGetAppImagesResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getAppImages", params, &response)
	return
}

// AppWidgetsGetGroupImageUploadServerResponse struct
type AppWidgetsGetGroupImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetGroupImageUploadServer returns a URL for uploading
// a photo to the community collection for community app widgets.
//
// https://vk.com/dev/appWidgets.getGroupImageUploadServer
func (vk *VK) AppWidgetsGetGroupImageUploadServer(params map[string]string) (response AppWidgetsGetGroupImageUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getGroupImageUploadServer", params, &response)
	return
}

// AppWidgetsGetGroupImagesResponse struct
type AppWidgetsGetGroupImagesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Images struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
}

// AppWidgetsGetGroupImages returns a community collection of images for community app widgets.
//
// https://vk.com/dev/appWidgets.getGroupImages
func (vk *VK) AppWidgetsGetGroupImages(params map[string]string) (response AppWidgetsGetGroupImagesResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getGroupImages", params, &response)
	return
}

// AppWidgetsGetImagesByIDResponse struct
type AppWidgetsGetImagesByIDResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsGetImagesByID returns an image for community app widgets by its ID.
//
// https://vk.com/dev/appWidgets.getImagesById
func (vk *VK) AppWidgetsGetImagesByID(params map[string]string) (response AppWidgetsGetImagesByIDResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getImagesById", params, &response)
	return
}

// AppWidgetsSaveAppImageResponse struct
type AppWidgetsSaveAppImageResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsSaveAppImage allows to save image into app collection for community app widgets.
//
// https://vk.com/dev/appWidgets.saveAppImage
func (vk *VK) AppWidgetsSaveAppImage(params map[string]string) (response AppWidgetsSaveAppImageResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.saveAppImage", params, &response)
	return
}

// AppWidgetsSaveGroupImageResponse struct
type AppWidgetsSaveGroupImageResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsSaveGroupImage allows to save image into community collection for community app widgets.
//
// https://vk.com/dev/appWidgets.saveGroupImage
func (vk *VK) AppWidgetsSaveGroupImage(params map[string]string) (response AppWidgetsSaveGroupImageResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.saveGroupImage", params, &response)
	return
}

// AppWidgetsUpdate allows to update community app widget.
//
// https://vk.com/dev/appWidgets.update
func (vk *VK) AppWidgetsUpdate(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("appWidgets.update", params, &response)

	return
}
