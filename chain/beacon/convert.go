package beacon

import (
	"github.com/Giordyfish/drand/chain"
	proto "github.com/Giordyfish/drand/protobuf/drand"
)

func beaconToProto(b *chain.Beacon) *proto.BeaconPacket {
	return &proto.BeaconPacket{
		PreviousSig: b.PreviousSig,
		Round:       b.Round,
		Signature:   b.Signature,
	}
}

func protoToBeacon(p *proto.BeaconPacket) *chain.Beacon {
	return &chain.Beacon{
		Round:       p.GetRound(),
		Signature:   p.GetSignature(),
		PreviousSig: p.GetPreviousSig(),
	}
}
