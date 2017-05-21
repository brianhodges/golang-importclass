package location

type Country struct {
    Name string
    States []string
    Population string
	Best_city City
}

type City struct {
	Name string
	Population string
}