Sinasco
=======

Sinasco is a standardization framework for Infrastructure as Code. Designed with cloud computing principals in mind, it allows users to cross validate their IaC scripts against Sinasco managed or user defined policies and configuration rules to enable a well managed and secure cloud infrastructure, without compromising the benefits an organization can harness in a cloud centric infrastructure ecosystem.


Sinasco can be used to manage [Terraform][1], a popular Infrastructure as Code tool., scripts for infrastructure provisioning on [AWS][2], [Azure][3], and [GCP][4].
<br>
<br>

The key features of Sinasco
---------------------------------------
:cloud: #### <i class="icon-cloud "></i> Multi-Cloud Supportability
With industry moving towards multi-cloud, it is vital to be CSP-independent and move cloud services between different environments with standardized service definitions.
Sinasco can be used to verify multi-cloud supportability of your code, by providing cross-validation on the policy docements against Terraform scripts written for any major cloud services provider.
<br>

:link: #### <i class="icon-th-large"></i> Terraform Module Validations
Validate whether the approved Terraform modules are used in the code. Modules which available over the internet could compromised the infrastructure built by violating the organizational standards.
Sinasco can be used to force organizational modules for the scripts.
<br>

:page_facing_up: #### <i class="icon-file-text-alt"></i> Organizational policy validation
Sinasco policies are written in simple YAML configuration files that enable users to specify policies on a resource type, purpose and environment. These policy files are CSP independent and can be used on Terraform script written for any CSP.
<br>

:lock: #### <i class="icon-lock "></i> Infrastructure drift validations
Detect drift on an entire stack or on a particular resource by comparing the current stack configuration to the one specified in the template that was used to create or update the stack.
<br>

:chart_with_upwards_trend: #### <i class="icon-lock "></i> Measure Health & Custom Quality Gates
Measure the health of the Terraform script depending on the environment and create a quality report for the build. Further, weight and measure compliance & policy violations on severity, environment and the impact to generate the quality score. The quality score can be used to determine whether the code passes the custom quality gates built for the project, thus reject or create the infrastructure on a fully-automated manner.

<br>

----------


License
-------------

Sinasco is licensed under the [Mozilla Public License v2.0][5]

  [1]: https://www.terraform.io/
  [2]: https://aws.amazon.com/
  [3]: https://azure.microsoft.com/en-us/
  [4]: https://cloud.google.com/
  [5]: https://github.com/Udaara/sinasco/blob/main/LICENSE
