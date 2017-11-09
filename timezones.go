package ics

// CommonLocaleDataRepositoryTranslator maps between the CLDR timezone to the first IANA territory in said zone.
// See http://unicode.org/repos/cldr/trunk/common/supplemental/windowsZones.xml;
var CommonLocaleDataRepositoryTranslator = map[string]string{
	"(UTC-12:00) International Date Line West":                      "Etc/GMT+12",
	"(UTC-11:00) Coordinated Universal Time-11":                     "Etc/GMT+11",
	"(UTC-10:00) Aleutian Islands":                                  "America/Adak",
	"(UTC-10:00) Hawaii":                                            "Pacific/Honolulu",
	"(UTC-09:30) Marquesas Islands":                                 "Pacific/Marquesas",
	"(UTC-09:00) Alaska":                                            "America/Anchorage",
	"(UTC-09:00) Coordinated Universal Time-09":                     "Etc/GMT+9",
	"(UTC-08:00) Baja California":                                   "America/Tijuana",
	"(UTC-08:00) Coordinated Universal Time-08":                     "Etc/GMT+8",
	"(UTC-08:00) Pacific Time (US & Canada)":                        "America/Los_Angeles",
	"(UTC-07:00) Arizona":                                           "America/Phoenix",
	"(UTC-07:00) Chihuahua, La Paz, Mazatlan":                       "America/Chihuahua",
	"(UTC-07:00) Mountain Time (US & Canada)":                       "America/Denver",
	"(UTC-06:00) Central America":                                   "America/Guatemala",
	"(UTC-06:00) Central Time (US & Canada)":                        "America/Chicago",
	"(UTC-06:00) Easter Island":                                     "Pacific/Easter",
	"(UTC-06:00) Guadalajara, Mexico City, Monterrey":               "America/Mexico_City",
	"(UTC-06:00) Saskatchewan":                                      "America/Regina",
	"(UTC-05:00) Bogota, Lima, Quito, Rio Branco":                   "America/Bogota",
	"(UTC-05:00) Chetumal":                                          "America/Cancun",
	"(UTC-05:00) Eastern Time (US & Canada)":                        "America/New_York",
	"(UTC-05:00) Haiti":                                             "America/Port-au-Prince",
	"(UTC-05:00) Havana":                                            "America/Havana",
	"(UTC-05:00) Indiana (East)":                                    "America/Indianapolis",
	"(UTC-04:00) Asuncion":                                          "America/Asuncion",
	"(UTC-04:00) Atlantic Time (Canada)":                            "America/Halifax",
	"(UTC-04:00) Caracas":                                           "America/Caracas",
	"(UTC-04:00) Cuiaba":                                            "America/Cuiaba",
	"(UTC-04:00) Georgetown, La Paz, Manaus, San Juan":              "America/La_Paz",
	"(UTC-04:00) Santiago":                                          "America/Santiago",
	"(UTC-04:00) Turks and Caicos":                                  "America/Grand_Turk",
	"(UTC-03:30) Newfoundland":                                      "America/St_Johns",
	"(UTC-03:00) Araguaina":                                         "America/Araguaina",
	"(UTC-03:00) Brasilia":                                          "America/Sao_Paulo",
	"(UTC-03:00) Cayenne, Fortaleza":                                "America/Cayenne",
	"(UTC-03:00) City of Buenos Aires":                              "America/Buenos_Aires",
	"(UTC-03:00) Greenland":                                         "America/Godthab",
	"(UTC-03:00) Montevideo":                                        "America/Montevideo",
	"(UTC-03:00) Saint Pierre and Miquelon":                         "America/Miquelon",
	"(UTC-03:00) Salvador":                                          "America/Bahia",
	"(UTC-02:00) Coordinated Universal Time-02":                     "Etc/GMT+2",
	"(UTC-01:00) Azores":                                            "Atlantic/Azores",
	"(UTC-01:00) Cabo Verde Is.":                                    "Atlantic/Cape_Verde",
	"(UTC) Coordinated Universal Time":                              "Etc/GMT",
	"(UTC+00:00) Casablanca":                                        "Africa/Casablanca",
	"(UTC+00:00) Dublin, Edinburgh, Lisbon, London":                 "Europe/London",
	"(UTC+00:00) Monrovia, Reykjavik":                               "Atlantic/Reykjavik",
	"(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna":  "Europe/Berlin",
	"(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague": "Europe/Budapest",
	"(UTC+01:00) Brussels, Copenhagen, Madrid, Paris":               "Europe/Paris",
	"(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb":                  "Europe/Warsaw",
	"(UTC+01:00) West Central Africa":                               "Africa/Lagos",
	"(UTC+01:00) Windhoek":                                          "Africa/Windhoek",
	"(UTC+02:00) Amman":                                             "Asia/Amman",
	"(UTC+02:00) Athens, Bucharest":                                 "Europe/Bucharest",
	"(UTC+02:00) Beirut":                                            "Asia/Beirut",
	"(UTC+02:00) Cairo":                                             "Africa/Cairo",
	"(UTC+02:00) Chisinau":                                          "Europe/Chisinau",
	"(UTC+02:00) Damascus":                                          "Asia/Damascus",
	"(UTC+02:00) Gaza, Hebron":                                      "Asia/Hebron",
	"(UTC+02:00) Harare, Pretoria":                                  "Africa/Johannesburg",
	"(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius":     "Europe/Kiev",
	"(UTC+02:00) Jerusalem":                                         "Asia/Jerusalem",
	"(UTC+02:00) Kaliningrad":                                       "Europe/Kaliningrad",
	"(UTC+02:00) Tripoli":                                           "Africa/Tripoli",
	"(UTC+03:00) Baghdad":                                           "Asia/Baghdad",
	"(UTC+03:00) Istanbul":                                          "Europe/Istanbul",
	"(UTC+03:00) Kuwait, Riyadh":                                    "Asia/Riyadh",
	"(UTC+03:00) Minsk":                                             "Europe/Minsk",
	"(UTC+03:00) Moscow, St. Petersburg, Volgograd":                 "Europe/Moscow",
	"(UTC+03:00) Nairobi":                                           "Africa/Nairobi",
	"(UTC+03:30) Tehran":                                            "Asia/Tehran",
	"(UTC+04:00) Abu Dhabi, Muscat":                                 "Asia/Dubai",
	"(UTC+04:00) Astrakhan, Ulyanovsk":                              "Europe/Astrakhan",
	"(UTC+04:00) Baku":                                              "Asia/Baku",
	"(UTC+04:00) Izhevsk, Samara":                                   "Europe/Samara",
	"(UTC+04:00) Port Louis":                                        "Indian/Mauritius",
	"(UTC+04:00) Tbilisi":                                           "Asia/Tbilisi",
	"(UTC+04:00) Yerevan":                                           "Asia/Yerevan",
	"(UTC+04:30) Kabul":                                             "Asia/Kabul",
	"(UTC+05:00) Ashgabat, Tashkent":                                "Asia/Tashkent",
	"(UTC+05:00) Ekaterinburg":                                      "Asia/Yekaterinburg",
	"(UTC+05:00) Islamabad, Karachi":                                "Asia/Karachi",
	"(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi":               "Asia/Calcutta",
	"(UTC+05:30) Sri Jayawardenepura":                               "Asia/Colombo",
	"(UTC+05:45) Kathmandu":                                         "Asia/Katmandu",
	"(UTC+06:00) Astana":                                            "Asia/Almaty",
	"(UTC+06:00) Dhaka":                                             "Asia/Dhaka",
	"(UTC+06:00) Omsk":                                              "Asia/Omsk",
	"(UTC+06:30) Yangon (Rangoon)":                                  "Asia/Rangoon",
	"(UTC+07:00) Bangkok, Hanoi, Jakarta":                           "Asia/Bangkok",
	"(UTC+07:00) Barnaul, Gorno-Altaysk":                            "Asia/Barnaul",
	"(UTC+07:00) Hovd":                                              "Asia/Hovd",
	"(UTC+07:00) Krasnoyarsk":                                       "Asia/Krasnoyarsk",
	"(UTC+07:00) Novosibirsk":                                       "Asia/Novosibirsk",
	"(UTC+07:00) Tomsk":                                             "Asia/Tomsk",
	"(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi":             "Asia/Shanghai",
	"(UTC+08:00) Irkutsk":                                           "Asia/Irkutsk",
	"(UTC+08:00) Kuala Lumpur, Singapore":                           "Asia/Singapore",
	"(UTC+08:00) Perth":                                             "Australia/Perth",
	"(UTC+08:00) Taipei":                                            "Asia/Taipei",
	"(UTC+08:00) Ulaanbaatar":                                       "Asia/Ulaanbaatar",
	"(UTC+08:30) Pyongyang":                                         "Asia/Pyongyang",
	"(UTC+08:45) Eucla":                                             "Australia/Eucla",
	"(UTC+09:00) Chita":                                             "Asia/Chita",
	"(UTC+09:00) Osaka, Sapporo, Tokyo":                             "Asia/Tokyo",
	"(UTC+09:00) Seoul":                                             "Asia/Seoul",
	"(UTC+09:00) Yakutsk":                                           "Asia/Yakutsk",
	"(UTC+09:30) Adelaide":                                          "Australia/Adelaide",
	"(UTC+09:30) Darwin":                                            "Australia/Darwin",
	"(UTC+10:00) Brisbane":                                          "Australia/Brisbane",
	"(UTC+10:00) Canberra, Melbourne, Sydney":                       "Australia/Sydney",
	"(UTC+10:00) Guam, Port Moresby":                                "Pacific/Port_Moresby",
	"(UTC+10:00) Hobart":                                            "Australia/Hobart",
	"(UTC+10:00) Vladivostok":                                       "Asia/Vladivostok",
	"(UTC+10:30) Lord Howe Island":                                  "Australia/Lord_Howe",
	"(UTC+11:00) Bougainville Island":                               "Pacific/Bougainville",
	"(UTC+11:00) Chokurdakh":                                        "Asia/Srednekolymsk",
	"(UTC+11:00) Magadan":                                           "Asia/Magadan",
	"(UTC+11:00) Norfolk Island":                                    "Pacific/Norfolk",
	"(UTC+11:00) Sakhalin":                                          "Asia/Sakhalin",
	"(UTC+11:00) Solomon Is., New Caledonia":                        "Pacific/Guadalcanal",
	"(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky":                  "Asia/Kamchatka",
	"(UTC+12:00) Auckland, Wellington":                              "Pacific/Auckland",
	"(UTC+12:00) Coordinated Universal Time+12":                     "Etc/GMT-12",
	"(UTC+12:00) Fiji":                                              "Pacific/Fiji",
	"(UTC+12:45) Chatham Islands":                                   "Pacific/Chatham",
	"(UTC+13:00) Nuku'alofa":                                        "Pacific/Tongatapu",
	"(UTC+13:00) Samoa":                                             "Pacific/Apia",
	"(UTC+14:00) Kiritimati Island":                                 "Pacific/Kiritimati",
	// territory="001"
	"Dateline Standard Time":  "Etc/GMT+12",
	"UTC-11":                  "Etc/GMT+11",
	"Aleutian Standard Time":  "America/Adak",
	"Hawaiian Standard Time":  "Pacific/Honolulu",
	"Marquesas Standard Time": "Pacific/Marquesas",
	"Alaskan Standard Time":   "America/Anchorage",
	"UTC-09":                  "Etc/GMT+9",
	"Pacific Standard Time (Mexico)": "America/Tijuana",
	"UTC-08":                          "Etc/GMT+8",
	"Pacific Standard Time":           "America/Los_Angeles",
	"US Mountain Standard Time":       "America/Phoenix",
	"Mountain Standard Time (Mexico)": "America/Chihuahua",
	"Mountain Standard Time":          "America/Denver",
	"Central America Standard Time":   "America/Guatemala",
	"Central Standard Time":           "America/Chicago",
	"Easter Island Standard Time":     "Pacific/Easter",
	"Central Standard Time (Mexico)":  "America/Mexico_City",
	"Canada Central Standard Time":    "America/Regina",
	"SA Pacific Standard Time":        "America/Bogota",
	"Eastern Standard Time (Mexico)":  "America/Cancun",
	"Eastern Standard Time":           "America/New_York",
	"Haiti Standard Time":             "America/Port-au-Prince",
	"Cuba Standard Time":              "America/Havana",
	"US Eastern Standard Time":        "America/Indianapolis",
	"Paraguay Standard Time":          "America/Asuncion",
	"Atlantic Standard Time":          "America/Halifax",
	"Venezuela Standard Time":         "America/Caracas",
	"Central Brazilian Standard Time": "America/Cuiaba",
	"SA Western Standard Time":        "America/La_Paz",
	"Pacific SA Standard Time":        "America/Santiago",
	"Turks And Caicos Standard Time":  "America/Grand_Turk",
	"Newfoundland Standard Time":      "America/St_Johns",
	"Tocantins Standard Time":         "America/Araguaina",
	"E. South America Standard Time":  "America/Sao_Paulo",
	"SA Eastern Standard Time":        "America/Cayenne",
	"Argentina Standard Time":         "America/Buenos_Aires",
	"Greenland Standard Time":         "America/Godthab",
	"Montevideo Standard Time":        "America/Montevideo",
	"Saint Pierre Standard Time":      "America/Miquelon",
	"Bahia Standard Time":             "America/Bahia",
	"UTC-02":                          "Etc/GMT+2",
	"Azores Standard Time":            "Atlantic/Azores",
	"Cape Verde Standard Time":        "Atlantic/Cape_Verde",
	"UTC": "Etc/GMT",
	"Morocco Standard Time":           "Africa/Casablanca",
	"GMT Standard Time":               "Europe/London",
	"Greenwich Standard Time":         "Atlantic/Reykjavik",
	"W. Europe Standard Time":         "Europe/Berlin",
	"Central Europe Standard Time":    "Europe/Budapest",
	"Romance Standard Time":           "Europe/Paris",
	"Central European Standard Time":  "Europe/Warsaw",
	"W. Central Africa Standard Time": "Africa/Lagos",
	"Namibia Standard Time":           "Africa/Windhoek",
	"Jordan Standard Time":            "Asia/Amman",
	"GTB Standard Time":               "Europe/Bucharest",
	"Middle East Standard Time":       "Asia/Beirut",
	"Egypt Standard Time":             "Africa/Cairo",
	"E. Europe Standard Time":         "Europe/Chisinau",
	"Syria Standard Time":             "Asia/Damascus",
	"West Bank Standard Time":         "Asia/Hebron",
	"South Africa Standard Time":      "Africa/Johannesburg",
	"FLE Standard Time":               "Europe/Kiev",
	"Israel Standard Time":            "Asia/Jerusalem",
	"Kaliningrad Standard Time":       "Europe/Kaliningrad",
	"Libya Standard Time":             "Africa/Tripoli",
	"Arabic Standard Time":            "Asia/Baghdad",
	"Turkey Standard Time":            "Europe/Istanbul",
	"Arab Standard Time":              "Asia/Riyadh",
	"Belarus Standard Time":           "Europe/Minsk",
	"Russian Standard Time":           "Europe/Moscow",
	"E. Africa Standard Time":         "Africa/Nairobi",
	"Iran Standard Time":              "Asia/Tehran",
	"Arabian Standard Time":           "Asia/Dubai",
	"Astrakhan Standard Time":         "Europe/Astrakhan",
	"Azerbaijan Standard Time":        "Asia/Baku",
	"Russia Time Zone 3":              "Europe/Samara",
	"Mauritius Standard Time":         "Indian/Mauritius",
	"Georgian Standard Time":          "Asia/Tbilisi",
	"Caucasus Standard Time":          "Asia/Yerevan",
	"Afghanistan Standard Time":       "Asia/Kabul",
	"West Asia Standard Time":         "Asia/Tashkent",
	"Ekaterinburg Standard Time":      "Asia/Yekaterinburg",
	"Pakistan Standard Time":          "Asia/Karachi",
	"India Standard Time":             "Asia/Calcutta",
	"Sri Lanka Standard Time":         "Asia/Colombo",
	"Nepal Standard Time":             "Asia/Katmandu",
	"Central Asia Standard Time":      "Asia/Almaty",
	"Bangladesh Standard Time":        "Asia/Dhaka",
	"Omsk Standard Time":              "Asia/Omsk",
	"Myanmar Standard Time":           "Asia/Rangoon",
	"SE Asia Standard Time":           "Asia/Bangkok",
	"Altai Standard Time":             "Asia/Barnaul",
	"W. Mongolia Standard Time":       "Asia/Hovd",
	"North Asia Standard Time":        "Asia/Krasnoyarsk",
	"N. Central Asia Standard Time":   "Asia/Novosibirsk",
	"Tomsk Standard Time":             "Asia/Tomsk",
	"China Standard Time":             "Asia/Shanghai",
	"North Asia East Standard Time":   "Asia/Irkutsk",
	"Singapore Standard Time":         "Asia/Singapore",
	"W. Australia Standard Time":      "Australia/Perth",
	"Taipei Standard Time":            "Asia/Taipei",
	"Ulaanbaatar Standard Time":       "Asia/Ulaanbaatar",
	"North Korea Standard Time":       "Asia/Pyongyang",
	"Aus Central W. Standard Time":    "Australia/Eucla",
	"Transbaikal Standard Time":       "Asia/Chita",
	"Tokyo Standard Time":             "Asia/Tokyo",
	"Korea Standard Time":             "Asia/Seoul",
	"Yakutsk Standard Time":           "Asia/Yakutsk",
	"Cen. Australia Standard Time":    "Australia/Adelaide",
	"AUS Central Standard Time":       "Australia/Darwin",
	"E. Australia Standard Time":      "Australia/Brisbane",
	"AUS Eastern Standard Time":       "Australia/Sydney",
	"West Pacific Standard Time":      "Pacific/Port_Moresby",
	"Tasmania Standard Time":          "Australia/Hobart",
	"Vladivostok Standard Time":       "Asia/Vladivostok",
	"Lord Howe Standard Time":         "Australia/Lord_Howe",
	"Bougainville Standard Time":      "Pacific/Bougainville",
	"Russia Time Zone 10":             "Asia/Srednekolymsk",
	"Magadan Standard Time":           "Asia/Magadan",
	"Norfolk Standard Time":           "Pacific/Norfolk",
	"Sakhalin Standard Time":          "Asia/Sakhalin",
	"Central Pacific Standard Time":   "Pacific/Guadalcanal",
	"Russia Time Zone 11":             "Asia/Kamchatka",
	"New Zealand Standard Time":       "Pacific/Auckland",
	"UTC+12":                          "Etc/GMT-12",
	"Fiji Standard Time":              "Pacific/Fiji",
	"Chatham Islands Standard Time":   "Pacific/Chatham",
	"Tonga Standard Time":             "Pacific/Tongatapu",
	"Samoa Standard Time":             "Pacific/Apia",
	"Line Islands Standard Time":      "Pacific/Kiritimati",

	//Zoom HTML invite cases
	"Midway Island, Samoa":            "Etc/GMT+11",
	"Pago Pago":                       "Etc/GMT+11",
	"Hawaii":                          "Pacific/Honolulu",
	"Alaska":                          "America/Anchorage",
	"Vancouver":                       "Etc/GMT+9",
	"Pacific Time (US and Canada)":    "America/Los_Angeles",
	"Tijuana":                         "America/Tijuana",
	"Edmonton":                        "Etc/GMT+7",
	"Mountain Time (US and Canada)":   "America/Denver",
	"Arizona":                         "America/Phoenix",
	"Mazatlan":                        "America/Chihuahua",
	"Winnipeg":                        "Etc/GMT+6",
	"Saskatchewan":                    "America/Regina",
	"Central Time (US and Canada)":    "America/Chicago",
	"Mexico City":                     "America/Mexico_City",
	"Guatemala":                       "Etc/GMT+6",
	"El Salvador":                     "Etc/GMT+6",
	"Managua":                         "Etc/GMT+6",
	"Costa Rica":                      "Etc/GMT+6",
	"Tegucigalpa":                     "Etc/GMT+6",
	"Montreal":                        "Etc/GMT+5",
	"Eastern Time (US and Canada)":    "America/New_York",
	"Indiana (East)":                  "America/Indianapolis",
	"Panama":                          "Etc/GMT+5",
	"Bogot":                           "Etc/GMT+5",
	"Lima":                            "Etc/GMT+5",
	"Halifax":                         "Etc/GMT+4",
	"Puerto Rico":                     "Etc/GMT+4",
	"Caracas":                         "America/Caracas",
	"Atlantic Time (Canada)":          "America/Halifax",
	"Newfoundland and Labrador":       "America/St_Johns",
	"Santiago":                        "America/Santiago",
	"Montevideo":                      "America/Montevideo",
	"Brasilia":                        "America/Sao_Paulo",
	"Buenos Aires, Georgetown":        "America/Buenos_Aires",
	"Greenland":                       "America/Godthab",
	"Sao Paulo":                       "Etc/GMT+2",
	"Azores":                          "Atlantic/Azores",
	"Cabo Verde Islands":              "Atlantic/Cape_Verde",
	"Universal Time UTC":              "Etc/GMT",
	"Greenwich Mean Time":             "Etc/GMT",
	"Reykjavik":                       "Atlantic/Reykjavik",
	"Dublin":                          "Europe/London",
	"London":                          "Europe/London",
	"Lisbon":                          "Europe/London",
	"Casablanca":                      "Africa/Casablanca",
	"Nouakchott":                      "Etc/GMT",
	"Belgrade, Bratislava, Ljubljana": "Europe/Budapest",
	"Sarajevo, Skopje, Zagreb":        "Europe/Warsaw",
	"Oslo":       "Etc/GMT+1",
	"Copenhagen": "Europe/Paris",
	"Brussels":   "Europe/Paris",
	"Amsterdam, Berlin, Rome, Stockholm, Vienna": "Europe/Berlin",
	"Amsterdam":                    "Europe/Berlin",
	"Rome":                         "Europe/Berlin",
	"Stockholm":                    "Europe/Berlin",
	"Vienna":                       "Europe/Berlin",
	"Luxembourg":                   "Europe/Paris",
	"Paris":                        "Europe/Paris",
	"Zurich":                       "Europe/Paris",
	"Madrid":                       "Europe/Paris",
	"West Central Africa":          "Africa/Lagos",
	"Algiers":                      "Etc/GMT+1",
	"Tunis":                        "Etc/GMT+1",
	"Warsaw":                       "Etc/GMT+1",
	"Prague Bratislava":            "Etc/GMT+1",
	"Budapest":                     "Etc/GMT+1",
	"Helsinki":                     "Europe/Kiev",
	"Harare, Pretoria":             "Africa/Johannesburg",
	"Sofia":                        "Europe/Kiev",
	"Athens":                       "Europe/Bucharest",
	"Bucharest":                    "Europe/Bucharest",
	"Nicosia":                      "Europe/Kiev",
	"Beirut":                       "Asia/Beirut",
	"Damascus":                     "Asia/Damascus",
	"Jerusalem":                    "Asia/Jerusalem",
	"Amman":                        "Asia/Amman",
	"Tripoli":                      "Africa/Tripoli",
	"Cairo":                        "Africa/Cairo",
	"Johannesburg":                 "Africa/Johannesburg",
	"Kiev":                         "Europe/Kiev",
	"Nairobi":                      "Africa/Nairobi",
	"Istanbul":                     "Europe/Istanbul",
	"Moscow":                       "Europe/Moscow",
	"Baghdad":                      "Asia/Baghdad",
	"Kuwait":                       "Asia/Riyadh",
	"Riyadh":                       "Asia/Riyadh",
	"Bahrain":                      "Asia/Riyadh",
	"Qatar":                        "Etc/GMT+3",
	"Aden":                         "Etc/GMT+3",
	"Khartoum":                     "Etc/GMT+3",
	"Djibouti":                     "Etc/GMT+3",
	"Mogadishu":                    "Etc/GMT+3",
	"Tehran":                       "Asia/Tehran",
	"Dubai":                        "Asia/Dubai",
	"Muscat":                       "Asia/Dubai",
	"Baku, Tbilisi, Yerevan":       "Etc/GMT+4",
	"Kabul":                        "Asia/Kabul",
	"Yekaterinburg":                "Etc/GMT+5",
	"Islamabad, Karachi, Tashkent": "Asia/Karachi",
	"India": "Asia/Calcutta",
	"Mumbai, Kolkata, New Delhi":     "Asia/Calcutta",
	"Kathmandu":                      "Asia/Katmandu",
	"Almaty":                         "Asia/Almaty",
	"Dacca":                          "Asia/Dhaka",
	"Astana, Dhaka":                  "Asia/Dhaka",
	"Novosibirsk":                    "Asia/Novosibirsk",
	"Krasnoyarsk":                    "Asia/Krasnoyarsk",
	"Bangkok":                        "Asia/Bangkok",
	"Vietnam":                        "Etc/GMT+7",
	"Jakarta":                        "Asia/Bangkok",
	"Irkutsk, Ulaanbaatar":           "Asia/Irkutsk",
	"Beijing, Shanghai":              "Asia/Shanghai",
	"Hong Kong":                      "Asia/Shanghai",
	"Taipei":                         "Asia/Taipei",
	"Kuala Lumpur":                   "Asia/Singapore",
	"Singapore":                      "Asia/Singapore",
	"Perth":                          "Australia/Perth",
	"Yakutsk":                        "Asia/Yakutsk",
	"Seoul":                          "Asia/Seoul",
	"Osaka, Sapporo, Tokyo":          "Asia/Tokyo",
	"Darwin":                         "Australia/Darwin",
	"Vladivostok":                    "Asia/Vladivostok",
	"Guam, Port Moresby":             "Pacific/Port_Moresby",
	"Brisbane":                       "Australia/Brisbane",
	"Adelaide":                       "Australia/Adelaide",
	"Canberra, Melbourne, Sydney":    "Australia/Sydney",
	"Hobart":                         "Australia/Hobart",
	"Magadan":                        "Asia/Magadan",
	"Solomon Islands":                "Pacific/Guadalcanal",
	"New Caledonia":                  "Pacific/Guadalcanal",
	"Kamchatka":                      "Asia/Kamchatka",
	"Fiji Islands, Marshall Islands": "Etc/GMT+13",
	"Auckland, Wellington":           "Etc/GMT+13",
	"Independent State of Samoa":     "Etc/GMT+14",
}

