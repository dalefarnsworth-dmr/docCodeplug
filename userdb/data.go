// Copyright 2017-2019 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of UserDB.
//
// UserDB is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU Lesser General Public
// License as published by the Free Software Foundation.
//
// UserDB is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with UserDB.  If not, see <http://www.gnu.org/licenses/>.

package userdb

var countryAbbreviations = map[string]string{
	"Andorra":                "AD",
	"Argentina":              "AR",
	"Australia":              "AU",
	"Austria":                "AT",
	"Barbados":               "BB",
	"Belgium":                "BE",
	"Belize":                 "BZ",
	"Bosnia and Hercegovina": "BA",
	"Brazil":                 "BR",
	"Bulgaria":               "BG",
	"Canada":                 "CA",
	"Chile":                  "CL",
	"China":                  "CN",
	"Colombia":               "CO",
	"Costa Rica":             "CR",
	"Croatia":                "HR",
	"Cyprus":                 "CY",
	"Czech Republic":         "CZ",
	"Denmark":                "DK",
	"Dominica":               "DM",
	"Dominican Republic":     "DO",
	"Ecuador":                "EC",
	"Estonia":                "EE",
	"Finland":                "FI",
	"France":                 "FR",
	"Germany":                "DE",
	"Greece":                 "GR",
	"Guatemala":              "GT",
	"Haiti":                  "HT",
	"Hong Kong":              "HK",
	"Hungary":                "HU",
	"India":                  "IN",
	"Indonesia":              "ID",
	"Ireland":                "IE",
	"Israel":                 "IL",
	"Italy":                  "IT",
	"Japan":                  "JP",
	"Korea":                  "KR",
	"Kuwait":                 "KW",
	"Latvia":                 "LV",
	"Liechtenstein":          "LI",
	"Luxembourg":             "LU",
	"Macedonia":              "MK",
	"Malaysia":               "MY",
	"Malta":                  "MT",
	"Mexico":                 "MX",
	"Montenegro":             "ME",
	"Netherlands Antilles":   "AN",
	"Netherlands":            "NL",
	"New Zealand":            "NZ",
	"Norway":                 "NO",
	"Panama":                 "PA",
	"Paraguay":               "PY",
	"Philippines":            "PH",
	"Poland":                 "PL",
	"Portugal":               "PT",
	"Qatar":                  "QA",
	"Reunion":                "RE",
	"Romania":                "RO",
	"Russia":                 "RU",
	"Serbia":                 "RS",
	"Singapore":              "SG",
	"Slovakia":               "SK",
	"Slovenia":               "SI",
	"South Africa":           "ZA",
	"Spain":                  "ES",
	"Sweden":                 "SE",
	"Switzerland":            "CH",
	"Taiwan":                 "TW",
	"Thailand":               "TH",
	"Trinidad and Tobago":    "TT",
	"Turkey":                 "TR",
	"Ukraine":                "UA",
	"United Kingdom":         "UK",
	"United States":          "US",
	"Uruguay":                "UY",
	"Venezuela":              "VE",
}

var extraCountryAbbreviations = map[string]string{
	"Argentina Republic":   "AR",
	"Bosnia and Hercegovi": "BA",
	"Korea S":              "KR",
	"Luxemburg":            "LU",
}

var stateAbbreviationsByCountry = map[string]map[string]string{
	"United States": {
		"Alabama":              "AL",
		"Alaska":               "AK",
		"Arizona":              "AZ",
		"Arkansas":             "AR",
		"California":           "CA",
		"Colorado":             "CO",
		"Connecticut":          "CT",
		"Delaware":             "DE",
		"District of Columbia": "DC",
		"Florida":              "FL",
		"Georgia":              "GA",
		"Hawaii":               "HI",
		"Idaho":                "ID",
		"Illinois":             "IL",
		"Indiana":              "IN",
		"Iowa":                 "IA",
		"Kansas":               "KS",
		"Kentucky":             "KY",
		"Louisiana":            "LA",
		"Maine":                "ME",
		"Maryland":             "MD",
		"Massachusetts":        "MA",
		"Michigan":             "MI",
		"Minnesota":            "MN",
		"Mississippi":          "MS",
		"Missouri":             "MO",
		"Montana":              "MT",
		"Nebraska":             "NE",
		"Nevada":               "NV",
		"New Hampshire":        "NH",
		"New Jersey":           "NJ",
		"New Mexico":           "NM",
		"New York":             "NY",
		"North Carolina":       "NC",
		"North Dakota":         "ND",
		"Ohio":                 "OH",
		"Oklahoma":             "OK",
		"Oregon":               "OR",
		"Pennsylvania":         "PA",
		"Puerto Rico":          "PR",
		"Rhode Island":         "RI",
		"South Carolina":       "SC",
		"South Dakota":         "SD",
		"Tennessee":            "TN",
		"Texas":                "TX",
		"Utah":                 "UT",
		"Vermont":              "VT",
		"Virginia":             "VA",
		"Washington":           "WA",
		"West Virginia":        "WV",
		"Wisconsin":            "WI",
	},
	"Canada": {
		"Alberta":              "AB",
		"British Columbia":     "BC",
		"Manitoba":             "MB",
		"New Brunswick":        "NB",
		"Newfoundland":         "NL",
		"Nova Scotia":          "NS",
		"Ontario":              "ON",
		"Prince Edward Island": "PE",
		"Quebec":               "QC",
		"Saskatchewan":         "SK",
		"Yukon":                "YT",
	},
	"Australia": {
		"Australian Capital Territory": "ACT",
		"New South Wales":              "NSW",
		"Northern Territory":           "NT",
		"Queensland":                   "QLD",
		"South Australia":              "SA",
		"Tasmania":                     "TAS",
		"Victoria":                     "VIC",
		"Western Australia":            "WAU",
	},
	"Germany": {
		"Baden-Wuerttemberg":     "BW",
		"Bavaria":                "BY",
		"Berlin":                 "BE",
		"Brandenburg":            "BB",
		"Bremen":                 "HB",
		"Hamburg":                "HH",
		"Hessen":                 "HE",
		"Lower Saxony":           "NI",
		"Mecklenburg-Vorpommern": "MV",
		"North Rhine-Westphalia": "NW",
		"Rhineland-Palatinate":   "RP",
		"Saarland":               "SL",
		"Saxony-Anhalt":          "ST",
		"Saxony":                 "SN",
		"Schleswig-Holstein":     "SH",
		"Thuringia":              "TH",
	},
	"Netherlands": {
		"Drenthe":         "DR",
		"Flevoland":       "FD",
		"Friesland":       "FR",
		"Gelderland":      "GE",
		"Groningen":       "GR",
		"Limburg":         "LI",
		"North Brabant":   "N-B",
		"North Holland":   "N-H",
		"Oost-Vlaanderen": "O-V",
		"Overijssel":      "OV",
		"South Holland":   "ZH",
		"Utrecht":         "UTR",
		"Zeeland":         "ZE",
	},
	"Belgium": {
		"Antwerp":         "VAN",
		"East Flanders":   "VOV",
		"Flemish Brabant": "VBR",
		"Hainaut":         "WHT",
		"Ile-de-France":   "IF",
		"Liege":           "WLG",
		"Limburg":         "VLI",
		"Luxembourg":      "WLX",
		"Namur":           "WNA",
		"Walloon Brabant": "WBR",
		"West Flanders":   "VWV",
	},
}

