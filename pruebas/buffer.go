package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 4096*4096)

	ntemp, err := readLine(reader)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	intTemp, err := strconv.ParseInt(ntemp, 10, 64)
	checkError(err)
	n := int32(intTemp)

	m := make(map[string]string)

	for i := 0; i < int(n); i++ {
		stemp, err := readLine(reader)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		arrTemp := strings.Split(stemp, " ")
		m[arrTemp[0]] = arrTemp[1]
	}

	for {
		r, err := readLine(reader)
		if err != nil {
			fmt.Println("no mas busquedas")
			return
		}
		if m[r] == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", r, m[r])
		}

	}

}

func readLine(reader *bufio.Reader) (string, error) {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return "", err
	}

	return strings.TrimRight(string(str), "\r\n"), nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
