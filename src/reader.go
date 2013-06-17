/**
 * Created with IntelliJ IDEA.
 * User: jonhaddad
 * Date: 6/16/13
 * Time: 4:21 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"os"
	"encoding/binary"
	"bytes"
	"io"
)

func main() {
	fmt.Printf("Hello world!")
	fp, _ := os.Open("/tmp/numbers.binary")
	buf := make([]byte, 8)

	for true {
		_, err := fp.Read(buf)

		if err == io.EOF {
			break
		}

		fmt.Println(buf)

		buffer := bytes.NewBuffer(buf)
		var k int64
		binary.Read(buffer, binary.LittleEndian, &k)
		fmt.Println(k)
	}
	fmt.Println("DONE")
}
