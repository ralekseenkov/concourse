package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/concourse/concourse/atc"
)

// CheckResponse contains a default space and a list of resource versions. This response is returned from the check of the resource v2 interface. The default space can be empty for resources that do not have a default space (ex. PR resource).
type CheckResponse struct {
	DefaultSpace string `json:"default_space"`
	Versions     []ResourceVersion
}

type ResourceVersion struct {
	Space    string              `json:"space"`
	Version  atc.Version         `json:"version"`
	Metadata []atc.MetadataField `json:"metadata"`
}

type checkRequest struct {
	// XXX: make the names consistent
	Source       atc.Source                `json:"config"`
	From         map[atc.Space]atc.Version `json:"from"`
	ResponsePath string                    `json:"response_path"`
}

func (r *resource) Check(ctx context.Context, source atc.Source, from map[atc.Space]atc.Version) (atc.Spaces, error) {
	var spaces atc.Spaces

	tmpfile, err := ioutil.TempFile("", "response")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name())

	path := r.info.Artifacts.Check
	input := checkRequest{source, from, tmpfile.Name()}
	output := &spaces

	request, err := json.Marshal(input)
	if err != nil {
		return err
	}

	stderr := new(bytes.Buffer)

	processIO := garden.ProcessIO{
		Stdin:  bytes.NewBuffer(request),
		Stderr: stderr,
	}

	process, err := r.container.Run(garden.ProcessSpec{
		Path: path,
	}, processIO)
	if err != nil {
		return err
	}

	processExited := make(chan struct{})

	var processStatus int
	var processErr error

	go func() {
		processStatus, processErr = process.Wait()
		close(processExited)
	}()

	select {
	case <-processExited:
		if processErr != nil {
			return processErr
		}

		if processStatus != 0 {
			return ErrResourceScriptFailed{
				Path:       path,
				ExitStatus: processStatus,

				Stderr: stderr.String(),
			}
		}

	case <-ctx.Done():
		r.container.Stop(false)
		<-processExited
		return ctx.Err()
	}

	var spaces []Space

	fileReader, err := os.Open(tmpfile)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(fileReader)

	var spaces []atc.Space
	var versions []atc.SpaceVersion

	if decoder.More() {
		var defaultSpace atc.DefaultSpace
		err := decoder.Decode(&defaultSpace)
		if err != nil {
			return err
		}
	}

	for decoder.More() {
		var version atc.SpaceVersion
		err := decoder.Decode(&version)
		if err != nil {
			return err
		}
		if len(spaces) == 0 || spaces[len(spaces)-1] != version.Space {
			spaces = append(spaces, version.Space)
		}
		versions = append(versions, version)
	}

	return atc.Spaces{defaultSpace, spaces, versions}, nil
}
