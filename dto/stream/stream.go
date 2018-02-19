package stream

import "time"

// Streams represent a list of graylog stream
type Streams struct {
	Total   int      `json:"total"`
	Streams []Stream `json:"streams"`
}

// Stream represent a graylog stream
type Stream struct {
	ID                             *string    `json:"id,omitempty"`
	CreatorUserID                  *string    `json:"creator_user_id,omitempty"`
	MatchingType                   string     `json:"matching_type"`
	Description                    string     `json:"description"`
	CreatedAt                      *time.Time `json:"created_at,omitempty"`
	Disabled                       *bool      `json:"disabled,omitempty"`
	Rules                          []Rule     `json:"rules"`
	Title                          string     `json:"title"`
	RemoveMatchesFromDefaultStream bool       `json:"remove_matches_from_default_stream"`
	IndexSetID                     string     `json:"index_set_id"`
}

// Rule represent a graylog rule
type Rule struct {
	Field       string  `json:"field"`
	StreamID    *string `json:"stream_id,omitempty"`
	Description string  `json:"description"`
	ID          *string `json:"id,omitempty"`
	Type        int     `json:"type"`
	Inverted    bool    `json:"inverted"`
	Value       string  `json:"value"`
}
