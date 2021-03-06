/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// GroupVersion contains the \"group/version\" and \"version\" string of a version. It is made a struct to keep extensibility.
type V1GroupVersionForDiscovery struct {

	// groupVersion specifies the API group and version in the form \"group/version\"
	GroupVersion string `json:"groupVersion"`

	// version specifies the version in the form of \"version\". This is to save the clients the trouble of splitting the GroupVersion.
	Version string `json:"version"`
}
