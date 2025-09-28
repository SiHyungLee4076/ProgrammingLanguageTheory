package main

import (
	"fmt"
	"strings"
)

// max 함수 정의
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min 함수 정의
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 입력 문자열을 파싱하여 함수명과 인수를 추출하는 함수
func parseInput(input string) (string, int, int, error) {
	// 입력 문자열에서 공백 제거
	input = strings.TrimSpace(input)

	// 함수명과 인수 추출
	var funcName string
	var a, b int
	n, err := fmt.Sscanf(input, "%3s(%d,%d)", &funcName, &a, &b)
	if n != 3 || err != nil {
		return "", 0, 0, fmt.Errorf("잘못된 입력 형식입니다")
	}
	return funcName, a, b, nil
}

func main() {
	var input string

	fmt.Print("함수를 입력하세요 (예: max(20,30) 또는 min(10,50)): ")
	fmt.Scanln(&input)

	funcName, a, b, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch funcName {
	case "max":
		fmt.Printf("max(%d, %d) = %d\n", a, b, max(a, b))
	case "min":
		fmt.Printf("min(%d, %d) = %d\n", a, b, min(a, b))
	default:
		fmt.Println("지원하지 않는 함수입니다. max 또는 min 함수를 사용하세요.")
	}
}
