package sinasco.aws.organization.tags
import input as tfplan

# Total score for the validation
quality_gate = 5

# Marks assigned for validations
quality_values = {
    "aws_instance": {"tags": 10},
    "aws_security_group": {"tags": 10},
    "aws_s3_bucket": {"tags": 10}
}

# Cloud resources measured in the validation
resource_types = {"aws_instance","aws_security_group","aws_s3_bucket"}

# Required Tag keys
minimum_tags = {"appID", "appCode", "sysCode", "sysID", "infraOwner"}

# Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

# Compute the score for the terraform gold module usage
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            tags := crud["tags"] * validate_tags[resource_type];
            res := tags
    ]
    eval := sum(all)
}

# List all resources json objects
resources[resource_type] = all {
    some resource_type
    resource_types[resource_type]
    all := [name |
        name:= tfplan.resource_changes[_]
        name.type == resource_type
    ]
}

# Error message to display on a violation
violation["One or more required tags are missing. Please add all mandatory tags"] {
    validate_tags[resource_types[_]] > 0
}

# Enforce the mandatory tags
validate_tags[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    modifies := [res |  res:= all[_]; not tags_contain_proper_keys(res.change.after.tags)]
    num := count(modifies)
}
tags_contain_proper_keys(tags) {
    keys := {key | tags[key]}
    leftover := minimum_tags - keys
    leftover == set()
}
