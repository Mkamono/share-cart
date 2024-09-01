module "cloud-dns" {
  source  = "terraform-google-modules/cloud-dns/google"
  version = "5.2.0"

  project_id = local.project_id
  name       = "share-cart-com"
  domain     = "share-cart.com."
  type       = "public"

  dnssec_config = {
    state = "on"
  }

  recordsets = [
    {
      name = ""
      type = "A"
      ttl  = 300
      records = [
        module.lb-http.external_ip,
      ]
    },
  ]
}
