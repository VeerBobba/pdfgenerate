package main

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	err := m.OutputFileAndClose("pdfs/veeratest.pdf")
	if err != nil {
		fmt.Println("Error while saving the PDF")
	}
}
