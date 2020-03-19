# Copyright 2020 The Envjector Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

REPO = us-docker.pkg.dev/ratchet-builder/envjector/envjector

docker:
	@docker build -t ${REPO} .
	@docker push ${REPO}

test:
	@go test -count=1 -parallel=40 ./...
.PHONY: test
