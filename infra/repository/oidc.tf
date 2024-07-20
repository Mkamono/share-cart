module "gh_oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  project_id  = local.project_id
  pool_id     = "gha-pool"
  provider_id = "gha-provider"
  sa_mapping = {
    (google_service_account.gha_wif_sa.account_id) = {
      sa_name   = google_service_account.gha_wif_sa.name
      attribute = "attribute.repository/${github_repository.share-cart.full_name}"
    }
  }
}
