provider "github" {
  owner = "Mkamono"
}

provider "google" {
  project = "coastal-fiber-430001-s0"
  region  = "asia-northeast1"
}

data "google_project" "project" {
}

provider "google-beta" {
  project = "coastal-fiber-430001-s0"
  region  = "asia-northeast1"
}
