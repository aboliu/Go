package phone

type Phone interface {
	Owner() string
	ContactPerson() string
	BlackList() string
}

func CallOut(p Phone) {
	println(p.Owner(), "要打電話！撥打電話給", p.ContactPerson())
}

func CallIn(p Phone) {
	println(p.Owner(), ":嘟嘟嘟～有電話！是", p.BlackList(), "打來的！！快掛掉！")
}
