# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

variable "project_id" {
description = "The project ID to host the crypto key."
type = string
default = "foobar"
}

variable "organization_id" {
description = "The organization ID."
type = string
default = "12345"
}

provider "google" {
project = "{{.Provider.project}}"
}

resource "google_kms_key_ring" "gg-asset-43537-eb9a" {
name = "gg-asset-43537-eb9a"
location = "us-central1"
}

resource "google_kms_crypto_key" "gg-asset-43537-eb9a" {
name = "gg-asset-43537-eb9a"
key_ring = google_kms_key_ring.gg-asset-43537-eb9a.id
purpose = "ASYMMETRIC_SIGN"
version_template {
algorithm = "EC_SIGN_P256_SHA256"
protection_level = "HSM"
}
labels = {
"example-label" = "example-value"
}
skip_initial_version_creation = true
}
