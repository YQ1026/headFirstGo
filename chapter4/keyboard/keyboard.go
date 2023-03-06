// 从键盘读取用户输入
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 从键盘读一个浮点数
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
