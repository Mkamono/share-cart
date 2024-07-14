resource "google_service_account" "wif_sa" {
  account_id   = "wif-sa"
  display_name = "WIF Service Account"
}

resource "google_service_account_iam_member" "wif_sa_member" {
  service_account_id = google_service_account.wif_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.gh_identity_pool.name}/attribute.repository/${github_repository.share-cart.full_name}"
}
