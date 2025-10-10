resource "google_vmwareengine_network" "acc_network" {
  name        = "pc-nw"
  location    = "global"
  type        = "STANDARD"
  description = "PC network description."
}

resource "google_vmwareengine_private_cloud" "acc_pc" {
  location    = "us-central1-a"
  name        = "acc-pc"
  description = "PC description."
  network_config {
    management_cidr = "192.168.0.0/24"
    vmware_engine_network = google_vmwareengine_network.acc_network.id
  }
}

resource "google_vmwareengine_external_address" "acc_addr" {
  name        = "acc-addr"
  parent      = google_vmwareengine_private_cloud.acc_pc.id
  internal_ip = "192.168.0.10"
  description = "Sample external address"
}
