# gradebook
Package gradebook implements simple functions of a real student gradebook.

The package allows you to add a grade to a specific gradebook and calculate the average of a student's gradebook. There are also functions to provide a the complete list of grades and final grade of a specific subject.

Example usage:

```
Janez := Student{"Janez", "Novak", []int{5,6,7}}
studenti := map[string]Student{"1": Janez}

dodajOceno(studenti, "1", 10);
fmt.Printf("%.2f\n",povprecje(studenti, "1"))
izpisRedovalnice(studenti)
izpisiKoncniUspeh(studenti)
```


