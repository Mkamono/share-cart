# resource "github_actions_secret" "wif_provider_id" {
#   repository  = github_repository.share-cart.name
#   secret_name = "WIF_PROVIDER_NAME"
# }

# resource "github_actions_secret" "wif_sa_email" {
#   repository      = github_repository.share-cart.name
#   secret_name     = "WIF_SA_EMAIL"
#   plaintext_value = google_service_account.gha_wif_sa.email
# }
