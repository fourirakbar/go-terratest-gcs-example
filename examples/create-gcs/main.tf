module "default" {
  source = "../../"

  # Bucket detail
  name          = "testing-terratest-gcs"
  location      = "asia-southeast1"
  project       = "oing-experiment"
  storage_class = "REGIONAL"
  force_destroy = false
}
