terraform {
  provider "tfe" {
    hostname = "app.terraform.io"
    token = "eM6P0JyKaZzJaA.atlasv1.U0oMslxMeZtNlTMM5owqnPg0MomT8GIHa2GhyrTNl0zgE7vzksFfzCqHF2Wjzyuaflw"
  }
  cloud {
    organization = "My-Organization394"

    workspaces {
      name = "test-terratest"
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
