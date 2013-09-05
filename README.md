permutation
===========

A Go Pkg to permutate slices and 2D slices.

Example:

<code>
import (
	"fmt"
	"github.com/itcraftsman/permutation"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	result := permutation.PermutateSlice(list)
	for _, v := range result {
		fmt.Println(v)
	}
	fmt.Println(len(result))
	
	list2d := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}
	result2d := permutation.Permutate2dSlice(list2d)
	for _, v := range result2d {
		fmt.Println(v)
	}
	fmt.Println(len(result2d))
}
</code>


