package types

import "github.com/ryssapp/backend/src/go/common/pb"

type Location struct {
	Address   string  `json:"address"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (l *Location) ToProto() *pb.Location {
	return &pb.Location{
		Address: l.Address,
		City: l.City,
		ZipCode: l.ZipCode,
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func LocationFromProto(l *pb.Location) *Location {
	return &Location{
		Address: l.GetAddress(),
		City: l.GetCity(),
		ZipCode: l.GetZipCode(),
		Latitude:  l.GetLatitude(),
		Longitude: l.GetLongitude(),
	}
}

type Pagination struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (p *Pagination) ToProto() *pb.Pagination {
	return &pb.Pagination{
		Limit:  p.Limit,
		Offset: p.Offset,
	}
}

func PaginationFromProto(p *pb.Pagination) *Pagination {
	return &Pagination{
		Limit:  p.GetLimit(),
		Offset: p.GetOffset(),
	}
}