var extraStateAbbreviationsByCountry = map[string]map[string]string{
	"Germany": {
		"Baden-Wurttemberg": "BW",
	},
	"Netherlands": {
		"Noord-Brabant": "N-B",
		"Noord-Holland": "N-H",
		"Zuid-Holland":  "ZH",
	},
	"Belgium": {
		"Antwerpen":       "VAN",
		"Henegouwen":      "WHT",
		"le-de-France":    "IF",
		"Namen":           "WNA",
		"Vlaams-Brabant":  "VBR",
		"West-Vlaanderen": "VWV",
	},
}

var directionAbbreviations = map[string]string{
	"North": "N.",
	"South": "S.",
	"East":  "E.",
	"West":  "W.",
}

var titleCaseWords = []string{
	"AALBORG",
	"AARON",
	"ABEL",
	"ABY",
	"ACACIO",
	"ACHAZ",
	"ACHILLEAS",
	"AD",
	"ADA",
	"ADAM",
	"ADELIO",
	"ADRIAN",
	"ADRIANO",
	"AGUSTIN",
	"AHOSKIE",
	"AICHI",
	"AITOR",
	"AKAA",
	"ALAIN",
	"ALAM",
	"ALAN",
	"ALBERT",
	"ALBERTO",
	"ALBUQUERQUE",
	"ALCORCON",
	"ALDO",
	"ALEJANDRO",
	"ALESSANDRA",
	"ALESSANDRO",
	"ALEX",
	"ALEXANDRE",
	"ALEXANDROS",
	"ALFA",
	"ALFONSO",
	"ALFRED",
	"ALFREDO",
	"ALI",
	"ALISTER",
	"ALL",
	"ALLAN",
	"ALLIGAND",
	"ALONSO",
	"ALOSHYN",
	"ALTO",
	"ALVADOR",
	"ALVARO",
	"AMARO",
	"AMATEUR",
	"AMBERG",
	"AMEDEO",
	"AMIRED",
	"ANAHEIM",
	"ANASTASIOS",
	"ANCHISE",
	"ANDERSEN",
	"ANDERSON",
	"ANDRE",
	"ANDREA",
	"ANDREAS",
	"ANDRES",
	"ANDREW",
	"ANDRZEJ",
	"ANDY",
	"ANGEL",
	"ANGELES",
	"ANGELO",
	"ANGELOS",
	"ANIBAL",
	"ANJEL",
	"ANNA",
	"ANTHONY",
	"ANTONINO",
	"ANTONIO",
	"ANTONIOS",
	"ANTONY",
	"AOSTA",
	"ARA",
	"ARACAJU",
	"ARALA",
	"ARANDAS",
	"ARC",
	"AREA",
	"AREC",
	"ARES",
	"ARI",
	"ARIS",
	"ARISTIDIS",
	"ARITZ",
	"ARLINGTON",
	"ARNAUD",
	"ARON",
	"ARROW",
	"ARTHUR",
	"ARTUR",
	"ARTURO",
	"ASCANIO",
	"ASHEBORO",
	"ASHEVILLE",
	"ASIS",
	"ASOCCULTURAL",
	"ASSN",
	"ATA",
	"ATILLA",
	"ATLANTIC",
	"AUCKLAND",
	"AUGUSTO",
	"AURORA",
	"AUXCOMM",
	"AZORES",
	"BAGARMOSSEN",
	"BAKERSVILLE",
	"BALANICI",
	"BALBINO",
	"BALDWIN",
	"BAR",
	"BARBARA",
	"BARCELONA",
	"BARIS",
	"BARRY",
	"BARTLESVILLE",
	"BARTOLOMEO",
	"BARTOSZ",
	"BAUD",
	"BAUTISTA",
	"BAY",
	"BEACH",
	"BEAR",
	"BEAT",
	"BEAVER",
	"BECHMANN",
	"BELLINGHAM",
	"BELLMORE",
	"BELO",
	"BEN",
	"BENEDETTO",
	"BENITO",
	"BENJAMIN",
	"BENNAS",
	"BENNY",
	"BENOIT",
	"BERKLEY",
	"BERNARD",
	"BERNARDINA",
	"BERNARDO",
	"BERND",
	"BERT",
	"BETHANY",
	"BEX",
	"BIAGIO",
	"BIRMINGHAM",
	"BISHOPS",
	"BIXBY",
	"BLOCKED",
	"BOCA",
	"BODEN",
	"BOELL",
	"BOJAN",
	"BOLINGBROOK",
	"BOLL",
	"BOLOGNA",
	"BOLZANO",
	"BORTA",
	"BOSSY",
	"BOWMANVILLE",
	"BRAMPTON",
	"BRANFORD",
	"BRETT",
	"BRIAN",
	"BRIDGETON",
	"BRISTOL",
	"BROKEN",
	"BRONX",
	"BROOKEVILLE",
	"BROOKLYN",
	"BROWNS",
	"BROWNSBURG",
	"BRUNO",
	"BRYAN",
	"BULL",
	"BULLI",
	"BUNN",
	"BURCAK",
	"BURHAN",
	"BURLINGTON",
	"BUSAN",
	"BUSH",
	"CADET",
	"CAETANO",
	"CALDAS",
	"CALDEROLI",
	"CALGARY",
	"CALHOUN",
	"CAMARILLO",
	"CAMBRIDGE",
	"CAMDEN",
	"CAMP",
	"CAMPAZI",
	"CANADA",
	"CANCUN",
	"CANDIDO",
	"CAPENA",
	"CARAGIULO",
	"CARDONA",
	"CARIDAD",
	"CARINOLA",
	"CARLA",
	"CARLO",
	"CARLOS",
	"CARMELO",
	"CARPINELLI",
	"CARROLLTON",
	"CARUARU",
	"CARUGATE",
	"CASANOVAS",
	"CASCADE",
	"CASTAIC",
	"CASTALIAN",
	"CASTOR",
	"CATALIN",
	"CATANIA",
	"CATHERINE",
	"CAXIAS",
	"CD",
	"CEDRIC",
	"CELESTINO",
	"CENK",
	"CENTENNIAL",
	"CESAR",
	"CESARE",
	"CHANDLER",
	"CHANG",
	"CHARLES",
	"CHAVES",
	"CHENG",
	"CHEONG",
	"CHESTER",
	"CHEYENNE",
	"CHIA",
	"CHICAGO",
	"CHINO",
	"CHOCTAW",
	"CHRISOSTOMOS",
	"CHRISTCHURCH",
	"CHRISTIAN",
	"CHRISTOPHE",
	"CHRISTOPHER",
	"CHRISTOS",
	"CHULA",
	"CHURCH",
	"CHYUAN",
	"CIESZYN",
	"CINCO",
	"CIRO",
	"CITY",
	"CIVERA",
	"CLACKAMAS",
	"CLARKSVILLE",
	"CLAUDE",
	"CLAUDIO",
	"CLAUS",
	"CLETO",
	"CLIFTON",
	"CLINTON",
	"CLIVE",
	"CLUB",
	"CO",
	"COAST",
	"CODINA",
	"COLIN",
	"COLOGNE",
	"COLON",
	"COLORADO",
	"COLUMBIA",
	"COMANO",
	"COMO",
	"CONCORD",
	"CONQUENSES",
	"CONSUELO",
	"CONTAGEM",
	"COQUIMBO",
	"CORDIGNANO",
	"CORDSVILLE",
	"CORNER",
	"CORONA",
	"COSIMO",
	"COURSE",
	"COVENTRY",
	"COX",
	"CRAIG",
	"CREEK",
	"CREWE",
	"CRISNA",
	"CRISTIAN",
	"CRISTIANO",
	"CRISTOBAL",
	"CROSS",
	"CROWN",
	"CRUZ",
	"CTY",
	"CUNEYT",
	"CUTLER",
	"DAEGU",
	"DAI",
	"DALDOSS",
	"DALLAS",
	"DALTON",
	"DAMIANO",
	"DAMON",
	"DANIEL",
	"DANIELE",
	"DANNEMORA",
	"DANTE",
	"DANY",
	"DARIO",
	"DARREN",
	"DAVE",
	"DAVID",
	"DAVIDE",
	"DAVY",
	"DE",
	"DEAN",
	"DEBRAY",
	"DEKALB",
	"DELSBO",
	"DELTA",
	"DEMETRIO",
	"DEMIS",
	"DEMO",
	"DENIS",
	"DENNI",
	"DENNIS",
	"DEON",
	"DERECK",
	"DEREK",
	"DES",
	"DI",
	"DIAMOND",
	"DIEGO",
	"DIFONA",
	"DIMITRIOS",
	"DIMITRIS",
	"DIMOS",
	"DINNINGTON",
	"DIVINOPOLIS",
	"DNEPROPETROVSK",
	"DO",
	"DOBBS",
	"DOLOMITES",
	"DOMENEC",
	"DOMENICO",
	"DOMINGA",
	"DOMINGO",
	"DOMINIC",
	"DONATO",
	"DONCASTER",
	"DORIN",
	"DORME",
	"DORR",
	"DOSQUET",
	"DOWNINGTOWN",
	"DRUMMONDVILLE",
	"DUEPRE",
	"DUNCAN",
	"DUNEDIN",
	"DUQUE",
	"DURHAM",
	"DURSUN",
	"DX",
	"EAST",
	"EASTLAKE",
	"EDITH",
	"EDMONTON",
	"EDOARDO",
	"EDUARDO",
	"EDWARD",
	"EFRAIM",
	"EFSTRATIOS",
	"EG",
	"EGUNA",
	"EL",
	"ELETHERIA",
	"ELIGIO",
	"ELIZABETH",
	"ELKHART",
	"ELMIRA",
	"ELOEM",
	"ELSINORE",
	"EMANUELE",
	"EMBRO",
	"EMILE",
	"EMILIANO",
	"EMILIO",
	"EMMANOUIL",
	"ENMORE",
	"ENNIO",
	"ENON",
	"ENRIC",
	"ENRICO",
	"ENRIQUE",
	"ENZO",
	"EOC",
	"EPSOM",
	"ERCOLE",
	"ERGUL",
	"ERHAN",
	"ERIC",
	"ERIKSSON",
	"ERIO",
	"ERMELLINO",
	"EROL",
	"ERRANTE",
	"ERVO",
	"ERWIN",
	"ESBJERG",
	"ESCUELA",
	"ESER",
	"ESLOV",
	"ESSEX",
	"ESTEBAN",
	"ESTER",
	"ETTERS",
	"ETTORE",
	"EUCLID",
	"EUGENIO",
	"EVAGGELOS",
	"EVANSVILLE",
	"EVERGREEN",
	"EVGENI",
	"EWA",
	"EZIO",
	"FABIO",
	"FABRETTI",
	"FABRICE",
	"FABRIZIO",
	"FALCON",
	"FALLS",
	"FARMINGVILLE",
	"FARSTA",
	"FAUSTO",
	"FCO",
	"FEDERICA",
	"FEDERICO",
	"FELIPE",
	"FELIX",
	"FERDINANDO",
	"FERDY",
	"FERNANDO",
	"FERRARA",
	"FERRY",
	"FIDEL",
	"FIDENZA",
	"FILIPPO",
	"FILOMENO",
	"FINSPNG",
	"FL",
	"FLAT",
	"FLAVIO",
	"FLEN",
	"FLIX",
	"FLORENCE",
	"FLORENTINO",
	"FLORIANO",
	"FOKION",
	"FORMIGINE",
	"FORSHAGA",
	"FORT",
	"FORTUNATO",
	"FOTIOS",
	"FOUNTAIN",
	"FRACCHIOLLA",
	"FRANCA",
	"FRANCESCA",
	"FRANCESCO",
	"FRANCIS",
	"FRANCISACO",
	"FRANCISCO",
	"FRANCO",
	"FRANCOIS",
	"FRANK",
	"FRANKLIN",
	"FRANKLINVILLE",
	"FRED",
	"FREDERIC",
	"FREITAS",
	"FRESNO",
	"FRIEDEN",
	"FUENTE",
	"FULTON",
	"FUNKERFREUNDE",
	"GABRIEL",
	"GABRIELE",
	"GALANTE",
	"GALBUSERA",
	"GALLATIN",
	"GARANHUNS",
	"GARCIA",
	"GARDNERVILLE",
	"GARETH",
	"GARIBALDI",
	"GARRY",
	"GARY",
	"GENNARO",
	"GEOFFREY",
	"GEOFFROY",
	"GEORGE",
	"GEORGIOS",
	"GERARD",
	"GERARDINO",
	"GERARDO",
	"GERBER",
	"GERLANDO",
	"GERMAN",
	"GERMANTOWN",
	"GERVASIO",
	"GIACINTO",
	"GIAN",
	"GIANCARLO",
	"GIANDOMENICO",
	"GIANFRANCO",
	"GIANLUCA",
	"GIANNI",
	"GIANNIS",
	"GILDO",
	"GILLES",
	"GINES",
	"GINO",
	"GIORGIO",
	"GIORGOS",
	"GIOVANNI",
	"GIRAY",
	"GIUBIASCO",
	"GIULIA",
	"GIUSEPPE",
	"GIVRINS",
	"GLASGOW",
	"GLASTONBURY",
	"GLIGA",
	"GLYN",
	"GODOWA",
	"GOFFREDO",
	"GONZALO",
	"GOOSE",
	"GORDON",
	"GOTHENBURG",
	"GOTTFRIED",
	"GRADO",
	"GRAHAM",
	"GRAND",
	"GRANDE",
	"GRANGE",
	"GRANITEVILLE",
	"GRAZIANO",
	"GREENBACKVILLE",
	"GREENSBORO",
	"GREGORIO",
	"GRIMSBY",
	"GRNQVIST",
	"GROUP",
	"GROVE",
	"GROVU",
	"GRUFTIZ",
	"GRUNDSUND",
	"GRUYERE",
	"GRZEGORZ",
	"GU",
	"GUADALAJARA",
	"GUANGLIN",
	"GUATEMALA",
	"GUAXUPE",
	"GUERINO",
	"GUIDO",
	"GUILLERMO",
	"GUIRGIO",
	"GUSTAVO",
	"GUTHRIE",
	"GUY",
	"HABRA",
	"HADONG",
	"HAINESVILLE",
	"HAKAN",
	"HAMDEN",
	"HAMILTON",
	"HAMMOND",
	"HAMPSTEAD",
	"HARBOR",
	"HARCO",
	"HARMAN",
	"HARRISON",
	"HARRY",
	"HARTFORD",
	"HATYAI",
	"HAWTHORNE",
	"HAYES",
	"HECTOR",
	"HEEB",
	"HELDER",
	"HELLAS",
	"HENDERSON",
	"HENDRIK",
	"HENRICO",
	"HENRYK",
	"HERMINIO",
	"HERMON",
	"HERVE",
	"HEYNSMANS",
	"HICKORY",
	"HILL",
	"HILLIARD",
	"HILLS",
	"HILO",
	"HJELM",
	"HONG",
	"HOPEDALE",
	"HOPEWELL",
	"HORIZONTE",
	"HOSCHTON",
	"HOUSTON",
	"HOVMANTORP",
	"HQ",
	"HRISTO",
	"HUDDINGE",
	"HUGHES",
	"HUGO",
	"HUGUES",
	"HUMBERTO",
	"HUNT",
	"HUNTINGTON",
	"IAIN",
	"IAN",
	"ICHUAN",
	"IEROTHEOS",
	"IGINO",
	"IGNACIO",
	"IINGNACIO",
	"ILION",
	"ILKESTON",
	"ILLAPEL",
	"IMPERO",
	"IMRE",
	"INDIANAPOLIS",
	"INN",
	"IOAN",
	"IOANNIS",
	"ION",
	"IONIA",
	"IOSEBA",
	"IPER",
	"IRAKLIS",
	"IRENEO",
	"IRVINE",
	"ISAAC",
	"ISIDORO",
	"ISIDRO",
	"ISLIP",
	"ISTANBUL",
	"IVAN",
	"IVANO",
	"IVO",
	"IVOR",
	"JACEK",
	"JACK",
	"JACKSON",
	"JACKSONVILLE",
	"JACKY",
	"JACQUES",
	"JACQUOT",
	"JAIME",
	"JAMES",
	"JAMIE",
	"JANAUBA",
	"JANEIRO",
	"JANEZ",
	"JANOS",
	"JASON",
	"JAUME",
	"JAVIER",
	"JEAN",
	"JEJU",
	"JENKINSVILLE",
	"JENSEN",
	"JEROME",
	"JESS",
	"JESUS",
	"JIM",
	"JOAN",
	"JOAO",
	"JOAQUIM",
	"JOAQUIN",
	"JOE",
	"JOEL",
	"JOHAN",
	"JOHANN",
	"JOHN",
	"JOHNNY",
	"JOISE",
	"JON",
	"JORDI",
	"JORGE",
	"JOSE",
	"JOSEBA",
	"JOSEF",
	"JOSEP",
	"JOSEPH",
	"JOSU",
	"JOZSEF",
	"JP",
	"JR",
	"JUAN",
	"JUANMA",
	"JUDICAEL",
	"JULIAN",
	"JULIEN",
	"JULIO",
	"JUSTO",
	"KADIR",
	"KAESER",
	"KAILUA",
	"KALIX",
	"KANATA",
	"KANE",
	"KAPITI",
	"KARA",
	"KARG",
	"KARL",
	"KARLSBORG",
	"KARST",
	"KATARZYNA",
	"KEANSBURG",
	"KEITH",
	"KELVIN",
	"KEMPISCHE",
	"KEN",
	"KENT",
	"KEVIN",
	"KEY",
	"KIL",
	"KINGSTON",
	"KINGWOOD",
	"KIRKLAND",
	"KISSIMMEE",
	"KLAASSE",
	"KLAUS",
	"KLUB",
	"KNOXVILLE",
	"KOKOMO",
	"KONG",
	"KONSTANTINOS",
	"KORAY",
	"KORIA",
	"KOSTAS",
	"KOUMALA",
	"KRIS",
	"KROTKOFALOWCOW",
	"KRZYSZTOF",
	"KYIV",
	"KYOTO",
	"LA",
	"LAFAYETTE",
	"LAGUNA",
	"LAILA",
	"LAKE",
	"LAKELAND",
	"LAKEWOOD",
	"LANDER",
	"LANDVETTER",
	"LANOKA",
	"LAPEER",
	"LARA",
	"LAS",
	"LATISANA",
	"LAURENT",
	"LAURO",
	"LAWRENCE",
	"LAWTON",
	"LAZARO",
	"LEANDRO",
	"LEDYARD",
	"LEEN",
	"LEFTERIS",
	"LEGNAGO",
	"LELLI",
	"LENOIR",
	"LEON",
	"LEONARD",
	"LEONARDI",
	"LEONARDO",
	"LESLIE",
	"LESO",
	"LEWISBERRY",
	"LEXINGTON",
	"LIANI",
	"LILLIAMO",
	"LINCOLN",
	"LINDENHURST",
	"LINO",
	"LISSAIOS",
	"LITTLE",
	"LIVONIA",
	"LODI",
	"LOMBARD",
	"LORENZO",
	"LOS",
	"LOUDONVILLE",
	"LOUISBURG",
	"LOUISVILLE",
	"LOYOLA",
	"LUBLIN",
	"LUC",
	"LUCA",
	"LUCIA",
	"LUCIANO",
	"LUCIO",
	"LUETHY",
	"LUIGI",
	"LUIS",
	"LUISA",
	"LUKASZ",
	"LULEA",
	"LURATI",
	"LYE",
	"MADISON",
	"MADISONVILLE",
	"MAERTENS",
	"MAGNIER",
	"MAHWAH",
	"MALCOLM",
	"MANAHAWKIN",
	"MANCHESTER",
	"MANFRED",
	"MANHATTAN",
	"MANOLO",
	"MANOS",
	"MANTORP",
	"MANUEL",
	"MARACAIBO",
	"MARC",
	"MARCEL",
	"MARCELLA",
	"MARCELLO",
	"MARCIN",
	"MARCO",
	"MARCOS",
	"MARCUS",
	"MARGATE",
	"MARIA",
	"MARIAN",
	"MARIE",
	"MARINA",
	"MARINO",
	"MARIO",
	"MARK",
	"MARKUS",
	"MARRA",
	"MARTIAL",
	"MARTIN",
	"MARTINEZ",
	"MARTINO",
	"MARYS",
	"MASSAPEQUA",
	"MASSIMILIANO",
	"MASSIMO",
	"MASTIC",
	"MATEO",
	"MATOUS",
	"MATS",
	"MATTEO",
	"MATTIA",
	"MAUMEE",
	"MAURIZIO",
	"MAURO",
	"MAXI",
	"MAYQUEL",
	"MC",
	"MCDONOUGH",
	"MCLOUGHLIN",
	"MCNAIRY",
	"MEDFORD",
	"MEGERDICH",
	"MEHMET",
	"MEINHARD",
	"MELBOURNE",
	"MEMPHIS",
	"MERIDA",
	"MERIDIAN",
	"MERRICK",
	"MERRILLVILLE",
	"MESA",
	"MEZIMESTI",
	"MIAMI",
	"MICHAEL",
	"MICHALIS",
	"MICHEL",
	"MICHELE",
	"MICK",
	"MIDDLETON",
	"MIDLAND",
	"MIDLOTHIAN",
	"MIDWEST",
	"MIGUEL",
	"MIKAEL",
	"MIKE",
	"MIKEL",
	"MILANO",
	"MILES",
	"MILFORD",
	"MILLS",
	"MILOCIN",
	"MILTIADIS",
	"MINDEN",
	"MINEOLA",
	"MINOZZI",
	"MIRA",
	"MIROSLAW",
	"MISHAWAKA",
	"MITCHELL",
	"MITROI",
	"MOMCILO",
	"MONROE",
	"MONTEBELLO",
	"MONTGOMERY",
	"MOR",
	"MORE",
	"MORENO",
	"MORGANTON",
	"MORRISTOWN",
	"MORUMBI",
	"MOUNT",
	"MOUNTAIN",
	"MOUNTVILLE",
	"MURALTO",
	"MUSKOGEE",
	"MUSTANG",
	"MUZAFFER",
	"MYERS",
	"NA",
	"NAGY",
	"NANDO",
	"NAPOLI",
	"NARINO",
	"NATCHEZ",
	"NATHALIE",
	"NATIONAL",
	"NAZZARENO",
	"NEILS",
	"NEW",
	"NEWARK",
	"NEWINGTON",
	"NEWTOWN",
	"NG",
	"NICHELINO",
	"NICK",
	"NICOLA",
	"NICOLAOS",
	"NIGEL",
	"NIKO",
	"NIKOLAOS",
	"NIKOLAS",
	"NIKOS",
	"NISOLLE",
	"NITRO",
	"NJURUNDA",
	"NO",
	"NOEL",
	"NOELLE",
	"NOKOMIS",
	"NORBERTO",
	"NORSBORG",
	"NORTH",
	"NORWICH",
	"NSW",
	"NUEVE",
	"NUNZIO",
	"NURI",
	"NWS",
	"NYLUND",
	"OAK",
	"OAKLEY",
	"OCEANA",
	"OCEANSIDE",
	"OESCH",
	"OGDENSBURG",
	"OGICA",
	"OH",
	"OKLAHOMA",
	"OLATHE",
	"OLE",
	"OLIVERO",
	"OLIVIER",
	"OLNEY",
	"OLYMPIA",
	"OLYPHANT",
	"OMAR",
	"ON",
	"ONANCOCK",
	"ONZ",
	"OO",
	"ORESTE",
	"ORIAGO",
	"ORLAND",
	"ORLANDO",
	"ORSA",
	"ORTIZ",
	"OSBY",
	"OSCAR",
	"OSLO",
	"OSNY",
	"OSWEGO",
	"OTAKI",
	"OTTAVIO",
	"OTTAWA",
	"OULU",
	"OVALLE",
	"OVERLAND",
	"OVIDIU",
	"OVIEDO",
	"PABLO",
	"PACIFICO",
	"PACO",
	"PALHOCA",
	"PALM",
	"PALMARES",
	"PANAGIOTIS",
	"PANAMA",
	"PANTELEIMON",
	"PAOLA",
	"PAOLO",
	"PARK",
	"PARMA",
	"PARMELE",
	"PARRAVICINI",
	"PARVIZ",
	"PASADENA",
	"PASCAL",
	"PASCHALIS",
	"PASCUAL",
	"PASQUALE",
	"PASQUALINO",
	"PASTOR",
	"PATRICK",
	"PATXI",
	"PAUL",
	"PAULINE",
	"PAULO",
	"PAWEL",
	"PEDRO",
	"PENANG",
	"PENDERGRASS",
	"PER",
	"PEREZ",
	"PERRY",
	"PERTH",
	"PERUGIA",
	"PETER",
	"PETERSBURG",
	"PHILADELPHIA",
	"PHILIP",
	"PHILIPPE",
	"PHOENIX",
	"PICKENS",
	"PICOT",
	"PIERANTONIO",
	"PIERGIORGIO",
	"PIERLUIGI",
	"PIERO",
	"PIERPAOLO",
	"PIERRE",
	"PIETRO",
	"PILAR",
	"PINE",
	"PINK",
	"PIOTR",
	"PISA",
	"PLAMEN",
	"PLANO",
	"PLEASANT",
	"POCOS",
	"POINT",
	"POLICHRONIS",
	"PONCA",
	"PORET",
	"PORT",
	"PORTLAND",
	"PORTSMOUTH",
	"PORUM",
	"POSSESSION",
	"POTOMAC",
	"POULSEN",
	"POWELL",
	"PRAIA",
	"PRAIRIE",
	"PRIMITIVO",
	"PRIMO",
	"PRINCETON",
	"PROBST",
	"PROMETHEUS",
	"PUYO",
	"PYEONG",
	"QUINTUS",
	"QUIQUE",
	"RADEK",
	"RADIO",
	"RADIOCLUB",
	"RAFA",
	"RAFAEL",
	"RALEIGH",
	"RAMIRO",
	"RAMON",
	"RANIERI",
	"RAOUL",
	"RAPIDS",
	"RAT",
	"RATON",
	"RAUL",
	"RAY",
	"RAYMORE",
	"READING",
	"RED",
	"REG",
	"REINE",
	"REMIGIUSZ",
	"REMO",
	"RENATO",
	"RENE",
	"RENO",
	"RHOME",
	"RICARDO",
	"RICCARDO",
	"RICHARD",
	"RICK",
	"RIDGE",
	"RIGHINI",
	"RINO",
	"RIO",
	"RISPOLI",
	"RIVER",
	"RIZZO",
	"ROBERT",
	"ROBERTO",
	"ROBIN",
	"ROCCO",
	"ROCHELLE",
	"ROCHESTER",
	"ROCIO",
	"ROCK",
	"ROCKY",
	"ROD",
	"RODERICK",
	"RODOLFO",
	"RODRIGUE",
	"ROGELIO",
	"ROGER",
	"ROGERSVILLE",
	"ROLF",
	"ROMA",
	"ROMAN",
	"ROMANO",
	"ROMARIC",
	"ROMEO",
	"RON",
	"RONALD",
	"RONG",
	"RONNIE",
	"ROSA",
	"ROSARIO",
	"ROSOLINO",
	"ROTTERDAM",
	"ROVERETO",
	"ROY",
	"RUBEN",
	"RUTI",
	"RYSZARD",
	"SABATER",
	"SABET",
	"SAFAK",
	"SAINT",
	"SALEMI",
	"SALVADO",
	"SALVADOR",
	"SALVATORE",
	"SAMUELE",
	"SAN",
	"SANCHEZ",
	"SANDRO",
	"SANOK",
	"SANTA",
	"SANTE",
	"SANTIAGO",
	"SANTO",
	"SANTOS",
	"SAO",
	"SARASOTA",
	"SASKATOON",
	"SASTAMALA",
	"SAURO",
	"SAVANNH",
	"SAVYTSKYI",
	"SCHMID",
	"SCHOERG",
	"SCOTTSBURG",
	"SCUOLA",
	"SEABROOK",
	"SEAFORD",
	"SEAN",
	"SEBASTIAN",
	"SEBASTIANO",
	"SECONDARIA",
	"SECONDO",
	"SEDAT",
	"SEOUL",
	"SERAFIN",
	"SERENA",
	"SERGIO",
	"SERGIPE",
	"SERTAC",
	"SETAUKET",
	"SETTLE",
	"SEVERINO",
	"SEYMOUR",
	"SH",
	"SHAH",
	"SHAKOPEE",
	"SHAWNEE",
	"SHEFFIELD",
	"SHENZHEN",
	"SHERBROOKE",
	"SHILOH",
	"SHIRLEY",
	"SILENT",
	"SILVANO",
	"SILVIO",
	"SIMON",
	"SIMPLICIO",
	"SIMSBURY",
	"SLU",
	"SNEADS",
	"SNORRE",
	"SOERENSEN",
	"SOFIANOS",
	"SOMERS",
	"SONOMA",
	"SORBOLO",
	"SOUTH",
	"SOUTHGATE",
	"SP",
	"SPARKS",
	"SPEZIA",
	"SPIROS",
	"SPORS",
	"SPRINGFIELD",
	"SPRINGS",
	"SPYROS",
	"SQUARE",
	"ST",
	"STAFFANSTORP",
	"STAMATIS",
	"STATION",
	"STAVROS",
	"STEFAN",
	"STEFANO",
	"STEPHANE",
	"STEPHEN",
	"STEVANI",
	"STEVE",
	"STEVEN",
	"STEWART",
	"STIIG",
	"STORTFORD",
	"STRATFORD",
	"STUART",
	"SUL",
	"SUMMERVILLE",
	"SUNDSBRUK",
	"SUNDSVALL",
	"SURFSIDE",
	"SUSAN",
	"SUSO",
	"SUWON",
	"SVARTSJOE",
	"SVENDBORG",
	"SWANSEA",
	"SYDNEY",
	"SZCZECIN",
	"SZOKE",
	"TAGLIACOZZO",
	"TAIPEI",
	"TALLASEN",
	"TAMAROA",
	"TAMER",
	"TAMPERE",
	"TAURANGA",
	"TED",
	"TEKIN",
	"TEMPE",
	"TEMPERANCE",
	"TEODORO",
	"TERENCE",
	"TERNI",
	"TERRY",
	"TERVO",
	"TESTE",
	"THEODOROS",
	"THERESE",
	"THIELLLS",
	"THIERRY",
	"THOMANN",
	"THOMAS",
	"TILICI",
	"TIM",
	"TIMOTHY",
	"TITO",
	"TITUSVILLE",
	"TIZAPAN",
	"TOLGA",
	"TOM",
	"TOMAS",
	"TOMASZ",
	"TOMMASO",
	"TOMS",
	"TONAMI",
	"TONI",
	"TONINO",
	"TONY",
	"TOP",
	"TORE",
	"TORI",
	"TORINO",
	"TORNIO",
	"TORONTO",
	"TOWNSHIP",
	"TRAIL",
	"TREMONT",
	"TREMOSNA",
	"TREVOR",
	"TRIANTAFYLLOS",
	"TRISTAN",
	"TRUGLIA",
	"TRUMBULL",
	"TRUQUET",
	"TUCSON",
	"TULLIO",
	"TULSA",
	"TURITTO",
	"TWP",
	"TYRESO",
	"ULSAN",
	"UMBERTO",
	"UNIDAD",
	"UPPSALA",
	"URCE",
	"URE",
	"URLE",
	"URPLACOM",
	"USER",
	"VA",
	"VALAIS",
	"VALENCIA",
	"VALENTIN",
	"VALENTINO",
	"VALERIO",
	"VALLAMAND",
	"VALLSTA",
	"VANCOUVER",
	"VANIER",
	"VANJA",
	"VANTAA",
	"VARENNES",
	"VARGARDA",
	"VASILEIOS",
	"VASILIS",
	"VAXHOLM",
	"VEDAT",
	"VEGAS",
	"VENICE",
	"VENTIMIGLIA",
	"VERON",
	"VERONA",
	"VESSKO",
	"VHF",
	"VIC",
	"VICENTE",
	"VICTOR",
	"VINCENZO",
	"VIRGILIO",
	"VIRNA",
	"VISALIA",
	"VISTA",
	"VITO",
	"VITTORIO",
	"VLADIMIR",
	"VLADISLAVS",
	"VOLKAN",
	"VOLTERRA",
	"VORDINGBORG",
	"VUITEL",
	"WALLINGFORD",
	"WALNUT",
	"WALTER",
	"WAPPINGER",
	"WARREN",
	"WARSAW",
	"WASHINGTON",
	"WATERTOWN",
	"WATERVLIET",
	"WAUKESHA",
	"WAYCASSY",
	"WAYNE",
	"WEBB",
	"WEST",
	"WESTBURY",
	"WESTON",
	"WHITE",
	"WHITESTONE",
	"WILKESBORO",
	"WILLIAM",
	"WILLIAMSTOWN",
	"WINDHAM",
	"WINDSOR",
	"WINTER",
	"WINTERPORT",
	"WLADIMIRO",
	"WLODZIMIERZ",
	"WOJCIECH",
	"WOODBRIDGE",
	"WOODLAND",
	"WOODSTOCK",
	"WYEE",
	"WYLIE",
	"XABIER",
	"XAVIER",
	"YANG",
	"YANNICK",
	"YAPHANK",
	"YONKERS",
	"YUMA",
	"YURECUARO",
	"YVAN",
	"YVES",
	"ZARAGOZA",
	"ZBIGNIEW",
	"ZCARLOS",
	"ZDENKO",
	"ZDRAVKO",
	"ZEVIO",
	"ZH",
}

