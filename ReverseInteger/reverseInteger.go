package ReverseInteger

import (
	"fmt"
	"math"
	"strconv"
)

//strconv.Itoa int to string
func reverse1(x int) int {
	var s string
	if x < 0 {
		s = strconv.Itoa(-x)
	} else {
		s = strconv.Itoa(x)
	}

	bs := []byte(s)
	head := 0
	tail := len(bs) - 1
	for head < tail {
		bs[head], bs[tail] = bs[tail], bs[head]
		head++
		tail--
	}
	rs := string(bs)
	if x < 0 {
		rs = "-" + rs
	}
	y, err := strconv.Atoi(rs)
	if err != nil {
		fmt.Println(err, rs)
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

func reverse2(x int) int {
	var question int
	var y = 0
	question = x
	if x < 0 {
		question = -x
	}
	for question > 0 {
		r := question % 10
		y = y*10 + r
		question = question / 10
	}
	if x < 0 {
		y = -y
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

// input: 321
// first: r= 1 , y=1, question = 32
// second r= 2 , y=1*10+2 question = 3
// third r= 3, y = 12*10 + 3 question = 0
//output = 123

//systeam32 會發生溢位 可以用以下方式判斷 怕加玩R 發生溢位 可以y*10+r 後再進行扣Ｒ/10 檢查是否等於前一次的值

func reverse3(x int) int {
	var question int
	var y = 0
	var pre = 0
	question = x
	if x < 0 {
		question = -x
	}
	for question > 0 {
		r := question % 10
		y = y*10 + r
		if (y-r)/10 != pre {
			return 0
		}
		pre = y
		question = question / 10
	}
	if x < 0 {
		y = -y
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}
