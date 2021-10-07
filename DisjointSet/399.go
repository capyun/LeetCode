package DisjointSet

/**
@author Shitanford
@create 2021-10-06 10:29
*/

func main() {

}

type node struct {
	Parent string
	Factor float64
}

type set map[string]*node

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	varSet := make(set)
	for i := 0; i < len(values); i++ {
		c, p := equations[i][0], equations[i][1]
		varSet.union(c, p, values[i])
	}

	var ans []float64
	for i := 0; i < len(queries); i++ {
		c, q := queries[i][0], queries[i][1]
		ans = append(ans, varSet.divide(c, q))
	}
	return ans
}

func (this set) find(child string) (string, float64) {
	if _, ok := this[child]; !ok {
		this[child] = &node{
			Parent: child,
			Factor: -1,
		}
		return child, 1
	}
	return this.findHelper(child)
}
func (this set) findHelper(child string) (string, float64) {
	n := this[child]
	if n.Factor < 0 {
		return n.Parent, 1
	}
	p, f := this.findHelper(n.Parent)
	f *= n.Factor
	n.Parent = p
	n.Factor = f
	return p, f
}

func (this set) union(child string, parent string, factor float64)  {
	c, fc := this.find(child)
	p, fp := this.find(parent)
	//两个在一组里
	if c == p {
		return
	}
	if this.weight(c) > this.weight(p) {
		c, p = p, c
		fc, fp = fp, fc
	}
	//这一步必做
	this[p].Factor += this[c].Factor
	this[c] = &node{
		Parent: p,
		Factor: factor * fp / fc,
	}
}

func (this set) divide(child string, parent string) float64 {
	//变量不存在，则返回-1
	if _, ok := this[child]; !ok {
		return -1
	}
	if _, ok := this[parent]; !ok {
		return -1
	}
	//变量存在的情况
	pc, fc := this.find(child)
	pp, fp := this.find(parent)
	if pc != pp {
		return -1
	}
	return fc / fp
}

func (this set) weight(root string) float64 {
	return -this[root].Factor
}