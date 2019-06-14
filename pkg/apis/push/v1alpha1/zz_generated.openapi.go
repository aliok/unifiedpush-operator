// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariant":          schema_pkg_apis_push_v1alpha1_AndroidVariant(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantSpec":      schema_pkg_apis_push_v1alpha1_AndroidVariantSpec(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantStatus":    schema_pkg_apis_push_v1alpha1_AndroidVariantStatus(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariant":              schema_pkg_apis_push_v1alpha1_IOSVariant(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantSpec":          schema_pkg_apis_push_v1alpha1_IOSVariantSpec(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantStatus":        schema_pkg_apis_push_v1alpha1_IOSVariantStatus(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplication":         schema_pkg_apis_push_v1alpha1_PushApplication(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationSpec":     schema_pkg_apis_push_v1alpha1_PushApplicationSpec(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationStatus":   schema_pkg_apis_push_v1alpha1_PushApplicationStatus(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServer":       schema_pkg_apis_push_v1alpha1_UnifiedPushServer(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerSpec":   schema_pkg_apis_push_v1alpha1_UnifiedPushServerSpec(ref),
		"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerStatus": schema_pkg_apis_push_v1alpha1_UnifiedPushServerStatus(ref),
	}
}

func schema_pkg_apis_push_v1alpha1_AndroidVariant(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AndroidVariant is the Schema for the androidvariants API",
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
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantSpec", "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.AndroidVariantStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_push_v1alpha1_AndroidVariantSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AndroidVariantSpec defines the desired state of AndroidVariant",
				Properties: map[string]spec.Schema{
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a human friendly description for the variant.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serverKey": {
						SchemaProps: spec.SchemaProps{
							Description: "ServerKey is the key from the Firebase Console of a project which has been enabled for FCM.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"senderId": {
						SchemaProps: spec.SchemaProps{
							Description: "SenderId is the \"Google Project Number\" from the API Console. It is *not* needed for sending push messages, but it is a convenience to \"see\" it on the UnifiedPush Server Admin UI as well, since the Android applications require it (called Sender ID there). That way all information is stored on the same object.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pushApplicationId": {
						SchemaProps: spec.SchemaProps{
							Description: "PushApplicationId is the Id of the Application that this Variant corresponds to in the UnifiedPush Server admin UI.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"serverKey", "senderId", "pushApplicationId"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_AndroidVariantStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AndroidVariantStatus defines the observed state of AndroidVariant",
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"variantId": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"ready"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_IOSVariant(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IOSVariant is the Schema for the iosvariants API",
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
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantSpec", "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.IOSVariantStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_push_v1alpha1_IOSVariantSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IOSVariantSpec defines the desired state of IOSVariant",
				Properties: map[string]spec.Schema{
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a human friendly description for the variant.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"certificate": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate defines the base64 encoded APNs certificate that is needed to establish a connection to Apple's APNs Push Servers.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"passphrase": {
						SchemaProps: spec.SchemaProps{
							Description: "Passphrase defines the APNs passphrase that is needed to establish a connection to any of Apple's APNs Push Servers.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"production": {
						SchemaProps: spec.SchemaProps{
							Description: "Production defines if a connection to production APNS server should be used. If false, a connection to Apple's Sandbox/Development APNs server will be established for this iOS variant.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"pushApplicationId": {
						SchemaProps: spec.SchemaProps{
							Description: "PushApplicationId is the Id of the Application that this Variant corresponds to in the UnifiedPush Server admin UI.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"certificate", "passphrase", "production", "pushApplicationId"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_IOSVariantStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IOSVariantStatus defines the observed state of IOSVariant",
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"variantId": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"ready"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_PushApplication(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PushApplication is the Schema for the pushapplications API",
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
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationSpec", "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.PushApplicationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_push_v1alpha1_PushApplicationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PushApplicationSpec defines the desired state of PushApplication",
				Properties: map[string]spec.Schema{
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a description of the app to be displayed in the UnifiedPush Server admin UI",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"description"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_PushApplicationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PushApplicationStatus defines the observed state of PushApplication",
				Properties: map[string]spec.Schema{
					"pushApplicationId": {
						SchemaProps: spec.SchemaProps{
							Description: "PushApplicationId is an identifer used to register Variants with this PushApplication",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"masterSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "MasterSecret is a master password, used for sending message to this PushApplication, or it's Variant(s)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"variants": {
						SchemaProps: spec.SchemaProps{
							Description: "Variants is a slice of variant (AndroidVariant or IOSVariant, in this package) names associated with this Application",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"pushApplicationId", "masterSecret"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_push_v1alpha1_UnifiedPushServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UnifiedPushServer is the Schema for the unifiedpushservers API",
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
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerSpec", "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_push_v1alpha1_UnifiedPushServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UnifiedPushServerSpec defines the desired state of UnifiedPushServer",
				Properties: map[string]spec.Schema{
					"backups": {
						SchemaProps: spec.SchemaProps{
							Description: "Backups is an array of configs that will be used to create CronJob resource instances",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerBackup"),
									},
								},
							},
						},
					},
					"useMessageBroker": {
						SchemaProps: spec.SchemaProps{
							Description: "Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1.UnifiedPushServerBackup"},
	}
}

func schema_pkg_apis_push_v1alpha1_UnifiedPushServerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UnifiedPushServerStatus defines the observed state of UnifiedPushServer",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{},
	}
}
