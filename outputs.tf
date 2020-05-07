output "public_subnets" {
  value = module.generic_vpc.primary_subnet_map
}

output "private_subnets" {
  value = module.generic_vpc.secondary_subnet_map
}

output "core_vpc_id" {
  value = module.generic_vpc.vpc_id
}