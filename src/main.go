/**
 * Created with IntelliJ IDEA.
 * User: jhaddad
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"encoding/binary"
	"bytes"
)

func main() {
	fmt.Println("Hello world!")
	t := time.Now().Nanosecond()
	fp, _ := os.Create("/tmp/numbers.binary")

	rand.Seed(int64(t))

	buf := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		k := rand.Int63()
		fmt.Println(k)

		binary.Write(buf, binary.LittleEndian, k)
		// write the buffer

		fp.Write(buf.Bytes())
	}
	fmt.Println("goodbye")
	fp.Close()
}
