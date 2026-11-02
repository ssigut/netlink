package nl

// Attributes of an LWTUNNEL_ENCAP_IP route encapsulation, from
// include/uapi/linux/lwtunnel.h.
const (
	LWTUNNEL_IP_UNSPEC = iota
	LWTUNNEL_IP_ID
	LWTUNNEL_IP_DST
	LWTUNNEL_IP_SRC
	LWTUNNEL_IP_TTL
	LWTUNNEL_IP_TOS
	LWTUNNEL_IP_FLAGS
	LWTUNNEL_IP_PAD
	LWTUNNEL_IP_OPTS // not implemented
	__LWTUNNEL_IP_MAX
)
