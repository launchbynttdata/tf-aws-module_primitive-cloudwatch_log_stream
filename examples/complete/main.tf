// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module "cloudwatch_log_stream" {
  source = "../.."

  name                      = module.resource_names["log_stream"].standard
  cloudwatch_log_group_name = module.resource_names["log_group"].standard

  //Without adding this explicit depedency, creation of log stream results into random error
  //Error: creating CloudWatch Logs Log Stream (<<log stream name>>): ResourceNotFoundException: The specified log group does not exist.
  depends_on = [module.cloudwatch_log_group]
}

module "cloudwatch_log_group" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/cloudwatch_log_group/aws"
  version = "~> 1.0"

  name = module.resource_names["log_group"].standard
}

module "resource_names" {
  source = "github.com/launchbynttdata/tf-launch-module_library-resource_name.git?ref=1.0.1"

  for_each = var.resource_names_map

  logical_product_family  = "terratest"
  logical_product_service = "cloudwatch"
  region                  = join("", split("-", var.region))
  class_env               = var.environment
  cloud_resource_type     = each.value.name
  instance_env            = var.environment_number
  instance_resource       = var.resource_number
  maximum_length          = each.value.max_length
}
