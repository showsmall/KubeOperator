package dto

import (
	"github.com/KubeOperator/KubeOperator/pkg/errorf"
	"github.com/KubeOperator/KubeOperator/pkg/model"
)

type Host struct {
	model.Host
	ClusterName string `json:"clusterName"`
	ProjectName string `json:"projectName"`
	ZoneName    string `json:"zoneName"`
}

type HostCreate struct {
	Name         string           `json:"name" validate:"required"`
	Ip           string           `json:"ip" validate:"required"`
	Port         int              `json:"port" validate:"required"`
	CredentialID string           `json:"credentialId"`
	Credential   CredentialCreate `json:"credential"`
}

type HostPage struct {
	Items []Host `json:"items"`
	Total int    `json:"total"`
}

type HostOp struct {
	Operation string `json:"operation" validate:"required"`
	Items     []Host `json:"items" validate:"required"`
}

type HostSync struct {
	HostName   string `json:"hostName"`
	HostStatus string `json:"hostStatus"`
}

type ImportHostResponse struct {
	Errs    errorf.CErrFs `json:"errs"`
	Success bool          `json:"success"`
}
