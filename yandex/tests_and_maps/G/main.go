package g

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])

	// create table
	table := make([]string, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		table[i] = line
	}

	// logic
	ans := fiveRow(n, m, table)
	// output
	writer.WriteString(ans)
	writer.WriteByte('\n')
}

func fiveRow1(n, m int, table []string) string {
	// метод основан не на массивах с коэффициентами сдвига
	// из каждой точки в 4 направления (с помощью коэф сдвига) пытаемся идти на 5 клеток вперед.
	// если вышли за пределы поля или фигура не совпадает -- выходим и идем дальше
	// получается 4 вложенных цикла внутри, но размер поля всего 1000*1000

	// можно еще кстати использовать идею динамического программирования, про которое рассказывалось в след лекции
	return ""
}

func fiveRow(n, m int, table []string) string {
	// тупое и топорное решение
	var ans bool

	lineSet := make(map[int][]rune, n)

	// transform input to 2D array
	for i, row := range table {
		lineSet[i] = make([]rune, m)
		for j, sym := range row {
			lineSet[i][j] = sym
		}
	}

	// check rows
	for i := 0; i < len(lineSet) && !ans; i++ {
		runeCount := 0
		negativeRune := '.'
		for j := 0; j < len(lineSet[i]) && !ans; j++ {
			if negativeRune == '.' {
				negativeRune = negativeSwitch(lineSet[i][j])
			}
			if lineSet[i][j] == negativeRune || lineSet[i][j] == '.' {
				// so we drop check sequence only in case of matching rune with negative rune
				if lineSet[i][j] == '.' {
					runeCount = 0
				} else {
					runeCount = 1
				}
				negativeRune = negativeSwitch(lineSet[i][j])
				continue
			}
			runeCount += 1
			if runeCount == 5 {
				ans = true
			}
		}
	}

	// check columns
	rows := len(lineSet)
	cols := len(lineSet[0])

	for j := 0; j < cols && !ans; j++ {
		runeCount := 0
		negativeRune := '.'
		for i := 0; i < rows && !ans; i++ {
			if negativeRune == '.' {
				negativeRune = negativeSwitch(lineSet[i][j])
			}
			if lineSet[i][j] == negativeRune || lineSet[i][j] == '.' {
				// so we drop check sequence only in case of matching rune with negative rune
				if lineSet[i][j] == '.' {
					runeCount = 0
				} else {
					runeCount = 1
				}
				negativeRune = negativeSwitch(lineSet[i][j])
				continue
			}
			runeCount += 1
			if runeCount == 5 {
				ans = true
			}
		}
	}
	if !ans {
		runeCount := 0
		negativeRune := '.'
		for i := 0; !ans && i < rows && i < cols; i++ {
			if negativeRune == '.' {
				negativeRune = negativeSwitch(lineSet[i][i])
			}
			if lineSet[i][i] == negativeRune || lineSet[i][i] == '.' {
				// so we drop check sequence only in case of matching rune with negative rune
				if lineSet[i][i] == '.' {
					runeCount = 0
				} else {
					runeCount = 1
				}
				negativeRune = negativeSwitch(lineSet[i][i])
				continue
			}
			runeCount += 1
			if runeCount == 5 {
				ans = true
			}
		}

		runeCount = 0
		negativeRune = '.'
		for i := 0; !ans && i < rows && i < cols; i++ {
			if negativeRune == '.' {
				negativeRune = negativeSwitch(lineSet[i][i])
			}
			if lineSet[i][i] == negativeRune || lineSet[i][i] == '.' {
				// so we drop check sequence only in case of matching rune with negative rune
				if lineSet[i][i] == '.' {
					runeCount = 0
				} else {
					runeCount = 1
				}
				negativeRune = negativeSwitch(lineSet[i][i])
				continue
			}
			runeCount += 1
			if runeCount == 5 {
				ans = true
			}
		}
	}
	// check diags
	if !ans {
		for d := 0; !ans && d < rows+cols-1; d++ {
			runeCount := 0
			negativeRune := '.'
			for i := 0; !ans && i < rows; i++ {
				j := d - i
				if j >= 0 && j < cols {
					if negativeRune == '.' {
						negativeRune = negativeSwitch(lineSet[i][j])
					}
					if lineSet[i][j] == negativeRune || lineSet[i][j] == '.' {
						// so we drop check sequence only in case of matching rune with negative rune
						if lineSet[i][j] == '.' {
							runeCount = 0
						} else {
							runeCount = 1
						}
						negativeRune = negativeSwitch(lineSet[i][j])
						continue
					}
					runeCount += 1
					if runeCount == 5 {
						ans = true
					}
				}
			}
		}

		for d := 0; !ans && d < rows+cols-1; d++ {
			runeCount := 0
			negativeRune := '.'
			for i := 0; !ans && i < rows; i++ {
				j := (cols - 1) - (d - i)
				if j >= 0 && j < cols {
					if negativeRune == '.' {
						negativeRune = negativeSwitch(lineSet[i][j])
					}
					if lineSet[i][j] == negativeRune || lineSet[i][j] == '.' {
						// so we drop check sequence only in case of matching rune with negative rune
						if lineSet[i][j] == '.' {
							runeCount = 0
						} else {
							runeCount = 1
						}
						negativeRune = negativeSwitch(lineSet[i][j])
						continue
					}
					runeCount += 1
					if runeCount == 5 {
						ans = true
					}
				}
			}
		}
	}

	if ans {
		return "Yes"
	}
	return "No"
}

func negativeSwitch(elem rune) rune {
	switch elem {
	case 'X':
		return 'O'
	case 'O':
		return 'X'
	default:
		// special case
		return '.'
	}
}
