package sinasco.aws.organization.naming
import input as tfplan

# Total score for the validation
quality_gate = 5

# Marks assigned for validations
quality_values = {
    "aws_instance": {"naming":10}
}

# Cloud resources measured in the validation
resource_types = {"aws_instance"}

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
            ec2_naming := crud["naming"] * ec2_naming_validation[resource_type];
            res := ec2_naming
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
violation["EC2 naming standard violated. Please refer udaara.confluence.com/org_naming for the standard documentation"] {
    ec2_naming_validation[resource_types[_]] > 0
}

# Enforce the instance naming standard
ec2_naming_validation[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    val = true
    creates := [res | res:= all[_]; not val = (glob.match("aue1[l,w][d,q,s,p][a-z][a-z][a-z][0-9][0-9]", [], res.change.after.tags.Name))];
    num := count(creates)
}
