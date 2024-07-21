provider "github" {
  owner = "Mkamono"
}

provider "google" {
  project = "coastal-fiber-430001-s0"
  region  = "asia-northeast1"
}

data "google_project" "share-cart" {
}

provider "google-beta" {
  project = data.google_project.share-cart.project_id
  region  = "asia-northeast1"
}
