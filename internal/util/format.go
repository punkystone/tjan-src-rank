//nolint:mnd // ignore
package util

import (
	"fmt"
	"strings"
)

func FormatSeconds(seconds float64) string {
	totalMilliseconds := int64(seconds * 1000)

	hours := totalMilliseconds / (3600 * 1000)
	remaining := totalMilliseconds % (3600 * 1000)

	minutes := remaining / (60 * 1000)
	remaining %= (60 * 1000)

	secs := remaining / 1000
	millis := remaining % 1000

	var parts []string
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	if secs > 0 || millis > 0 || (hours == 0 && minutes == 0) {
		parts = append(parts, fmt.Sprintf("%ds", secs))
	}
	if millis > 0 {
		parts = append(parts, fmt.Sprintf("%dms", millis))
	}

	return strings.Join(parts, " ")
}
