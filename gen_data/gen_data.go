package main

import (
	"fmt"
	"os"
	"math/rand"
	// "errors"
)
const (
	WORKDIR		string 	= "../iptrie/"
	INPUTSIZE	int 	= 2000000
	TESTSIZE	int 	= 10000
	INPUTFILE	string 	= WORKDIR + "input.txt"
	TESTFILE 	string 	= WORKDIR + "test.txt"
)

func checkFileIsExist(filename string) (bool) {
	var exist = true;
 	if _, err := os.Stat(filename); os.IsNotExist(err) {
  	exist = false;
 	}
 	return exist;
}

func gen_data(filename string) error {
	var (
		file 	*os.File
		err 	error
	)
	if checkFileIsExist(filename) {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
		if err != nil {
		return err
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
		return err
		}
	}
	
	defer file.Close()

	for i := 0; i < INPUTSIZE; i++ {
		fmt.Fprintf(file, "%d.%d.%d.0, 24, %v\n", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(100))
	}
	// for i := 0; i < 256; i++ {
	// 	for j := 0; j < 256; j++ {
	// 		for k := 0; k < 256; k++ {
	// 			fmt.Fprintf(file, "%d.%d.%d.0, 24, %v\n", i, j, k, rand.Intn(100))
	// 		}
	// 	}
	// }
	return nil
}

func gen_test(NUMBER int) error {
	testfile, err := os.Create(TESTFILE)
	if err != nil {
		return err
	}
	fmt.Fprintf(testfile, "%v\n", NUMBER)
	for i := 0; i < NUMBER; i++ {
		fmt.Fprintf(testfile, "%d.%d.%d.%d\n", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(256))
	}
	return nil
}


func main() {
	gen_data(INPUTFILE)
	gen_test(TESTSIZE)
}
