package attestation

import (
	"context"

	cr "cloud.google.com/go/run/apiv2"
	crpb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

type RunRevision struct {
	ServiceName  string
	RevisionName string
	ArtifactURI  string
}

func GetServices(ctx context.Context, parent string) ([]*crpb.Service, error) {
	var out []*crpb.Service

	c, err := cr.NewServicesClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	req := &crpb.ListServicesRequest{
		Parent: parent,
	}
	it := c.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return out, err
		}

		//fmt.Println("Service", resp)
		out = append(out, resp)
	}

	return out, nil
}

func GetRevisions(ctx context.Context, service *crpb.Service) ([]*crpb.Revision, error) {
	var out []*crpb.Revision

	c, err := cr.NewRevisionsClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	req := &crpb.ListRevisionsRequest{
		Parent: service.Name,
	}
	it := c.ListRevisions(ctx, req)
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return out, err
		}

		//fmt.Println("Revision", resp)
		out = append(out, resp)
	}

	return out, nil
}
