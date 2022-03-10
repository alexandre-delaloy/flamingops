
variable "database_user" {
  type    = string
  default = "root"
}

variable "database_password" {
  type    = string
  default = "root"
}
variable "sqs_name" {
  type    = string
  default = "flamingops-sqs"
}