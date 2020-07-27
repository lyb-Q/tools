package main

import (
	"bufio"
	"io"
	"os"
)

//流处理
func ReadFile(filePath string, handle func(string)) (err error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		handle(string(line))
	}
}
func handle(deal string) {
	//do something
}

func main() {
	path := ""
	ReadFile(path, handle)
}

//分片处理
func ReadBigFile(fileName string, handle func([]byte)) (err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	s := make([]byte, 4096)
	for {
		switch n, err := f.Read(s[:]); true {
		case n < 0:
			os.Exit(1)
		case n == 0: // EOF
			return nil
		case n > 0:
			handle(s[0:n])
		}
	}
	return nil
}
