resource "google_vmwareengine_network" "pc-nw" {
  name     = "pc-nw-{{.RANDOM_SUFFIX}}"
  location = "global"
  type     = "STANDARD"
}

resource "google_vmwareengine_private_cloud" "acc-pc" {
  location    = "us-central1-a"
  name        = "pc-acc-{{.RANDOM_SUFFIX}}"
  description = "PC description."

  network_config {
    management_cidr       = "192.168.0.0/24"
    vmware_engine_network = google_vmwareengine_network.pc-nw.id
  }

  management_cluster {
    cluster_id = "management"
    node_type_configs {
      node_type_id = "standard-72"
      node_count   = 3
    }
  }
}
