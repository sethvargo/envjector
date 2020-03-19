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
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", usage)
		flag.PrintDefaults()
	}

	flagFile := flag.String("file", "", "path to file with envvars")
	flag.Parse()

	if err := realMain(*flagFile, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func realMain(filename string, args []string) error {
	if filename == "" {
		return fmt.Errorf("missing -file")
	}

	if len(args) == 0 {
		return fmt.Errorf("missing command to run")
	}

	binary, err := exec.LookPath(args[0])
	if err != nil {
		return fmt.Errorf("failed to find %q: %w", args[0], err)
	}

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", filename, err)
	}
	defer f.Close()

	env, err := ParseEnv(f)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Append saved environment to the provided environment at the end, so our
	// file takes precendence.
	currentEnv := syscall.Environ()
	env = append(currentEnv, env...)

	if err := syscall.Exec(binary, args, env); err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}

	return nil
}

// ParseEnv parses the given environment and returns each line as a slice
// member. Empty lines are skipped. If a line does not contain an "=", an error
// is returned. Any errors in parsing are also returned.
func ParseEnv(r io.Reader) ([]string, error) {
	var env []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.TrimSpace(t) == "" {
			continue
		}

		if !strings.Contains(t, "=") {
			return nil, fmt.Errorf("missing '=' in %q", t)
		}

		env = append(env, t)
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return nil, err
	}

	return env, nil
}

const usage = `Usage of envjector:

  Envjector reads an environment file and then execs a child process with the
  given environment. The child process is specified after "--".

Example:

  $ envjector -file /path/to/envfile -- myapp -flag1

Flags:

`
