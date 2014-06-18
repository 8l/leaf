package prt

type Printable interface {
	PrintTo(p Iface)
}
