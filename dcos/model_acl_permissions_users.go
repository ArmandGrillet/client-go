/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type AclPermissionsUsers struct {
	Uid     string   `json:"uid"`
	Userurl string   `json:"userurl"`
	Actions []Action `json:"actions"`
}
