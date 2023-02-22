package docxlib

import (
	"encoding/xml"
	"io"
	"strconv"
)

const (
	// A4_EMU_MAX_WIDTH is the max display width of an A4 paper
	A4_EMU_MAX_WIDTH = 5274310
)

const (
	XMLNS_DRAWINGML_MAIN    = `http://schemas.openxmlformats.org/drawingml/2006/main`
	XMLNS_DRAWINGML_PICTURE = `http://schemas.openxmlformats.org/drawingml/2006/picture`
)

// Drawing element contains photos
type Drawing struct {
	XMLName xml.Name `xml:"w:drawing,omitempty"`
	Inline  *WPInline
	Anchor  *WPAnchor
}

func (r *Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "inline":
				r.Inline = new(WPInline)
				d.DecodeElement(r.Inline, &tt)
			case "anchor":
				r.Anchor = new(WPAnchor)
				d.DecodeElement(r.Anchor, &tt)
			default:
				continue
			}
		}

	}
	return nil

}

// WPInline wp:inline
type WPInline struct {
	XMLName xml.Name `xml:"wp:inline,omitempty"`
	DistT   int64    `xml:"distT,attr"`
	DistB   int64    `xml:"distB,attr"`
	DistL   int64    `xml:"distL,attr"`
	DistR   int64    `xml:"distR,attr"`
	// AnchorID string   `xml:"wp14:anchorId,attr,omitempty"`
	// EditID   string   `xml:"wp14:editId,attr,omitempty"`

	Extent            *WPExtent
	EffectExtent      *WPEffectExtent
	DocPr             *WPDocPr
	CNvGraphicFramePr *WPCNvGraphicFramePr
	Graphic           *AGraphic
}

func (r *WPInline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "distT":
			r.DistT, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return
			}
		case "distB":
			r.DistB, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return
			}
		case "distL":
			r.DistL, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return
			}
		case "distR":
			r.DistR, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return
			}
		default:
			// ignore other attributes
		}
	}
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
			switch tt.Name.Local {
			case "extent":
				r.Extent = new(WPExtent)
				r.Extent.CX, err = strconv.ParseInt(getAtt(tt.Attr, "cx"), 10, 64)
				if err != nil {
					return err
				}
				r.Extent.CY, err = strconv.ParseInt(getAtt(tt.Attr, "cy"), 10, 64)
				if err != nil {
					return err
				}
			case "effectExtent":
				r.EffectExtent = new(WPEffectExtent)
				r.EffectExtent.L, err = strconv.ParseInt(getAtt(tt.Attr, "l"), 10, 64)
				if err != nil {
					return err
				}
				r.EffectExtent.T, err = strconv.ParseInt(getAtt(tt.Attr, "t"), 10, 64)
				if err != nil {
					return err
				}
				r.EffectExtent.R, err = strconv.ParseInt(getAtt(tt.Attr, "r"), 10, 64)
				if err != nil {
					return err
				}
				r.EffectExtent.B, err = strconv.ParseInt(getAtt(tt.Attr, "b"), 10, 64)
				if err != nil {
					return err
				}
			case "docPr":
				r.DocPr = new(WPDocPr)
				d.DecodeElement(r.DocPr, &tt)
			case "cNvGraphicFramePr":
				var value WPCNvGraphicFramePr
				d.DecodeElement(&value, &tt)
				r.CNvGraphicFramePr = &value
			case "graphic":
				var value AGraphic
				d.DecodeElement(&value, &tt)
				r.Graphic = &value
			default:
				continue
			}
		}

	}
	return nil

}

// WPExtent represents the extent of a drawing in a Word document.
//
//	CX CY 's unit is English Metric Units, which is 1/914400 inch
type WPExtent struct {
	XMLName xml.Name `xml:"wp:extent,omitempty"`
	CX      int64    `xml:"cx,attr"`
	CY      int64    `xml:"cy,attr"`
}

func (r *WPExtent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "cx":
			r.CX, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		case "cy":
			r.CY, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		}
	}
	// Consume the end element
	_, err = d.Token()
	if err != nil {
		return err
	}
	return nil
}

