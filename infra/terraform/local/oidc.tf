module "gh_oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  project_id  = data.google_project.share-cart.project_id
  pool_id     = "gh-actions-pool"
  provider_id = "gh-actions-pool-provider"
  sa_mapping = {
    (google_service_account.gha_wif_sa.account_id) = {
      sa_name   = google_service_account.gha_wif_sa.name
      attribute = "attribute.repository/${github_repository.share-cart.full_name}"
    }
  }
}

data "google_iam_workload_identity_pool_provider" "gha_wif_provider" {
  provider                           = google-beta
  workload_identity_pool_id          = "gh-actions-pool"
  workload_identity_pool_provider_id = "gh-actions-pool-provider"
}
