package rocket

type Rocket interface {
	Launch()
}

func Launch(r Rocket) {
	r.Launch()
}
