package sinasco.aws.security.sg
import input as tfplan

#// Total score for the validation
quality_gate = 5

#// Marks assigned for validations
quality_values = {
    "aws_security_group_rule": {"ingress":10}
}

#// Cloud resources measured in the validation
resource_types = {"aws_security_group_rule"}

#// Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

#// Compute the score for Security Group Ingress Rules
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            ingress_source := crud["ingress"] * validate_sg_ingress[resource_type];
            res := ingress_source
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
violation["One or more security group ingress rules are exposed to internet. Please change the source to 10.0.0.0/8"] {
    validate_sg_ingress[resource_types[_]] > 0
}

#// Enforce ingress sources to organization intranet
validate_sg_ingress[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    creates := [res | res:= all[_]; res.change.after.type == "ingress"; res.change.after.cidr_blocks[_] == "0.0.0.0/0"];
    num := count(creates)
}