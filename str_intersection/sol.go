package str_intersection

import "strings"

// FindIntersection(strArr) read the array of strings stored in strArr which will contain 2 elements:
// the first element will represent a list of comma-separated numbers sorted in ascending order,
// the second element will represent a second list of comma-separated numbers (also sorted).
// Your goal is to return a comma-separated string containing the numbers that occur in elements of strArr in sorted order.
// If there is no intersection, return the string false.

func FindIntersection(strArr []string) string {
	type Empty struct{}
	m := map[string]Empty{}

	a1 := strings.Split(strArr[0], ", ")
	for _, val := range a1 {
		m[val] = Empty{}
	}

	a2 := strings.Split(strArr[1], ", ")
	var is []string
	for _, val := range a2 {
		if _, ok := m[val]; ok {
			is = append(is, val)
		}
	}
	if len(is) == 0 {
		return "false"
	}

	return strings.Join(is, ", ")
}
