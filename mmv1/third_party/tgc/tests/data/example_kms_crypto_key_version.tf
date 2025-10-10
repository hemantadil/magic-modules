resource "google_kms_key_ring" "keyring" {
  name     = "keyring-example"
  location = "global"
}

resource "google_kms_crypto_key" "cryptokey" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.id
  rotation_period = "7776000s"
}

resource "google_kms_crypto_key_version" "example-key" {
  crypto_key = google_kms_crypto_key.cryptokey.id
}
