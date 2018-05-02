package reflect

type Hash struct {
	Seed uint32
}
//产生Hash数值
func (hash *Hash) HashCode() uint32 {
	a := hash.Seed%2 + hash.Seed%7*2
	return a
}
