package i

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
	x, _ := strconv.Atoi(lineArr[0])
	y, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	f, _ := strconv.Atoi(lineArr[0])
	g, _ := strconv.Atoi(lineArr[1])

	// logic
	ans := route(x, y, f, g)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func route(x, y, f, g int) int {
	// движение прямо или направо -- это всегда 3 приема
	// только поворот налево -- один прием
	dx := abs(x, f)
	dy := abs(y, g)

	if dx == dy && dy == 0 {
		return 0
	}
	if dx == 0 || dy == 0 {
		// соответсвенно если надо двигаться с севера на юг или с запада на восток
		// то мы всегда движемся или прямо или направо (которые всегда за три приема)
		return (max(dx, dy) - 1) * 3
	}
	// вычитаем тут 2, тк левый поворот за один прием
	return (dx+dy-2)*3 + 1
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
