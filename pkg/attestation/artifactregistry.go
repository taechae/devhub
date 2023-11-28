package attestation

import (
	"context"
	"fmt"

	ar "cloud.google.com/go/artifactregistry/apiv1"
	arpb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func GetImagesByTag(ctx context.Context, parent string, tag string) ([]string, error) {
	var out []string

	repos, err := GetRepositories(ctx, parent)
	if err != nil {
		return out, errors.Wrap(err, "error listing repositories")
	}

	c, err := ar.NewClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	for _, repo := range repos {
		req := &arpb.ListDockerImagesRequest{
			Parent: repo,
		}
		it := c.ListDockerImages(ctx, req)
		for {
			resp, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				return out, err
			}

			//fmt.Println("Image", resp)
			if matchTag(tag, resp.Tags) {
				out = append(out, resp.Uri)
			}
		}
	}

	return out, nil
}

func matchTag(tag string, imageTags []string) bool {
	m1 := fmt.Sprintf("short-%s", tag)
	m2 := fmt.Sprintf("long-%s", tag)

	for _, imageTag := range imageTags {
		if imageTag == m1 || imageTag == m2 {
			return true
		}
	}
	return false
}

func matchImages(searchImages []string, images []string) bool {
	for _, img := range images {
		for _, searchImg := range searchImages {
			if searchImg == img {
				return true
			}
		}
	}
	return false
}

func GetRepositories(ctx context.Context, parent string) ([]string, error) {
	var out []string

	c, err := ar.NewClient(ctx)
	if err != nil {
		return out, errors.Wrap(err, "error creating client")
	}
	defer c.Close()

	req := &arpb.ListRepositoriesRequest{
		Parent: parent,
	}
	it := c.ListRepositories(ctx, req)
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return out, err
		}

		//fmt.Println("Repository", resp)
		out = append(out, resp.Name)
	}

	return out, nil
}
