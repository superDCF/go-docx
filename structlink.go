package docxlib

import (
	"encoding/xml"
	"io"
)

// Hyperlink element contains links
type Hyperlink struct {
	XMLName xml.Name `xml:"w:hyperlink,omitempty"`
	ID      string   `xml:"r:id,attr"`
	Run     Run
}

func (r *Hyperlink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tt := t.(type) {
		case xml.StartElement:
			if tt.Name.Local == "r" {
				d.DecodeElement(&r.Run, &tt)
			} else {
				continue
			}
		}

	}
	return nil

}