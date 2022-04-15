package domains

import (
	"time"

	"github.com/google/uuid"
)

type SurveyMethodologyEnum int

type SurveyStatusEnum int

const (
	ACTIVE SurveyStatusEnum = iota + 1
	INACTIVE
)

const (
	CSAT SurveyMethodologyEnum = iota + 1
	NPS
	CES
	CUSTOM
)

type Survey struct {
	Id                  uuid.UUID `json:"id"`
	Brand               string    `json:"brand"`
	Template            uuid.UUID
	Name                string                `json:"name"`
	Mothodology         SurveyMethodologyEnum `json:"methodology"`
	Status              SurveyStatusEnum      `json:"status"`
	Description         string                `json:"description"`
	ResponseRedirectUrl string                `json:"response_redirect_url"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
