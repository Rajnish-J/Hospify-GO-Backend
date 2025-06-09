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
		ProviderID:      i.ProviderID,
		PolicyNumber:    i.PolicyNumber,
		CoverageDetails: i.CoverageDetail,
	}
}

func FromInsuranceModel(m *models.Insurance) *InsuranceDTO {
	return &InsuranceDTO{
		ID:             m.InsuranceID,
		ProviderID:     m.ProviderID,
		PolicyNumber:   m.PolicyNumber,
		CoverageDetail: m.CoverageDetails,
	}
}
