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

// Student predstavlja enega študenta v redovalnici.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno za podanega študenta.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
		fmt.Println("Ocena ni pravilna")
		return
	}

	stud, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študenta ni na seznamu")
		return
	}

	stud.Ocene = append(stud.Ocene, ocena)
	studenti[vpisnaStevilka] = stud
}

// povprecje izračuna povprečje ocen študenta.
// Funkcija je skrita (neizvožena).
func povprecje(stud Student) float64 {
	if len(stud.Ocene) == 0 {
		return 0.0
	}

	sum := 0
	for _, o := range stud.Ocene {
		sum += o
	}
	return float64(sum) / float64(len(stud.Ocene))
}

// IzpisVsehOcen izpiše vse študente in njihove ocene.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for id, stud := range studenti {
		fmt.Printf("%s - %s %s: %v\n", id, stud.Ime, stud.Priimek, stud.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh študentov.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
	fmt.Println("KONČNE OCENE:")

	for _, stud := range studenti {
		if len(stud.Ocene) < stOcen {
			fmt.Printf("%s %s: premalo ocen\n", stud.Ime, stud.Priimek)
			continue
		}

		p := povprecje(stud)

		opis := "Neuspešen študent"
		if p >= 9 {
			opis = "Odličen študent!"
		} else if p >= 6 {
			opis = "Povprečen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.2f -> %s\n",
			stud.Ime, stud.Priimek, p, opis)
	}
}

