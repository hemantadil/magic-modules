resource "google_dataproc_cluster" "simple" {
  name   = "dproctest-simple"
  region = "us-central1"
}
