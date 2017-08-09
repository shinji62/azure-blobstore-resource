package keyvault

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ActionType enumerates the values for action type.
type ActionType string

const (
	// AutoRenew specifies the auto renew state for action type.
	AutoRenew ActionType = "AutoRenew"
	// EmailContacts specifies the email contacts state for action type.
	EmailContacts ActionType = "EmailContacts"
)

// JSONWebKeyEncryptionAlgorithm enumerates the values for json web key
// encryption algorithm.
type JSONWebKeyEncryptionAlgorithm string

const (
	// RSA15 specifies the rsa15 state for json web key encryption algorithm.
	RSA15 JSONWebKeyEncryptionAlgorithm = "RSA1_5"
	// RSAOAEP specifies the rsaoaep state for json web key encryption
	// algorithm.
	RSAOAEP JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
)

// JSONWebKeyOperation enumerates the values for json web key operation.
type JSONWebKeyOperation string

const (
	// Decrypt specifies the decrypt state for json web key operation.
	Decrypt JSONWebKeyOperation = "decrypt"
	// Encrypt specifies the encrypt state for json web key operation.
	Encrypt JSONWebKeyOperation = "encrypt"
	// Sign specifies the sign state for json web key operation.
	Sign JSONWebKeyOperation = "sign"
	// UnwrapKey specifies the unwrap key state for json web key operation.
	UnwrapKey JSONWebKeyOperation = "unwrapKey"
	// Verify specifies the verify state for json web key operation.
	Verify JSONWebKeyOperation = "verify"
	// WrapKey specifies the wrap key state for json web key operation.
	WrapKey JSONWebKeyOperation = "wrapKey"
)

// JSONWebKeySignatureAlgorithm enumerates the values for json web key
// signature algorithm.
type JSONWebKeySignatureAlgorithm string

const (
	// RS256 specifies the rs256 state for json web key signature algorithm.
	RS256 JSONWebKeySignatureAlgorithm = "RS256"
	// RS384 specifies the rs384 state for json web key signature algorithm.
	RS384 JSONWebKeySignatureAlgorithm = "RS384"
	// RS512 specifies the rs512 state for json web key signature algorithm.
	RS512 JSONWebKeySignatureAlgorithm = "RS512"
	// RSNULL specifies the rsnull state for json web key signature algorithm.
	RSNULL JSONWebKeySignatureAlgorithm = "RSNULL"
)

// JSONWebKeyType enumerates the values for json web key type.
type JSONWebKeyType string

const (
	// EC specifies the ec state for json web key type.
	EC JSONWebKeyType = "EC"
	// Oct specifies the oct state for json web key type.
	Oct JSONWebKeyType = "oct"
	// RSA specifies the rsa state for json web key type.
	RSA JSONWebKeyType = "RSA"
	// RSAHSM specifies the rsahsm state for json web key type.
	RSAHSM JSONWebKeyType = "RSA-HSM"
)

// KeyUsageType enumerates the values for key usage type.
type KeyUsageType string

const (
	// CRLSign specifies the crl sign state for key usage type.
	CRLSign KeyUsageType = "cRLSign"
	// DataEncipherment specifies the data encipherment state for key usage
	// type.
	DataEncipherment KeyUsageType = "dataEncipherment"
	// DecipherOnly specifies the decipher only state for key usage type.
	DecipherOnly KeyUsageType = "decipherOnly"
	// DigitalSignature specifies the digital signature state for key usage
	// type.
	DigitalSignature KeyUsageType = "digitalSignature"
	// EncipherOnly specifies the encipher only state for key usage type.
	EncipherOnly KeyUsageType = "encipherOnly"
	// KeyAgreement specifies the key agreement state for key usage type.
	KeyAgreement KeyUsageType = "keyAgreement"
	// KeyCertSign specifies the key cert sign state for key usage type.
	KeyCertSign KeyUsageType = "keyCertSign"
	// KeyEncipherment specifies the key encipherment state for key usage type.
	KeyEncipherment KeyUsageType = "keyEncipherment"
	// NonRepudiation specifies the non repudiation state for key usage type.
	NonRepudiation KeyUsageType = "nonRepudiation"
)

// Action is the action that will be executed.
type Action struct {
	ActionType ActionType `json:"action_type,omitempty"`
}

