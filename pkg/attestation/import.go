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

	return GetAAVulnerabilities(ctx, parent, filter)
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
				out = append(out, RunRevision{
					ServiceName:  service.Name,
					RevisionName: revision.Name,
					ArtifactURI:  containerImages[0],
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
