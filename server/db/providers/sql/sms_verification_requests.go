package sql

import (
	"context"
	"time"

	"github.com/authorizerdev/authorizer/server/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// UpsertSMSRequest adds/updates SMS verification request
func (p *provider) UpsertSMSRequest(ctx context.Context, smsRequest *models.SMSVerificationRequest) (*models.SMSVerificationRequest, error) {
	if smsRequest.ID == "" {
		smsRequest.ID = uuid.New().String()
	}
	smsRequest.CreatedAt = time.Now().Unix()
	smsRequest.UpdatedAt = time.Now().Unix()
	res := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "phone_number"}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "code_expires_at", "updated_at"}),
	}).Create(smsRequest)
	if res.Error != nil {
		return nil, res.Error
	}
	return smsRequest, nil
}

// GetCodeByPhone to get code for a given phone number
func (p *provider) GetCodeByPhone(ctx context.Context, phoneNumber string) (*models.SMSVerificationRequest, error) {
	var sms_verification_request models.SMSVerificationRequest

	result := p.db.Where("phone_number = ?", phoneNumber).First(&sms_verification_request)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sms_verification_request, nil
}

// DeleteSMSRequest to delete SMS verification request
func (p *provider) DeleteSMSRequest(ctx context.Context, smsRequest *models.SMSVerificationRequest) error {
	result := p.db.Delete(&models.SMSVerificationRequest{
		ID: smsRequest.ID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
