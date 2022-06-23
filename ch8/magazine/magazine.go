package magazine

type Park struct {
	Desc  string
	Count int
}

func ApplyCount2(p *Park) {
	p.Count = 400
}

type Employee struct {
	Name string
	Age  int
	Address
}

type Address struct {
	City string
}
