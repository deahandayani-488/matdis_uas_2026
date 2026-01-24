package data

type Node struct {
	Name      string
	Neighbors map[*Node]float64 // cost / jarak
}

// helper buat bikin node
func NewNode(name string) *Node {
	return &Node{
		Name:      name,
		Neighbors: make(map[*Node]float64),
	}
}

// data wilayah + relasi
func LoadGraph() (*Node, *Node) {
	a := NewNode("Kebayoran Lama")
	b := NewNode("Kebayoran Center")
	c := NewNode("Radio Dalam")
	d := NewNode("Mampang")
	p := NewNode("Agen Ban")

	// relasi atau jarak setiap node atau kantor
	// semakin besar angkanya semakin jauh
	// 3 dekat
	// 5 sedang
	// 7 jauh
	// 1 paling dekat
	// 0 tidak ada jarak

	a.Neighbors[b] = 5
	a.Neighbors[c] = 3
	b.Neighbors[d] = 4
	c.Neighbors[d] = 2
	d.Neighbors[p] = 1
	p.Neighbors[d] = 1

	return a, p //  tujuan mampang  (3 lokasi)
}
