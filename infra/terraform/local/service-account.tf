resource "google_service_account" "gha_wif_sa" {
  account_id   = "gha-wif-sa"
  display_name = "WIF Service Account for GitHub Actions"
}

resource "google_storage_bucket_iam_member" "gha_wif_sa_iam" {
  bucket = google_storage_bucket.terraform_state_bucket.name
  role   = "roles/storage.objectUser"
  member = "serviceAccount:${google_service_account.gha_wif_sa.email}"
}
