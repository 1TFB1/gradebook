//
//# gradebook
//Package gradebook implements simple functions of a real student gradebook.
//
//The package allows you to add a grade to a specific gradebook and calculate the average of a student's gradebook. There are also functions to provide a the complete list of grades and final grade of a specific subject.
//
//Example usagee:
//
//Janez := Student{"Janez", "Novak", []int{5,6,7}}
//studenti := map[string]Student{"1": Janez}
//
//dodajOceno(studenti, "1", 10);
//fmt.Printf("%.2f\n",povprecje(studenti, "1"))
//izpisRedovalnice(studenti)
//izpisiKoncniUspeh(studenti)

package grade

import "fmt"

//Student represents a single student in a gradebook
type Student struct {
    ime     string
    priimek string
    ocene   []int
}

//DodajOceno add a new grade into the gradebook for a specific student
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int){
	if ocena >=0 && ocena <=10 {
		stud, status := studenti[vpisnaStevilka]
		if status {
			stud.ocene = append(stud.ocene, ocena)
			studenti[vpisnaStevilka] = stud
		} else { fmt.Println("Studenta ni na seznamu") }
	} else { fmt.Println("Ocena ni pravilna")}
}

//Povprecje calculates the average of all the grades in the gradebook for a specific student
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

//IzpisRedovalnica prints all of the grades from the gradebook for a specific student
func IzpisRedovalnice(studenti map[string]Student){
	fmt.Println("REDOVALNICA:")
	for vpisnaStevilka, stud := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, stud.ime, stud.priimek, stud.ocene)
	}

}

//IzpisiKoncniUspeh print the final grade from all the grades of a specific student
func IzpisiKoncniUspeh(studenti map[string]Student){
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
}

