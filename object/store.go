package object

// StoreProductType type.
type StoreProductType string

// Possible values.
const (
	StoreProductStickers StoreProductType = "stickers"
)

// StoreProduct struct.
type StoreProduct struct {
	ID           int         `json:"id"`
	Type         string      `json:"type"`
	Purchased    BaseBoolInt `json:"purchased"`
	Active       BaseBoolInt `json:"active"`
	PurchaseDate int         `json:"purchase_date,omitempty"`
}

// StoreProductExtended struct.
type StoreProductExtended struct {
	StoreProduct
	Title        string         `json:"title"`
	Stickers     []StoreSticker `json:"stickers"`
	Icon         []BaseImage    `json:"icon"`
	Previews     []BaseImage    `json:"previews"`
	HasAnimation BaseBoolInt    `json:"has_animation"`
}

// StoreSticker struct.
type StoreSticker struct {
	StickerID            int         `json:"sticker_id"`
	IsAllowed            BaseBoolInt `json:"is_allowed"`
	Images               []BaseImage `json:"images"`
	ImagesWithBackground []BaseImage `json:"images_with_background"`
	AnimationURL         string      `json:"animation_url"`
}
