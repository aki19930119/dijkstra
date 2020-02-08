package graph
//結合の構造
type Node struct {
	Id   int
	Dist int
}
//関数　垂直の力　水平の力
type PQ []Node

//関数　差し込む
func Insert(pq *PQ, node Node) {
	*pq = append(*pq, node)
	i := len(*pq) - 1
	for i > 1 && (*pq)[i/2].Dist > (*pq)[i].Dist {
		tmp := (*pq)[i]
		(*pq)[i] = (*pq)[i/2]
		(*pq)[i/2] = tmp
		i = i / 2
	}
}
//関数　引き抜く
func Extract(pq *PQ) Node {
	node := (*pq)[1]
	(*pq)[1] = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	update(*pq, 1)
	return node
}
//関数　更新
func update(pq PQ, index int) {
	var min int
	l := 2 * index
	r := 2*index + 1

	if l <= len(pq)-1 && pq[l].Dist < pq[index].Dist {
		min = l
	} else {
		min = index
	}
	if r <= len(pq)-1 && pq[r].Dist < pq[min].Dist {
		min = r
	}
	if min != index {
		tmp := pq[index]
		pq[index] = pq[min]
		pq[min] = tmp
		update(pq, min)
	}
}