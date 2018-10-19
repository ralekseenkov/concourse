package v1

import (
	"context"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/resource/source"
)

type checkRequest struct {
	Source  atc.Source  `json:"source"`
	Version atc.Version `json:"version"`
}

func (r *Resource) Check(ctx context.Context, src atc.Source, fromVersion atc.Version) (atc.Spaces, error) {
	var versions []atc.Version

	err := source.RunScript(
		ctx,
		"/opt/resource/check",
		nil,
		checkRequest{src, fromVersion},
		&versions,
		nil,
		false,
		r.container,
	)
	if err != nil {
		return atc.Spaces{}, err
	}

	spaceVersions := []atc.SpaceVersion{}
	for _, v := range versions {
		spaceVersions = append(spaceVersions, atc.SpaceVersion{
			Space:   "v1space",
			Version: v,
		})
	}

	return atc.Spaces{
		DefaultSpace: "v1space",
		AllSpaces:    []atc.Space{"v1space"},
		Versions:     spaceVersions,
	}, nil
}
