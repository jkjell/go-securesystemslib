package dsse

import (
	"encoding/base64"
)

/*
Envelope captures an envelope as described by the DSSE specification. See here:
https://github.com/secure-systems-lab/dsse/blob/master/envelope.md
*/
type Envelope struct {
	PayloadType string      `json:"payloadType"`
	Payload     string      `json:"payload"`
	Signatures  []Signature `json:"signatures"`
}

/*
DecodeB64Payload returns the serialized body, decoded from the envelope's
payload field. A flexible decoder is used, first trying standard base64, then
URL-encoded base64.
*/
func (e *Envelope) DecodeB64Payload() ([]byte, error) {
	return b64Decode(e.Payload)
}

/*
Signature represents a generic in-toto signature that contains the identifier
of the key which was used to create the signature.
The used signature scheme has to be agreed upon by the signer and verifer
out of band.
The signature is a base64 encoding of the raw bytes from the signature
algorithm.
*/
type Signature struct {
	KeyID string `json:"keyid"`
	Sig   string `json:"sig"`
}

/*
Both standard and url encoding are allowed:
https://github.com/secure-systems-lab/dsse/blob/master/envelope.md
*/
func b64Decode(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		b, err = base64.URLEncoding.DecodeString(s)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
