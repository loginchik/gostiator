package models

import "fmt"

func StringPageRange(start string, end string) string {
	if (start == end) && (start == "") {
		return ""
	}
	if start == end {
		return start
	}
	if !(start == "") && !(end == "") {
		return fmt.Sprintf("%s-%s", start, end)
	}
	if start == "" {
		return end
	} else {
		return start
	}

}
