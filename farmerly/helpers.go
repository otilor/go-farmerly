package farmerly


func hasEmptyValues(values map[string][]string) bool{
	if len(values) == 0{
		return true
	}
	return false
}