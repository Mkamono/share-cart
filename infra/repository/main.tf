terraform {
  required_version = ">= 1.9.2"

  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "5.37.0"
    }
  }

  backend "gcs" {
    bucket = "kamono-terraform-state-bucket"
    prefix = "terraform/state/repository"
  }
}
