package zlog

type LogEntry struct {
	ID        string `json:"id,omitempty"`
	Level     Level  `json:"level,omitempty"`
	Service   string `json:"service,omitempty"`
	Message   string `json:"message,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
