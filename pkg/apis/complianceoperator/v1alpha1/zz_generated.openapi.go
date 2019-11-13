// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScan":       schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScan(ref),
		"github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanSpec":   schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScanSpec(ref),
		"github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanStatus": schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScanStatus(ref),
	}
}

func schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScan(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComplianceScan is the Schema for the compliancescans API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanSpec", "github.com/openshift/compliance-operator/pkg/apis/complianceoperator/v1alpha1.ComplianceScanStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScanSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComplianceScanSpec defines the desired state of ComplianceScan",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_complianceoperator_v1alpha1_ComplianceScanStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ComplianceScanStatus defines the observed state of ComplianceScan",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
