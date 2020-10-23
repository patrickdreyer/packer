/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.2
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// KeypairCreated Information about the created keypair.
type KeypairCreated struct {
	// If you create a keypair, the SHA-1 digest of the DER encoded private key.<br /> If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.
	KeypairFingerprint string `json:"KeypairFingerprint,omitempty"`
	// The name of the keypair.
	KeypairName string `json:"KeypairName,omitempty"`
	// The private key.
	PrivateKey string `json:"PrivateKey,omitempty"`
}