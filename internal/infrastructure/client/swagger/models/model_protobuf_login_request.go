/*
 * internal/infrastructure/protobuf/contracts.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ProtobufLoginRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
