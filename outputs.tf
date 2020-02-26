output "name" {
  value       = google_storage_bucket.development.name
  description = "Name of the bucket"
}

output "location" {
  value       = google_storage_bucket.development.location
  description = "Location of the bucker"
}

output "project" {
  value       = google_storage_bucket.development.project
  description = "Project of the bucker"
}

output "storage_class" {
  value       = google_storage_bucket.development.storage_class
  description = "Storage class of the bucket"
}

output "force_destroy" {
  value       = google_storage_bucket.development.force_destroy
  description = "Force destroy of the bucket"
}

output "bucket_url" {
  value       = google_storage_bucket.development.url
  description = "Bucket URL of the bucket"
}
