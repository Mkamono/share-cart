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
    bucket = "share-cart-terraform-state-bucket"
    prefix = "terraform/state/local"
  }
}

resource "google_storage_bucket" "terraform_state_bucket" {
  name          = "share-cart-terraform-state-bucket"
  location      = "ASIA-NORTHEAST1"
  storage_class = "STANDARD"
}
