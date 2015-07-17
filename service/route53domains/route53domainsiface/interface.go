// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53domainsiface provides an interface for the Amazon Route 53 Domains.
package route53domainsiface

import (
	"github.com/aws/aws-sdk-go/service/route53domains"
)

// Route53DomainsAPI is the interface type for route53domains.Route53Domains.
type Route53DomainsAPI interface {
	CheckDomainAvailability(*route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)

	DeleteTagsForDomain(*route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)

	DisableDomainAutoRenew(*route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)

	DisableDomainTransferLock(*route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)

	EnableDomainAutoRenew(*route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)

	EnableDomainTransferLock(*route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)

	GetDomainDetail(*route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)

	GetOperationDetail(*route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)

	ListDomains(*route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)

	ListOperations(*route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)

	ListTagsForDomain(*route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)

	RegisterDomain(*route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)

	RetrieveDomainAuthCode(*route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)

	TransferDomain(*route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)

	UpdateDomainContact(*route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)

	UpdateDomainContactPrivacy(*route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)

	UpdateDomainNameservers(*route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)

	UpdateTagsForDomain(*route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
}
