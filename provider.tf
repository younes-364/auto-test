terraform {

  # cloud {
  #   organization = "My-Organization394"
  #   hostname = "app.terraform.io"
  #   workspaces {
  #     name = "test-auto"
  #   }
  # }
   backend "remote" {
    organization = "My-Organization394"

    workspaces {
      name = "my-app-prod"
    }
  }

  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.48.0"
    }
  }
}
provider "aws" {
  access_key = "AKIAVPGBT54OGVNNC3XX"
  secret_key = "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"
  region     = "eu-central-1"
}