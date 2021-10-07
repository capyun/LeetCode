package DisjointSet

/**
@author Shitanford
@create 2021-10-07 9:46
*/

func main() {

}

type islandSet struct {
	set map[int]int
	num int
}

func numIslands(grid [][]byte) int {
	set := islandSet{
		set: make(map[int]int),
		num: 0,
	}
	cols := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				set.add(i, j, cols)
			}
		}
	}
	return set.num
}

func (this *islandSet) find(index int) (int, bool) {
	//if not exists, then
	if _, ok := this.set[index]; !ok {
		return 0, false
	} else {
		return this.findHelper(index), true
	}
}
func (this *islandSet) findHelper(index int) int {
	//set[index]是index的父 的位置，或者是根的数量
	if p, _ := this.set[index]; p < 0 {
		return index
	}
	root := this.findHelper(this.set[index])
	this.set[index] = root
	return root
}

func (this *islandSet) union(a, b int)  {
	//a, b exists
	rc, _ := this.find(a)
	rp, _ := this.find(b)
	if rc == rp {
		return
	}
	//if this.height(rc) > this.height(rp) {
	//	rc, rp = rp, rc
	//}
	//this.set[rp] += this.set[rc]
	this.set[rc] = rp
	this.num--
}

func (this *islandSet) add(x, y, cols int)  {
	i := index(x, y, cols)
	this.set[i] = -1
	this.num++

	//can't use i-1, if i is the head of row
	if _, ok := this.set[i-1]; y != 0 && ok {
		this.union(i, i-1)
	}
	if _, ok := this.set[i-cols]; ok {
		this.union(i, i-cols)
	}
}

func index(x, y, cols int) int {
	return x * cols + y
}

func (this *islandSet) height(i int) int {
	return -this.set[i]
}