package types

type (
	AppointmentID int
)

type (
	AppointmentStatus string
)

type (
	Customer struct {
		ID        *UserID     `json:"id,omitempty"`
		AvatarURL *ImageURL   `json:"avatar_url,omitempty"`
		Name      Name        `json:"name"`
		Contact   ContactInfo `json:"contact"`
	}

	Executor struct {
		ID        *UserID     `json:"id,omitempty"`
		AvatarURL *ImageURL   `json:"avatar_url,omitempty"`
		Name      Name        `json:"name"`
		Contact   ContactInfo `json:"contact"`
	}

	Appointment struct {
		ID                 AppointmentID     `json:"id"`
		Customer           Customer          `json:"customer"`
		Status             AppointmentStatus `json:"status"`
		Services           []Service         `json:"services"`
		Schedule           Interval          `json:"schedule"`
		TotalServices      uint              `json:"total_services"`
		TotalExecutionTime uint              `json:"total_execution_time"`
		BasePrice          uint              `json:"base_price"`
		TotalPrice         uint              `json:"total_price"`
		TotalDiscount      uint              `json:"total_discount"`
		PromoCode          *PromoCode        `json:"promo_code,omitempty"`
		Discount           *Discount         `json:"discount,omitempty"`
		Notes              String500         `json:"notes,omitempty"`
		Executor           Executor          `json:"executor"`
		Location           Location          `json:"location"`
		Audit              Audit             `json:"audit"`
	}
)