// AdministratorDetails is details of the organization administrator of the
// certificate issuer.
type AdministratorDetails struct {
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	EmailAddress *string `json:"email,omitempty"`
	Phone        *string `json:"phone,omitempty"`
}

// Attributes is the object attributes managed by the KeyVault service.
type Attributes struct {
	Enabled   *bool          `json:"enabled,omitempty"`
	NotBefore *date.UnixTime `json:"nbf,omitempty"`
	Expires   *date.UnixTime `json:"exp,omitempty"`
	Created   *date.UnixTime `json:"created,omitempty"`
	Updated   *date.UnixTime `json:"updated,omitempty"`
}

// BackupKeyResult is the backup key result, containing the backup blob.
type BackupKeyResult struct {
	autorest.Response `json:"-"`
	Value             *string `json:"value,omitempty"`
}

// CertificateAttributes is the certificate management attributes.
type CertificateAttributes struct {
	Enabled   *bool          `json:"enabled,omitempty"`
	NotBefore *date.UnixTime `json:"nbf,omitempty"`
	Expires   *date.UnixTime `json:"exp,omitempty"`
	Created   *date.UnixTime `json:"created,omitempty"`
	Updated   *date.UnixTime `json:"updated,omitempty"`
}

// CertificateBundle is a certificate bundle consists of a certificate (X509)
// plus its attributes.
type CertificateBundle struct {
	autorest.Response `json:"-"`
	ID                *string                `json:"id,omitempty"`
	Kid               *string                `json:"kid,omitempty"`
	Sid               *string                `json:"sid,omitempty"`
	X509Thumbprint    *string                `json:"x5t,omitempty"`
	Policy            *CertificatePolicy     `json:"policy,omitempty"`
	Cer               *[]byte                `json:"cer,omitempty"`
	ContentType       *string                `json:"contentType,omitempty"`
	Attributes        *CertificateAttributes `json:"attributes,omitempty"`
	Tags              *map[string]*string    `json:"tags,omitempty"`
}

// CertificateCreateParameters is the certificate create parameters.
type CertificateCreateParameters struct {
	CertificatePolicy     *CertificatePolicy     `json:"policy,omitempty"`
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`
	Tags                  *map[string]*string    `json:"tags,omitempty"`
}

// CertificateImportParameters is the certificate import parameters.
type CertificateImportParameters struct {
	Base64EncodedCertificate *string                `json:"value,omitempty"`
	Password                 *string                `json:"pwd,omitempty"`
	CertificatePolicy        *CertificatePolicy     `json:"policy,omitempty"`
	CertificateAttributes    *CertificateAttributes `json:"attributes,omitempty"`
	Tags                     *map[string]*string    `json:"tags,omitempty"`
}

// CertificateIssuerItem is the certificate issuer item containing certificate
// issuer metadata.
type CertificateIssuerItem struct {
	ID       *string `json:"id,omitempty"`
	Provider *string `json:"provider,omitempty"`
}

// CertificateIssuerListResult is the certificate issuer list result.
type CertificateIssuerListResult struct {
	autorest.Response `json:"-"`
	Value             *[]CertificateIssuerItem `json:"value,omitempty"`
	NextLink          *string                  `json:"nextLink,omitempty"`
}

// CertificateIssuerListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client CertificateIssuerListResult) CertificateIssuerListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// CertificateIssuerSetParameters is the certificate issuer set parameters.
type CertificateIssuerSetParameters struct {
	Provider            *string              `json:"provider,omitempty"`
	Credentials         *IssuerCredentials   `json:"credentials,omitempty"`
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`
	Attributes          *IssuerAttributes    `json:"attributes,omitempty"`
}

// CertificateIssuerUpdateParameters is the certificate issuer update
// parameters.
type CertificateIssuerUpdateParameters struct {
	Provider            *string              `json:"provider,omitempty"`
	Credentials         *IssuerCredentials   `json:"credentials,omitempty"`
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`
	Attributes          *IssuerAttributes    `json:"attributes,omitempty"`
}

// CertificateItem is the certificate item containing certificate metadata.
type CertificateItem struct {
	ID             *string                `json:"id,omitempty"`
	Attributes     *CertificateAttributes `json:"attributes,omitempty"`
	Tags           *map[string]*string    `json:"tags,omitempty"`
	X509Thumbprint *string                `json:"x5t,omitempty"`
}

// CertificateListResult is the certificate list result.
type CertificateListResult struct {
	autorest.Response `json:"-"`
	Value             *[]CertificateItem `json:"value,omitempty"`
	NextLink          *string            `json:"nextLink,omitempty"`
}

// CertificateListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client CertificateListResult) CertificateListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// CertificateMergeParameters is the certificate merge parameters
type CertificateMergeParameters struct {
	X509Certificates      *[][]byte              `json:"x5c,omitempty"`
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`
	Tags                  *map[string]*string    `json:"tags,omitempty"`
}

