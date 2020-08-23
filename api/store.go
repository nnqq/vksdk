package api

import "github.com/SevereCloud/vksdk/object"

// StoreActivateProduct method.
//
// https://vk.com/dev/store.activateProduct
func (vk *VK) StoreActivateProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.activateProduct", params, &response)
	return
}

// StoreGetProductsResponse struct.
type StoreGetProductsResponse struct {
	Count int                   `json:"count"`
	Items []object.StoreProduct `json:"items"`
}

// StoreGetProducts method.
//
// extended=0
//
// https://vk.com/dev/store.getProducts
func (vk *VK) StoreGetProducts(params Params) (response StoreGetProductsResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("store.getProducts", params, &response)

	return
}

// StoreGetProductsExtendedResponse struct.
type StoreGetProductsExtendedResponse struct {
	Count int                           `json:"count"`
	Items []object.StoreProductExtended `json:"items"`
}

// StoreGetProductsExtended method.
//
// extended=1
//
// https://vk.com/dev/store.getProducts
func (vk *VK) StoreGetProductsExtended(params Params) (response StoreGetProductsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("store.getProducts", params, &response)

	return
}
