package sinasco.aws.security.ami
import input as tfplan

# Total score for the validation
quality_gate = 5

# Marks assigned for validations
quality_values = {
    "aws_instance": {"ami":10}
}

# Cloud resources measured in the validation
resource_types = {"aws_instance"}

# Gold AMIs
approved_amis = {"ami-0d5eff06f840b45e9"}

# Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

# Compute the score for the golden AMI usage
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            ami := crud["ami"] * ami_validation[resource_type];
            res := ami
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
violation["Planning to provision one or more EC2 servers with unauthorized AMIs. Please use a Gold AMI"] {
    ami_validation[resource_types[_]] > 0
}

# Enforce to use Gold AMIs for EC2 provisioning
ami_validation[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    val = false
    creates := [res |  res:= all[_]; not val = (res.change.after.ami != approved_amis[_])]
    num := count(creates)
}
