// Copyright 2020 The Envjector Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_ParseEnv(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input string
		exp   []string
		err   bool
	}{
		{
			name: "empty",
		},
		{
			name: "newlines",
			input: `
FOO=bar

ZIP=zap`,
			exp: []string{"FOO=bar", "ZIP=zap"},
		},
		{
			name: "missing =",
			input: `
FOO
BAR`,
			err: true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := ParseEnv(strings.NewReader(tc.input))
			if err != nil {
				if !tc.err {
					t.Fatal(err)
				}
				return
			}

			if act, exp := result, tc.exp; !reflect.DeepEqual(act, exp) {
				t.Errorf("expected %q to be %q", act, exp)
			}
		})
	}
}
