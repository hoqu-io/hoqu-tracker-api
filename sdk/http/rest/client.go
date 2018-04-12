package rest

import (
    "net/url"
    "net/http"
    "bytes"
    "encoding/json"
    "hoqu-tracker-api/sdk/models"
    "fmt"
    "github.com/sirupsen/logrus"
    "crypto/tls"
    "time"
    "math/rand"
    "strconv"
    "hoqu-tracker-api/sdk/libs"
)

const AUTH_HEADER_DEFAULT = "Authorization"

type Client struct {
    httpClient *http.Client
    BaseURL   *url.URL
    Project string
    AuthHeader string
    Auth string
}

func NewClient(baseURL *url.URL, secure bool, project string, authHeader, auth string) *Client {
    httpClient := http.DefaultClient
    httpClient.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: !secure,
        },
    }

    if authHeader == "" {
        authHeader = AUTH_HEADER_DEFAULT
    }

    return &Client {
        httpClient: httpClient,
        BaseURL: baseURL,
        Project: project,
        AuthHeader: authHeader,
        Auth: auth,
    }
}

func (c *Client) DoRest(method, path string, requestModel, responseModel interface{}) (*models.RestRequest, error) {
    respModel := &models.RestResponse{
        Data: responseModel,
    }

    reqData, err := c.DoJson(method, path, requestModel, respModel)
    if err != nil {
        return nil, err
    }

    if respModel.Error.Code != "NO_ERROR" {
        return nil, fmt.Errorf("%s: %s", respModel.Error.Code, respModel.Error.Message)
    }

    return reqData, nil
}

func (c *Client) DoJson(method, path string, requestModel, responseModel interface{}) (*models.RestRequest, error) {
    reqData := &models.RestRequest{
        ReqId: c.generateRequestId(),
        RawRequest: new(bytes.Buffer),
    }

    rel, err := url.ParseRequestURI(path)
    if err != nil {
        return nil, fmt.Errorf("[%s] cannot parse url: %s", reqData.ReqId, err.Error())
    }

    u := c.BaseURL.ResolveReference(rel)

    if requestModel != nil {
        err := json.NewEncoder(reqData.RawRequest).Encode(requestModel)
        if err != nil {
            return nil, fmt.Errorf("[%s] cannot marshal request model: %s", reqData.ReqId, err.Error())
        }
    }

    req, err := http.NewRequest(method, u.String(), reqData.RawRequest)
    if err != nil {
        return nil, fmt.Errorf("[%s] cannot create http request: %s", reqData.ReqId, err.Error())
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.Project + " REST client")

    logrus.Debugf("[%s] Request -- %s %s", reqData.ReqId, method, u.String())
    if requestModel != nil {
        reqJson, _ := json.Marshal(requestModel)
        logrus.Debugf("[%s] Body: %s", reqData.ReqId, string(reqJson))

        req.Header.Set("Content-Type", "application/json")
    }

    if c.Auth != "" {
        req.Header.Set(c.AuthHeader, c.Auth)
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("[%s] http request error: %s", reqData.ReqId, err.Error())
    }
    defer resp.Body.Close()

    reqData.StatusCode = resp.StatusCode
    reqData.RawResponse = libs.StreamToBuffer(resp.Body)
    logrus.Debugf("[%s] Response %d -- Body: %s", reqData.ReqId, reqData.StatusCode, reqData.RawResponse.String())

    err = json.NewDecoder(bytes.NewReader(reqData.RawResponse.Bytes())).Decode(responseModel)
    if err != nil {
        return nil, fmt.Errorf("[%s] cannot unmarshal response to model: %s", reqData.ReqId, err.Error())
    }

    return reqData, nil
}

func (c *Client) generateRequestId() string {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    return fmt.Sprintf("%d", int32(time.Now().Unix())) +"_" + strconv.Itoa(r1.Intn(100))
}
