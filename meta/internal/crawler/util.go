package crawler

func contains(s []string, item string) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

func deleteDuplicates(s []string) []string {
	var arr []string
	for _, item := range s {
		if !contains(arr, item) {
			arr = append(arr, item)
		}
	}
	return arr
}
