package dcu

func appendStringToSliceIfMissing(slice []string, i string) []string {
	if sliceContainsString(slice, i) {
		return slice
	}
	return append(slice, i)
}

func sliceContainsString(slice []string, i string) bool {
	for _, ele := range slice {
		if ele == i {
			return true
		}
	}
	return false
}
