package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSslVpnClientCert invokes the vpc.DescribeSslVpnClientCert API synchronously
// api document: https://help.aliyun.com/api/vpc/describesslvpnclientcert.html
func (client *Client) DescribeSslVpnClientCert(request *DescribeSslVpnClientCertRequest) (response *DescribeSslVpnClientCertResponse, err error) {
	response = CreateDescribeSslVpnClientCertResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSslVpnClientCertWithChan invokes the vpc.DescribeSslVpnClientCert API asynchronously
// api document: https://help.aliyun.com/api/vpc/describesslvpnclientcert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSslVpnClientCertWithChan(request *DescribeSslVpnClientCertRequest) (<-chan *DescribeSslVpnClientCertResponse, <-chan error) {
	responseChan := make(chan *DescribeSslVpnClientCertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSslVpnClientCert(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSslVpnClientCertWithCallback invokes the vpc.DescribeSslVpnClientCert API asynchronously
// api document: https://help.aliyun.com/api/vpc/describesslvpnclientcert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSslVpnClientCertWithCallback(request *DescribeSslVpnClientCertRequest, callback func(response *DescribeSslVpnClientCertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSslVpnClientCertResponse
		var err error
		defer close(result)
		response, err = client.DescribeSslVpnClientCert(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSslVpnClientCertRequest is the request struct for api DescribeSslVpnClientCert
type DescribeSslVpnClientCertRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SslVpnClientCertId   string           `position:"Query" name:"SslVpnClientCertId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSslVpnClientCertResponse is the response struct for api DescribeSslVpnClientCert
type DescribeSslVpnClientCertResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	RegionId           string `json:"RegionId" xml:"RegionId"`
	SslVpnClientCertId string `json:"SslVpnClientCertId" xml:"SslVpnClientCertId"`
	Name               string `json:"Name" xml:"Name"`
	SslVpnServerId     string `json:"SslVpnServerId" xml:"SslVpnServerId"`
	CaCert             string `json:"CaCert" xml:"CaCert"`
	ClientCert         string `json:"ClientCert" xml:"ClientCert"`
	ClientKey          string `json:"ClientKey" xml:"ClientKey"`
	ClientConfig       string `json:"ClientConfig" xml:"ClientConfig"`
	CreateTime         int64  `json:"CreateTime" xml:"CreateTime"`
	EndTime            int64  `json:"EndTime" xml:"EndTime"`
	Status             string `json:"Status" xml:"Status"`
}

// CreateDescribeSslVpnClientCertRequest creates a request to invoke DescribeSslVpnClientCert API
func CreateDescribeSslVpnClientCertRequest() (request *DescribeSslVpnClientCertRequest) {
	request = &DescribeSslVpnClientCertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSslVpnClientCert", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSslVpnClientCertResponse creates a response to parse from DescribeSslVpnClientCert response
func CreateDescribeSslVpnClientCertResponse() (response *DescribeSslVpnClientCertResponse) {
	response = &DescribeSslVpnClientCertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
