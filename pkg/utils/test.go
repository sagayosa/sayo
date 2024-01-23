package utils

import "fmt"

func ComparisonFailure(ans interface{}, wrongAns interface{}) error {
	return fmt.Errorf("comparison failure:\n [correct result]: %v\n [incorrect result]: %v", ans, wrongAns)
}

func MissingFailure(miss interface{}) error {
	return fmt.Errorf("missing failure:\n [missing]: %v", miss)
}
