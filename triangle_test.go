package workspace

import (
	"testing"
	"fmt"
)

func TestIsOk(t *testing.T) {
	rs :=IsOk(10,20,29)
	if rs {
		fmt.Println("10,20,29 ok")
	} else {
		fmt.Println("10,20,29 not ok")
	}
	rs = IsOk(0,1,1)
	if rs {
		fmt.Println("0,1,1 ok")
	} else {
		fmt.Println("0,1,1 not ok")
	}
}

func BenchmarkIsOk(b *testing.B) {
	for i:=0;i<b.N;i++ {
		rs :=IsOk(10,20,29)
		if rs {
			fmt.Println("ok")
		} else {
			fmt.Println("not ok")
		}
	}
}

func TestTriangles(t *testing.T) {
	rs := Triangles(99,88,170)
	fmt.Printf("%d,%d,%d %s\n",99,88,170,rs)
	rs = Triangles(99,99,99)
	fmt.Printf("%d,%d,%d %s\n",99,99,99,rs)
	rs = Triangles(88,88,99)
	fmt.Printf("%d,%d,%d %s\n",88,88,99,rs)
	rs = Triangles(-99,99,99)
	fmt.Printf("%d,%d,%d %s\n",-99,99,99,rs)
}

func BenchmarkTriangles(b *testing.B) {
	for i:=0;i<b.N;i++ {
		rs := Triangles(i,i+1,i+2)
		fmt.Printf("%d,%d,%d %s\n",i,i+1,i+2,rs)
	}
}