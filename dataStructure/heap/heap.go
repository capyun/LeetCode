package heap

/**
@author Shitanford
@create 2021-05-13 9:41
*/

type Heap struct {
	arr []int
	size int
}

func NewHeap(arr []int) Heap {
	heapArr := make([]int, 10+len(arr))
	for i:=0; i<len(arr); i++ {
		heapArr[i+1] = arr[i]
	}
	h := Heap{arr: heapArr,
		size: len(arr)}
	h.buildHeap()
	return h
}

func (h Heap) GetSize() int {
	return h.size
}

func (h *Heap) Push(v int)  {
	h.size++
	hole := h.size
	for ; hole>1 && h.arr[hole/2]<v; hole /= 2 {
		h.arr[hole] = h.arr[hole/2]
	}
	h.arr[hole] = v
}

func (h *Heap) Pop() int {
	v := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.percolateDown(1)
	return v
}

func (h Heap) percolateDown(hole int)  {
	var child int
	tmp := h.arr[hole]
	for ; hole*2 <= h.size ; hole=child {
		child = hole*2
		if child!=h.size && h.arr[child]<h.arr[child+1] {
			child++
		}
		if h.arr[child] > tmp {
			h.arr[hole] = h.arr[child]
		} else {
			break
		}
	}
	h.arr[hole] = tmp
}

func (h Heap) buildHeap()  {
	for i:=h.size/2; i>0; i-- {
		h.percolateDown(i)
	}
}

