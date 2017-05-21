package main
import (
    "strconv"
    "fmt"
    "strings"
    "bufio"
    "os"
    "./location"
)

const nl = "\n"

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    countries := []location.Country{
        {
            Name: "USA",
            States: []string{"New York", "Virginia", "California", "Florida", "Hawaii"},
            Population: "350 million",
			Best_city: location.City{
				Name: "Los Angeles",
				Population: "4 million",
			},
        },
		{
			Name: "Mexico",
            States: []string{"Durango", "Puebla", "Guerrero", "Quintana Roo"},
            Population: "127 million",
			Best_city: location.City{
				Name: "Cancun",
				Population: "700,000",
			},
		},
    }
	
	//Write to .txt File
    f, err := os.Create("./log.txt")
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    
    _, err = fmt.Fprintf(w, nl)
    check(err)
	for _, c := range countries {
		_, err = fmt.Fprintf(w, "Name: " + c.Name + nl)
        check(err)
		_, err = fmt.Fprintf(w, "Population: " + c.Population + nl)
        check(err)
		_, err = fmt.Fprintf(w, "Best States: " + strings.Join(c.States, ", ") + nl)
        check(err)
		_, err = fmt.Fprintf(w, "Top State Count: " + strconv.Itoa(len(c.States)) + nl)
        check(err)
		_, err = fmt.Fprintf(w, "Best City: " + c.Best_city.Name + nl)
        check(err)
		_, err = fmt.Fprintf(w, c.Best_city.Name + " Population: " + c.Best_city.Population + nl)
        check(err)
		_, err = fmt.Fprintf(w, nl)
        check(err)
	}
    w.Flush()
}