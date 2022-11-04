variable "zone" {
  type = string
}

variable "region" {
  type = string
}

variable "env" {
  type = string
}

variable "project_id" {
  type        = string
  description = "Your project ID"
}

variable "access_key" {
    type = string
    description = "Scaleway authtentication token used in the function"
}

variable "secret_key" {
    type = string
}

variable "namespace_id" {
    type = string
}