package main

func isValidSerialization(preorder string) bool {
	slots := 1
	for i := 0; i < len(preorder); {
		if slots == 0 {
			return false
		}
		switch preorder[i] {
		case ',':
			i++
		case '#':
			slots--
			i++
		default:
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
			slots++
		}
	}
	return slots == 0
}
