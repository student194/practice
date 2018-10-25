package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

///////////////////定义Point类型/////////////////////////////////////
type Point struct {
	x, y int
}

func (p *Point) Show() {
	fmt.Println("the x-at is", p.x, "the y at is", p.y)
}
func (p *Point) Distance(q *Point) float64 {
	var a1, b1 = float64(p.x), float64(p.y)
	var a2, b2 = float64(q.x), float64(q.y)
	return math.Sqrt(math.Pow(a1-a2, 2) + math.Pow(b1-b2, 2))
}

///////////////////////Point类定义终止//////////////////////////////
type Point_que []Point

func (p Point_que) Len() int {
	return len(p)
}
func (p Point_que) Less(i, j int) bool {
	if p[i].x != p[j].x {
		return p[i].x < p[j].x
	} else {
		return p[i].y < p[j].y
	}
}
func (p Point_que) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Point_que) reverse() []Point {
	length := len(p)
	ans := make([]Point, length, length)
	for i := 0; i < length; i++ {
		ans[length-i-1] = p[i]
	}
	return ans
}

func (p *Point_que) Convexll() Point_que {
	sort.Sort((*p))
	var lUpper = make([]Point, 0, 5)
	var lLower = make([]Point, 0, 5)
	lUpper = append(lUpper, (*p)[0], (*p)[1])
	for _, k := range (*p)[2:] {
		lUpper = append(lUpper, k)
		for len(lUpper) >= 3 && !check(lUpper[len(lUpper)-1], lUpper[len(lUpper)-2], lUpper[len(lUpper)-3]) {
			temp := lUpper[len(lUpper)-1]
			lUpper = lUpper[:len(lUpper)-2]
			lUpper = append(lUpper, temp)
		}
	}
	rque := p.reverse()
	lLower = append(lLower, rque[0], rque[1])
	for _, v := range rque[2:] {
		lLower = append(lLower, v)
		for len(lLower) >= 3 && !check(lLower[len(lLower)-1], lLower[len(lLower)-2], lLower[len(lLower)-3]) {
			temp := lLower[len(lLower)-1]
			lLower = lLower[:len(lLower)-2]
			lLower = append(lLower, temp)
		}
	}
	for _, v := range lLower[1 : len(lLower)-1] {
		lUpper = append(lUpper, v)
	}
	return Point_que(lUpper)
}
func check(c, b, a Point) bool {
	var x1, y1 = float64(a.x) - float64(c.x), float64(a.y) - float64(c.y)
	var x2, y2 = float64(b.x) - float64(c.x), float64(b.y) - float64(c.y)
	temp := x1*y2 - y1*x2
	if temp > 0 {
		return true
	} else {
		return false
	}
}

////////////////////////Point_que类型终止/////////////////////////////////////
func createPoint(size int) Point_que {
	temp := make([]Point, 0, size)
	for i := 0; i < size; i++ {
		temp = append(temp, Point{rand.Int() % 100, rand.Int() % 100}) //添加随机生成点
	}
	return Point_que(temp)
}