//UtcOffsetToTimezoneMap translates the UTC time offset to one of the timezones that can ber parsed by ICS parsers.
var UtcOffsetToTimezoneMap = map[string]string{
	"-1200": "Etc/GMT+12",
	"-1100": "Etc/GMT+11",
	"-1000": "America/Adak",
	"-0930": "Pacific/Marquesas",
	"-0900": "America/Anchorage",
	"-0800": "America/Los_Angeles",
	"-0700": "America/Phoenix",
	"-0600": "America/Guatemala",
	"-0500": "America/New_York",
	"-0400": "America/Santiago",
	"-0330": "America/St_Johns",
	"-0300": "America/Buenos_Aires",
	"-0200": "Etc/GMT+2",
	"-0100": "Atlantic/Cape_Verde",
	"+0000": "Europe/London",
	"0000":  "Europe/London",
	"+0100": "Europe/Berlin",
	"+0200": "Africa/Cairo",
	"+0300": "Europe/Moscow",
	"+0400": "Asia/Dubai",
	"+0430": "Asia/Kabul",
	"+0500": "Asia/Karachi",
	"+0530": "Asia/Calcutta",
	"+0545": "Asia/Katmandu",
	"+0600": "Asia/Dhaka",
	"+0630": "Asia/Rangoon",
	"+0700": "Asia/Bangkok",
	"+0800": "Asia/Shanghai",
	"+0830": "Asia/Pyongyang",
	"+0845": "Australia/Eucla",
	"+0900": "Asia/Tokyo",
	"+0930": "Australia/Darwin",
	"+1000": "Australia/Sydney",
	"+1030": "Australia/Lord_Howe",
	"+1100": "Pacific/Norfolk",
	"+1200": "Pacific/Fiji",
	"+1245": "Pacific/Chatham",
	"+1300": "Pacific/Tongatapu",
	"+1400": "Pacific/Kiritimati",
}
