package array

func UniqueNoNullSlice(slice ...string) (newSlice []string) {
	found := make(map[string]bool)
	for _, val := range slice {
		if val == "" {
			continue
		}
		if _, ok := found[val]; !ok {
			found[val] = true
			newSlice = append(newSlice, val)
		}
	}
	return
}

func StringArraySplit(slice []string, limit int) [][]string {
	var count = len(slice)
	if count <= limit {
		return [][]string{slice}
	}
	var begin = 0
	var end = limit
	var result [][]string
	for {
		if end > count {
			end = count
		}
		if begin >= count {
			break
		}
		result = append(result, slice[begin:end])

		begin += limit
		end += limit
	}
	return result
}

func StringArrayDifference(old []string, new []string) (additions []string, deletions []string) {
	additionsMap := make(map[string]struct{})
	deletionsMap := make(map[string]struct{})
	for _, s := range old {
		if len(s) > 0 {
			deletionsMap[s] = struct{}{}
		}
	}
	for _, s := range new {
		if len(s) > 0 {
			additionsMap[s] = struct{}{}
		}
	}
	for s, _ := range additionsMap {
		if _, ok := deletionsMap[s]; !ok {
			additions = append(additions, s)
		}
	}
	for s, _ := range deletionsMap {
		if _, ok := additionsMap[s]; !ok {
			deletions = append(deletions, s)
		}
	}
	return
}

func StringSetToArray(set map[string]bool) []string {
	result := make([]string, 0, len(set))
	for s, ok := range set {
		if ok {
			result = append(result, s)
		}
	}
	return result
}

func StringArrayToSet(slice []string) map[string]bool {
	result := make(map[string]bool, len(slice))
	for _, s := range slice {
		result[s] = true
	}
	return result
}
