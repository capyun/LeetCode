package DisjointSet

/**
@author Shitanford
@create 2021-10-06 19:12
*/

func main() {

}

const offset int = 10e9
type consecutiveSet map[int]int

func longestConsecutive(nums []int) int {
	set := make(consecutiveSet)
	for _, v := range nums {
		set.add(v)
	}
	ans := offset
	for key, _ := range set {
		ans = maxInt(ans, set.height(key))
	}
	return ans - offset
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (this consecutiveSet) height(child int) int {
	if root, ok := this.find(child); ok {
		return this[root]
	}
	return 0
}
func (this consecutiveSet) find(child int) (int, bool) {
	if _, ok := this[child]; !ok {
		return 0, false
	}
	return this.findHelper(child), true
}

func (this consecutiveSet) findHelper(child int) int {
	if this[child] > offset {
		return child
	}
	root, _ := this.find(this[child])
	this[child] = root
	return root
}

func (this consecutiveSet) add(child int)  {
	if _, ok := this[child]; ok {
		return
	}
	this[child] = offset+1
	root := this.union(child, child-1)
	this.union(root, child+1)
}

func (this consecutiveSet) union(c, y int) int {
	//c is root already
	//r exists
	if r, ok := this.find(y); ok {
		//c's weight > r's weight
		if this[c] > this[r] {
			c, r = r, c
		}
		this[r] += this[c]-offset
		this[c] = r
		return r
	}
	return c
}