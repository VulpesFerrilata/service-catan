package cache

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Harbor struct {
	ID              int                  `json:"id"`
	Q               int                  `json:"q"`
	R               int                  `json:"r"`
	Type            datamodel.HarborType `json:"type"`
	HarborDirection HarborDirection      `json:"harborDirection"`
}

type HarborDirection string

const (
	TD_TOP_LEFT     = "TOP_LEFT"
	TD_MIDDLE_LEFT  = "MIDDLE_LEFT"
	TD_BOTTOM_LEFT  = "BOTTOM_LEFT"
	TD_TOP_RIGHT    = "TOP_RIGHT"
	TD_MIDDLE_RIGHT = "MIDDLE_RIGHT"
	TD_BOTTOM_RIGHT = "BOTTOM_RIGHT"
)
