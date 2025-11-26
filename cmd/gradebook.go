package main

import "fmt"

func main() {

	Janez := Student{"Janez", "Novak", []int{5,6,7}}
	Mojca := Student{"Mojca", "Novak", []int{10,10,10,10,10,10}}
	studenti := map[string]Student{"1": Janez, "2": Mojca, "3": Student{"Karmen", "Pekarmen", []int{7,8,9,10,9,8,7}}}

	fmt.Println(studenti)
	fmt.Println("\n")

	dodajOceno(studenti, "1", 10);
	fmt.Println(studenti) // map[1:{Janez Novak [5 6 7 10]} 2:{Mojca Novak [10 10 10 10 10 10]} 3:{Karmen Pekarmen [7 8 9 10 9 8 7]}]
	fmt.Println("\n")
	dodajOceno(studenti, "2", 11); // Ocena ni pravilna
	dodajOceno(studenti, "4", 7); // Studenta ni na seznamu

  //fmt.Printf("%.2f\n",povprecje(studenti, "1")), skrita funkcija

	izpisRedovalnice(studenti)

	izpisiKoncniUspeh(studenti)
  
}
