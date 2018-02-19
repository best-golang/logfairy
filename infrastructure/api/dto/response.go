package dto

// StreamCreation represents a positive response of a stream creation
type StreamCreation struct {
	StreamID string `json:"stream_id"`
}

// DashboardCreation represents a positive response of a dashboard creation
type DashboardCreation struct {
	DashboardID string `json:"dashboard_id"`
}

// WidgetCreation represents a positive response of a widget creation
type WidgetCreation struct {
	WidgetID string `json:"widget_id"`
}

// ErrorResponse represents a negative response from graylog
type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
