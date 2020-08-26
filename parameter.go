package gutils

import "sort"

// IdsUniqueFitter ids去重
func IdsUniqueFitter(ids []int) []int {
	sort.Ints(ids)
	var newIds []int
	var lastID int
	for i, id := range ids {
		if i == 0 {
			lastID = id
			newIds = append(newIds, id)
		} else if id != lastID {
			lastID = id
			newIds = append(newIds, id)
		}
	}
	return newIds
}