// WPEffectExtent represents the effect extent of a drawing in a Word document.
type WPEffectExtent struct {
	XMLName xml.Name `xml:"wp:effectExtent,omitempty"`
	L       int64    `xml:"l,attr"`
	T       int64    `xml:"t,attr"`
	R       int64    `xml:"r,attr"`
	B       int64    `xml:"b,attr"`
}

func (r *WPEffectExtent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "l":
			r.L, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		case "t":
			r.T, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		case "r":
			r.R, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		case "b":
			r.B, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		}
	}
	// Consume the end element
	_, err = d.Token()
	if err != nil {
		return err
	}
	return nil
}

// WPDocPr represents the document properties of a drawing in a Word document.
type WPDocPr struct {
	XMLName xml.Name `xml:"wp:docPr,omitempty"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr,omitempty"`
}

func (r *WPDocPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			id, err := strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
			r.ID = id
		case "name":
			r.Name = attr.Value

		default:
			// ignore other attributes
		}
	}
	// Consume the end element
	_, err := d.Token()
	if err != nil {
		return err
	}
	return nil
}

// WPCNvGraphicFramePr represents the non-visual properties of a graphic frame.
type WPCNvGraphicFramePr struct {
	XMLName xml.Name `xml:"wp:cNvGraphicFramePr,omitempty"`
	Locks   *AGraphicFrameLocks
}

func (w *WPCNvGraphicFramePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "graphicFrameLocks":
				var value AGraphicFrameLocks
				d.DecodeElement(&value, &tt)
				value.NoChangeAspect, err = strconv.Atoi(getAtt(tt.Attr, "noChangeAspect"))
				if err != nil {
					return err
				}
				w.Locks = &value
			default:
				continue
			}
		}

	}
	return nil
}

// AGraphicFrameLocks represents the locks applied to a graphic frame.
type AGraphicFrameLocks struct {
	XMLName        xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main graphicFrameLocks,omitempty"`
	NoChangeAspect int      `xml:"noChangeAspect,attr"`
}

// AGraphic represents a graphic in a Word document.
type AGraphic struct {
	XMLName     xml.Name `xml:"a:graphic,omitempty"`
	XMLA        string   `xml:"xmlns:a,attr,omitempty"`
	GraphicData *AGraphicData
}

func (a *AGraphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "a":
			a.XMLA = attr.Value
		default:
			// ignore other attributes
		}
	}
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
			switch tt.Name.Local {
			case "graphicData":
				var value AGraphicData
				d.DecodeElement(&value, &tt)
				value.URI = getAtt(tt.Attr, "uri")
				a.GraphicData = &value
			default:
				continue
			}
		}

	}
	return nil
}

// AGraphicData represents the data of a graphic in a Word document.
type AGraphicData struct {
	XMLName xml.Name `xml:"a:graphicData,omitempty"`
	URI     string   `xml:"uri,attr"`
	Pic     *PICPic
}

func (a *AGraphicData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "pic":
				var value PICPic
				d.DecodeElement(&value, &tt)
				value.XMLPIC = getAtt(tt.Attr, "pic")
				a.Pic = &value
			default:
				continue
			}
		}

	}
	return nil
}

// PICPic represents a picture in a Word document.
type PICPic struct {
	XMLName                xml.Name `xml:"pic:pic,omitempty"`
	XMLPIC                 string   `xml:"xmlns:pic,attr,omitempty"`
	NonVisualPicProperties *PICNonVisualPicProperties
	BlipFill               *PICBlipFill
	SpPr                   *PICSpPr
}

func (p *PICPic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "nvPicPr":
				var value PICNonVisualPicProperties
				d.DecodeElement(&value, &tt)
				p.NonVisualPicProperties = &value
			case "blipFill":
				var value PICBlipFill
				d.DecodeElement(&value, &tt)
				p.BlipFill = &value
			case "spPr":
				var value PICSpPr
				d.DecodeElement(&value, &tt)
				p.SpPr = &value
			default:
				continue
			}
		}

	}
	return nil
}

