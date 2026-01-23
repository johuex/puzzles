package binaryindexletter

func nextGreatestLetter(letters []byte, target byte) byte {

	l := 0
	r := len(letters) - 1

	for l <= r {
		m := l + ((r - l) / 2) // because we can't devide by 0
		tmpN := letters[m]
		if tmpN < target {
			l = m + 1
		} else if tmpN > target {
			r = m - 1
		} else {
			// need to find first greater than target elem
			idx := m
			for letters[m] == letters[idx] {
				idx++
				if idx > len(letters)-1 {
					// when character is last in array
					// it means that no nearest char that greate than it
					// return first elem
					return letters[0]
				}
			}
			return letters[idx]
		}
	}

	if l >= len(letters) || l <= 0 {
		// 'l' and 'r' crossed
		// so targer not found, return first elem by default
		return letters[0]
	}
	return letters[l]
}
