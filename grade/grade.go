/*
Program zakasnjeno-izvajanje prikazuje uporabo stavka defer v jeziku go
*/
package main

import "fmt"

type Student struct {
    ime     string
    priimek string
    ocene   []int
}

func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int){
	if ocena >=0 && ocena <=10 {
		stud, status := studenti[vpisnaStevilka]
		if status {
			stud.ocene = append(stud.ocene, ocena)
			studenti[vpisnaStevilka] = stud
		} else { fmt.Println("Studenta ni na seznamu") }
	} else { fmt.Println("Ocena ni pravilna")}
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64{
	stud, status := studenti[vpisnaStevilka]
	if status {
		var povprecje float64 = 0.0
		var i float64 = 0
		for _, ocena := range stud.ocene {
			povprecje += float64(ocena)
			i++
		}
		if i >= 6 {return (povprecje/i)
		} else {return 0.0}
	}
	return -1.0;
}

func izpisRedovalnice(studenti map[string]Student){
	fmt.Println("REDOVALNICA:")
	for vpisnaStevilka, stud := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, stud.ime, stud.priimek, stud.ocene)
	}

}

func izpisiKoncniUspeh(studenti map[string]Student){
	fmt.Println("KONCNE OCENE:")
	for vpisnaStevilka, stud := range studenti {
		var povprecje float64 = povprecje(studenti,vpisnaStevilka)
		var s string = ""
		if povprecje >= 9.0 { s = "Odličen študent!"
		} else if povprecje >= 6.0 && povprecje < 9.0{ s = "Povprečen študent"
		} else { s = "Neuspešen študent"}
		fmt.Printf("%s %s: povprečna ocena %.2f -> %s\n", stud.ime, stud.priimek, povprecje, s)
	}

}

func main() {

	Janez := Student{"Janez", "Novak", []int{5,6,7}}
	Mojca := Student{"Mojca", "Novak", []int{10,10,10,10,10,10}}
	studenti := map[string]Student{"1": Janez, "2": Mojca, "3": Student{"Karmen", "Pekarmen", []int{7,8,9,10,9,8,7}}}

	fmt.Println(studenti)
	fmt.Println("\n")
	/*
	 map[1:{Janez Novak [5 6 7]} 2:{Mojca Novak [10 10 10 10 10 10]} 3:{Karmen Pekarmen [7 8 9 10 9 8 7]}]
	*/

	dodajOceno(studenti, "1", 10);
	fmt.Println(studenti) // map[1:{Janez Novak [5 6 7 10]} 2:{Mojca Novak [10 10 10 10 10 10]} 3:{Karmen Pekarmen [7 8 9 10 9 8 7]}]
	fmt.Println("\n")
	dodajOceno(studenti, "2", 11); // Ocena ni pravilna
	dodajOceno(studenti, "4", 7); // Studenta ni na seznamu

	fmt.Printf("%.2f\n",povprecje(studenti, "1")) //0.00
	fmt.Printf("%.2f\n",povprecje(studenti, "2")) //10.00
	fmt.Printf("%.2f\n",povprecje(studenti, "3")) //8.29
	fmt.Printf("%.2f\n",povprecje(studenti, "4")) //-1.00

	izpisRedovalnice(studenti)
	/*
	REDOVALNICA:
		1 - Janez Novak: [5 6 7 10]
		2 - Mojca Novak: [10 10 10 10 10 10]
		3 - Karmen Pekarmen: [7 8 9 10 9 8 7]
	*/

	izpisiKoncniUspeh(studenti)
	/*
	KONCNE OCENE:
		Janez Novak: povprečna ocena 0.00 -> Neuspešen študent
		Mojca Novak: povprečna ocena 10.00 -> Odličen študent!
		Karmen Pekarmen: povprečna ocena 8.29 -> Povprečen študent
	*/


}

