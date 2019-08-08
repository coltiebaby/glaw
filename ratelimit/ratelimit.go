package ratelimit

type Limiter interface {
	Activate()
	Deactivate()

	Add(int, Limit)
	Take(int)
}

type Limit interface {
	Recharge()
	Take()
}
