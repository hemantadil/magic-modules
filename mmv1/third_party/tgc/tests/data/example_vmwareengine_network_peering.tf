resource "google_compute_network" "network-peering-vpc" {
  name = "network-peering-vpc-{{.RANDOM_SUFFIX}}"
}

resource "google_vmwareengine_network" "network-peering-standard-nw" {
  name     = "default-standard-nw-np-{{.RANDOM_SUFFIX}}"
  location = "global"
  type     = "STANDARD"
}

resource "google_vmwareengine_network_peering" "acc-peering" {
  name                  = "network-peering-peering-{{.RANDOM_SUFFIX}}"
  vmware_engine_network = google_vmwareengine_network.network-peering-standard-nw.id
  peer_network          = google_compute_network.network-peering-vpc.id
  peer_network_type     = "STANDARD"
}
