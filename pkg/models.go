package models

import (
		"encoding/xml"
		 "time"
	)

// User is a simple user struct.
type User struct {
	Usr string
	Pss string
	Sigla string
	Ip string
}

// Msg is a simple message struct
type Msg struct {
	Title string
	Desc  string
}

//DRC2 Person by Email
type PersonByEmail struct {
	Content  int64  `json:"content"`
	HasValue bool `json:"hasValue"`
	Errors   []struct {
		MemberNames  []string `json:"memberNames"`
		ErrorMessage string   `json:"errorMessage"`
	} `json:"errors"`
	HasException bool `json:"hasException"`
	Exception    struct {
	} `json:"exception"`
	ExceptionMessage      string    `json:"exceptionMessage"`
	InnerExceptionMessage string    `json:"innerExceptionMessage"`
	ExecutionStart        time.Time `json:"executionStart"`
	ExecutionEnd          time.Time `json:"executionEnd"`
	ExecutionTime         string    `json:"executionTime"`
	StatusCode            string    `json:"statusCode"`
	HasErrors             bool      `json:"hasErrors"`
	Expires               time.Time `json:"expires"`
	IsCached              bool      `json:"isCached"`
}

//DRC2 User by PersonId
type UserByPerson struct {
	Content struct {
		ID                          int64       `json:"id"`
		LdapCustomerConfigurationID int64       `json:"ldapCustomerConfigurationID"`
		PersonID                    int64       `json:"personID"`
		AccessTypeID                int64       `json:"accessTypeID"`
		CustomerID                  int64       `json:"customerID"`
		AuthenticationTypeID        int64       `json:"authenticationTypeID"`
		Login                       string    `json:"login"`
		Password                    string    `json:"password"`
		PasswordExpiration          time.Time `json:"passwordExpiration"`
		Enabled                     int64       `json:"enabled"`
		ValidateLogin               int64       `json:"validateLogin"`
		SecurityQuestion            string    `json:"securityQuestion"`
		SecurityAnswer              string    `json:"securityAnswer"`
		AuditUser                   string    `json:"auditUser"`
		AuditUpdateDate             time.Time `json:"auditUpdateDate"`
		Sha                         string    `json:"sha"`
		LdapLogin                   string    `json:"ldapLogin"`
		PasswordAttempts            int64       `json:"passwordAttempts"`
		UseIMConnect                int64       `json:"useIMConnect"`
	} `json:"content"`
	HasValue bool `json:"hasValue"`
	Errors   []struct {
		MemberNames  []string `json:"memberNames"`
		ErrorMessage string   `json:"errorMessage"`
	} `json:"errors"`
	HasException bool `json:"hasException"`
	Exception    struct {
	} `json:"exception"`
	ExceptionMessage      string    `json:"exceptionMessage"`
	InnerExceptionMessage string    `json:"innerExceptionMessage"`
	ExecutionStart        time.Time `json:"executionStart"`
	ExecutionEnd          time.Time `json:"executionEnd"`
	ExecutionTime         string    `json:"executionTime"`
	StatusCode            string    `json:"statusCode"`
	HasErrors             bool      `json:"hasErrors"`
	Expires               time.Time `json:"expires"`
	IsCached              bool      `json:"isCached"`
}

type UserModel struct {
	ID                          int64       `json:"id"`
	LdapCustomerConfigurationID int64       `json:"ldapCustomerConfigurationID"`
	PersonID                    int64       `json:"personID"`
	AccessTypeID                int64       `json:"accessTypeID"`
	CustomerID                  int64       `json:"customerID"`
	AuthenticationTypeID        int64       `json:"authenticationTypeID"`
	Login                       string    `json:"login"`
	Password                    string    `json:"password"`
	PasswordExpiration          time.Time `json:"passwordExpiration"`
	Enabled                     int64       `json:"enabled"`
	ValidateLogin               int64       `json:"validateLogin"`
	SecurityQuestion            string    `json:"securityQuestion"`
	SecurityAnswer              string    `json:"securityAnswer"`
	AuditUser                   string    `json:"auditUser"`
	AuditUpdateDate             time.Time `json:"auditUpdateDate"`
	Sha                         string    `json:"sha"`
	LdapLogin                   string    `json:"ldapLogin"`
	PasswordAttempts            int64       `json:"passwordAttempts"`
	UseIMConnect                int64       `json:"useIMConnect"`
}

/*
// XMLFragment container structure.
type XMLFragment struct {
	XMLName   xml.Name  `xml:"xml-fragment"`
	Candidate Candidate `xml:"candidate"`
}

// Candidate candidate resources to be created.
type Candidate struct {
	ContainerID string `xml:"container-id,attr"`
	Name        string `xml:"name,attr"`
	Label       string `xml:"label,attr"`
	LdapAlias   string `xml:"ldap-alias,attr"`
	LdapDn      string `xml:"ldap-dn,attr"`
}

// AddGroup contains the group guid that will be added.
type AddGroup struct {
	GUID string `json:"@guid"`
}

// RemoveGroup contains the guid of the group that will be removed.
type RemoveGroup struct {
	GUID string `json:"@guid"`
}

// Resource structure that contains the resource ID and the actions that will be performed.
type Resource struct {
	GUID        string        `json:"@guid"`
	AddGroup    []AddGroup    `json:"add-group,omitempty"`
	RemoveGroup []RemoveGroup `json:"remove-group,omitempty"`
}

// UpdateResource structure that contains the resource and org model identification.
type UpdateResource struct {
	ModelVersion string     `json:"@model-version,omitempty"`
	Resource     []Resource `json:"resource"`
}

// RootUpdateResource main structure, contains the UpdateResource structure.
type RootUpdateResource struct {
	UpdateResource UpdateResource `json:"updateResource"`
}

// RootEntity main structure, contains the Entity structure.
type RootEntity struct {
	Entity Entity `json:"entity"`
}

// Entity represents the return of the action of creating a resource.
type Entity struct {
	AlreadyPresent string `json:"@already-present"`
	GUID           string `json:"@guid"`
	Name           string `json:"@name"`
	Label          string `json:"@label"`
	ContainerID    string `json:"@container-id"`
	LdapAlias      string `json:"@ldap-alias"`
	LdapDn         string `json:"@ldap-dn"`
}
*/
