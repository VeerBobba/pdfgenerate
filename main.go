package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	buildHeader(m)
	buildInvoiceHeader(m)
	buildInvoiceTable(m)
	buildFooter(m)

	err := m.OutputFileAndClose("pdfs/veeratest.pdf")
	if err != nil {
		fmt.Println("error while saving the pdf.", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(20, func() {
				err := m.FileImage("images/logobig.jpg", props.Rect{
					Center:  true,
					Percent: 100,
				})
				if err != nil {
					fmt.Println("Image file not loaded")
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(10, func() {
			m.Text("Telscope IT Inc... \n Endless Possibilites...", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getRedColor(),
			})
		})
	})
}

func buildFooter(m pdf.Maroto) {

	m.RegisterFooter(func() {
		m.Row(7, func() {
			m.Col(10, func() {
				m.Text("HR Manager", props.Text{Align: consts.Left})
			})
		})
		m.Row(10, func() {
			m.Col(6, func() {
				m.Signature("Vikram Aditya", props.Font{
					Style: consts.Italic,
				})
			})
			m.Col(6, func() {
				m.Text(time.Now().Format("02-January-2006"), props.Text{Align: consts.Right})
			})
		})
	})
}
func buildInvoiceHeader(m pdf.Maroto) {
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
}

func buildInvoiceTable(m pdf.Maroto) {
	headerList := []string{"SNo", "ClientName", "ResourceName", "Hours", "BillRate", "Total"}
	contents := [][]string{{"1", "HUMANA INC.(101 GROPVE Street, Huson, NY-06074)", "Travis Derik", "70", "50.00", "3500.00"},
		{"2", "CGI INC.(22 JENGA Drive, Mephis, TN-03020)", "ANVENDHS Training", "72", "50.00", "3600.00"},
		{"", "", "", "", "Total", "7100.00"},
	}
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(headerList, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 3, 4, 1, 1, 2},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 3, 4, 1, 1, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 true,
		AlternatedBackground: &color.Color{Red: 200, Green: 230, Blue: 230},
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
	m.Line(10, props.Line{
		Width: 2,
		Color: color.NewBlack(),
		Style: consts.Solid,
	})
}

func getRedColor() color.Color {
	return color.Color{
		Red:   255,
		Blue:  0,
		Green: 0,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Blue:  166,
		Green: 166,
	}
}
