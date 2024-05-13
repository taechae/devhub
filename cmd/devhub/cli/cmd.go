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
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/taechae/devhub/pkg/attestation"
	"github.com/taechae/devhub/pkg/types"
	c "github.com/urfave/cli/v2"
)

var (
	runCmd = &c.Command{
		Name:    "runtimes",
		Aliases: []string{"run"},
		Usage:   "runtimes",
		Action:  runCmdAction,
		Flags: []c.Flag{
			projectFlag,
			locationFlag,
		},
	}

	vulnCmd = &c.Command{
		Name:    "vulnerabilities",
		Aliases: []string{"vuln"},
		Usage:   "vulnerabilities",
		Action:  vulnCmdAction,
		Flags: []c.Flag{
			projectFlag,
			cveFlag,
			limitFlag,
		},
	}
)

func runCmdAction(c *c.Context) error {
	opt := &types.RunOptions{
		Project:  c.String(projectFlag.Name),
		Location: c.String(locationFlag.Name),
		Quiet:    isQuiet(c),
	}

	out, err := attestation.GetRunRevisions(c.Context, opt)
	if err != nil {
		return errors.Wrap(err, "error executing command")
	}

	b, _ := json.Marshal(out)
	fmt.Println(string(b))

	return nil
}

func vulnCmdAction(c *c.Context) error {
	opt := &types.VulnOptions{
		Project: c.String(projectFlag.Name),
		Cve:     c.String(cveFlag.Name),
		Limit:   c.Int(limitFlag.Name),
		Quiet:   isQuiet(c),
	}

	//out, err := attestation.GetVulnerabilities(c.Context, opt)
	out, err := attestation.GetTopVulnerabilities(c.Context, opt)
	if err != nil {
		return errors.Wrap(err, "error executing command")
	}

	b, _ := json.Marshal(out)
	fmt.Println(string(b))

	return nil
}
