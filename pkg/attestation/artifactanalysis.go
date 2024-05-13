package attestation

import (
	"context"
	"fmt"
	"time"

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

type TopVulnerabilities struct {
	Cve       string
	Severity  string
	Cvss      float32
	Artifacts []string
}

type Package struct {
	Artifact string
	Package  string
	Version  string
}

type Build struct {
	Date      time.Time
	Artifact  string
	BuilderId string
	BuildId   string
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

func GetAAPackages(ctx context.Context, parent string, filter string) ([]Package, error) {
	var out []Package

	c, err := ca.NewClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	reqFilter := "kind=\"PACKAGE\""
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

		v := Package{
			Artifact: resp.ResourceUri,
			Package:  resp.GetPackage().Name,
			Version:  resp.GetPackage().Version.Name,
		}
		out = append(out, v)
	}

	return out, nil
}

func GetAABuilds(ctx context.Context, parent string, filter string) ([]Build, error) {
	var out []Build

	c, err := ca.NewClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	reqFilter := "kind=\"BUILD\""
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

		if resp.GetBuild().InTotoSlsaProvenanceV1 != nil {
			p := resp.GetBuild().InTotoSlsaProvenanceV1

			//fmt.Println("dddd", p.Predicate.RunDetails.Builder.Id, p.Predicate.RunDetails.Metadata.InvocationId)
			v := Build{
				Date:      resp.CreateTime.AsTime(),
				Artifact:  resp.ResourceUri,
				BuilderId: p.Predicate.RunDetails.Builder.Id,
				BuildId:   p.Predicate.RunDetails.Metadata.InvocationId,
			}
			out = append(out, v)
		}

	}

	return out, nil
}
