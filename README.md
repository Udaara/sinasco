Sinasco
=======

<p align="center">
<img alt="Sinasco" src="https://github.com/Udaara/sinasco/blob/main/assets/images/sinasco.png" width="200px">
</p>

Sinasco is a standardization framework for Infrastructure as Code. Designed with cloud computing principals in mind, it allows users to cross validate their IaC scripts against organizational and security policies and configuration rules to enable a well managed and secure cloud infrastructure, without compromising the benefits an organization can harness in a cloud centric infrastructure ecosystem. Sinasco includes a library of rules written in Golang alogn with a wrapper written for Open Policy Agent (OPA) binaries.


Sinasco can be used to manage [Terraform][1], a popular Infrastructure as Code tool, scripts for infrastructure provisioning on [AWS][2], [Azure][3], and [GCP][4].
<br>
<br>

Getting Started with Sinasco
-----------------------------
#### Install Requirements

Sinasco requires the following tools to function:
- [OPA][6]
- [Terraform][7]

Installation can be verified as below
    
    opa version
      Version: 0.28.0
      Build Commit: 3fbcd71
      Build Timestamp: 2021-04-27T13:51:21Z
      Build Hostname: c8a0b3ab05bf
      Go Version: go1.15.8

    terraform -v
      Terraform v0.14.4

Once OPA and Terraform are set, navigate to the Sinasco [releases][8] and download the latest binary (sinasco.zip)
<b>bin</b> directory contains the sinasco wrapper, which is used to run the framework. <b>lib</b> directory contains Sinasco managed policies for unit tests, organizational compliance and security, for major cloud providers which can be used to evaluate and standardize the terraform code. Additional policies can be added to here

#### Evaluating your IaC Code

Once the steps on the <b>Install Requirements</b> are done, we can proceed to evaluate the Terraform code. Sinasco requires 3 user inputs to evalaute the code
| Input Flag | Description                               | Sample Input                                                                     |
|------------|-------------------------------------------|----------------------------------------------------------------------------------|
| -d         | Sinasco rule file to evalaute the code    | lib/aws/security/datastore.go                                                    | 
| -i         | Directory with Terraform Code             | RP-Code/nonprod/                                                                 | 
| -f         | Evaluation output format                  | <b>score</b> - Quality Gate Evaluation or <b>violation</b> - Violated policies   | 
<br>

Sinasco can be used to show the violated rule through `violation` flag

    sinasco.sh -d lib/aws/security/datastore.go -i RP-Code/nonprod/ -f violation
      +--------------------------------+
      |           violations           |
      +--------------------------------+
      | ["One or more S3 Buckets are   |
      | public. Please change the ACL  |
      | to Private"]                   |
      +--------------------------------+

Sinasco can be used to show the cumulative marks assigned to violates rules through `score` flag

    sinasco.sh -d lib/aws/security/datastore.go -i RP-Code/nonprod/ -f score
      +------------+
      | score      |
      +------------+
      | 10         |
      +------------+

Sinasco can be used to measure whether given resource stack passed the custom quality  gate through `quality_gate_passed` flag

    sinasco.sh -d lib/aws/security/datastore.go -i RP-Code/nonprod/ -f quality_gate_passed
      +---------------------+
      | Quality Gate Passed |
      +---------------------+
      | false               |
      +---------------------+

<br>
The key features of Sinasco
---------------------------------------
#### :cloud: Multi-Cloud Supportability
With industry moving towards multi-cloud, it is vital to be CSP-independent and move cloud services between different environments with standardized service definitions.
Sinasco can be used to verify multi-cloud supportability of your code, by providing cross-validation on the policy docements against Terraform scripts written for any major cloud services provider.
<br>

#### :link: Terraform Module Validations
Validate whether the approved Terraform modules are used in the code. Modules which available over the internet could compromised the infrastructure built by violating the organizational standards.
Sinasco can be used to force organizational modules for the scripts.
<br>

#### :page_facing_up: Organizational Policy Validation
Sinasco policies are written in simple YAML configuration files that enable users to specify policies on a resource type, purpose and environment. These policy files are CSP independent and can be used on Terraform script written for any CSP.
<br>

#### :lock: Infrastructure Drift Validations
Detect drift on an entire stack or on a particular resource by comparing the current stack configuration to the one specified in the template that was used to create or update the stack.
<br>

#### :chart_with_upwards_trend: Measure Health & Custom Quality Gates
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
  [6]: https://www.openpolicyagent.org/docs/latest/#1-download-opa
  [7]: https://www.terraform.io/downloads.html
  [8]: https://github.com/Udaara/sinasco/releases
