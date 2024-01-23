package types

type Log struct {
	ProxyProtocolAddr      string `json:"proxy_protocol_addr"`
	RemoteAddr             string `json:"remote_addr"`
	RemoteUser             string `json:"remote_user"`
	TimeLocal              string `json:"time_local"`
	Request                string `json:"request"`
	Status                 string `json:"status"`
	BodyBytesSent          string `json:"body_bytes_sent"`
	HTTPReferer            string `json:"http_referer"`
	HTTPUserAgent          string `json:"http_user_agent"`
	RequestTime            string `json:"request_time"`
	RequestLength          string `json:"request_length"`
	ProxyUpstreamName      string `json:"proxy_upstream_name"`
	UpstreamAddr           string `json:"upstream_addr"`
	UpstreamStatus         string `json:"upstream_status"`
	UpStreamResponseTime   string `json:"upstream_response_time"`
	UpStreamResponseLength string `json:"upstream_connect_time"`
}
