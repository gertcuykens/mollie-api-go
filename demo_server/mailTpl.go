package server

// SERVER
const tpl = `<!DOCTYPE html>
<html>

<head>
  <title>Demo</title>
  <style>
    table,
    td {
      border-collapse: collapse;
      padding: 0;
      margin: 0;
    }
  </style>
</head>

<body>
  <h3>{{.Event.Title}}</h3>

  <p>
    Id: {{.ID}}
    <br/>Organisator: {{.Event.Organizer}}
    <br/>Datum: {{.Event.Date}}
    <br/>IBAN: {{.Event.Iban}}
  </p>

  <table>
    <tr>
      <td>Bedrag:</td>
      <td style="text-align:right; padding-left:5px;">{{.Randonneur.Price}}</td>
      <td>&euro;</td>
    </tr>
    <tr>
      <td>Afstands medaille:</td>
      <td style="text-align:right; padding-left:5px;">{{.Randonneur.Medaille}}</td>
      <td>&euro;</td>
    </tr>
    <tr>
      <td>Korting NTFU (en andere):</td>
      <td style="text-align:right; padding-left:5px;">{{.Randonneur.Discount1}}</td>
      <td>&euro;</td>
    </tr>
    <tr>
      <td>Korting ERN lid:</td>
      <td style="text-align:right; padding-left:5px;">{{.Randonneur.Discount2}}</td>
      <td>&euro;</td>
    </tr>
    <tr>
      <td>Korting voorinschrijving:</td>
      <td style="text-align:right; padding-left:5px;">{{.Randonneur.Discount3}}</td>
      <td>&euro;</td>
    </tr>
    <tr>
      <td></td>
      <td style="border-top:1px solid black; text-align:right; padding-left:5px; min-width:50px;">
	    <a style="text-decoration:none;" href="` + SERVER + `/status.html?id={{.ID}}&nr={{.Nr}}">{{.Randonneur.Total}}</a>
      </td>
      <td>&euro;</td>
    </tr>
  </table>
  
  <br/>
  
  <table>
    <tr>
      <td>Naam:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Lastname}}</td>
    </tr>
    <tr>
      <td>Voornaam:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Firstname}} ({{.Randonneur.Gender}})</td>
    </tr>
    <tr>
      <td>Straat:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Street}}</td>
    </tr>
    <tr>
      <td>Postcode:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Zipcode}}</td>
    </tr>
    <tr>
      <td>Plaats:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.City}}</td>
    </tr>
    <tr>
      <td>Land:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Country}}</td>
    </tr>
    <tr>
      <td>Geboortedatum:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Birthday}}</td>
    </tr>
    <tr>
      <td>e-mail:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Email}}</td>
    </tr>
    <tr>
      <td>Mobiel:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Mobile}}</td>
    </tr>
    <tr>
      <td>Noodgeval:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Emergency}}</td>
    </tr>
    <tr>
      <td>Clubnaam:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.Club}}</td>
    </tr>
    <tr>
      <td>ACP:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.ACP}}</td>
    </tr>
    <tr>
      <td>NTFU:</td>
      <td style="text-align:left; padding-left:5px;">{{.Randonneur.NTFU}}</td>
    </tr>
    <tr>
      <td>Status:</td>
      <td style="text-align:left; padding-left:5px;">
        <a href="` + SERVER + `/status.html?id={{.ID}}&nr={{.Nr}}">{{.Randonneur.Status}}</a>
      </td>
    </tr>
  </table>

  <p>
    {{.Randonneur.Comment}}
    <br/>{{.Randonneur.Start}}
  </p>

  <p>
    Ik heb kennis genomen van het BRM (Brevet Randonneurs Mondiaux) <a href="http://www.randonneurs.nl/reglementen/brm-brevets-randonneurs-mondiaux/" target="_blank">reglement</a> en ga accoord met de gestelde regels en voorwaarden.
  </p>

  <p>
    Ik verklaar dat ik een passende verzekering heb afgesloten voor aansprakelijkheid bij ongevallen en andere gebeurtenissen.
  </p>

  <p>
    Begrippen en Informatie voor Beginnende <a href="http://randonneurs.nl/reglementen/begrippen-en-informatie-voor-beginnende-randonneurs/">Randonneurs</a>.
  </p>

</body>

</html>`
