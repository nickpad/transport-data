package transportdata

type unixTime int64

type Database map[string]StopNode

type StopNode struct {
	Name string
	Hops []HopEdge
}

type HopEdge struct {
	From      StopNode
	To        StopNode
	DepartsAt unixTime
	ArrivesAt unixTime
}
