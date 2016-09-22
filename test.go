package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func init() {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		Size: gofpdf.SizeType{100, 100},
	})
	pdf.SetCellMargin(8)
	pdf.AddPage()

	info := pdf.RegisterImage("work/fairy-tail-7919635.jpg", "")
	pdf.ImageOptions("work/fairy-tail-7919635.jpg", 0, 0, info.Height(), info.Width(), false, gofpdf.ImageOptions{
		ImageType: "JPG",
	}, 0, "")

	err := pdf.OutputFileAndClose("test.pdf")

	if err != nil {
		log.Fatalln(err)
	}
}
