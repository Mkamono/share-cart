provider "google" {
  project = "coastal-fiber-430001-s0"
  region  = "asia-northeast1"
}

data "google_project" "share-cart" {
}
