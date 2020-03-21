package common

import "github.com/ryssapp/backend/src/go/common/pb"

type Location struct {
	Latitude float32
	Longitude float32
}

func (l *Location) ToProto() *pb.Location {
	return &pb.Location{
		Latitude: l.Latitude,
		Longitude: l.Longitude,
	}
}