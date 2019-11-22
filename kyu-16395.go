package main

import "fmt"

func main(){
	var n,k int
	_,_ = fmt.Scan(&n,&k)
	var temp [][]int = generate(n+1)
	fmt.Println(temp[n-1][k-1])
}

func generate(numRows int) [][]int {
	rtn := make([][]int,0,numRows)
	if numRows != 0{
		firstArray :=[]int{1}
		rtn = append(rtn, firstArray)
		for i:=1;i<numRows;i++{
			arrayToAdd := make([]int,i+1)
			arrayToAdd[0]=1
			arrayToAdd[i]=1
			for k:=1;k<i;k++{
				arrayToAdd[k]=rtn[i-1][k-1]+rtn[i-1][k]
			}
			rtn = append(rtn, arrayToAdd)
		}
	}
	return rtn
}

// 이렇게 풀면 틀렸다 함 ㄷㄷ
//func main() {
//	var n, k int
//	var rtn int = 1
//	_, _ = fmt.Scan(&n, &k)
//	for i := k; i < n; i++ {
//		rtn = rtn * i
//	}
//	for j := 1; j <= n-k; j++ {
//		rtn = rtn / j
//	}
//	fmt.Println(rtn)
//}