// CertificateOperation is a certificate operation is returned in case of
// asynchronous requests.
type CertificateOperation struct {
	autorest.Response     `json:"-"`
	ID                    *string           `json:"id,omitempty"`
	IssuerParameters      *IssuerParameters `json:"issuer,omitempty"`
	Csr                   *[]byte           `json:"csr,omitempty"`
	CancellationRequested *bool             `json:"cancellation_requested,omitempty"`
	Status                *string           `json:"status,omitempty"`
	StatusDetails         *string           `json:"status_details,omitempty"`
	Error                 *Error            `json:"error,omitempty"`
	Target                *string           `json:"target,omitempty"`
	RequestID             *string           `json:"request_id,omitempty"`
}

// CertificateOperationUpdateParameter is the certificate operation update
// parameters.
type CertificateOperationUpdateParameter struct {
	CancellationRequested *bool `json:"cancellation_requested,omitempty"`
}

// CertificatePolicy is management policy for a certificate.
type CertificatePolicy struct {
	autorest.Response         `json:"-"`
	ID                        *string                    `json:"id,omitempty"`
	KeyProperties             *KeyProperties             `json:"key_props,omitempty"`
	SecretProperties          *SecretProperties          `json:"secret_props,omitempty"`
	X509CertificateProperties *X509CertificateProperties `json:"x509_props,omitempty"`
	LifetimeActions           *[]LifetimeAction          `json:"lifetime_actions,omitempty"`
	IssuerParameters          *IssuerParameters          `json:"issuer,omitempty"`
	Attributes                *CertificateAttributes     `json:"attributes,omitempty"`
}

// CertificateUpdateParameters is the certificate update parameters.
type CertificateUpdateParameters struct {
	CertificatePolicy     *CertificatePolicy     `json:"policy,omitempty"`
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`
	Tags                  *map[string]*string    `json:"tags,omitempty"`
}

// Contact is the contact information for the vault certificates.
type Contact struct {
	EmailAddress *string `json:"email,omitempty"`
	Name         *string `json:"name,omitempty"`
	Phone        *string `json:"phone,omitempty"`
}

// Contacts is the contacts for the vault certificates.
type Contacts struct {
	autorest.Response `json:"-"`
	ID                *string    `json:"id,omitempty"`
	ContactList       *[]Contact `json:"contacts,omitempty"`
}

// Error is the key vault server error.
type Error struct {
	Code       *string `json:"code,omitempty"`
	Message    *string `json:"message,omitempty"`
	InnerError *Error  `json:"innererror,omitempty"`
}

// ErrorType is the key vault error exception.
type ErrorType struct {
	Error *Error `json:"error,omitempty"`
}

// IssuerAttributes is the attributes of an issuer managed by the Key Vault
// service.
type IssuerAttributes struct {
	Enabled *bool          `json:"enabled,omitempty"`
	Created *date.UnixTime `json:"created,omitempty"`
	Updated *date.UnixTime `json:"updated,omitempty"`
}

// IssuerBundle is the issuer for Key Vault certificate.
type IssuerBundle struct {
	autorest.Response   `json:"-"`
	ID                  *string              `json:"id,omitempty"`
	Provider            *string              `json:"provider,omitempty"`
	Credentials         *IssuerCredentials   `json:"credentials,omitempty"`
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`
	Attributes          *IssuerAttributes    `json:"attributes,omitempty"`
}

// IssuerCredentials is the credentials to be used for the certificate issuer.
type IssuerCredentials struct {
	AccountID *string `json:"account_id,omitempty"`
	Password  *string `json:"pwd,omitempty"`
}

// IssuerParameters is parameters for the issuer of the X509 component of a
// certificate.
type IssuerParameters struct {
	Name            *string `json:"name,omitempty"`
	CertificateType *string `json:"cty,omitempty"`
}

