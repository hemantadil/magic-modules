terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.54.0"
    }
  }
}

provider "google" {
  project = "{{.Provider.project}}"
}

resource "google_iam_workload_identity_pool" "pool_1" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-50421-6718"
  display_name              = "gg-asset-50421-6718"
  description               = "Workload Identity Pool for gg-asset-50421-6718"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "provider_1" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_1.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-50421-6718"
  display_name                       = "gg-asset-50421-6718"
  description                        = "Workload Identity Pool Provider for gg-asset-50421-6718"

  aws {
    account_id = "111111111111"
  }
}

resource "google_iam_workload_identity_pool" "pool_2" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-53029-6165-pool"
  display_name              = "gg-asset-53029-6165-pool"
  description               = "Workload Identity Pool for gg-asset-53029-6165"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "provider_2" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_2.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-53029-6165"
  display_name                       = "gg-asset-53029-6165"
  description                        = "OIDC provider for gg-asset-53029-6165"
  attribute_mapping = {
    "google.subject" = "assertion.sub"
  }
  oidc {
    issuer_uri = "https://accounts.google.com"
  }
}

resource "google_iam_workload_identity_pool" "pool_3" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-pool-50569-cb71"
  display_name              = "gg-asset-pool-50569-cb71"
  description               = "Workload Identity Pool for gg-asset-50569-cb71"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "provider_3" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_3.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-provider-50569-cb71"
  display_name                       = "gg-asset-provider-50569-cb71"
  description                        = "AWS provider for gg-asset-50569-cb71"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.aws_role"   = "assertion.arn"
    "attribute.repository" = "assertion.repository"
  }
  attribute_condition = "assertion.repository_owner == 'my-owner'"

  aws {
    account_id = "111122223333"
  }
}

resource "google_iam_workload_identity_pool" "pool_4" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-50716-eec2"
  display_name              = "gg-asset-50716-eec2"
  description               = "A sample workload identity pool."
}

resource "google_iam_workload_identity_pool_provider" "provider_4" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_4.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-50716-eec2"
  display_name                       = "gg-asset-50716-eec2"
  description                        = "A sample workload identity pool provider."

  aws {
    account_id = "111122223333"
  }
}

resource "google_iam_workload_identity_pool" "pool_5" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-50882-8747"
  display_name              = "gg-asset-50882-8747"
  description               = "Workload Identity Pool for gg-asset-50882-8747"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "provider_5" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_5.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-50882-8747"
  display_name                       = "gg-asset-50882-8747"
  description                        = "Workload Identity Pool Provider for gg-asset-50882-8747"
  disabled                           = true

  aws {
    account_id = "111122223333"
  }
}

resource "google_iam_workload_identity_pool" "pool_6" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-51035-89e8"
  display_name              = "gg-asset-51035-89e8"
  description               = "Workload Identity Pool for gg-asset-51035-89e8"
}

resource "google_iam_workload_identity_pool_provider" "provider_6" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_6.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-51035-89e8"
  display_name                       = "gg-asset-51035-89e8"
  description                        = "Workload Identity Pool Provider for gg-asset-51035-89e8"
  oidc {
    issuer_uri = "https://accounts.google.com"
  }
  attribute_mapping = {
    "google.subject" = "assertion.sub"
  }
}

resource "google_iam_workload_identity_pool" "pool_7" {
  project                   = "{{.Provider.project}}"
  workload_identity_pool_id = "gg-asset-pool-51207-ce6d"
  display_name              = "gg-asset-pool-51207-ce6d"
  description               = "Workload Identity Pool for gg-asset-51207-ce6d"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "provider_7" {
  project                            = "{{.Provider.project}}"
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool_7.workload_identity_pool_id
  workload_identity_pool_provider_id = "gg-asset-provider-51207-ce6d"
  display_name                       = "gg-asset-provider-51207-ce6d"
  description                        = "OIDC provider for gg-asset-51207-ce6d"
  attribute_mapping = {
    "google.subject"       = "assertion.sub",
    "attribute.actor"      = "assertion.actor",
    "attribute.aud"        = "assertion.aud",
    "attribute.repository" = "assertion.repository"
  }
  oidc {
    issuer_uri        = "https://oidc.gg-asset.com/51207-ce6d"
    allowed_audiences = ["gg-asset-51207-ce6d-audience"]
  }
}
