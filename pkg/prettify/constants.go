package prettify

var countries = map[string]string{
	"Ascension Island":         "\U0001f1e6\U0001f1e8",
	"Andorra":                  "\U0001f1e6\U0001f1e9",
	"United Arab Emirates":     "\U0001f1e6\U0001f1ea",
	"Afghanistan":              "\U0001f1e6\U0001f1eb",
	"Antigua And Barbuda":      "\U0001f1e6\U0001f1ec",
	"Anguilla":                 "\U0001f1e6\U0001f1ee",
	"Albania":                  "\U0001f1e6\U0001f1f1",
	"Armenia":                  "\U0001f1e6\U0001f1f2",
	"Angola":                   "\U0001f1e6\U0001f1f4",
	"Antarctica":               "\U0001f1e6\U0001f1f6",
	"Argentina":                "\U0001f1e6\U0001f1f7",
	"American Samoa":           "\U0001f1e6\U0001f1f8",
	"Austria":                  "\U0001f1e6\U0001f1f9",
	"Australia":                "\U0001f1e6\U0001f1fa",
	"Aruba":                    "\U0001f1e6\U0001f1fc",
	"Aland Islands":            "\U0001f1e6\U0001f1fd",
	"Azerbaijan":               "\U0001f1e6\U0001f1ff",
	"Bosnia And Herzegovina":   "\U0001f1e7\U0001f1e6",
	"Barbados":                 "\U0001f1e7\U0001f1e7",
	"Bangladesh":               "\U0001f1e7\U0001f1e9",
	"Belgium":                  "\U0001f1e7\U0001f1ea",
	"Burkina Faso":             "\U0001f1e7\U0001f1eb",
	"Bulgaria":                 "\U0001f1e7\U0001f1ec",
	"Bahrain":                  "\U0001f1e7\U0001f1ed",
	"Burundi":                  "\U0001f1e7\U0001f1ee",
	"Benin":                    "\U0001f1e7\U0001f1ef",
	"St Barthelemy":            "\U0001f1e7\U0001f1f1",
	"Bermuda":                  "\U0001f1e7\U0001f1f2",
	"Brunei":                   "\U0001f1e7\U0001f1f3",
	"Bolivia":                  "\U0001f1e7\U0001f1f4",
	"Caribbean Netherlands":    "\U0001f1e7\U0001f1f6",
	"Brazil":                   "\U0001f1e7\U0001f1f7",
	"Bahamas":                  "\U0001f1e7\U0001f1f8",
	"Bhutan":                   "\U0001f1e7\U0001f1f9",
	"Bouvet Island":            "\U0001f1e7\U0001f1fb",
	"Botswana":                 "\U0001f1e7\U0001f1fc",
	"Belarus":                  "\U0001f1e7\U0001f1fe",
	"Belize":                   "\U0001f1e7\U0001f1ff",
	"Canada":                   "\U0001f1e8\U0001f1e6",
	"Cocos Keeling Islands":    "\U0001f1e8\U0001f1e8",
	"Congo Kinshasa":           "\U0001f1e8\U0001f1e9",
	"Central African Republic": "\U0001f1e8\U0001f1eb",
	"Congo Brazzaville":        "\U0001f1e8\U0001f1ec",
	"Switzerland":              "\U0001f1e8\U0001f1ed",
	"Cote DIvoire":             "\U0001f1e8\U0001f1ee",
	"Cook Islands":             "\U0001f1e8\U0001f1f0",
	"Chile":                    "\U0001f1e8\U0001f1f1",
	"Cameroon":                 "\U0001f1e8\U0001f1f2",
	"China":                    "\U0001f1e8\U0001f1f3",
	"Colombia":                 "\U0001f1e8\U0001f1f4",
	"Clipperton Island":        "\U0001f1e8\U0001f1f5",
	"Costa Rica":               "\U0001f1e8\U0001f1f7",
	"Cuba":                     "\U0001f1e8\U0001f1fa",
	"Cape Verde":               "\U0001f1e8\U0001f1fb",
	"Curacao":                  "\U0001f1e8\U0001f1fc",
	"Christmas Island":         "\U0001f1e8\U0001f1fd",
	"Cyprus":                   "\U0001f1e8\U0001f1fe",
	"Czechia":                  "\U0001f1e8\U0001f1ff",
	"Germany":                  "\U0001f1e9\U0001f1ea",
	"Diego Garcia":             "\U0001f1e9\U0001f1ec",
	"Djibouti":                 "\U0001f1e9\U0001f1ef",
	"Denmark":                  "\U0001f1e9\U0001f1f0",
	"Dominica":                 "\U0001f1e9\U0001f1f2",
	"Dominican Republic":       "\U0001f1e9\U0001f1f4",
	"Algeria":                  "\U0001f1e9\U0001f1ff",
	"Ceuta And Melilla":        "\U0001f1ea\U0001f1e6",
	"Ecuador":                  "\U0001f1ea\U0001f1e8",
	"Estonia":                  "\U0001f1ea\U0001f1ea",
	"Egypt":                    "\U0001f1ea\U0001f1ec",
	"Western Sahara":           "\U0001f1ea\U0001f1ed",
	"Eritrea":                  "\U0001f1ea\U0001f1f7",
	"Spain":                    "\U0001f1ea\U0001f1f8",
	"Ethiopia":                 "\U0001f1ea\U0001f1f9",
	"European Union":           "\U0001f1ea\U0001f1fa",
	"Finland":                  "\U0001f1eb\U0001f1ee",
	"Fiji":                     "\U0001f1eb\U0001f1ef",
	"Falkland Islands":         "\U0001f1eb\U0001f1f0",
	"Micronesia":               "\U0001f1eb\U0001f1f2",
	"Faroe Islands":            "\U0001f1eb\U0001f1f4",
	"France":                   "\U0001f1eb\U0001f1f7",
	"Gabon":                    "\U0001f1ec\U0001f1e6",
	"United Kingdom":           "\U0001f1ec\U0001f1e7",
	"Grenada":                  "\U0001f1ec\U0001f1e9",
	"Georgia":                  "\U0001f1ec\U0001f1ea",
	"French Guiana":            "\U0001f1ec\U0001f1eb",
	"Guernsey":                 "\U0001f1ec\U0001f1ec",
	"Ghana":                    "\U0001f1ec\U0001f1ed",
	"Gibraltar":                "\U0001f1ec\U0001f1ee",
	"Greenland":                "\U0001f1ec\U0001f1f1",
	"Gambia":                   "\U0001f1ec\U0001f1f2",
	"Guinea":                   "\U0001f1ec\U0001f1f3",
	"Guadeloupe":               "\U0001f1ec\U0001f1f5",
	"Equatorial Guinea":        "\U0001f1ec\U0001f1f6",
	"Greece":                   "\U0001f1ec\U0001f1f7",
	"South Georgia And South Sandwich Islands": "\U0001f1ec\U0001f1f8",
	"Guatemala":                      "\U0001f1ec\U0001f1f9",
	"Guam":                           "\U0001f1ec\U0001f1fa",
	"Guinea Bissau":                  "\U0001f1ec\U0001f1fc",
	"Guyana":                         "\U0001f1ec\U0001f1fe",
	"Hong Kong Sar China":            "\U0001f1ed\U0001f1f0",
	"Heard And Mcdonald Islands":     "\U0001f1ed\U0001f1f2",
	"Honduras":                       "\U0001f1ed\U0001f1f3",
	"Croatia":                        "\U0001f1ed\U0001f1f7",
	"Haiti":                          "\U0001f1ed\U0001f1f9",
	"Hungary":                        "\U0001f1ed\U0001f1fa",
	"Canary Islands":                 "\U0001f1ee\U0001f1e8",
	"Indonesia":                      "\U0001f1ee\U0001f1e9",
	"Ireland":                        "\U0001f1ee\U0001f1ea",
	"Israel":                         "\U0001f1ee\U0001f1f1",
	"Isle Of Man":                    "\U0001f1ee\U0001f1f2",
	"India":                          "\U0001f1ee\U0001f1f3",
	"British Indian Ocean Territory": "\U0001f1ee\U0001f1f4",
	"Iraq":                           "\U0001f1ee\U0001f1f6",
	"Iran":                           "\U0001f1ee\U0001f1f7",
	"Iceland":                        "\U0001f1ee\U0001f1f8",
	"Italy":                          "\U0001f1ee\U0001f1f9",
	"Jersey":                         "\U0001f1ef\U0001f1ea",
	"Jamaica":                        "\U0001f1ef\U0001f1f2",
	"Jordan":                         "\U0001f1ef\U0001f1f4",
	"Japan":                          "\U0001f1ef\U0001f1f5",
	"Kenya":                          "\U0001f1f0\U0001f1ea",
	"Kyrgyzstan":                     "\U0001f1f0\U0001f1ec",
	"Cambodia":                       "\U0001f1f0\U0001f1ed",
	"Kiribati":                       "\U0001f1f0\U0001f1ee",
	"Comoros":                        "\U0001f1f0\U0001f1f2",
	"St Kitts And Nevis":             "\U0001f1f0\U0001f1f3",
	"North Korea":                    "\U0001f1f0\U0001f1f5",
	"South Korea":                    "\U0001f1f0\U0001f1f7",
	"Kuwait":                         "\U0001f1f0\U0001f1fc",
	"Cayman Islands":                 "\U0001f1f0\U0001f1fe",
	"Kazakhstan":                     "\U0001f1f0\U0001f1ff",
	"Laos":                           "\U0001f1f1\U0001f1e6",
	"Lebanon":                        "\U0001f1f1\U0001f1e7",
	"St Lucia":                       "\U0001f1f1\U0001f1e8",
	"Liechtenstein":                  "\U0001f1f1\U0001f1ee",
	"Sri Lanka":                      "\U0001f1f1\U0001f1f0",
	"Liberia":                        "\U0001f1f1\U0001f1f7",
	"Lesotho":                        "\U0001f1f1\U0001f1f8",
	"Lithuania":                      "\U0001f1f1\U0001f1f9",
	"Luxembourg":                     "\U0001f1f1\U0001f1fa",
	"Latvia":                         "\U0001f1f1\U0001f1fb",
	"Libya":                          "\U0001f1f1\U0001f1fe",
	"Morocco":                        "\U0001f1f2\U0001f1e6",
	"Monaco":                         "\U0001f1f2\U0001f1e8",
	"Moldova":                        "\U0001f1f2\U0001f1e9",
	"Montenegro":                     "\U0001f1f2\U0001f1ea",
	"St Martin":                      "\U0001f1f2\U0001f1eb",
	"Madagascar":                     "\U0001f1f2\U0001f1ec",
	"Marshall Islands":               "\U0001f1f2\U0001f1ed",
	"North Macedonia":                "\U0001f1f2\U0001f1f0",
	"Mali":                           "\U0001f1f2\U0001f1f1",
	"Myanmar Burma":                  "\U0001f1f2\U0001f1f2",
	"Mongolia":                       "\U0001f1f2\U0001f1f3",
	"Macao Sar China":                "\U0001f1f2\U0001f1f4",
	"Northern Mariana Islands":       "\U0001f1f2\U0001f1f5",
	"Martinique":                     "\U0001f1f2\U0001f1f6",
	"Mauritania":                     "\U0001f1f2\U0001f1f7",
	"Montserrat":                     "\U0001f1f2\U0001f1f8",
	"Malta":                          "\U0001f1f2\U0001f1f9",
	"Mauritius":                      "\U0001f1f2\U0001f1fa",
	"Maldives":                       "\U0001f1f2\U0001f1fb",
	"Malawi":                         "\U0001f1f2\U0001f1fc",
	"Mexico":                         "\U0001f1f2\U0001f1fd",
	"Malaysia":                       "\U0001f1f2\U0001f1fe",
	"Mozambique":                     "\U0001f1f2\U0001f1ff",
	"Namibia":                        "\U0001f1f3\U0001f1e6",
	"New Caledonia":                  "\U0001f1f3\U0001f1e8",
	"Niger":                          "\U0001f1f3\U0001f1ea",
	"Norfolk Island":                 "\U0001f1f3\U0001f1eb",
	"Nigeria":                        "\U0001f1f3\U0001f1ec",
	"Nicaragua":                      "\U0001f1f3\U0001f1ee",
	"Netherlands":                    "\U0001f1f3\U0001f1f1",
	"Norway":                         "\U0001f1f3\U0001f1f4",
	"Nepal":                          "\U0001f1f3\U0001f1f5",
	"Nauru":                          "\U0001f1f3\U0001f1f7",
	"Niue":                           "\U0001f1f3\U0001f1fa",
	"New Zealand":                    "\U0001f1f3\U0001f1ff",
	"Oman":                           "\U0001f1f4\U0001f1f2",
	"Panama":                         "\U0001f1f5\U0001f1e6",
	"Peru":                           "\U0001f1f5\U0001f1ea",
	"French Polynesia":               "\U0001f1f5\U0001f1eb",
	"Papua New Guinea":               "\U0001f1f5\U0001f1ec",
	"Philippines":                    "\U0001f1f5\U0001f1ed",
	"Pakistan":                       "\U0001f1f5\U0001f1f0",
	"Poland":                         "\U0001f1f5\U0001f1f1",
	"St Pierre And Miquelon":         "\U0001f1f5\U0001f1f2",
	"Pitcairn Islands":               "\U0001f1f5\U0001f1f3",
	"Puerto Rico":                    "\U0001f1f5\U0001f1f7",
	"Palestinian Territories":        "\U0001f1f5\U0001f1f8",
	"Portugal":                       "\U0001f1f5\U0001f1f9",
	"Palau":                          "\U0001f1f5\U0001f1fc",
	"Paraguay":                       "\U0001f1f5\U0001f1fe",
	"Qatar":                          "\U0001f1f6\U0001f1e6",
	"Reunion":                        "\U0001f1f7\U0001f1ea",
	"Romania":                        "\U0001f1f7\U0001f1f4",
	"Serbia":                         "\U0001f1f7\U0001f1f8",
	"Russia":                         "\U0001f1f7\U0001f1fa",
	"Rwanda":                         "\U0001f1f7\U0001f1fc",
	"Saudi Arabia":                   "\U0001f1f8\U0001f1e6",
	"Solomon Islands":                "\U0001f1f8\U0001f1e7",
	"Seychelles":                     "\U0001f1f8\U0001f1e8",
	"Sudan":                          "\U0001f1f8\U0001f1e9",
	"Sweden":                         "\U0001f1f8\U0001f1ea",
	"Singapore":                      "\U0001f1f8\U0001f1ec",
	"St Helena":                      "\U0001f1f8\U0001f1ed",
	"Slovenia":                       "\U0001f1f8\U0001f1ee",
	"Svalbard And Jan Mayen":         "\U0001f1f8\U0001f1ef",
	"Slovakia":                       "\U0001f1f8\U0001f1f0",
	"Sierra Leone":                   "\U0001f1f8\U0001f1f1",
	"San Marino":                     "\U0001f1f8\U0001f1f2",
	"Senegal":                        "\U0001f1f8\U0001f1f3",
	"Somalia":                        "\U0001f1f8\U0001f1f4",
	"Suriname":                       "\U0001f1f8\U0001f1f7",
	"South Sudan":                    "\U0001f1f8\U0001f1f8",
	"Sao Tome And Principe":          "\U0001f1f8\U0001f1f9",
	"El Salvador":                    "\U0001f1f8\U0001f1fb",
	"Sint Maarten":                   "\U0001f1f8\U0001f1fd",
	"Syria":                          "\U0001f1f8\U0001f1fe",
	"Eswatini":                       "\U0001f1f8\U0001f1ff",
	"Tristan Da Cunha":               "\U0001f1f9\U0001f1e6",
	"Turks And Caicos Islands":       "\U0001f1f9\U0001f1e8",
	"Chad":                           "\U0001f1f9\U0001f1e9",
	"French Southern Territories":    "\U0001f1f9\U0001f1eb",
	"Togo":                           "\U0001f1f9\U0001f1ec",
	"Thailand":                       "\U0001f1f9\U0001f1ed",
	"Tajikistan":                     "\U0001f1f9\U0001f1ef",
	"Tokelau":                        "\U0001f1f9\U0001f1f0",
	"Timor Leste":                    "\U0001f1f9\U0001f1f1",
	"Turkmenistan":                   "\U0001f1f9\U0001f1f2",
	"Tunisia":                        "\U0001f1f9\U0001f1f3",
	"Tonga":                          "\U0001f1f9\U0001f1f4",
	"Turkey":                         "\U0001f1f9\U0001f1f7",
	"Trinidad And Tobago":            "\U0001f1f9\U0001f1f9",
	"Tuvalu":                         "\U0001f1f9\U0001f1fb",
	"Taiwan":                         "\U0001f1f9\U0001f1fc",
	"Tanzania":                       "\U0001f1f9\U0001f1ff",
	"Ukraine":                        "\U0001f1fa\U0001f1e6",
	"Uganda":                         "\U0001f1fa\U0001f1ec",
	"Us Outlying Islands":            "\U0001f1fa\U0001f1f2",
	"United Nations":                 "\U0001f1fa\U0001f1f3",
	"USA":                            "\U0001f1fa\U0001f1f8",
	"Uruguay":                        "\U0001f1fa\U0001f1fe",
	"Uzbekistan":                     "\U0001f1fa\U0001f1ff",
	"Vatican City":                   "\U0001f1fb\U0001f1e6",
	"St Vincent And Grenadines":      "\U0001f1fb\U0001f1e8",
	"Venezuela":                      "\U0001f1fb\U0001f1ea",
	"British Virgin Islands":         "\U0001f1fb\U0001f1ec",
	"Us Virgin Islands":              "\U0001f1fb\U0001f1ee",
	"Vietnam":                        "\U0001f1fb\U0001f1f3",
	"Vanuatu":                        "\U0001f1fb\U0001f1fa",
	"Wallis And Futuna":              "\U0001f1fc\U0001f1eb",
	"Samoa":                          "\U0001f1fc\U0001f1f8",
	"Kosovo":                         "\U0001f1fd\U0001f1f0",
	"Yemen":                          "\U0001f1fe\U0001f1ea",
	"Mayotte":                        "\U0001f1fe\U0001f1f9",
	"South Africa":                   "\U0001f1ff\U0001f1e6",
	"Zambia":                         "\U0001f1ff\U0001f1f2",
	"Zimbabwe":                       "\U0001f1ff\U0001f1fc",
}
