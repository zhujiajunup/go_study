package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	const (
		i = 1 << iota // 1 << 0 = 1
		j = 3 << iota // 3 << 1 = 6
		k             // 3 << 2 = 12
		l             // 3 << 3 = 24
	)
	fmt.Println(i, j, k, l)
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprint(&b, " World!")
	io.Copy(os.Stdout, &b)

}