// JSONWebKey is as of
// http://tools.ietf.org/html/draft-ietf-jose-json-web-key-18
type JSONWebKey struct {
	Kid    *string        `json:"kid,omitempty"`
	Kty    JSONWebKeyType `json:"kty,omitempty"`
	KeyOps *[]string      `json:"key_ops,omitempty"`
	N      *string        `json:"n,omitempty"`
	E      *string        `json:"e,omitempty"`
	D      *string        `json:"d,omitempty"`
	DP     *string        `json:"dp,omitempty"`
	DQ     *string        `json:"dq,omitempty"`
	QI     *string        `json:"qi,omitempty"`
	P      *string        `json:"p,omitempty"`
	Q      *string        `json:"q,omitempty"`
	K      *string        `json:"k,omitempty"`
	T      *string        `json:"key_hsm,omitempty"`
}

// KeyAttributes is the attributes of a key managed by the key vault service.
type KeyAttributes struct {
	Enabled   *bool          `json:"enabled,omitempty"`
	NotBefore *date.UnixTime `json:"nbf,omitempty"`
	Expires   *date.UnixTime `json:"exp,omitempty"`
	Created   *date.UnixTime `json:"created,omitempty"`
	Updated   *date.UnixTime `json:"updated,omitempty"`
}

// KeyBundle is a KeyBundle consisting of a WebKey plus its attributes.
type KeyBundle struct {
	autorest.Response `json:"-"`
	Key               *JSONWebKey         `json:"key,omitempty"`
	Attributes        *KeyAttributes      `json:"attributes,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Managed           *bool               `json:"managed,omitempty"`
}

// KeyCreateParameters is the key create parameters.
type KeyCreateParameters struct {
	Kty           JSONWebKeyType         `json:"kty,omitempty"`
	KeySize       *int32                 `json:"key_size,omitempty"`
	KeyOps        *[]JSONWebKeyOperation `json:"key_ops,omitempty"`
	KeyAttributes *KeyAttributes         `json:"attributes,omitempty"`
	Tags          *map[string]*string    `json:"tags,omitempty"`
}

// KeyImportParameters is the key import parameters.
type KeyImportParameters struct {
	Hsm           *bool               `json:"Hsm,omitempty"`
	Key           *JSONWebKey         `json:"key,omitempty"`
	KeyAttributes *KeyAttributes      `json:"attributes,omitempty"`
	Tags          *map[string]*string `json:"tags,omitempty"`
}

// KeyItem is the key item containing key metadata.
type KeyItem struct {
	Kid        *string             `json:"kid,omitempty"`
	Attributes *KeyAttributes      `json:"attributes,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Managed    *bool               `json:"managed,omitempty"`
}

