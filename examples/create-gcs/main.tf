module "development" {
  source = "../../"

  # Bucket detail
  name          = var.name
  location      = var.location
  project       = var.project
  storage_class = var.storage_class
  force_destroy = var.force_destroy
}
