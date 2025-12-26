package main

func findNameById(name map[int]string, id int) (string, bool) {
	value, ok := name[id]
	if ok {
		return value, ok
	}
	return "", ok
}
