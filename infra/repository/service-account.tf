resource "google_service_account" "wif_sa" {
  account_id   = "wif-sa"
  display_name = "WIF Service Account"
}
