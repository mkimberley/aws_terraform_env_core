variable "environment" {
}
variable "vpc_name" {
}
variable "vpc_cidr" {
}
variable "vpc_instance_tenancy" {
}
variable "vpc_enable_dns_hostnames" {
}
variable "vpc_enable_dns_support" {
}

variable "public_subnets" {
    description = "A list of subnets"
    type        = list(string)
    default     = []
}

variable "private_subnets" {
    description = "A list of subnets"
    type        = list(string)
    default     = []
}

variable "backend_subnets" {
    description = "A list of subnets"
    type        = list(string)
    default     = []
}

variable "vpc_tags" {
    default = { Name = "Built via terraform"}
}

variable "primary_subnet_name" {
}

variable "secondary_subnet_name" {
}
variable "backend_subnet_name" {
}
variable "vpc_primary_public_ip_on_launch" {
}

variable "vpc_secondary_public_ip_on_launch" {
}

variable "sub_azs" {
    description = "A list of AZs"
    type        = list(string)
    default = ["eu-west-2a", "eu-west-2b", "eu-west-2c"]
}

