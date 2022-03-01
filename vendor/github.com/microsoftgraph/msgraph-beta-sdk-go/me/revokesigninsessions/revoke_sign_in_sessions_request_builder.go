package revokesigninsessions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// RevokeSignInSessionsRequestBuilder builds and executes requests for operations under \me\microsoft.graph.revokeSignInSessions
type RevokeSignInSessionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RevokeSignInSessionsRequestBuilderPostOptions options for Post
type RevokeSignInSessionsRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRevokeSignInSessionsRequestBuilderInternal instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewRevokeSignInSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeSignInSessionsRequestBuilder) {
    m := &RevokeSignInSessionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.revokeSignInSessions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRevokeSignInSessionsRequestBuilder instantiates a new RevokeSignInSessionsRequestBuilder and sets the default values.
func NewRevokeSignInSessionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RevokeSignInSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRevokeSignInSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action revokeSignInSessions
func (m *RevokeSignInSessionsRequestBuilder) CreatePostRequestInformation(options *RevokeSignInSessionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action revokeSignInSessions
func (m *RevokeSignInSessionsRequestBuilder) Post(options *RevokeSignInSessionsRequestBuilderPostOptions)(*bool, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "bool", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*bool), nil
}
