package main

import "fmt"

var source = []string{
	"APTZvA", "BddOIt", "ctuuYn", "BCd5js", "cVCuqR", "AQynrL", "AoZ62r", "BV9DXI", "cqkYj7", "ALSKpF", "CEkB4M", "By6jE3", "Aclr2o", "cLiix5", "AClM5o", "BN36oa", "BYj4K0", "cKtPyI", "BGOn7c", "BQreVu", "B7kQ15", "BHhAY0", "cbQBTI", "A2KDsf", "AwmbeJ", "BsNdy0", "BoIVCB", "C3pHMS", "CP9Wc6", "C6vyPb", "A6BTpf", "AguFNY", "AoeaF8", "AyQ3dP", "CzlhVY", "BkFrls", "C4WncK", "ASTebw", "CTpdJi", "BtGzKA", "cWtmeT", "BgLz5G", "A9Ohfh", "ASv3qg", "A4du4s", "BstIGr", "BSIkmq", "CKxdNR", "BgCF6g", "CWkjqZ",
}

func solve1() {
	source2 := make([]string, len(source))
	for i := 0; i < len(source); i++ {
		source2[i] = source[i][1:len(source[i])]
	}
	for _, v := range source2 {
		fmt.Println(v)
	}
}
func solve2() {
	for i := 0; i < len(source); i++ {
		source[i] = source[i][:3] + "A" + source[i][4:]
	}
	for _, v := range source {
		fmt.Println(v)
	}
}
func solve3() {
	dic := make(map[rune]int)
	for i := 0; i < len(source); i++ {
		dic[rune(source[i][0])]++
	}
	for k, v := range dic {
		fmt.Println(string(k), v)
	}
}

func main() {
	solve3()
}
