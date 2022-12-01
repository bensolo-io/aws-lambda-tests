package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

/*
---
apiVersion: v1
kind: Namespace
metadata:
  name:  ben
---
apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  name: aws-lambda
  namespace: ben
spec:
  aws:
    secretRef:
      name: aws-creds
      namespace: gloo-system
    lambdaFunctions:
    - lambdaFunctionName: golang-apigw-latest
      logicalName: golang-apigw-latest
    region: us-east-2
    roleArn: arn:aws:iam::931713665590:role/gloo-lambda-executor
---
apiVersion: v1
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIekNDQWdlZ0F3SUJBZ0lVVitvdCtyMWg4MVJ2d1FVNE0rd0NEVXlNdm9Nd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0h6RWRNQnNHQTFVRUF3d1VjR1YwYzNSdmNtVXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXhNREkyTVRnMApPVE0yV2hjTk1qTXhNREkyTVRnME9UTTJXakFmTVIwd0d3WURWUVFEREJSd1pYUnpkRzl5WlM1bGVHRnRjR3hsCkxtTnZiVENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFMVW5TUjh2WW5PSW1DejkKaG1UZE5LcXA3Y0RnYTE2VXFVYUxWVDhEQlFIU2ZVd1BkaStqbXE2R0xJMzE4NkxzWE9XbHZITVVoSlM2VGk1VApSZ2lBSk1oVWdiMFFMYndkejEvcTRnSjNVNGJDOGxaRFVtclJkd0xuSng5bkJZUmQ5U0dIQUNpODE2dEtRK2c1CmdFcnIybWg1ay8vME1sQjh0MnF2STlRZFBaSGxlMTBvb0NaSzhyc3RHQlZJVCt5aU81M2RNa09sVXF6dk9NU3MKU1VHd2ZpRUlSZXc2RWtwdFRWQWtkQWZpbW5xYjVZM1M2K1JpWXlCN0piekVYVFlBREM2aGRvUXVOL2J1bHFwVwpGWmppZld5K05heGR0SzJYZ25RTDV3bmY0aHpQV0VhVnRtNjBTUmRhUjkySWFBTlh3dDdENXRxSFprQVQ5SExIClJzUFRzVGtDQXdFQUFhTlRNRkV3SFFZRFZSME9CQllFRkhyaWk3K0lHUlZjdE42elBrZVpBdUNEMnhoOE1COEcKQTFVZEl3UVlNQmFBRkhyaWk3K0lHUlZjdE42elBrZVpBdUNEMnhoOE1BOEdBMVVkRXdFQi93UUZNQU1CQWY4dwpEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSXY0WmliUjRmVFV5UG5wNGV0blA5dmtXVTBiSk9IUmNVZlhtUHIvCkJPNHNIcWFKZTNQcWJHZDBDVFJidC94Zi8reWVYdG9hWjdTVlhZV25hSnhJSVNkbmVHd1lJVjZpVnpRYmhEZWMKSzhpeXVNaHgxWDJERUxJQ0l4SWZINWJVKzR3bEVYYzZJUjhPOGc0a0JJNkFUcXlCS29PVlB6TEFxenBieVA3bQp5U2hMMjdzNjcrWnRJdi9aVUJTdmQ1ZVhYMGRBWVVDNTJXQU5ZcTNEa3Y3aklraStrQTB3L3BBWGt1MHl4ZFpzCmNKZkFMRm1BaXExc1BvcW5BWWovRDFGSGM3WFNjMGNLUElPMUR3Ykl4Rlpsb0FtdHdZUloxQ1kyeVhJZEtsOUoKL2pqU0JEUytCTFBRVktmU3I2eDNmWU5BVHpKYVU4eUtkTnBKNmtMZklmNE11QnM9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  tls.key: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQzFKMGtmTDJKemlKZ3MKL1laazNUU3FxZTNBNEd0ZWxLbEdpMVUvQXdVQjBuMU1EM1l2bzVxdWhpeU45Zk9pN0Z6bHBieHpGSVNVdWs0dQpVMFlJZ0NUSVZJRzlFQzI4SGM5ZjZ1SUNkMU9Hd3ZKV1ExSnEwWGNDNXljZlp3V0VYZlVoaHdBb3ZOZXJTa1BvCk9ZQks2OXBvZVpQLzlESlFmTGRxcnlQVUhUMlI1WHRkS0tBbVN2SzdMUmdWU0Uvc29qdWQzVEpEcFZLczd6akUKckVsQnNINGhDRVhzT2hKS2JVMVFKSFFINHBwNm0rV04wdXZrWW1NZ2V5Vzh4RjAyQUF3dW9YYUVMamYyN3BhcQpWaFdZNG4xc3ZqV3NYYlN0bDRKMEMrY0ozK0ljejFoR2xiWnV0RWtYV2tmZGlHZ0RWOExldytiYWgyWkFFL1J5CngwYkQwN0U1QWdNQkFBRUNnZ0VBUTg5VTI4dU0rdHBpdkZTYWZoOGZtOUxJSUs3aEFpSjd5dUJWSktVb3Rxbm8KSlJTVWxQaVU2a2RuWGl4MjZzRzNkRzg1djRvcXB0R21la2pKcWx6dFo5L2FRTDlSbjEwNVJ4cFJSOE1MRFNMawpPODR2aDdTbzYvbEM5OHBFa212cFdvZUNVNzE4cmEyN0JhNmdpMnNGOHAydi9OdVlDZkRsWjVYdnYzTENuVjI4Ck1ucXF1N2tVQ0RkYnVTM25RNi9xc1p1bEJXcFFDT0ZWVktCdnBGUkZWMWdMc2htbEFleE5RN0RSV050QUMxMlcKVUJheVZPTmdQSlh4V2prY0J4V3p2OWxJK2hZQlU1Tm1JU1l3Z05WNUw0Y2hPaEE2cEk2ODZ1RmNmTU5tcVMrWApMUWxqWTQ2TmpPUjcwbzFaMlV0T0RybjBGTnNJZ0NxSHhib24zejNRMFFLQmdRRGJ3enNZUmpqNml5VGJGcWVBCmpjZjkwbE9SczVod280UkxiSDZuUnVReWgvWWdqbk0rVWR6TG9XOWVVc1JZWHJVVTNSQVVVRWhpWFJJdUpXTncKaDYrYjJXa1h4UXhKMXZuZHFacXEyS0FKVmNwNCs1M01GNzNzZExhcEU1dy9mc2EvZkE1WUFRbUN5Wi9QYzhTWQowMVgvR0o5eTdlMVErWXdSeUFTMnExV2ZDd0tCZ1FEVEJrSTJIdVNaWVNCZGpCMVI2OGV1OEttOW1vak1URVhyCjBrRGZhaitiL3ZKNWFUM1hrTEsrTlNKNmFYbi9GalV3SXJSZjA2ZnBlTlpQTkcwQmpMSEFDdk9TV0hXV0NMTEwKa25LdnJSOFQwdE9yK3o0clBQZlZRcTRJV0FZU21uWktHaCt0eE5nNExoRXpsM1dmbmRFRWZtT0pXWmtSTVhZbApYZm82clozclN3S0JnUUM3a2tJaE1PYTNLZ1pXSFZyd2haTTZXTWZOWjQzb0xoamZ5NFc2dnU1TkZ2RUR6ckljCmNnRFRxVUdHTDN5NHRIVTRqb3FIM0JJOEtwWTIzdUNtRHBuYm10QnhZbFZmdk9aZHhNSm5xaWZHYi93MkVRVVoKU3ZabkdTTkM0cU1OS3ViMlR5dHEvOCtmV3ZwVk5jbUthMjlPSVRVUEFuYjVFMVh6WTFacWw0aW9Dd0tCZ0hHRgpsR2o1QlpGZHBzT3NkTGwxVmQ3T3FRSE8rSGl2TDQ1RmRaQzYzNjFUNGExZTZGM25BY0ZCWkdMbUN6TW5CMFgxCjVZTUhvZlQvaElybmNSeThTNE04WVB3QmlvQkQvYXQyQlN4c3ZhTTBiNXE5ZGh4Y21CYXA0R1dzdE5lZE1MVjgKaUQ0Ni92WjZFZGJuUytlcVJwOWNQci9NNjROTVVIcVpxOXVWT3JjeEFvR0FGRnVQWUhlYjBKenRFMXdTSVE5VgpKSGtnQVhTSGFHam5ZNUlEbmxSeWVNQWhJcDVlMUpjKy9ZZGRPeVN6YktsMml2NFE1MjErSEtFUHhGb1JKU2MyCkljc1liSERUVkJjRzJkZUZuZzdkcmtlZXBJZUVkRmpKTTUyMEl2M0dxdUxQVmNnL3JUTzhocmRoUm13VkgrdkcKalViSFRSSVEzeFhpY3dHTzI3dC90QVk9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
kind: Secret
metadata:
  name: upstream-tls
  namespace: gloo-system
type: kubernetes.io/tls
---
apiVersion: gateway.solo.io/v1
kind: VirtualService
metadata:
  name: ben
  namespace: gloo-system
spec:
  sslConfig:
    secretRef:
      name: upstream-tls
      namespace: gloo-system
  virtualHost:
    domains:
    - '*'
    routes:
    - matchers:
      - prefix: /
      options:
        timeout: 35s  # default value is 15s
        retries:
          retryOn: '5xx'
          numRetries: 3
          perTryTimeout: '10s'
      routeAction:
        single:
          destinationSpec:
            aws:
              logicalName: golang-apigw-latest
              wrapAsApiGateway: true
			  unwrapAsApiGateway: true
          upstream:
            name: aws-lambda
            namespace: ben

*/

func handleLambdaEvent(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// this function expects events.APIGatewayProxyRequest from aws golang SDK
	// you would use this function signature if you want this lambda code to be compatiable with
	//
	// AWS API Gateway - you do not need to use this signature to integrate with Gloo; however, if you want
	// the code to remain compatible with AWS API Gateway in the future then you can use this event type
	// and set wrapAsApiGateway on the route and Gloo will convert the http request to a events.APIGatewayProxyRequest
	spew.Dump(event)

	// this function returns events.APIGatewayProxyRequest from aws golang SDK
	// you would use this response type if you want this lambda code to be compatiable with
	//
	// AWS API Gateway - you do not need to use this signature to integrate with Gloo; however, if you want
	// the code to remain compatible with AWS API Gateway in the future then you can
	return events.APIGatewayProxyResponse{
		StatusCode:        401,
		Body:              `{"message": "bad request!"}`,
		Headers:           map[string]string{},
		MultiValueHeaders: make(map[string][]string),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(handleLambdaEvent)
}
