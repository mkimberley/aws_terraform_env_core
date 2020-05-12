provider "aws" {
  region = "eu-west-2"
}

module "generic_vpc" {
  source                           = "git::https://github.com/mkimberley/aws_terraform_mod_common//network/vpc"
  #source                           = "../aws_terraform_mod_common//network/vpc"
  vpc_environment                  = var.environment
  vpc_name                         = var.vpc_name
  vpc_cidr                         = var.vpc_cidr
  vpc_tags                         = var.vpc_tags
  vpc_instance_tenancy             = var.vpc_instance_tenancy
  vpc_enable_dns_hostnames         = var.vpc_enable_dns_hostnames
  vpc_enable_dns_support           = var.vpc_enable_dns_support
  vpc_primary_sub_cidr             = var.public_subnets
  vpc_primary_az                   = var.sub_azs
  vpc_secondary_az                 = var.sub_azs
  vpc_secondary_sub_cidr           = var.private_subnets
  vpc_primary_subnet_name          = var.primary_subnet_name
  vpc_secondary_subnet_name        = var.secondary_subnet_name
  vpc_primary_public_ip_on_launch   = var.vpc_primary_public_ip_on_launch
  vpc_secondary_public_ip_on_launch = var.vpc_secondary_public_ip_on_launch
}

module "backend-subnets" {
  source                           = "git::https://github.com/mkimberley/aws_terraform_mod_common//network/subnet"
  #source                           = "../aws_terraform_mod_common//network/subnet"
  sub_cidr_block                   = var.backend_subnets
  sub_public_ip_on_launch          = false
  sub_vpc_id                       = module.generic_vpc.vpc_id
  sub_environment                  = var.environment
  sub_name                         = var.backend_subnet_name
  sub_azs                          = var.sub_azs
  subnets                          = var.backend_subnets
  sub_tags = {
    Name                           = "Back-end Subnet"
    Description                    = "Automation Services"
  }
}

