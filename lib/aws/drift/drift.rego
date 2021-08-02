package sinasco.aws.drift
import input as tfplan

# Total score for the validation
quality_gate = 5

# Marks assigned for validations
quality_values = {
    "aws_s3_bucket": {"violation": 10},
    "aws_security_group": {"violation": 10},
    "aws_security_group_rule": {"violation": 10},
    "aws_instance": {"violation": 10}
}

# Cloud resources measured in the validation
resource_types = {"aws_s3_bucket","aws_security_group","aws_security_group_rule","aws_instance"}

# Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

# Compute the score for encryption
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            violation := crud["violation"] * evaluate_drift[resource_type];
            res := violation
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
violation["Drift detected!"] {
    evaluate_drift[resource_types[_]] > 0
}

# Enforce ingress sources to organization intranet
evaluate_drift[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    modifies := [res |  res:= all[_]; res.change.actions[_] == "update"];
    num := count(modifies)
}
