package bulk

import (
	"github.com/uniplaces/logfairy/dto/dashboard"
	"github.com/uniplaces/logfairy/dto/stream"
)

// Bulk represents the structure that can be handled in a bulk creation/update action
type Bulk struct {
	Streams    []stream.Stream       `json:"streams"`
	Dashboards []dashboard.Dashboard `json:"dashboards"`
}
