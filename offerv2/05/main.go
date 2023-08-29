package main

func replaceSpace(s string) string {
	var count int
	for _, r := range s {
		if r == ' ' {
			count += 3
		} else {
			count++
		}
	}

	res := make([]rune, count)
	var idx int
	for _, r := range s {
		if r == ' ' {
			res[idx] = '%'
			res[idx+1] = '2'
			res[idx+2] = '0'
			idx += 3
		} else {
			res[idx] = r
			idx++
		}
	}

	return string(res)
}
