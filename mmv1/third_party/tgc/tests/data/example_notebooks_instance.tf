resource "google_notebooks_instance" "instance" {
  name         = "notebooks-instance"
  location     = "us-central1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-gpu"
  }
}