var upperCaseWords = []string{
	"AA",
	"ACCART",
	"ACT",
	"AFTT",
	"AFZ",
	"AIRE",
	"AJT",
	"AK",
	"ALA",
	"ANFR",
	"APRS",
	"AQ",
	"AR",
	"ARAC",
	"ARACA",
	"AREX",
	"ARG",
	"ARRL",
	"ARRS",
	"ARS",
	"ARSGM",
	"AS",
	"ASORL",
	"ASRR",
	"ASTT",
	"ATL",
	"AU",
	"AWR",
	"AXPE",
	"AZ",
	"BC",
	"BGD",
	"BL",
	"BRU",
	"BSG",
	"BXE",
	"CARG",
	"CAS",
	"CDARC",
	"CIS",
	"CK",
	"CM",
	"CMARS",
	"CN",
	"CRD",
	"DARC",
	"DARES",
	"DC",
	"DFW",
	"DK",
	"DLR",
	"DLZA",
	"DMR",
	"DOK",
	"DRC",
	"DSTAR",
	"DSTAT",
	"DT",
	"EDF",
	"EDR",
	"ELDA",
	"ELZAS",
	"EMA",
	"EMC",
	"EMRE",
	"ERA",
	"FBI",
	"FFBG",
	"FM",
	"FMLOG",
	"FRA",
	"FRAG",
	"FRAT",
	"FRO",
	"FUAT",
	"GA",
	"GBHARC",
	"GBN",
	"GE",
	"GFB",
	"GM",
	"GP",
	"GT",
	"GV",
	"GW",
	"HFP",
	"HK",
	"HKG",
	"HKI",
	"HOLO",
	"HORK",
	"HP",
	"HRO",
	"HRS",
	"HS",
	"ICOM",
	"ID",
	"IG",
	"IGBF",
	"IL",
	"IPHA",
	"IPR",
	"JAM",
	"JB",
	"JEZ",
	"JF",
	"JH",
	"JM",
	"JMARC",
	"JOTA",
	"KBH",
	"KLN",
	"KPGK",
	"KS",
	"KTK",
	"LARU",
	"LLV",
	"LRRA",
	"MAD",
	"MBG",
	"MERT",
	"MLARS",
	"MLB",
	"MM",
	"MOKPO",
	"NASA",
	"NC",
	"NCD",
	"NDR",
	"NJECT",
	"NOZ",
	"NRRL",
	"NT",
	"NV",
	"NY",
	"NYC",
	"OE",
	"OEARC",
	"OEVSV",
	"OT",
	"OV",
	"PA",
	"PBL",
	"PCRE",
	"PZK",
	"RA",
	"RAAI",
	"RACCW",
	"RACW",
	"RAGLAN",
	"RC",
	"RCA",
	"RCL",
	"RCV",
	"REC",
	"RECEP",
	"REEC",
	"REM",
	"RSF",
	"RSGB",
	"RST",
	"RWEG",
	"RWTH",
	"SAR",
	"SAS",
	"SBARC",
	"SC",
	"SCH",
	"SCS",
	"SES",
	"SJ",
	"SK",
	"SL",
	"SM",
	"SNARS",
	"SORE",
	"SSA",
	"SSRA",
	"START",
	"SV",
	"SW",
	"SZ",
	"SZCZEPAN",
	"TAA",
	"TC",
	"TDOTA",
	"TG",
	"THW",
	"TJ",
	"TV",
	"TX",
	"UBA",
	"UK",
	"UR",
	"URC",
	"URCAT",
	"URO",
	"URP",
	"USKA",
	"VCDRC",
	"VD",
	"WG",
	"WIMO",
	"WLBR",
	"WX",
	"YOTA",
}
