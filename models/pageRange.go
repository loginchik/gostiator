package models

import "fmt"

type PageRange struct {
	Start string
	End   string
}

func (pr PageRange) StringRange() string {
	if (pr.Start == pr.End) && (pr.Start == "") {
		return ""
	}

	if pr.Start == pr.End {
		return fmt.Sprintf("%s", pr.Start)
	} else {
		if !(pr.Start == "") && !(pr.End == "") {
			return fmt.Sprintf("%s-%s", pr.Start, pr.End)
		} else if pr.Start == "" {
			return fmt.Sprintf("%s", pr.End)
		} else {
			return fmt.Sprintf("%s", pr.Start)
		}
	}
}
