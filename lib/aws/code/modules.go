package sinasco.aws.code.module
import input as tfplan

#// Total score for the validation
quality_gate = 5

#// Marks assigned for validations
quality_values = {
    "module_source": {"gold_modules":10}
}

#// Cloud resources measured in the validation
resource_types = {"module_source"}

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
            gold_modules := crud["gold_modules"] * validate_modules[resource_type];
            res := gold_modules
    ]
    eval := sum(all)
}

#// List all module source json objects
modules[resource_type] = all {
    some resource_type
    resource_types[resource_type]
    all := [name |
        name:= tfplan.configuration[_]
    ]
}

#// Error message to display on a violation
violation["Unauthorized Terraform module(s) detected. Please use the Gold modules"] {
   validate_modules[resource_types[_]] > 0
}

#// Validating the modules
validate_modules[resource_type] = num {
    some resource_type
    resource_types[resource_type]
    all := modules[resource_type]
    val = true
    creates := [res | res:= all[_]; not val = (glob.match("**Udaara/rp-code-modules.git**", [], res.module_calls[_].source))];
    num := count(creates)
}
