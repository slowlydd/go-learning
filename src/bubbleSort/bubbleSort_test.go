package bubbleSort

import (
	"fmt"
	"testing"
)

func bubbleSort(n [5]int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				temp := n[j]
				n[j] = n[j+1]
				n[j+1] = temp
			}
		}
	}
	fmt.Println(n)
}

func TestBubbleSort(t *testing.T) {
	arr := [...]int{2, 5, 3, 1, 8}
	bubbleSort(arr)
	t.Log(arr)
}
