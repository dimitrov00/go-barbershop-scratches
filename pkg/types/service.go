package types

type (
	ServiceID string
)

type (
	Service struct {
		ID            ServiceID `json:"id"`
		Name          String50  `json:"name"`
		Description   String500 `json:"description"`
		Price         uint      `json:"price"`
		ExecutionTime uint      `json:"execution_time"`
	}
)
