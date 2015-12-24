package iptrie

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func TestFromFile(t *testing.T) {
	var trie IPTrie
	trie.root = NewIPTrie()
	var (
		outf *os.File
		err  error
	)

	// open ot create log file
	if checkFileIsExist("./log.out") {
		outf, err = os.OpenFile("./log.out", os.O_WRONLY, 0666)
		if err != nil {
			t.Error()
		}
	} else {
		outf, err = os.Create("./log.out")
		if err != nil {
			t.Error()
		}
	}
	defer outf.Close()

	start := time.Now()
	if checkFileIsExist("./input.txt") {
		fmt.Println("Loading from file: ./input.txt")
		trie.LoadFromFile("./input.txt")
	} else {
		t.Error()
		return
	}

	duration := time.Since(start)
	fmt.Fprintf(outf, "Load time duration: %v\n", duration.Seconds())

	testf, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("Test file wrong open")
		t.Error()
		return
	}
	defer testf.Close()

	var num int
	fmt.Fscanf(testf, "%d", &num)
	fmt.Fprintln(outf, num)

	fmt.Println("Reading test file: ./test.txt")
	hit := 0
	los := 0
	start = time.Now()
	for i := 0; i < num; i++ {
		var str string
		argc, err := fmt.Fscanf(testf, "%s", &str)
		if err != nil || argc == 0 {
			t.Error()
		}
		// fmt.Fprintf(outf, "%s\n", str)

		match, ok := trie.Get(str)
		if !ok {
			los++
			fmt.Fprintln(outf, str, "Not found")
		} else {
			hit++
			fmt.Fprintln(outf, match)
		}
	}

	testf.Close()

	fmt.Println("Test by loaded file, which all have been in the Trie")
	testf, err = os.Open("./input.txt")
	if err != nil {
		t.Error()
	}

	for {
		var (
			str    string
			length int
			cwnd   int
		)

		argc, err := fmt.Fscanf(testf, "%s%d,%d", &str, &length, &cwnd)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Error()
			}
		}
		if argc > 0 && argc != 3 {
			fmt.Println(str, "Format not right")
		}

		str = str[0 : len(str)-1]

		match, ok := trie.Get(str)
		if !ok {
			los++
			fmt.Fprintf(outf, "Not found\n")
		} else {
			hit++
			fmt.Fprintln(outf, match)
		}
	}

	duration = time.Since(start)
	num += hit
	fmt.Fprintf(outf, "Queries hit: %v, Queries los: %v\nQuery total time: %vs, average time: %vs\n", 
		hit, los,
		duration.Seconds(), float64(duration.Seconds())/float64(hit+los))
	fmt.Printf("Queries hit: %v, Queries los: %v\nQuery total time: %vs, average time: %vs\n", 
		hit, los,
		duration.Seconds(), float64(duration.Seconds())/float64(hit+los))
}
