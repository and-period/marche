module "dns_and_period_jp" {
  source = "./../../../modules/route53"

  #####################################################################
  # Common
  #####################################################################
  tags     = var.tags

  #####################################################################
  # Route53 (zone)
  #####################################################################
  domain  = "and-period.jp"
  comment = "&. 本番用"

  subdomains = []
}
