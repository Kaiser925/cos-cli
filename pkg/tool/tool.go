package tool

import (
	"fmt"
)

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
)

func GetByteSize(in int64) string {
	i := float64(in)
	switch {
	case i >= KB && i < MB:
		return fmt.Sprintf("%.1fK", i/KB)
	case i >= MB && i < GB:
		return fmt.Sprintf("%.1fM", i/MB)
	case i >= GB && i < TB:
		return fmt.Sprintf("%.1fG", i/GB)
	case i >= TB:
		return fmt.Sprintf("%.1fT", i/TB)
	default:
		return fmt.Sprintf("%.0fB", i)
	}
}
