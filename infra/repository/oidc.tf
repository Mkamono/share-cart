resource "google_iam_workload_identity_pool" "gh_identity_pool" {
  workload_identity_pool_id = "github-actions-pool"
  display_name              = "github-actions-pool"
  disabled                  = false
}

resource "google_iam_workload_identity_pool_provider" "pool-provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.gh_identity_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-actions-provider"
  display_name                       = "github-actions-provider"
  description                        = "Github Actions OIDC Provider"
  disabled                           = false

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.actor"      = "assertion.actor"
    "attribute.aud"        = "assertion.aud"
    "attribute.repository" = "assertion.repository"
  }

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}