// KeyListResult is the key list result.
type KeyListResult struct {
	autorest.Response `json:"-"`
	Value             *[]KeyItem `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// KeyListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client KeyListResult) KeyListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// KeyOperationResult is the key operation result.
type KeyOperationResult struct {
	autorest.Response `json:"-"`
	Kid               *string `json:"kid,omitempty"`
	Result            *string `json:"value,omitempty"`
}

// KeyOperationsParameters is the key operations parameters.
type KeyOperationsParameters struct {
	Algorithm JSONWebKeyEncryptionAlgorithm `json:"alg,omitempty"`
	Value     *string                       `json:"value,omitempty"`
}

// KeyProperties is properties of the key pair backing a certificate.
type KeyProperties struct {
	Exportable *bool   `json:"exportable,omitempty"`
	KeyType    *string `json:"kty,omitempty"`
	KeySize    *int32  `json:"key_size,omitempty"`
	ReuseKey   *bool   `json:"reuse_key,omitempty"`
}

// KeyRestoreParameters is the key restore parameters.
type KeyRestoreParameters struct {
	KeyBundleBackup *string `json:"value,omitempty"`
}

// KeySignParameters is the key operations parameters.
type KeySignParameters struct {
	Algorithm JSONWebKeySignatureAlgorithm `json:"alg,omitempty"`
	Value     *string                      `json:"value,omitempty"`
}

// KeyUpdateParameters is the key update parameters.
type KeyUpdateParameters struct {
	KeyOps        *[]JSONWebKeyOperation `json:"key_ops,omitempty"`
	KeyAttributes *KeyAttributes         `json:"attributes,omitempty"`
	Tags          *map[string]*string    `json:"tags,omitempty"`
}

// KeyVerifyParameters is the key verify parameters.
type KeyVerifyParameters struct {
	Algorithm JSONWebKeySignatureAlgorithm `json:"alg,omitempty"`
	Digest    *string                      `json:"digest,omitempty"`
	Signature *string                      `json:"value,omitempty"`
}

// KeyVerifyResult is the key verify result.
type KeyVerifyResult struct {
	autorest.Response `json:"-"`
	Value             *bool `json:"value,omitempty"`
}

// LifetimeAction is action and its trigger that will be performed by Key Vault
// over the lifetime of a certificate.
type LifetimeAction struct {
	Trigger *Trigger `json:"trigger,omitempty"`
	Action  *Action  `json:"action,omitempty"`
}

// OrganizationDetails is details of the organization of the certificate
// issuer.
type OrganizationDetails struct {
	ID           *string                 `json:"id,omitempty"`
	AdminDetails *[]AdministratorDetails `json:"admin_details,omitempty"`
}

// PendingCertificateSigningRequestResult is the pending certificate signing
// request result.
type PendingCertificateSigningRequestResult struct {
	Value *string `json:"value,omitempty"`
}

// SecretAttributes is the secret management attributes.
type SecretAttributes struct {
	Enabled   *bool          `json:"enabled,omitempty"`
	NotBefore *date.UnixTime `json:"nbf,omitempty"`
	Expires   *date.UnixTime `json:"exp,omitempty"`
	Created   *date.UnixTime `json:"created,omitempty"`
	Updated   *date.UnixTime `json:"updated,omitempty"`
}

// SecretBundle is a secret consisting of a value, id and its attributes.
type SecretBundle struct {
	autorest.Response `json:"-"`
	Value             *string             `json:"value,omitempty"`
	ID                *string             `json:"id,omitempty"`
	ContentType       *string             `json:"contentType,omitempty"`
	Attributes        *SecretAttributes   `json:"attributes,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Kid               *string             `json:"kid,omitempty"`
	Managed           *bool               `json:"managed,omitempty"`
}

// SecretItem is the secret item containing secret metadata.
type SecretItem struct {
	ID          *string             `json:"id,omitempty"`
	Attributes  *SecretAttributes   `json:"attributes,omitempty"`
	Tags        *map[string]*string `json:"tags,omitempty"`
	ContentType *string             `json:"contentType,omitempty"`
	Managed     *bool               `json:"managed,omitempty"`
}

// SecretListResult is the secret list result.
type SecretListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SecretItem `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// SecretListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SecretListResult) SecretListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SecretProperties is properties of the key backing a certificate.
type SecretProperties struct {
	ContentType *string `json:"contentType,omitempty"`
}

// SecretSetParameters is the secret set parameters.
type SecretSetParameters struct {
	Value            *string             `json:"value,omitempty"`
	Tags             *map[string]*string `json:"tags,omitempty"`
	ContentType      *string             `json:"contentType,omitempty"`
	SecretAttributes *SecretAttributes   `json:"attributes,omitempty"`
}

// SecretUpdateParameters is the secret update parameters.
type SecretUpdateParameters struct {
	ContentType      *string             `json:"contentType,omitempty"`
	SecretAttributes *SecretAttributes   `json:"attributes,omitempty"`
	Tags             *map[string]*string `json:"tags,omitempty"`
}

// SubjectAlternativeNames is the subject alternate names of a X509 object.
type SubjectAlternativeNames struct {
	Emails   *[]string `json:"emails,omitempty"`
	DNSNames *[]string `json:"dns_names,omitempty"`
	Upns     *[]string `json:"upns,omitempty"`
}

// Trigger is a condition to be satisfied for an action to be executed.
type Trigger struct {
	LifetimePercentage *int32 `json:"lifetime_percentage,omitempty"`
	DaysBeforeExpiry   *int32 `json:"days_before_expiry,omitempty"`
}

// X509CertificateProperties is properties of the X509 component of a
// certificate.
type X509CertificateProperties struct {
	Subject                 *string                  `json:"subject,omitempty"`
	Ekus                    *[]string                `json:"ekus,omitempty"`
	SubjectAlternativeNames *SubjectAlternativeNames `json:"sans,omitempty"`
	KeyUsage                *[]KeyUsageType          `json:"key_usage,omitempty"`
	ValidityInMonths        *int32                   `json:"validity_months,omitempty"`
}
