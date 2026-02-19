package removeinvalidparantheses

func removeInvalidParentheses(s string) []string {
	res := []string{}
	visited := map[string]bool{s: true}
	queue := []string{s}
	found := false

	for len(queue) > 0 && !found {
		// обрабатываем весь текущий уровень BFS
		nextQueue := []string{}

		for _, cur := range queue {
			if isValid(cur) {
				res = append(res, cur)
				found = true // нашли минимальное удаление, глубже не идём
			}

			if !found { // генерируем только если ещё не нашли ответ
				for i := 0; i < len(cur); i++ {
					if cur[i] != '(' && cur[i] != ')' {
						continue
					}
					next := cur[:i] + cur[i+1:]
					if !visited[next] {
						visited[next] = true
						nextQueue = append(nextQueue, next)
					}
				}
			}
		}
		queue = nextQueue
	}

	if len(res) == 0 {
		return []string{""}
	}
	return res
}

func isValid(s string) bool {
	count := 0
	for _, ch := range s {
		switch ch {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}
