resource "google_vmwareengine_network" "cluster-nw" {
  name     = "cluster-nw-{{.RANDOM_SUFFIX}}"
  location = "global"
  type     = "STANDARD"
}

resource "google_vmwareengine_private_cloud" "cluster-pc" {
  location    = "us-central1-a"
  name        = "cluster-pc-{{.RANDOM_SUFFIX}}"
  description = "PC for cluster test."

  network_config {
    management_cidr       = "192.168.0.0/24"
    vmware_engine_network = google_vmwareengine_network.cluster-nw.id
  }

  management_cluster {
    cluster_id = "management"
    node_type_configs {
      node_type_id = "standard-72"
      node_count   = 3
    }
  }
}

resource "google_vmwareengine_cluster" "acc-cluster" {
  name   = "cluster-acc-{{.RANDOM_SUFFIX}}"
  parent = google_vmwareengine_private_cloud.cluster-pc.id
  node_type_configs {
    node_type_id = "standard-72"
    node_count   = 3
  }
}
