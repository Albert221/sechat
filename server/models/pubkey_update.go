package models

type OtherPublicKeyUpdate struct {
	PublicKeyUpdate []byte `json:"otherClientPublicKey"`
}

func NewOtherPublicKeyUpdate(publicKey []byte) OtherPublicKeyUpdate {
	return OtherPublicKeyUpdate{
		PublicKeyUpdate: publicKey,
	}
}

func (p *OtherPublicKeyUpdate) UpdateStruct() map[string]interface{} {
	return map[string]interface{}{
		"type":    TypeOtherPublicKey,
		"payload": p,
	}
}
