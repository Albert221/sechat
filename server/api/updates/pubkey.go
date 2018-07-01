package updates

import "encoding/base64"

type OtherPublicKeyUpdate struct {
	publicKeyUpdate []byte
}

func NewOtherPublicKeyUpdate(publicKey []byte) OtherPublicKeyUpdate {
	return OtherPublicKeyUpdate{
		publicKeyUpdate: publicKey,
	}
}

func (p *OtherPublicKeyUpdate) UpdateStruct() map[string]interface{} {
	return map[string]interface{}{
		"type":    TypeOtherPublicKey,
		"payload": base64.StdEncoding.EncodeToString(p.publicKeyUpdate),
	}
}
