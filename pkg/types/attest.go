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

package types

type RunOptions struct {
	// Project is the ID of the project to import the report into.
	Project string

	// Region
	Location string

	// Quiet suppresses output
	Quiet bool
}

// Validate validates the options.
func (o *RunOptions) Validate() error {
	if o.Project == "" {
		return ErrMissingProject
	}

	if o.Location == "" {
		return ErrMissingSource
	}

	return nil
}

type VulnOptions struct {
	// Project is the ID of the project to import the report into.
	Project string

	// Region
	Cve string

	ArtifactURI string

	Limit int

	// Quiet suppresses output
	Quiet bool
}

// Validate validates the options.
func (o *VulnOptions) Validate() error {
	if o.Project == "" {
		return ErrMissingProject
	}

	return nil
}

type PkgOptions struct {
	// Project is the ID of the project to import the report into.
	Project string

	Package string

	Limit int

	// Quiet suppresses output
	Quiet bool
}

// Validate validates the options.
func (o *PkgOptions) Validate() error {
	if o.Project == "" {
		return ErrMissingProject
	}

	return nil
}

type BuildOptions struct {
	// Project is the ID of the project to import the report into.
	Project string

	Limit int

	// Quiet suppresses output
	Quiet bool
}

// Validate validates the options.
func (o *BuildOptions) Validate() error {
	if o.Project == "" {
		return ErrMissingProject
	}

	return nil
}
