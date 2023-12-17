package kvstore

import "time"

type GetOption func(o *GetOptionContext)
type SetOption func(o *SetOptionContext)
type IncrOption func(o *IncrOptionContext)

type GetOptionContext struct {
	TTL  time.Duration
	ExAt time.Time
	Del  bool
}

type SetOptionContext struct {
	TTL     time.Duration
	ExAt    time.Time
	Mode    string
	Get     bool
	KeepTTL bool
}

type IncrOptionContext struct {
	TTL     time.Duration
	ExAt    time.Time
	KeepTTL bool
}

func WithGetTTL(ttl time.Duration) GetOption {
	return func(o *GetOptionContext) {
		o.TTL = ttl
		o.ExAt = time.Time{}
		o.Del = false
	}
}

func WithGetExpireAt(eat time.Time) GetOption {
	return func(o *GetOptionContext) {
		o.ExAt = eat
		o.TTL = 0
		o.Del = false
	}
}

func WithDel() GetOption {
	return func(o *GetOptionContext) {
		o.Del = true
		o.TTL = 0
		o.ExAt = time.Time{}
	}
}

func WithSetOnlyIfNotExists() SetOption {
	return func(o *SetOptionContext) {
		o.Mode = "NX"
	}
}

func WithSetOnlyIfAlreadyExists() SetOption {
	return func(o *SetOptionContext) {
		o.Mode = "XX"
	}
}

func WithRetrieveDisabled() SetOption {
	return func(o *SetOptionContext) {
		o.Get = false
	}
}

func WithSetTTL(ttl time.Duration) SetOption {
	return func(o *SetOptionContext) {
		o.TTL = ttl
		o.ExAt = time.Time{}
		o.KeepTTL = false
	}
}

func WithSetExpireAt(eat time.Time) SetOption {
	return func(o *SetOptionContext) {
		o.ExAt = eat
		o.TTL = 0
		o.KeepTTL = false
	}
}

func WithIncrTTL(ttl time.Duration, keepTTL bool) IncrOption {
	return func(o *IncrOptionContext) {
		o.TTL = ttl
		o.KeepTTL = keepTTL
		o.ExAt = time.Time{}
	}
}

func WithIncrExpireAt(eat time.Time) IncrOption {
	return func(o *IncrOptionContext) {
		o.ExAt = eat
		o.TTL = 0
		o.KeepTTL = false
	}
}
