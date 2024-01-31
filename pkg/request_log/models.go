package request_log

type QueueMessage struct {
	Request             string
	Response            string
	UserId              string
	Version             string
	Endpoint            string
	Domain              string
	Headers             map[string]string
	Method              string
	QueryParams         map[string]string
	HttpMethod          string
	HttpResponse        int
	HttpResponseHeaders map[string]string
	StartedDateTime     string
	Time                int
}
