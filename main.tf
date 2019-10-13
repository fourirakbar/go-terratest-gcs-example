terraform {
  required_version = "~> 0.12.0"
}

resource "google_storage_bucket" "development" {
  name          = var.name
  location      = var.location
  project       = var.project
  storage_class = var.storage_class
  force_destroy = var.force_destroy
}
