terraform {
    backend "local" {

    }
    required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

variable "qr_request_queue_name" {
  default = "qr-request"
}

resource "aws_sqs_queue" "qr-request" {
  name  = var.qr_request_queue_name
}

output "qr_request_queue_name" {
  value = var.qr_request_queue_name
}