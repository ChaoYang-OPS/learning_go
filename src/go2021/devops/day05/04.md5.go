package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base64
	// 通常说的base64  0-9a-zA-Z+/ 64
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("幸运女神在微笑"))) // 5bm46L+Q5aWz56We5Zyo5b6u56yR
	message, _ := base64.StdEncoding.DecodeString("5bm46L+Q5aWz56We5Zyo5b6u56yR")
	fmt.Println(string(message)) // 幸运女神在微笑
	// 在URL中+_特殊字符，base64url (+(-)/(_)替换)
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("幸运女神在微笑"))) // 5bm46L-Q5aWz56We5Zyo5b6u56yR
	message, _ = base64.URLEncoding.DecodeString("5bm46L-Q5aWz56We5Zyo5b6u56yR")
	fmt.Println(string(message)) // 幸运女神在微笑

	// 标准
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte("幸运女神在微笑")))
	// 5bm46L+Q5aWz56We5Zyo5b6u56yR
	// url
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("幸运女神在微笑")))
	// 5bm46L-Q5aWz56We5Zyo5b6u56yR
}
