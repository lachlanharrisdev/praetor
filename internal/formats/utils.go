package formats

import (
	"fmt"
	"strings"
	"time"
)

func shortenPath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return path
	}
	start := parts[0]
	if start == "" && len(parts) > 1 {
		start = "/" + parts[1]
	}
	end := parts[len(parts)-1]
	shortened := start + "/.../" + end
	if len(shortened) > maxLen {
		return "..." + end
	}
	return shortened
}

func formatFields(fields map[string]any) string {
	parts := make([]string, 0, len(fields))
	for k, v := range fields {
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(parts, " ")
}

func formatTime(ts string) string {
	parsed, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		return ts
	}
	return parsed.Format("2006-01-02 15:04")
}
