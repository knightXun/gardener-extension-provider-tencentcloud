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

// DisableVpcClassicLink invokes the vpc.DisableVpcClassicLink API synchronously
// api document: https://help.aliyun.com/api/vpc/disablevpcclassiclink.html
func (client *Client) DisableVpcClassicLink(request *DisableVpcClassicLinkRequest) (response *DisableVpcClassicLinkResponse, err error) {
	response = CreateDisableVpcClassicLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DisableVpcClassicLinkWithChan invokes the vpc.DisableVpcClassicLink API asynchronously
// api document: https://help.aliyun.com/api/vpc/disablevpcclassiclink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableVpcClassicLinkWithChan(request *DisableVpcClassicLinkRequest) (<-chan *DisableVpcClassicLinkResponse, <-chan error) {
	responseChan := make(chan *DisableVpcClassicLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableVpcClassicLink(request)
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

// DisableVpcClassicLinkWithCallback invokes the vpc.DisableVpcClassicLink API asynchronously
// api document: https://help.aliyun.com/api/vpc/disablevpcclassiclink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableVpcClassicLinkWithCallback(request *DisableVpcClassicLinkRequest, callback func(response *DisableVpcClassicLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableVpcClassicLinkResponse
		var err error
		defer close(result)
		response, err = client.DisableVpcClassicLink(request)
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

// DisableVpcClassicLinkRequest is the request struct for api DisableVpcClassicLink
type DisableVpcClassicLinkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// DisableVpcClassicLinkResponse is the response struct for api DisableVpcClassicLink
type DisableVpcClassicLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableVpcClassicLinkRequest creates a request to invoke DisableVpcClassicLink API
func CreateDisableVpcClassicLinkRequest() (request *DisableVpcClassicLinkRequest) {
	request = &DisableVpcClassicLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DisableVpcClassicLink", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableVpcClassicLinkResponse creates a response to parse from DisableVpcClassicLink response
func CreateDisableVpcClassicLinkResponse() (response *DisableVpcClassicLinkResponse) {
	response = &DisableVpcClassicLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
