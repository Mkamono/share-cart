resource "google_service_account" "wif_sa" {
  account_id   = "wif-sa"
  display_name = "WIF Service Account"
}

# data "google_iam_policy" "admin" {
#   binding {
#     role = "roles/iam.serviceAccountUser"

#     members = []
#   }
# }

resource "google_service_account_iam_member" "wif_sa_iam" {
  service_account_id = google_service_account.wif_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.number}.svc.id.goog[default/gh-pool]"
}


# data "google_iam_policy" "admin" {
#   binding {
#     role = "roles/iam.serviceAccountUser"

#     members = [
#       "user:jane@example.com",
#     ]
#   }
# }

# resource "google_service_account" "sa" {
#   account_id   = "my-service-account"
#   display_name = "A service account that only Jane can interact with"
# }

# resource "google_service_account_iam_policy" "admin-account-iam" {
#   service_account_id = google_service_account.sa.name
#   policy_data        = data.google_iam_policy.admin.policy_data
# }
