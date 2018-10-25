package main

import (
	"fmt"
	//"strings"
)

func main() {
	var num_point, sizeP, pop_size, echo, flage int
	fmt.Println("生成点集大小：")
	fmt.Scanln(&num_point)
	fmt.Println("中心点数量：")
	fmt.Scanln(&sizeP)
	fmt.Println("种群大小：")
	fmt.Scanln(&pop_size)
	fmt.Println("迭代次数：")
	fmt.Scanln(&echo)
	mywork(num_point, sizeP, pop_size, echo)
	fmt.Scanln(&flage)
}
func mywork(num_point, sizeP, Pop_size, echo int) {
	var test = createPoint(num_point) //70) //生成任意点
	ans := test.Convexll()
	fmt.Println(len(test), " ", len(ans)) //凸包边界点
	Chrom_population := Population(createChrom(sizeP, Pop_size /*6, 500,*/, &ans, &test))
	fmt.Println("finish1")
	tans := Chrom_population.genetic(echo, &ans, &test) ///////出错了
	fmt.Println("finsh2")
	for _, v := range tans {
		fmt.Println(v.cost)
	}
	fmt.Println("end")
	/*var temp = strings.Split("1_2_3_15_", "_")
	for _, v := range temp {
		fmt.Println(v)
	}*/
}
