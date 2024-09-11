package main

import (
	"fmt"

	"github.com/fhir-fli/fhirpath-go/fhir"
	"github.com/fhir-fli/fhirpath-go/fhirpath"
	"github.com/fhir-fli/fhirpath-go/pkg/containedresource"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
)

func evaluateFHIRPathExpressions(resourceData []byte, searchParams []SearchParameter) (map[string]interface{}, error) {
	results := make(map[string]interface{})

	unmarshaller, err := jsonformat.NewUnmarshaller("UTC", fhirversion.R4)
	if err != nil {
		return nil, fmt.Errorf("failed to create unmarshaller: %w", err)
	}

	msg, err := unmarshaller.Unmarshal(resourceData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal FHIR resource: %w", err)
	}

	cr := containedresource.Wrap(msg.(fhir.Resource))
	unwrappedResource := containedresource.Unwrap(cr)

	for _, param := range searchParams {
		compiledExpr, err := fhirpath.Compile(param.Expression)
		if err != nil {
			return nil, fmt.Errorf("failed to compile FHIRPath expression: %w", err)
		}
		result, err := compiledExpr.Evaluate([]fhir.Resource{unwrappedResource})
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate FHIRPath expression: %w", err)
		}
		if len(result) > 0 {
			results[param.Code] = result[0]
		} else {
			results[param.Code] = nil
		}
	}
	return results, nil
}
