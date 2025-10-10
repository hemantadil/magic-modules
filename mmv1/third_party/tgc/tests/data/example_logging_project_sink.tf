resource "google_storage_bucket" "log-bucket" {
  name     = "project-logging-bucket-{{.RANDOM_SUFFIX}}"
  location = "US"
}

resource "google_logging_project_sink" "my-sink" {
  name        = "my-project-sink"
  project     = "{{.Provider.project}}"
  destination = "storage.googleapis.com/${google_storage_bucket.log-bucket.name}"
  filter      = "resource.type = gce_instance AND severity >= WARNING"
  unique_writer_identity = true
}
