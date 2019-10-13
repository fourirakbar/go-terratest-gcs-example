variable "name" {
  description = "The name of the bucket"
}

variable "location" {
  description = "Location of the bucket"
}

variable "project" {
  description = "Project of the bucket"
}

variable "storage_class" {
  description = "Storage class of the bucket"
}

variable "force_destroy" {
  description = "Force destroy when deleting bucker"
}
