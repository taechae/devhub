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

import "errors"

var (
	ErrorNotSupported        = errors.New("not supported")
	ErrMissingRequiredFields = errors.New("missing required fields")
	ErrMissingProject        = errors.New("missing project")
	ErrInvalidFormat         = errors.New("invalid format")
	ErrMissingPath           = errors.New("missing path")
	ErrMissingSource         = errors.New("missing source")
	ErrInvalidSource         = errors.New("invalid source")
	ErrInternal              = errors.New("internal error")
)
