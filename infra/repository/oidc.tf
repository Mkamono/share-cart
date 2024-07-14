module "gh_oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  project_id  = data.google_project.project.project_id
  pool_id     = "gh-pool"
  provider_id = "gh-provider"
  sa_mapping = {
    "foo-service-account" = {
      sa_name   = google_service_account.wif_sa.name
      attribute = "attribute.repository/${github_repository.share-cart.full_name}"
    }
  }
}

data "google_iam_workload_identity_pool_provider" "wif_pool_provider" {
  provider                           = google-beta
  workload_identity_pool_id          = "gh-pool"
  workload_identity_pool_provider_id = "gh-provider"
}
