terraform {
  backend "remote" {
    organization = "My-Organization394"

    workspaces {
      name = "my-app-prod"
    }
  }
}