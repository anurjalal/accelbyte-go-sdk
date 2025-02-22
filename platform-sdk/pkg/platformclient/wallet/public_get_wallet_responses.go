// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// PublicGetWalletReader is a Reader for the PublicGetWallet structure.
type PublicGetWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetWalletOK creates a PublicGetWalletOK with default headers values
func NewPublicGetWalletOK() *PublicGetWalletOK {
	return &PublicGetWalletOK{}
}

/*PublicGetWalletOK handles this case with default header values.

  successful operation
*/
type PublicGetWalletOK struct {
	Payload *platformclientmodels.WalletInfo
}

func (o *PublicGetWalletOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}][%d] publicGetWalletOK  %+v", 200, o.Payload)
}

func (o *PublicGetWalletOK) GetPayload() *platformclientmodels.WalletInfo {
	return o.Payload
}

func (o *PublicGetWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.WalletInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
