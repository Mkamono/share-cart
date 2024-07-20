resource "google_service_account" "gha_wif_sa" {
  account_id   = "gha-wif-sa"
  display_name = "WIF Service Account for GitHub Actions"
}
