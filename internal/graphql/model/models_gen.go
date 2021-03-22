// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddonList struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Endpoint string `json:"endpoint"`
}

type ControlPlane struct {
	Name    string                `json:"name"`
	Members []*ControlPlaneMember `json:"members"`
}

type ControlPlaneFilter struct {
	Type *MeshType `json:"type"`
}

type ControlPlaneMember struct {
	Name      string `json:"name"`
	Component string `json:"component"`
	Version   string `json:"version"`
	Namespace string `json:"namespace"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type NameSpace struct {
	Namespace string `json:"namespace"`
}

type OperatorControllerStatus struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  Status `json:"status"`
	Error   *Error `json:"error"`
}

type OperatorStatus struct {
	Status      Status                      `json:"status"`
	Version     string                      `json:"version"`
	Controllers []*OperatorControllerStatus `json:"controllers"`
	Error       *Error                      `json:"error"`
}

type MeshType string

const (
	MeshTypeIstio       MeshType = "ISTIO"
	MeshTypeLinkerd     MeshType = "LINKERD"
	MeshTypeConsul      MeshType = "CONSUL"
	MeshTypeOctarine    MeshType = "OCTARINE"
	MeshTypeTraefikmesh MeshType = "TRAEFIKMESH"
	MeshTypeOsm         MeshType = "OSM"
	MeshTypeKuma        MeshType = "KUMA"
	MeshTypeNginxsm     MeshType = "NGINXSM"
	MeshTypeNsm         MeshType = "NSM"
	MeshTypeCitrix      MeshType = "CITRIX"
)

var AllMeshType = []MeshType{
	MeshTypeIstio,
	MeshTypeLinkerd,
	MeshTypeConsul,
	MeshTypeOctarine,
	MeshTypeTraefikmesh,
	MeshTypeOsm,
	MeshTypeKuma,
	MeshTypeNginxsm,
	MeshTypeNsm,
	MeshTypeCitrix,
}

func (e MeshType) IsValid() bool {
	switch e {
	case MeshTypeIstio, MeshTypeLinkerd, MeshTypeConsul, MeshTypeOctarine, MeshTypeTraefikmesh, MeshTypeOsm, MeshTypeKuma, MeshTypeNginxsm, MeshTypeNsm, MeshTypeCitrix:
		return true
	}
	return false
}

func (e MeshType) String() string {
	return string(e)
}

func (e *MeshType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshType", str)
	}
	return nil
}

func (e MeshType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusEnabled    Status = "ENABLED"
	StatusDisabled   Status = "DISABLED"
	StatusProcessing Status = "PROCESSING"
	StatusUnknown    Status = "UNKNOWN"
)

var AllStatus = []Status{
	StatusEnabled,
	StatusDisabled,
	StatusProcessing,
	StatusUnknown,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusEnabled, StatusDisabled, StatusProcessing, StatusUnknown:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
