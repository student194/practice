package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type mlength struct {
	name   int
	length float64
}
type mlque []mlength

func (p mlque) Len() int {
	return len(p)
}
func (p mlque) Less(i, j int) bool {
	return p[i].length < p[j].length
}
func (p mlque) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

///////////////////////////////////////
type Chrom struct {
	body string
	cost float64
}

func createChrom(sizeP, groupSize int, setAt *Point_que, myMap *Point_que) []Chrom {
	length := len(*setAt)
	var temp string
	var value float64
	var newr int
	ans := make([]Chrom, 0, groupSize)
	for i := 0; i < groupSize; i++ {
		temp, value, newr = "", 0.0, 0
		var check = make(map[int]bool)
		for tt := 0; tt < sizeP; {
			newr = rand.Int() % length
			if !check[newr] {
				temp += (strconv.Itoa(newr) + "_")
				check[newr] = true
				tt++
			}
		}
		value = cost(temp, setAt, myMap) ///////////计算相应成本
		ans = append(ans, Chrom{temp, value})
	}
	return ans
}
func cost(aim string, setAt *Point_que, myMap *Point_que) float64 {
	var ans, temp = Qdistance(aim, setAt, (*myMap)[0]), 0.0
	for _, v := range *myMap {
		temp = Qdistance(aim, setAt, v)
		if ans < temp {
			ans = temp
		}
	}
	return ans
}
func (p *Chrom) ccost(setAt *Point_que, myMap *Point_que) {
	aim := p.body
	var ans, temp = Qdistance(aim, setAt, (*myMap)[0]), 0.0
	for _, v := range *myMap {
		temp = Qdistance(aim, setAt, v)
		if ans < temp {
			ans = temp
		}
	}
	p.cost = ans
}
func max3(a, b, c float64) float64 {
	if a > b && b > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}
func Qdistance(aaim string, setAt *Point_que, aimp Point) float64 {
	aim := strings.Split(aaim, "_")
	aim = aim[:len(aim)-1]
	var ptr, _ = strconv.Atoi(string(aim[0]))
	var temp float64 = 0.0
	var ans float64 = (*setAt)[ptr].Distance(&aimp)
	for _, v := range aim {
		ptr, _ = strconv.Atoi(string(v))
		temp = (*setAt)[ptr].Distance(&aimp)
		if ans > temp {
			ans = temp
		}
	}
	return ans
}

type Population []Chrom

func (p Population) Len() int {
	return len(p)
}
func (p Population) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}
func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Population) genetic(echo int, setAt *Point_que, myMap *Point_que) []Chrom {
	last_ans := make([]Chrom, 0, 100) ////////////////////////////////////////////////100echo
	changNum := len(*setAt) - len((*p)[0].body)
	var avg, good float64 = 0.0, 0.9
	var nowbest float64 = 0.0
	for i := 0; i < echo; i++ {
		for _, v := range *p {
			avg += v.cost
		}
		avg /= float64(len(*p))
		for j := 0; j < len(*p); j++ {
			switch {
			case (*p)[j].cost == nowbest:

			case (*p)[j].cost >= avg:
				if myRand1() < good {
					reOrganize(&((*p)[j]), setAt, myMap, changNum, i)
				}
			case (*p)[j].cost < avg:
				if myRand1() < myRand2(echo, (*p)[j].cost, avg) {
					reOrganize(&((*p)[j]), setAt, myMap, changNum, i)
				}
			default:
				fmt.Println("error")
			}
		}
		sa(p, setAt, myMap, echo, nowbest)
		sort.Sort(*p)
		nowbest = (*p)[0].cost
		last_ans = append(last_ans, (*p)[0])
		fmt.Println("small is", (*p)[0])
	}
	return last_ans
}
func reOrganize(aim *Chrom, setAt *Point_que, myMap *Point_que, changNum, echo int) { //重组算法
	dict := make(map[string]bool)
	table := make([]string, 0, 5)
	temp := strings.Split(aim.body, "_")
	temp = temp[:len(temp)-1]
	cNum := int(math.Min(float64(len(temp)-1), float64(changNum)))
	cNum = int(math.Min(float64(cNum), float64(choosef(echo))))
	for _, v := range temp {
		dict[v] = true
	}
	for i := 0; i < len(*setAt); i++ {
		if ttm := strconv.Itoa(i); !dict[ttm] {
			table = append(table, ttm)
		}
	}
	check1 := make(map[int]bool)
	check2 := make(map[int]bool)
	for i := 0; i < cNum; i++ {
		ptr1 := rand.Int31n(int32(len(temp)))
		ptr2 := rand.Int31n(int32(len(table)))
		for (!check1[int(ptr1)]) && (!check2[int(ptr2)]) {
			check1[int(ptr1)], check2[int(ptr2)] = true, true
			temp[ptr1] = table[ptr2]
		}
	}
	ntemp := strings.Join(temp, "_")
	ntemp += "_"
	aim.body = ntemp
	aim.ccost(setAt, myMap)
}
func myRand1() float64 {
	var a int32 = rand.Int31n(100)
	return float64(a) / 100.0
}
func myRand2(echo int, myf, avgf float64) float64 {
	return math.Exp((myf - avgf) / float64(echo))
}
func myRand3(echo int, myf, avgf float64) float64 {
	return math.Exp(float64(echo) / (avgf - myf))
}
func choosef(nowEcoh int) int {
	switch {
	case nowEcoh > 0 && nowEcoh <= 80:
		return 1
	case nowEcoh > 80 && nowEcoh <= 160:
		return 2
	case nowEcoh > 160 && nowEcoh <= 200:
		return 3
	default:
		return 0
	}
}

