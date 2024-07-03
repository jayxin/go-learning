package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

/*type error interface {
  Error() string
}*/

/*type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func Stat(name string) (fi FileInfo, err error) {
	var stat syscall.Stat_t
	err = syscall.Stat(name, &stat)
	if err != nil {
		return nil, &PathError{"stat", name, str}
	}
	return fileInfoFromStat(&stat, name), nil
}*/

func main() {
	n, err := Foo(1)
	if err != nil {
		fmt.Println("Error!")
		return
	} else {
		fmt.Println(n)
	}

	fi, err := os.Stat("aa.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			fmt.Println(e.Err.Error())
		}
	} else {
		fmt.Println(fi)
	}
}

func Foo(param int) (n int, err error) {
	if param < 0 {
		err = errors.New("Parameter should be a non-negative number.")
		return
	}
	return 1, nil
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

//func panic(interface{})
// panic(404)
// panic("network broken")
// panic(Error("File is not existed."))
//func recover() interface{}
/*defer func() {
  if r := recover(); r != nil{
    log.Printf("Runtime error caught: %v", r)
  }
}()*/
