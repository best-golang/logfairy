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

// Widget Dashboard represent a graylog Widget
type Widget struct {
	CreatorUserID *string `json:"creator_user_id,omitempty"`
	CacheTime     int     `json:"cache_time"`
	Description   string  `json:"description"`
	ID            *string `json:"id,omitempty"`
	Type          string  `json:"type"`
	Config        Config  `json:"config"`
}

// Config Dashboard represent a graylog widget's config
type Config struct {
	Timerange      Timerange `json:"timerange"`
	Interval       *string   `json:"interval,omitempty"`
	LowerIsBetter  bool      `json:"lower_is_better"`
	StreamID       string    `json:"stream_id"`
	Trend          bool      `json:"trend"`
	Query          string    `json:"query"`
	Field          *string   `json:"field,omitempty"`
	ShowPieChart   *bool     `json:"show_pie_chart,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	ShowDataTable  *bool     `json:"show_data_table,omitempty"`
	DataTableLimit *int      `json:"data_table_limit,omitempty"`
	SortOrder      *string   `json:"sort_order,omitempty"`
}

// Dashboard represent a graylog widget's timerange
type Timerange struct {
	Type  string `json:"type"`
	Range int    `json:"range"`
}

// GetByTitle return a stream by its title
func (dashboards Dashboards) GetByTitle(title string) (Dashboard, bool) {
	for _, dashboardFound := range dashboards.Dashboards {
		if dashboardFound.Title == title {
			return dashboardFound, true
		}
	}

	return Dashboard{}, false
}

// GetByDescription return a stream by its title
func (dashboard Dashboard) GetByDescription(description string) (Widget, bool) {
	for _, widgetFound := range dashboard.Widgets {
		if widgetFound.Description == description {
			return widgetFound, true
		}
	}

	return Widget{}, false
}
