package research.aws.security.encryption
import input as tfplan

#// Total score for the validation
quality_gate = 5

#// Marks assigned for validations
quality_values = {
    "aws_s3_bucket": {"sse": 10}
}

#// Cloud resources measured in the validation
resource_types = {"aws_s3_bucket"}

#// Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

#// Compute the score for encryption
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            s3_encryption := crud["sse"] * validate_s3_encryption[resource_type];
            res := s3_encryption
    ]
    eval := sum(all)
}

#// List all resources json objects
resources[resource_type] = all {
    some resource_type
    resource_types[resource_type]
    all := [name |
        name:= tfplan.resource_changes[_]
        name.type == resource_type
    ]
}

#// Error message to display on a violation
violation["One or more S3 Buckets are not encrypted"] {
    validate_s3_encryption[resource_types[_]] > 0
}

#// Enforce ingress sources to organization intranet
validate_s3_encryption[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    modifies := [res |  res:= all[_]; res.change.after.server_side_encryption_configuration[_].rule[_].apply_server_side_encryption_by_default[_].sse_algorithm != "AES256"]; #//AES256 / aws:kms
    num := count(modifies)
}
