resource "google_vmwareengine_network" "acc-nw" {
  name     = "network-policy-nw-{{.RANDOM_SUFFIX}}"
  location = "global"
  type     = "STANDARD"
}

resource "google_vmwareengine_network_policy" "acc-np" {
  name                  = "network-policy-np-{{.RANDOM_SUFFIX}}"
  location              = "us-central1"
  project               = "{{.Provider.project}}"
  vmware_engine_network = google_vmwareengine_network.acc-nw.id
  edge_services_cidr    = "192.168.30.0/26"

  internet_access {
    enabled = true
  }

  external_ip {
    enabled = true
  }
}
