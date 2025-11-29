package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/1TFB1/gradebook/grade"
)

func main() {
	cmd := &cli.Command{
		Name:  "gradebook",
		Usage: "Gradebook izpisuje in ocenjuje trenutne ocene študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "minimalna možna ocena za dodajanje",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "maksimalna možna ocena za dodajanje",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			min := cmd.Int("minOcena")
			max := cmd.Int("maxOcena")
			st := cmd.Int("stOcen")
			return run(ctx, st, min, max)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context, stOcen int, minOcena int, maxOcena int) error {
	// Primer podatkov
	janez := grade.Student{Ime: "Janez", Priimek: "Novak", Ocene: []int{5, 6, 7}}
	mojca := grade.Student{Ime: "Mojca", Priimek: "Novak", Ocene: []int{10, 10, 10, 10, 10, 10}}
	karmen := grade.Student{Ime: "Karmen", Priimek: "Pekarmen", Ocene: []int{7, 8, 9, 10, 9, 8, 7}}

	studenti := map[string]grade.Student{
		"1": janez,
		"2": mojca,
		"3": karmen,
	}

	// demonstracija funkcionalnosti
	grade.DodajOceno(studenti, "1", 10, minOcena, maxOcena) // OK
	grade.DodajOceno(studenti, "2", 11, minOcena, maxOcena) // Ocena ni pravilna
	grade.DodajOceno(studenti, "4", 7, minOcena, maxOcena)  // Študenta ni na seznamu

	grade.IzpisVsehOcen(studenti)
	grade.IzpisiKoncniUspeh(studenti, stOcen)

	return nil
}