// PICNonVisualPicProperties represents the non-visual properties of a picture in a Word document.
type PICNonVisualPicProperties struct {
	XMLName                    xml.Name `xml:"pic:nvPicPr,omitempty"`
	NonVisualDrawingProperties PICNonVisualDrawingProperties
	CNvPicPr                   struct{} `xml:"pic:cNvPicPr"`
}

func (p *PICNonVisualPicProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "cNvPr":
				p.NonVisualDrawingProperties.ID = getAtt(tt.Attr, "id")
				p.NonVisualDrawingProperties.Name = getAtt(tt.Attr, "name")
			default:
				continue
			}
		}

	}
	return nil
}

// PICNonVisualDrawingProperties represents the non-visual drawing properties of a picture in a Word document.
type PICNonVisualDrawingProperties struct {
	XMLName xml.Name `xml:"pic:cNvPr,omitempty"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// PICBlipFill represents the blip fill of a picture in a Word document.
type PICBlipFill struct {
	XMLName xml.Name `xml:"pic:blipFill,omitempty"`
	Blip    ABlip
	Stretch AStretch
}

func (p *PICBlipFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "blip":
				d.DecodeElement(&p.Blip, &tt)
			case "stretch":
				d.DecodeElement(&p.Stretch, &tt)
			default:
				continue
			}
		}

	}
	return nil
}

// ABlip represents the blip of a picture in a Word document.
type ABlip struct {
	XMLName     xml.Name `xml:"a:blip,omitempty"`
	Embed       string   `xml:"r:embed,attr"`
	Cstate      string   `xml:"cstate,attr"`
	AlphaModFix *AAlphaModFix
}

func (a *ABlip) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "embed":
			a.Embed = attr.Value
		case "cstate":
			a.Cstate = attr.Value
		default:
			// ignore other attributes
		}
	}
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
			switch tt.Name.Local {
			case "alphaModFix":
				var value AAlphaModFix
				value.Amount, err = strconv.Atoi(getAtt(tt.Attr, "amt"))
				if err != nil {
					return err
				}
				a.AlphaModFix = &value
			default:
				continue
			}
		}

	}
	return nil
}

type AAlphaModFix struct {
	XMLName xml.Name `xml:"a:alphaModFix,omitempty"`
	Amount  int      `xml:"amt,attr"`
}

// AStretch ...
type AStretch struct {
	XMLName  xml.Name `xml:"a:stretch,omitempty"`
	FillRect AFillRect
}

// AFillRect ...
type AFillRect struct {
	XMLName xml.Name `xml:"a:fillRect,omitempty"`
}

// PICSpPr is a struct representing the <pic:spPr> element in OpenXML,
// which describes the shape properties for a picture.
type PICSpPr struct {
	XMLName  xml.Name `xml:"pic:spPr,omitempty"`
	Xfrm     AXfrm
	PrstGeom APrstGeom
}

func (p *PICSpPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch tt.Name.Local {
			case "xfrm":
				d.DecodeElement(&p.Xfrm, &tt)
			case "prstGeom":
				d.DecodeElement(&p.PrstGeom, &tt)
				p.PrstGeom.Prst = getAtt(tt.Attr, "prst")
			default:
				continue
			}
		}

	}
	return nil
}

// AXfrm is a struct representing the <a:xfrm> element in OpenXML,
// which describes the position and size of a shape.
type AXfrm struct {
	XMLName xml.Name `xml:"a:xfrm,omitempty"`
	Rot     int64    `xml:"rot,attr,omitempty"`
	FlipH   int      `xml:"flipH,attr,omitempty"`
	FlipV   int      `xml:"flipV,attr,omitempty"`
	Off     AOff
	Ext     AExt
}

func (a *AXfrm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rot":
			a.Rot, err = strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
		case "flipH":
			a.FlipH, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
		case "flipV":
			a.FlipV, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
		default:
			// ignore other attributes
		}
	}
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
			switch tt.Name.Local {
			case "off":
				a.Off.X, err = strconv.ParseInt(getAtt(tt.Attr, "x"), 10, 64)
				if err != nil {
					return err
				}
				a.Off.Y, err = strconv.ParseInt(getAtt(tt.Attr, "y"), 10, 64)
				if err != nil {
					return err
				}
			case "ext":
				a.Ext.CX, err = strconv.ParseInt(getAtt(tt.Attr, "cx"), 10, 64)
				if err != nil {
					return err
				}
				a.Ext.CY, err = strconv.ParseInt(getAtt(tt.Attr, "cy"), 10, 64)
				if err != nil {
					return err
				}
			default:
				continue
			}
		}

	}
	return nil
}

// AOff is a struct representing the <a:off> element in OpenXML,
// which describes the offset of a shape from its original position.
type AOff struct {
	XMLName xml.Name `xml:"a:off,omitempty"`
	X       int64    `xml:"x,attr"`
	Y       int64    `xml:"y,attr"`
}

// AExt is a struct representing the <a:ext> element in OpenXML,
// which describes the size of a shape.
type AExt struct {
	XMLName xml.Name `xml:"a:ext,omitempty"`
	CX      int64    `xml:"cx,attr"`
	CY      int64    `xml:"cy,attr"`
}

// APrstGeom is a struct representing the <a:prstGeom> element in OpenXML,
// which describes the preset shape geometry for a shape.
type APrstGeom struct {
	XMLName xml.Name `xml:"a:prstGeom,omitempty"`
	Prst    string   `xml:"prst,attr"`
	AvLst   AAvLst
}

// AAvLst is a struct representing the <a:avLst> element in OpenXML,
// which describes the adjustments to the shape's preset geometry.
type AAvLst struct {
	XMLName xml.Name `xml:"a:avLst,omitempty"`
	RawXML  string   `xml:",innerxml"`
}

func (a *AAvLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content []byte

	if content, err = xml.Marshal(start); err != nil {
		return err
	}

	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if end, ok := t.(xml.EndElement); ok && end == start.End() {
			break
		}

		b, err := xml.Marshal(t)
		if err != nil {
			return err
		}

		content = append(content, b...)
	}

	a.RawXML = BytesToString(content)

	return nil
}

// WPAnchor wp:anchor
type WPAnchor struct {
	XMLName        xml.Name `xml:"wp:anchor,omitempty"`
	DistT          int64    `xml:"distT,attr"`
	DistB          int64    `xml:"distB,attr"`
	DistL          int64    `xml:"distL,attr"`
	DistR          int64    `xml:"distR,attr"`
	SimplePos      int      `xml:"simplePos,attr"`
	RelativeHeight int      `xml:"relativeHeight,attr"`
	BehindDoc      int      `xml:"behindDoc,attr"`
	Locked         int      `xml:"locked,attr"`
	LayoutInCell   int      `xml:"layoutInCell,attr"`
	AllowOverlap   int      `xml:"allowOverlap,attr"`

	SimplePosXY       *WPSimplePos
	PositionH         *WPPositionH
	PositionV         *WPPositionV
	Extent            *WPExtent
	EffectExtent      *WPEffectExtent
	WrapNone          *struct{} `xml:"wp:wrapNone,omitempty"`
	WrapSquare        *WPWrapSquare
	DocPr             *WPDocPr
	CNvGraphicFramePr *WPCNvGraphicFramePr
	Graphic           *AGraphic
}

func (r *WPAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, tt := range start.Attr {
		switch tt.Name.Local {
		case "distT":
			r.DistT, err = strconv.ParseInt(tt.Value, 10, 64)
			if err != nil {
				return err
			}
		case "distB":
			r.DistB, err = strconv.ParseInt(tt.Value, 10, 64)
			if err != nil {
				return err
			}
		case "distL":
			r.DistL, err = strconv.ParseInt(tt.Value, 10, 64)
			if err != nil {
				return err
			}
		case "distR":
			r.DistR, err = strconv.ParseInt(tt.Value, 10, 64)
			if err != nil {
				return err
			}
		case "simplePos":
			r.SimplePos, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		case "relativeHeight":
			r.RelativeHeight, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		case "behindDoc":
			r.BehindDoc, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		case "locked":
			r.Locked, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		case "layoutInCell":
			r.LayoutInCell, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		case "allowOverlap":
			r.AllowOverlap, err = strconv.Atoi(tt.Value)
			if err != nil {
				return err
			}
		}
	}
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
			switch tt.Name.Local {
			case "simplePos":
				r.SimplePosXY = new(WPSimplePos)
				r.SimplePosXY.X, err = strconv.ParseInt(getAtt(tt.Attr, "x"), 10, 64)
				if err != nil {
					return err
				}
				r.SimplePosXY.Y, err = strconv.ParseInt(getAtt(tt.Attr, "y"), 10, 64)
				if err != nil {
					return err
				}
			case "positionH":
				r.PositionH = new(WPPositionH)
				// r.PositionH.RelativeFrom = getAtt(tt.Attr, "relativeFrom")
				d.DecodeElement(&r.PositionH, &tt)
			case "positionV":
				r.PositionV = new(WPPositionV)
				// r.PositionV.RelativeFrom = getAtt(tt.Attr, "relativeFrom")
				d.DecodeElement(&r.PositionV, &tt)
			case "extent":
				r.Extent = new(WPExtent)
				d.DecodeElement(&r.Extent, &tt)
			case "effectExtent":
				r.EffectExtent = new(WPEffectExtent)
				d.DecodeElement(&r.EffectExtent, &tt)
			case "wrapNone":
				r.WrapNone = &struct{}{}
			case "wrapSquare":
				r.WrapSquare = new(WPWrapSquare)
				r.WrapSquare.WrapText = getAtt(tt.Attr, "wrapText")
			case "docPr":
				r.DocPr = new(WPDocPr)
				d.DecodeElement(r.DocPr, &tt)
			case "cNvGraphicFramePr":
				r.CNvGraphicFramePr = new(WPCNvGraphicFramePr)
				d.DecodeElement(r.CNvGraphicFramePr, &tt)
			case "graphic":
				r.Graphic = new(AGraphic)
				d.DecodeElement(&r.Graphic, &tt)
			default:
				continue
			}
		}
	}
	return nil
}

// WPSimplePos represents the position of an object in a Word document.
type WPSimplePos struct {
	XMLName xml.Name `xml:"wp:simplePos,omitempty"`
	X       int64    `xml:"x,attr"`
	Y       int64    `xml:"y,attr"`
}

// WPPositionH represents the horizontal position of an object in a Word document.
type WPPositionH struct {
	XMLName      xml.Name `xml:"wp:positionH,omitempty"`
	RelativeFrom string   `xml:"relativeFrom,attr"`
	PosOffset    int64    `xml:"wp:posOffset"`
}

func (r *WPPositionH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "relativeFrom":
			r.RelativeFrom = attr.Value
		}
	}
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
			switch tt.Name.Local {
			case "posOffset":
				err = d.DecodeElement(&r.PosOffset, &tt)
				if err != nil {
					return err
				}
			default:
				continue
			}
		}
	}
	return nil
}

// WPPositionV represents the vertical position of an object in a Word document.
type WPPositionV struct {
	XMLName      xml.Name `xml:"wp:positionV,omitempty"`
	RelativeFrom string   `xml:"relativeFrom,attr"`
	PosOffset    int64    `xml:"wp:posOffset"`
}

func (r *WPPositionV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "relativeFrom":
			r.RelativeFrom = attr.Value
		}
	}
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
			switch tt.Name.Local {
			case "posOffset":
				err = d.DecodeElement(&r.PosOffset, &tt)
				if err != nil {
					return err
				}
			default:
				continue
			}
		}
	}
	return nil
}

// WPWrapSquare represents the square wrapping of an object in a Word document.
type WPWrapSquare struct {
	XMLName  xml.Name `xml:"wp:wrapSquare,omitempty"`
	WrapText string   `xml:"wrapText,attr"`
}