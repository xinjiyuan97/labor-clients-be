package constants

type BrandAuthStatus string

const (
	BrandAuthStatusPending  BrandAuthStatus = "pending"
	BrandAuthStatusApproved BrandAuthStatus = "approved"
	BrandAuthStatusRejected BrandAuthStatus = "rejected"
)

type BrandAccountStatus string

const (
	BrandAccountStatusActive   BrandAccountStatus = "active"
	BrandAccountStatusDisabled BrandAccountStatus = "disabled"
	BrandAccountStatusFrozen   BrandAccountStatus = "frozen"
)
