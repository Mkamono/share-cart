resource "github_actions_secret" "workload_identity_pool_provider_id" {
  repository      = github_repository.share-cart.name
  secret_name     = "WIF_POOL_PROVIDER_ID"
  plaintext_value = google_iam_workload_identity_pool_provider.pool-provider.id
}
