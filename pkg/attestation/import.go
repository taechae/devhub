// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attestation

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/taechae/devhub/pkg/types"
)

func GetVulnerabilities(ctx context.Context, options types.Options) ([]Vulnerability, error) {
	var out []Vulnerability

	opt, ok := options.(*types.VulnOptions)
	if !ok || opt == nil {
		return out, errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return out, errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s", opt.Project)

	var filter string
	if len(opt.Cve) > 0 {
		filter = fmt.Sprintf("noteId=\"%s\"", opt.Cve)
	}

	if len(opt.ArtifactURI) > 0 {
		if len(filter) > 0 {
			filter += " AND "
		}
		filter += fmt.Sprintf("resourceUrl=\"%s\"", opt.ArtifactURI)
	}

	return GetAAVulnerabilities(ctx, parent, filter)
}

func GetTopVulnerabilities(ctx context.Context, options types.Options) ([]TopVulnerabilities, error) {
	vMap := make(map[string]TopVulnerabilities, 0)
	var out []TopVulnerabilities

	opt, ok := options.(*types.VulnOptions)
	if !ok || opt == nil {
		return out, errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return out, errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s", opt.Project)

	vulns, _ := GetAAVulnerabilities(ctx, parent, "")
	for _, v := range vulns {
		if _, ok := vMap[v.Cve]; !ok {
			vMap[v.Cve] = TopVulnerabilities{
				Cve:      v.Cve,
				Severity: v.Severity,
				Cvss:     v.Cvss,
			}

		}

		vEntry := vMap[v.Cve]
		vEntry.Artifacts = append(vEntry.Artifacts, v.ArtifactURI)
		vMap[v.Cve] = vEntry
	}

	keys := make([]string, 0, len(vMap))
	for _, k := range vMap {
		keys = append(keys, k.Cve)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return vMap[keys[i]].Cvss > vMap[keys[j]].Cvss
	})

	count := 0
	for _, key := range keys {
		out = append(out, vMap[key])
		//fmt.Println(vMap[key])

		count++
		if opt.Limit > 0 && count >= opt.Limit {
			break
		}
	}

	return out, nil
}

func GetRunRevisions(ctx context.Context, options types.Options) ([]RunRevision, error) {
	var out []RunRevision

	opt, ok := options.(*types.RunOptions)
	if !ok || opt == nil {
		return out, errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return out, errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", opt.Project, opt.Location)

	services, err := GetServices(ctx, parent)
	if err != nil {
		return out, errors.Wrap(err, "error listing services")
	}

	for _, service := range services {
		revisions, err := GetRevisions(ctx, service)
		if err != nil {
			return out, errors.Wrap(err, "error listing revisions")
		}

		for _, revision := range revisions {
			var containerImages []string
			for _, container := range revision.Containers {
				containerImages = append(containerImages, container.Image)
			}

			// TODO: What if service is splitting traffic across multiple revisions?
			// Is ready revision even guaranteed to be serving traffic?
			if revision.Name == service.LatestReadyRevision {
				//fmt.Printf("%s, %s, %s\n", service.Name, revision.Name, containerImages[0])
				serviceTokens := strings.Split(service.Name, "/")
				revisionTokens := strings.Split(revision.Name, "/")
				out = append(out, RunRevision{
					ServiceName:  serviceTokens[len(serviceTokens)-1],
					RevisionName: revisionTokens[len(revisionTokens)-1],
					ArtifactURI:  "https://" + containerImages[0],
				})
			}
		}
	}

	return out, nil
}

// Import imports attestation metadata from a source.
func Import(ctx context.Context, options types.Options) error {
	opt, ok := options.(*types.RunOptions)
	if !ok || opt == nil {
		return errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", opt.Project, opt.Location)
	commit := "zzz"

	fmt.Printf("Searching commit (%s)...\n\n", commit)

	images, err := GetImagesByTag(ctx, parent, commit)
	if err != nil {
		return errors.Wrap(err, "error listing images by tag")
	}

	if len(images) == 0 {
		fmt.Println("No images found for commit.")
		return nil
	}
	fmt.Println("Images found:")
	for _, image := range images {
		fmt.Printf("- %s\n", image)
	}

	var runtimes int
	services, err := GetServices(ctx, parent)
	if err != nil {
		return errors.Wrap(err, "error listing services")
	}

	for _, service := range services {
		revisions, err := GetRevisions(ctx, service)
		if err != nil {
			return errors.Wrap(err, "error listing revisions")
		}

		for _, revision := range revisions {
			var containerImages []string
			for _, container := range revision.Containers {
				containerImages = append(containerImages, container.Image)
			}
			if matchImages(images, containerImages) {
				if runtimes == 0 {
					fmt.Println("\nRuntimes found:")
					fmt.Println("- Cloud Run Revisions:")
				}
				runtimes++

				fmt.Printf("  - Name: %s\n", revision.Name)

				servingTraffic := false
				// TODO: What if service is splitting traffic across multiple revisions?
				// Is ready revision even guaranteed to be serving traffic?
				if revision.Name == service.LatestReadyRevision {
					servingTraffic = true
				}
				fmt.Printf("  - Serving Traffic: %t\n", servingTraffic)
			}
		}
	}

	if runtimes == 0 {
		fmt.Println("\nNo runtimes found.")
	}

	return nil
}

func GetPackages(ctx context.Context, options types.Options) ([]Package, error) {
	var out []Package

	opt, ok := options.(*types.PkgOptions)
	if !ok || opt == nil {
		return out, errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return out, errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s", opt.Project)

	pkgs, _ := GetAAPackages(ctx, parent, "")

	for _, p := range pkgs {
		if len(opt.Package) > 0 && strings.Contains(p.Package, opt.Package) {
			out = append(out, p)
		}
	}

	return out, nil
}

func GetBuilds(ctx context.Context, options types.Options) ([]Build, error) {
	var out []Build

	opt, ok := options.(*types.BuildOptions)
	if !ok || opt == nil {
		return out, errors.New("valid options required")
	}
	if err := options.Validate(); err != nil {
		return out, errors.Wrap(err, "error validating options")
	}

	parent := fmt.Sprintf("projects/%s", opt.Project)

	builds, _ := GetAABuilds(ctx, parent, "")
	sort.SliceStable(builds, func(i, j int) bool {
		return builds[i].Date.Compare(builds[j].Date) > 0
	})

	count := 0
	for _, b := range builds {
		count++
		if opt.Limit > 0 && count > opt.Limit {
			break
		}
		out = append(out, b)
	}

	return out, nil
}
