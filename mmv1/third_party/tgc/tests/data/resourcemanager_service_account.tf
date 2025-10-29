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

resource "google_service_account" "gg-asset-34338-63e0" {
  account_id = "gg-asset-34338-63e0"
}

resource "google_service_account" "gg-asset-34872-33bd" {
  project      = "{{.Provider.project}}"
  account_id   = "gg-asset-34872-33bd"
  display_name = "gg-asset-34872-33bd"
}

resource "google_service_account" "gg_asset_34952_20de" {
  account_id   = "gg-asset-34952-20de"
  display_name = "gg-asset-34952-20de"
  description  = "A service account with a description."
}

resource "google_service_account" "gg-asset-35048-7183" {
  project      = "{{.Provider.project}}"
  account_id   = "gg-asset-35048-7183"
  display_name = "gg-asset-35048-7183"
  description  = "A service account with a display name and description."
}

resource "google_service_account" "gg_asset_35376_f9a2" {
  account_id   = "gg-asset-35376-f9a2"
  display_name = "gg-asset-35376-f9a2"
  project      = "{{.Provider.project}}"
  disabled     = false
}

resource "google_service_account" "gg-asset-35461-a4e7" {
  project      = "{{.Provider.project}}"
  account_id   = "gg-asset-35461-a4e7"
  display_name = "gg-asset-35461-a4e7"
}

resource "google_service_account" "gg_asset_35564_1d71" {
  account_id   = "gg-asset-35564-1d71"
  display_name = "gg-asset-35564-1d71"
  project      = "{{.Provider.project}}"

  create_ignore_already_exists = true
}
