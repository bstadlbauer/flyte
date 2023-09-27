/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Refers to a the workflow the node is to execute.
type CoreWorkflowNode struct {
	// A globally unique identifier for the launch plan.
	LaunchplanRef *CoreIdentifier `json:"launchplan_ref,omitempty"`
	SubWorkflowRef *CoreIdentifier `json:"sub_workflow_ref,omitempty"`
}