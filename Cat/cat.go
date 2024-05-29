package cat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Cat(f *os.File, x int) {
	i := 1
	fileReader := bufio.NewReader(f)
	for {
		byteData, err := fileReader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			if len(byteData) != 0 {
				fmt.Fprintf(os.Stdout, "%5d  %s", i, byteData)
			}
			break
		}
		if x != -1 && i-1 == x {
			break
		}
		fmt.Fprintf(os.Stdout, "%5d  %s", i, byteData)
		i++
	}
}
