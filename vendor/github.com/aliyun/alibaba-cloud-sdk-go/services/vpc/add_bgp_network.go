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

// AddBgpNetwork invokes the vpc.AddBgpNetwork API synchronously
// api document: https://help.aliyun.com/api/vpc/addbgpnetwork.html
func (client *Client) AddBgpNetwork(request *AddBgpNetworkRequest) (response *AddBgpNetworkResponse, err error) {
	response = CreateAddBgpNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// AddBgpNetworkWithChan invokes the vpc.AddBgpNetwork API asynchronously
// api document: https://help.aliyun.com/api/vpc/addbgpnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBgpNetworkWithChan(request *AddBgpNetworkRequest) (<-chan *AddBgpNetworkResponse, <-chan error) {
	responseChan := make(chan *AddBgpNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBgpNetwork(request)
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

// AddBgpNetworkWithCallback invokes the vpc.AddBgpNetwork API asynchronously
// api document: https://help.aliyun.com/api/vpc/addbgpnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBgpNetworkWithCallback(request *AddBgpNetworkRequest, callback func(response *AddBgpNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBgpNetworkResponse
		var err error
		defer close(result)
		response, err = client.AddBgpNetwork(request)
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

// AddBgpNetworkRequest is the request struct for api AddBgpNetwork
type AddBgpNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterId             string           `position:"Query" name:"RouterId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	DstCidrBlock         string           `position:"Query" name:"DstCidrBlock"`
}

// AddBgpNetworkResponse is the response struct for api AddBgpNetwork
type AddBgpNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddBgpNetworkRequest creates a request to invoke AddBgpNetwork API
func CreateAddBgpNetworkRequest() (request *AddBgpNetworkRequest) {
	request = &AddBgpNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AddBgpNetwork", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddBgpNetworkResponse creates a response to parse from AddBgpNetwork response
func CreateAddBgpNetworkResponse() (response *AddBgpNetworkResponse) {
	response = &AddBgpNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}