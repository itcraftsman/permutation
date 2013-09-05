package permutation

func fac(n int) (int) {
	res := 1 
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}

func swap2d(slice [][]int, pos1, pos2 int) (newslice [][]int) {
	newslice = make([][]int, len(slice))
	newslice = slice
	if len(slice) > 1 {
		newslice[pos1], newslice[pos2] = slice[pos2], slice[pos1]
	}
	return
}

func swap(slice []int, pos1, pos2 int) (newslice []int) {
	newslice = make([]int, len(slice))
	newslice = slice
	if len(slice) > 1 {
		newslice[pos1], newslice[pos2] = slice[pos2], slice[pos1]
	}
	return
}

func splice2d(slice [][]int, pos int) (elem []int, newslice [][]int) {
	if len(slice) > 0 && pos < len(slice) {
		elem = slice[pos]
		newslice = make([][]int, 0, (len(slice) -1))
		newslice = append(newslice, slice[:pos]...)
		newslice = append(newslice, slice[pos +1:]...)
	}
	return
}

func splice(slice []int, pos int) (elem int, newslice []int) {
	if len(slice) > 0 && pos < len(slice) {
		elem = slice[pos]
		newslice = make([]int, 0, (len(slice) -1))
		newslice = append(newslice, slice[:pos]...)
		newslice = append(newslice, slice[pos +1:]...)
	}
	return
}

func Permutate2dSlice(slice [][]int) (permutations [][][]int){
	f := fac(len(slice))
	for i := 0; i < len(slice); i++ {
		elem, s := splice2d(slice, i)
		pos := 0
		for count := 0; count < (f / len(slice)); count++{
			if pos == (len(s) -1) {
				pos = 0
			}
			s = swap2d(s, pos, pos +1)
			permutation := make([][]int, len(slice))
			permutation = s
			permutation = append(permutation, elem)
			permutations = append(permutations, permutation)
			pos++
		} 
	} 
	return
}

func PermutateSlice(slice []int) (permutations [][]int){
	f := fac(len(slice))
	for i := 0; i < len(slice); i++ {
		elem, s := splice(slice, i)
		pos := 0
		for count := 0; count < (f / len(slice)); count++{
			if pos == (len(s) -1) {
				pos = 0
			}
			s = swap(s, pos, pos +1)
			permutation := make([]int, len(slice))
			permutation = s
			permutation = append(permutation, elem)
			permutations = append(permutations, permutation)
			pos++
		} 
	} 
	return
}
