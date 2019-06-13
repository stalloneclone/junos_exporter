package bgp

type BgpSession struct {
	IP               string
	ASN              string
	Description      string
	Up               bool
	ReceivedPrefixes float64
	AcceptedPrefixes float64
	RejectedPrefixes float64
	ActivePrefixes   float64
	InputMessages    float64
	OutputMessages   float64
	Flaps            float64
}
