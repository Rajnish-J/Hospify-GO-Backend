package dto

import "hospify/src/models"

type InsuranceDTO struct {
	ID             uint   `json:"id"`
	ProviderID     uint   `json:"provider_id"`
	PolicyNumber   string `json:"policy_number"`
	CoverageDetail string `json:"coverage_detail"`
}

func (i *InsuranceDTO) ToModel() *models.Insurance {
	return &models.Insurance{
		InsuranceID:     i.ID,
		ProviderID:      &i.ProviderID,
		PolicyNumber:    i.PolicyNumber,
		CoverageDetails: i.CoverageDetail,
	}
}

func FromInsuranceModel(m *models.Insurance) *InsuranceDTO {
	var providerID uint
	if m.ProviderID != nil {
		providerID = *m.ProviderID        // dereference pointer to get uint
	}
	return &InsuranceDTO{
		ID:             m.InsuranceID,
<<<<<<< HEAD
		ProviderID:     providerID,
=======
		ProviderID:     *m.ProviderID,
>>>>>>> cd4e49e0dc061abf54be13ed20e0a52daf7ac933
		PolicyNumber:   m.PolicyNumber,
		CoverageDetail: m.CoverageDetails,
	}
}
