resource "github_actions_secret" "wif_pool_provider_id" {
  repository      = github_repository.share-cart.name
  secret_name     = "WIF_POOL_PROVIDER_ID"
  plaintext_value = data.google_iam_workload_identity_pool_provider.wif_pool_provider.name
}

resource "github_actions_secret" "wif_sa_email" {
  repository      = github_repository.share-cart.name
  secret_name     = "WIF_SA_EMAIL"
  plaintext_value = google_service_account.wif_sa.email
}
