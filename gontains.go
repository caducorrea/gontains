package gontains

func Gontains(slice []string, item string) (has bool) {
	for _, i := range slice {
		if i == item {
			has = true
			return
		}
	}
	return
}
