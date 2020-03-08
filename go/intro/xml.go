package main

import (
    "encoding/xml"
    "fmt"
)

type Plant struct {
    XMLName  xml.Name `xml:"plant"`
    Id       int      `xml: "id, attr"`// Id field is an XML attribute rather than a nested element
    Name     string   `xml:"name"`
    Origin   []string `xml:"origin"`
}

func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)

}
func main() {
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    out, _ := xml.MarshalIndent(coffee, " ", " ")
    fmt.Println(string(out))
    /*
    <plant>
    <Id>27</Id>
    <name>Coffee</name>
    <origin>Ethiopia</origin>
    <origin>Brazil</origin>
    </plant>
    */

    fmt.Println(xml.Header + string(out))
    /*
    <?xml version="1.0" encoding="UTF-8"?>
     <plant>
      <Id>27</Id>
      <name>Coffee</name>
      <origin>Ethiopia</origin>
      <origin>Brazil</origin>
     </plant>
    */
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p) //Plant id=27, name=Coffee, origin=[Ethiopia Brazil]

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"` // tells the encoder to nest all plants under <parent><child>
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", " ")
    fmt.Println(string(out))
    /*
    <nesting>
     <parent>
      <child>
       <plant>
        <Id>27</Id>
        <name>Coffee</name>
        <origin>Ethiopia</origin>
        <origin>Brazil</origin>
       </plant>
       <plant>
        <Id>81</Id>
        <name>Tomato</name>
        <origin>Mexico</origin>
        <origin>California</origin>
       </plant>
      </child>
     </parent>
    </nesting>
    */

}
