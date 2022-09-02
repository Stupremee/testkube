/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// CRD based executor data
type Event struct {
	// UUID of event
	Id string `json:"id"`
	// DEPERECATED event uri remove after rewrite @TODO
	Uri       string     `json:"uri,omitempty"`
	Type_     *EventType `json:"type"`
	Execution *Execution `json:"execution,omitempty"`
}