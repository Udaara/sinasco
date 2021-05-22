package sinasco.aws.organization.region
import input as tfplan

#// Total score for the validation
quality_gate = 5

#// Marks assigned for validations
quality_values = {
    "aws_instance": {"region":10},
    "aws_security_group": {"region":10}
}

#// Cloud resources measured in the validation
resource_types = {"aws_instance","aws_security_group"}

#// Quality Gate Evaluation
default quality_gate_passed = false
quality_gate_passed {
    score < quality_gate
}

#// Compute the score for the terraform gold module usage
score = eval {
    all := [ res |
            some resource_type
            crud := quality_values[resource_type];
            ec2_region := crud["region"] * ec2_region_validation[resource_type];
            sg_region := crud["region"] * sg_region_validation[resource_type];
            res := ec2_region + sg_region
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
violation["EC2 are provisioned in unsupported region(s). Please use a VPC in us-east-1 (North Virginia)"] {
    ec2_region_validation[resource_types[_]] > 0
}

violation["Security Groups are provisioned in unsupported region(s). Please use a VPC in us-east-1 (North Virginia)"] {
    sg_region_validation[resource_types[_]] > 0
}

#// Enforce region to us-east-1
ec2_region_validation[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    creates := [res |  res:= all[_]; res.change.after.tags.VPC != "vpc-0a213e474abafbf7b"]
    num := count(creates)
}

sg_region_validation[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := resources[resource_type]
    creates := [res |  res:= all[_]; res.change.after.vpc_id != "vpc-0a213e474abafbf7b"]
    num := count(creates)
}
