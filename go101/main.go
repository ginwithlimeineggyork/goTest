package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
	"strings"
)

func mapp() map[string]int {
	m:=make(map[string]int)
	m["h"]=89
	m["k"]=12
	return m
}

func goroo() {
	fmt.Println("This is too late/go routine")
}

type Bmi struct {
	height,weight float32	
}

type Profit struct {
	old, current float32
}

func (p Profit) Output() float32 {
	return (p.current-p.old)/p.old
}

func (b Bmi) Output() float32 {
	h := math.Exp2(float64(b.height/100))
	return b.weight/float32(h)
}

type Output interface {
	Output() float32
}

type error_name interface {
	Error() string
}

//var numFans = 4

//type Fan struct {
//	age int
//	language string
//}

//fans :=[{
//	age:1,
//	language:"Mandarin"
//},{age:15,language:"Putonghua"}]

//func coutFans() {
//	for k:=0;k<numFans;k++ {
//		if fans[k].age > 2 {
//			numFans-=1
//			fmt.Println("num: ",numFans)
//		}
//	}
//}

func joinstring(element...string)string{
    return strings.Join(element, "-")
}

func main() {
	var t = mapp()
	fmt.Println(t)
	if map_val, ishere := t["h"]; ishere {
		fmt.Println(map_val)
	}
	if map_val_k, ishere := t["a"]; ishere {
		fmt.Println("ads",map_val_k)
	}
	fmt.Println(len(t))
	delete(t,"h")
	_,ishere:=t["h"]
	fmt.Println("is t map 'h' here: ",ishere)
	d := 1;
	var dd = &d;
	go goroo()
	fmt.Println("address: ",dd)
	time.Sleep(1 * time.Second)
	fmt.Println(*dd)
	rrr := make([]int,7)
	f:=&rrr[0]
	fmt.Println(f)
	rrr = append(rrr,4)
	fmt.Println(rrr)
	rrr = append(rrr,5,6,7)
	fmt.Println(rrr)
	fmt.Println(rrr[1:5])
	fmt.Println("cap: rrr slice 1:5 => %d",cap(rrr[1:5]))
	bmiObj := Bmi{height:188,weight:60}
	profitObj := Profit{old:10,current:14}
	var otu Output
	otu = bmiObj
	fmt.Println(otu.Output())
	otu = profitObj
	fmt.Println(otu.Output())
	ggg:=make([]string,5)
	hhh:=[]string{"112"}
	ggg=append(ggg,"1","e","r","5")
	fmt.Println(ggg)
	copy(ggg,hhh)
	fmt.Println(ggg)
	fk:=[]float64{12.233,14.55,12323.23,5.12}
	fmt.Println(fk)
	fe:="thisisit:%d"
	gr:=fmt.Sprintf(fe,2050)
	fmt.Println(gr)
	w:=struct{}{}
	fmt.Println(unsafe.Sizeof(w))
	new_map_obj := make(map[string]struct{})
	for _, j := range []string{"1","h","k"} {
		new_map_obj[j] = struct{}{}
	}
	fmt.Println(new_map_obj)
	maop:=map[string]bool{"i":true,"l":false,"1":true}
	map2:=make(map[string]bool)
	for k,u:=range maop {
		map2[k]=u
	}
	fmt.Println(map2)
	//coutFans()
	fmt.Println(joinstring("Interview", "Bit"))
    fmt.Println(joinstring("Golang", "Interview", "Questions"))
}
