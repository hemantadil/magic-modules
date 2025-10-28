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

resource "google_service_account" "test_account_1" {
  account_id   = "gg-asset-34338-63e0"
  display_name = "Test Service Account 1"
  description  = "A test service account."
}

resource "google_service_account" "test_account_2" {
  account_id   = "gg-asset-34872-33bd"
  display_name = "Test Service Account 2"
}

resource "google_service_account" "test_account_3" {
  account_id = "gg-asset-34952-20de"
}