///////////////////////////////////模拟退火部分
func findLarge(aim *Chrom, setAt *Point_que, myMap *Point_que) (float64, string) {
	temp := strings.Split(aim.body, "_")
	temp = temp[:len(temp)-1]
	var flage string
	var tep float64
	ptr, _ := strconv.Atoi(temp[0])
	ans := (*setAt)[ptr].Distance(&(*myMap)[0])
	flage = temp[0] + "_" + "0"
	nfla := ""
	for i := 0; i < len(*myMap); i++ {
		tep, nfla = dist_min((*myMap)[i], setAt, temp)
		nfla = strconv.Itoa(i) + "_" + nfla
		if ans < tep {
			ans = tep
			flage = nfla
		}
	}
	return ans, flage
} /// flage------》myMap p-setAt-》p
func dist_min(aim Point, setAt *Point_que, ptr []string) (float64, string) {
	nptr, _ := strconv.Atoi(ptr[0])
	ans := (*setAt)[nptr].Distance(&aim)
	flage := ptr[0]
	for _, i := range ptr {
		pptr, _ := strconv.Atoi(i)
		temp := (*setAt)[pptr].Distance(&aim)
		if ans > temp && temp != 0 {
			ans = temp
			flage = i
		}
	}
	return ans, flage
}
func find_min_p(aim *Chrom, point string, setAt *Point_que, myMap *Point_que) mlque {
	ans := make([]mlength, 0, 2)
	check := make(map[string]bool)
	table := make([]int, 0, 5)
	temp := strings.Split(aim.body, "_")
	temp = temp[:len(temp)-1]
	for _, k := range temp {
		check[k] = true
	}
	for i := 0; i < len(*setAt); i++ {
		t := strconv.Itoa(i)
		if !check[t] {
			table = append(table, i)
		}
	}
	tt := strings.Split(point, "_") //////
	if len(table) == 0 {
		return nil
	}
	npt, _ := strconv.Atoi(tt[0])
	for _, v := range table {

		tr := (*myMap)[npt].Distance(&(*setAt)[v])
		ans = append(ans, mlength{v, tr})
	}
	ans_t := mlque(ans)
	sort.Sort(ans_t)
	return ans_t
}
func find(table []string, aim string) int {
	i := -1
	for j := 0; j < len(table); j++ {
		if string(table[j]) == aim {
			i = j
			break
		}
	}
	return i
}
func saChange(aim *Chrom, setAt *Point_que, myMap *Point_que, T float64) {
	for T > 1 {
		_, pointQ := findLarge(aim, setAt, myMap) /// flage------》myMap p-setAt-》p
		//temp := strings.Split(pointQ, "_")
		search := find_min_p(aim, pointQ, setAt, myMap)
		if search == nil || len(search) == 0 {
			return
		}
		tbody := strings.Split(aim.body, "_")
		tbody = tbody[:len(tbody)-1]
		ptr := rand.Int31n(int32(len(search)))
		bptr := rand.Int31n(int32(len(tbody)))
		tbody[bptr] = strconv.Itoa(search[ptr].name)
		ttbody := strings.Join(tbody, "_")
		ttbody += "_"
		tchrom := Chrom{ttbody, 0}
		tchrom.ccost(setAt, myMap)
		dc := tchrom.cost - aim.cost
		if dc < 0 {
			(*aim) = tchrom
		} else if math.Exp(-dc/T) >= myRand1() {
			(*aim) = tchrom
		}
		T *= 0.9
	}
}
func sa(p *Population, setAt *Point_que, myMap *Point_que, echo int, global float64) {
	var avg, good float64 = 0.0, 0.9
	for _, v := range *p {
		avg += v.cost
	}
	avg /= float64(len(*p))
	for j := 0; j < len(*p); j++ {
		switch {
		case (*p)[j].cost == global:

		case (*p)[j].cost >= avg:
			if myRand1() < good {
				saChange(&((*p)[j]), setAt, myMap, 1000)
			}
		case (*p)[j].cost < avg:
			if myRand1() < myRand3(echo, (*p)[j].cost, avg) {
				saChange(&((*p)[j]), setAt, myMap, 1000)
			}
		default:
			fmt.Println("error")
		}
	}

}
