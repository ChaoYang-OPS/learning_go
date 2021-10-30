package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// hash算法===> 签名(不可逆)

	fmt.Printf("%x\n", md5.Sum([]byte("Kubernetes")))

	hasher := md5.New()
	hasher.Write([]byte("Lsunstack"))
	hasher.Write([]byte("Kubernetes"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil))) // 5a2522794107b07bda5e180fd1f3bbf2

	fmt.Printf("%x\n", []byte("您好"))
	fmt.Println(hex.EncodeToString([]byte("您好"))) // e682a8e5a5bd
	txt, _ := hex.DecodeString("e682a8e5a5bd")
	fmt.Println(string(txt)) // 您好
}
