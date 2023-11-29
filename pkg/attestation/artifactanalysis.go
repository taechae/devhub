package attestation

import (
	"context"
	"fmt"

	ca "cloud.google.com/go/containeranalysis/apiv1"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	g "google.golang.org/genproto/googleapis/grafeas/v1"
)

type Vulnerability struct {
	Cve         string
	Severity    string
	Cvss        float32
	ArtifactURI string
}

func GetAAVulnerabilities(ctx context.Context, parent string, filter string) ([]Vulnerability, error) {
	var out []Vulnerability

	c, err := ca.NewClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	reqFilter := "kind=\"VULNERABILITY\""
	if len(filter) > 0 {
		reqFilter = fmt.Sprintf("%s AND %s", reqFilter, filter)
	}

	req := &g.ListOccurrencesRequest{
		Parent:   parent,
		Filter:   reqFilter,
		PageSize: 1000,
	}

	it := c.GetGrafeasClient().ListOccurrences(ctx, req)
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return out, err
		}

		//fmt.Printf("%s, %s", resp.GetVulnerability().ShortDescription, resp.ResourceUri)

		v := Vulnerability{
			Cve:         resp.GetVulnerability().ShortDescription,
			Severity:    resp.GetVulnerability().Severity.String(),
			Cvss:        resp.GetVulnerability().CvssScore,
			ArtifactURI: resp.ResourceUri,
		}
		out = append(out, v)
	}

	return out, nil
}
