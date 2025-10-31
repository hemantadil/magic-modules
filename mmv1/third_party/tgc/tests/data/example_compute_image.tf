provider "google" {
  project = "{{.Provider.project}}"
}

resource "google_kms_key_ring" "gg_asset_08944_f101_key_ring" {
  name     = "gg-asset-08944-f101-key-ring"
  location = "us-central1"
}

resource "google_kms_crypto_key" "gg_asset_08944_f101_crypto_key" {
  name            = "gg-asset-08944-f101-crypto-key"
  key_ring        = google_kms_key_ring.gg_asset_08944_f101_key_ring.id
  rotation_period = "100000s"
}

resource "google_kms_crypto_key_iam_member" "gg_asset_08944_f101_iam_binding" {
  crypto_key_id = google_kms_crypto_key.gg_asset_08944_f101_crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-{{.Project.Number}}@compute-system.iam.gserviceaccount.com"
}

resource "google_kms_crypto_key_iam_member" "gce_sa_kms_binding" {
  crypto_key_id = google_kms_crypto_key.gg_asset_08944_f101_crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-{{.Project.Number}}@compute-system.iam.gserviceaccount.com"
}

resource "google_compute_disk" "gg_asset_08944_f101_disk" {
  name  = "gg-asset-08944-f101-disk"
  image = "debian-cloud/debian-11"
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "gg_asset_08944_f101_snapshot" {
  name        = "gg-asset-08944-f101-snapshot"
  source_disk = google_compute_disk.gg_asset_08944_f101_disk.id
  zone        = "us-central1-a"
  snapshot_encryption_key {
    kms_key_self_link = google_kms_crypto_key.gg_asset_08944_f101_crypto_key.id
  }
  depends_on = [google_kms_crypto_key_iam_member.gce_sa_kms_binding]
}

resource "google_compute_image" "gg_asset_08944_f101_image" {
  name            = "gg-asset-08944-f101-image"
  source_snapshot = google_compute_snapshot.gg_asset_08944_f101_snapshot.id
  guest_os_features {
    type = "MULTI_IP_SUBNET"
  }
  image_encryption_key {
    kms_key_self_link = google_kms_crypto_key.gg_asset_08944_f101_crypto_key.id
  }
  depends_on = [google_kms_crypto_key_iam_member.gce_sa_kms_binding]
}
