provider "github" {
  owner = "Mkamono"
}

provider "google" {
  project = "kamono-personal"
  region  = "asia-northeast1"
}

data "google_project" "project" {
}

# provider "google-beta" {
#   project = "kamono-personal"
#   region  = "asia-northeast1"
# }
