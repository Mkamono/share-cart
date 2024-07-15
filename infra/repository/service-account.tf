resource "google_service_account" "wif_sa" {
  account_id   = "wif-sa"
  display_name = "WIF Service Account"
}

# resource "google_service_account_iam_member" "wif_sa_iam_wif_user" {
#   service_account_id = google_service_account.wif_sa.id
#   role               = "roles/iam.workloadIdentityUser"
#   member             = "serviceAccount:${data.google_project.project.number}.svc.id.goog[default/gh-pool]"
# }

# resource "google_service_account_iam_member" "wif_sa_iam_editor" {
#   service_account_id = google_service_account.wif_sa.id
#   role               = "roles/editor"
#   member             = "serviceAccount:${data.google_project.project.number}.svc.id.goog[default/gh-pool]"
# }


# resource "google_service_account_iam_member" "wif_sa_iam" {
#   service_account_id = google_service_account.wif_sa.name
#   role               = "roles/iam.workloadIdentityUser"
#   member             = "serviceAccount:${data.google_project.project.number}.svc.id.goog[default/gh-pool]"
# }
