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

package cli

import (
	"github.com/pkg/errors"
	"github.com/taechae/devhub/pkg/attestation"
	"github.com/taechae/devhub/pkg/types"
	c "github.com/urfave/cli/v2"
)

var (
	attestCmd = &c.Command{
		Name:    "search-commit",
		Aliases: []string{"search"},
		Usage:   "import attestation metadata",
		Action:  attestationCmd,
		Flags: []c.Flag{
			projectFlag,
			locationFlag,
			commitFlag,
		},
	}
)

func attestationCmd(c *c.Context) error {
	opt := &types.AttestationOptions{
		Project:  c.String(projectFlag.Name),
		Location: c.String(locationFlag.Name),
		Commit:   c.String(commitFlag.Name),
		Quiet:    isQuiet(c),
	}

	printVersion(c)

	if err := attestation.Import(c.Context, opt); err != nil {
		return errors.Wrap(err, "error executing command")
	}

	return nil
}
