terraform {
  required_version = ">= 1.9.2"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.37.0"
    }
  }

  backend "gcs" {
    bucket = "share-cart-terraform-state-bucket"
    prefix = "terraform/state/remote"
  }
}
