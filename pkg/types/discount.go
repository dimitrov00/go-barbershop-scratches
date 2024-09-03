package types

type (
	PromoCodeID int
	DiscountID  int
)

type (
	DiscountType string
)

type (
	PromoCode struct {
		ID       PromoCodeID  `json:"id"`
		Code     string       `json:"code"`
		Discount float32      `json:"discount"`
		Type     DiscountType `json:"type"`
	}

	Discount struct {
		ID          DiscountID   `json:"id"`
		Name        String50     `json:"name"`
		Discount    float32      `json:"discount"`
		Type        DiscountType `json:"type"`
		Description String500    `json:"description"`
	}
)
