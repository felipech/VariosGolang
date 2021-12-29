package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	//var strs = []string {"flower","fkow"}
	//fmt.Println(longestCommonPrefix(strs))

	//var nums = []int {-1,0,1,2,-1,-4}
	//[-2,0,1,1,2]
	//var nums = []int {-2,0,1,1,2}
	//-2,0,0,2,2]
	//[3,0,-2,-1,1,2]
	//[-1,0,1,2,-1,-4]
	var nums = []int {-1,0,1,2,-1,-4}
	//var nums = []int {-2,0,0,2,2}
	//var nums = []int {-1,0,1,0}
	//[0,0,0,0]
	//var nums = []int {0,0,0,0}
	fmt.Println("arreglo ",nums)
	sum := threeSum(nums)
	fmt.Println("largo respuesta -> " , len(sum))
	fmt.Println("respuesta ", sum)
	fmt.Println("largo nums -> " , len(nums))

}


func testingSomthing(){
	values := [][]int{}

	// Step 2: these are the first two rows.
	// ... Append each row to the two-dimensional slice.
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: display first row, and second row.
	fmt.Println("Row ---------------- 1")
	fmt.Println(values[0])
	fmt.Println("Row ---------------- 2")
	fmt.Println(values[1])
}


func threeSum(nums []int) [][]int {
	var values [][]int

	sort.Ints(nums)
	if len(nums) == 0 || len(nums) < 3{
		return values
	}
	if len(nums) == 3 {
		if nums [0] + nums[1] + nums[2] == 0 {
			rows1 := make([]int, 3)
			rows1[0] = nums[0]
			rows1[1] = nums[1]
			rows1[2] = nums[2]
			values = append(values, rows1)
			return values
		}else {
			return values
		}
	}

	for iterador, a := range nums {
		//fmt.Println("valor i: " , iterador)
		if iterador > 0 && a == nums[iterador- 1] {
			continue
		}
		var l, r = iterador + 1, len(nums) - 1
		for l < r {
			var threeSum = a + nums[l] + nums[r]
			if threeSum > 0 {
				r -= 1

			}else if threeSum < 0{
				l += 1
			}else {
				rows1 := make([]int, 3)
				rows1[0] = a
				rows1[1] = nums[l]
				rows1[2] = nums[r]
				values = append(values, rows1)
				l += 1
				for nums[l] == nums[l - 1] && l < r {
					l += 1
				}
			}
		}
	}
	return values
}


//arreglo contiene dentro de otro arreglo
func contains(s [][]int, str []int) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, str) {
			return true
		}
	}

	return false
}


//let code problem
func longestCommonPrefix(strs []string) string {
	var menor = shortestWord(strs)
	fmt.Println("menor 1 " + menor)

	var prefix, m = combinations(menor)

	var resultado = ""
	var largoantiguo int
	var largoEncontrado int
	for _, s := range strs{
		largoEncontrado = FindPrefix(prefix, s)
		if largoEncontrado == 0 {
			largoantiguo = 0
			break
		}
		//println("largo encontrado 1: ", largoEncontrado)
		if largoantiguo == 0 {
			largoantiguo = largoEncontrado
		}else if largoantiguo > largoEncontrado  {
			largoantiguo = largoEncontrado
		}
		//println("largo encontrado 2: ", largoantiguo)
	}

	if largoantiguo == 0 {
		fmt.Println("resutlado " + "")
		return ""
	}else {
		resultado = m[largoantiguo]
		fmt.Println("resultado: ",resultado)
		return resultado
	}
}

func shortestWord(strs []string) string {
	var largo = len(strs[0])
	var menor = strs[0]
	for i := 1; i < len(strs); i++ {
		if largo > len(strs[i]) {
			largo = len(strs[i])
			menor = strs[i]
		}
	}
	return menor
}


func combinations (texto string)  ([]string, map[int]string){
	comb := make([]string, len(texto)+1)
	var m = make(map[int]string)
	for i := 0; i < len(texto) + 1; i++ {
		comb[i] = texto[:i]
		m[len(texto[:i])] = texto[:i]
	}

	return comb, m

}

func FindPrefix(prefijos [] string, plabra string)  int{
	var largo = len(prefijos[0])
	for i := 0; i < len(prefijos); i++ {
		if strings.Compare(prefijos[i], plabra[0:i]) == 0 && plabra[0:i] != "" {
			if largo < len(prefijos[i]){
				largo = len(prefijos[i])
			}
		}
	}
	return largo

}