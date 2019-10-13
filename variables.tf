variable "name" {
  default     = ""
  description = "The name of the bucket"
}

variable "location" {
  default     = "asia-southeast1"
  description = "Location of the bucket"
}

variable "project" {
  default     = ""
  description = "Project of the bucket"
}

variable "storage_class" {
  default     = ""
  description = "Storage class of the bucket"
}

variable "force_destroy" {
  default     = "false"
  description = "Force destroy when deleting bucker"
}
