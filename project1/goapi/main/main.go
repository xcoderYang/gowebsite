package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("yangxuechao")
	// fmt.Println(buf.Truncate(4))
	// _, err := buf.Read(p)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	p := []byte{'y', 'a', 'n', 'g'}
	q := "xuechao"
	n, err := buf.Write(p)

	if err != nil {
		fmt.Println(err)
	}

	m, err := buf.WriteString(q)

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(buf)
}
