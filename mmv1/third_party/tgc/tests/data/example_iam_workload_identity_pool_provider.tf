resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "wi-pool-{{.RANDOM_SUFFIX}}"
  display_name              = "Workload Identity Pool"
  description               = "Workload Identity Pool for testing"
}

resource "google_iam_workload_identity_pool_provider" "provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "wi-provider-{{.RANDOM_SUFFIX}}"
  display_name                       = "Workload Identity Pool Provider"
  description                        = "Workload Identity Pool Provider for testing"
  oidc {
    issuer_uri = "https://sts.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47/"
  }
}
