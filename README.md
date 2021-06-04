Sinasco
=======

<p align="center">
<img alt="Sinasco" src="https://github.com/Udaara/sinasco/blob/main/assets/images/sinasco.png" width="200px">
</p>

Sinasco is a standardization framework for Infrastructure as Code. Designed with cloud computing principals in mind, it allows users to cross validate their IaC scripts against organizational and security policies and configuration rules to enable a well managed and secure cloud infrastructure, without compromising the benefits an organization can harness in a cloud centric infrastructure ecosystem. Sinasco includes a library of rules written in Golang alogn with a wrapper written for Open Policy Agent (OPA) binaries.


Sinasco can be used to manage [Terraform][1], a popular Infrastructure as Code tool, scripts for infrastructure provisioning on [AWS][2], [Azure][3], and [GCP][4].
<br>
<br>


The key features of Sinasco
---------------------------------------
#### :link: Terraform Module Validations
Validate whether the approved Terraform modules are used in the code. Modules which available over the internet could compromised the infrastructure built by violating the organizational standards.
Sinasco can be used to force organizational modules for the scripts.
<br>

#### :lock: Terraform Syntax Validations
Validate whether the valid Terraform syntaxes are used in the code. As the validations conducts before initializing and building the code, it provides faster validation.
<br>

#### :page_facing_up: Organizational Policy Validation
Managed Sinasco rules can be utilized for standard infrastructure evaluations. However, users are encouraged to write their own rules, catered for the organizational standrads and policies.
<br>

#### :chart_with_upwards_trend: Measure Health & Custom Quality Gates
Measure the health of the Terraform script depending on the environment and create a quality report for the build. Further, weight and measure compliance & policy violations on severity, environment and the impact to generate the quality score. The quality score can be used to determine whether the code passes the custom quality gates built for the project, thus reject or create the infrastructure on a fully-automated manner.
<br>

#### :cloud: Multi-Cloud Supportability
With industry moving towards multi-cloud, it is vital to be CSP-independent and move cloud services between different environments with standardized service definitions.
Sinasco can be used to verify multi-cloud supportability of your code, by providing cross-validation on the policy docements against Terraform scripts written for any major cloud services provider.
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
| Input Flag | Description                               | Sample Input                                                                                                                                  |
|------------|-------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| -d         | Sinasco rule file to evalaute the code    | lib/aws/security/datastore.go                                                                                                                 | 
| -i         | Directory with Terraform Code             | RP-Code/nonprod/                                                                                                                              | 
| -f         | Evaluation output format                  | <b>score</b> - Quality Gate Evaluation or <b>violation</b> - Violated policies or <b>quality_gate_passed</b> - Final Evaluation of the Code   | 
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

Structure of the Sinasco Rule File
-----------------------------------

Setting the Quality Gateway. score is calculated through the summation of the violated rules multiplied by the weight assigned for the rule.

```
// Cumulative score for the validation
quality_gate = 5

// Weights assigned for validation rule
quality_values = {
    "aws_instance": {"naming":10}
}

// Compute the score for the terraform gold module usage
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            ec2_naming := crud["naming"] * ec2_naming_validation[resource_type];
            res := ec2_naming
    ]
    eval := sum(all)
}

// Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}
```

Evaluating the defined standardization rule and providing the output is handled by the below section
```
// Error message to display on a violation
violation["EC2 naming standard violated. Please refer udaara.confluence.com/org_naming for the standard documentation"] {
    ec2_naming_validation[resource_types[_]] > 0
}

// Enforce the instance naming standard
ec2_naming_validation[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    val = true
    creates := [res | res:= all[_]; not val = (glob.match("aue1[l,w][d,q,s,p][a-z][a-z][a-z][0-9][0-9]", [], res.change.after.tags.Name))];
    num := count(creates)
}
```
<br>

License
-------------

Sinasco is licensed under the [MIT License][5]

  [1]: https://www.terraform.io/
  [2]: https://aws.amazon.com/
  [3]: https://azure.microsoft.com/en-us/
  [4]: https://cloud.google.com/
  [5]: https://github.com/Udaara/sinasco/blob/main/LICENSE
  [6]: https://www.openpolicyagent.org/docs/latest/#1-download-opa
  [7]: https://www.terraform.io/downloads.html
  [8]: https://github.com/Udaara/sinasco/releases
