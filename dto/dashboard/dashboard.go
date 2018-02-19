package dashboard

import "time"

// Dashboards represent a list of graylog dashboard
type Dashboards struct {
	Total      int         `json:"total"`
	Dashboards []Dashboard `json:"dashboards"`
}

// Dashboard represent a graylog dashboard
type Dashboard struct {
	ID            *string    `json:"id,omitempty"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CreatorUserID *string    `json:"creatorUserId,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Widgets       []Widget   `json:"widgets,omitempty"`
}

type Widget struct {
	CreatorUserID *string `json:"creator_user_id,omitempty"`
	CacheTime     int     `json:"cache_time"`
	Description   string  `json:"description"`
	ID            *string `json:"id,omitempty"`
	Type          string  `json:"type"`
	Config        Config  `json:"config"`
}

type Config struct {
	Timerange     Timerange `json:"timerange"`
	LowerIsBetter bool      `json:"lower_is_better"`
	StreamID      string    `json:"stream_id"`
	Trend         bool      `json:"trend"`
	Query         string    `json:"query"`
}

type Timerange struct {
	Type  string `json:"type"`
	Range int    `json:"range"`
}
