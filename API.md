# Application Programming Interface Documentation

## List
| No | Web Service | Method | URL | Role | Authentication |
|----|-------------|--------|-----|------|----------------|
| 1 | [Register](#register) | POST | /bos/register, /operasional/register, /pengedar/register | Bos, Operasional, Pengedar | No |
| 2 | [Login](#login) | POST | /bos/login, /operasional/login, /pengedar/login | Bos, Operasional, Pengedar | No |
| 3 | [Profile](#profile) | GET | /bos/profile, /operasional/profile, /pengedar/profile | Bos, Operasional, Pengedar | Yes |
| 4 | [Logout](#logout) | PUT | /bos/logout, /operasional/logout, /pengedar/logout | Bos, Operasional, Pengedar | Yes |
| 5 | [Search all pokemon from pokeapi](#search-all-pokemon-from-pokeapi) | GET | /search/pokemons | Operasional | Yes |
| 6 | [Search pokemon by pokemon's name from pokeapi](#search-pokemon-by-pokemons-name-from-pokeapi) | GET | /search/pokemon | Operasional | Yes |
| 7 | [Create pokemon](#create-pokemon) | POST | /pokemon | Operasional | Yes |
| 8 | [Update pokemon](#update-pokemon) | PUT | /pokemon | Operasional | Yes |
| 9 | [Delete pokemon](#delete-pokemon) | DELETE | /pokemon | Operasional | Yes |
| 10 | [Get all pokemons](#get-all-pokemons) | GET | /pokemons | Operasional, Pengedar | Yes |
| 11 | [Search pokemon by id](#search-pokemon-by-id) | GET | /pokemon/id | Operasional | Yes |
| 12 | [Search pokemon by name](#search-pokemon-by-name) | GET | /pokemon/name | Operasional | Yes |
| 13 | [Create transaction](#create-transaction) | POST | /transaction | Pengedar | Yes |
| 14 | [Cancel transaction](#cancel-transaction) | PUT | /transaction | Pengedar | Yes |
| 15 | [Get transaction by id](#get-transaction-by-id) | GET | /transaction | Bos | Yes |
| 15 | [Get all success transaction](#get-all-success-transaction) | GET | /transactions/success | Bos | Yes |
| 15 | [Get all cancelled transaction](#get-all-cancelled-transaction) | GET | /transactions/cancelled | Bos | Yes |

## Register
### Register Bos
#### URL : `/bos/register`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "bos@blackmarket",
	"password" : "12345",
	"name": "Riska"
}
```

#### Body Response
```json
{
    "id": 6,
    "level": "Bos",
    "email": "bos@blackmarket",
    "name": "Riska"
}
```

### Register Operasional
#### URL : `/operasional/register`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "operasional@blackmarket",
	"password" : "12345",
	"name": "Riska"
}
```

#### Body Response
```json
{
    "id": 6,
    "level": "Operasional",
    "email": "operasional@blackmarket",
    "name": "Riska"
}
```

### Register Pengedar
#### URL : `/pengedar/register`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "pengedar@blackmarket",
	"password" : "12345",
	"name": "Riska"
}
```

#### Body Response
```json
{
    "id": 6,
    "level": "Pengedar",
    "email": "pengedar@blackmarket",
    "name": "Riska"
}
```

## Login
### Login Bos
#### URL : `/bos/login`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "bos1@blackmarket",
	"password" : "12345"
}
```

#### Body Response
```json
{
    "id": 1,
    "level": "Bos",
    "email": "bos1@blackmarket",
    "name": "Riska",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzE2NDQyNDcsInVzZXJJZCI6MX0.h7l4OhDR5EMWT1NG0NaJEpIAfVhf84V3pR_FSWWjUqY"
}
```

### Login Operasional
#### URL : `/operasional/login`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "operasional1@blackmarket",
	"password" : "12345"
}
```

#### Body Response
```json
{
    "id": 1,
    "level": "Operasional",
    "email": "operasional1@blackmarket",
    "name": "Riska",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzE2NDQyNDcsInVzZXJJZCI6MX0.h7l4OhDR5EMWT1NG0NaJEpIAfVhf84V3pR_FSWWjUqY"
}
```

### Login Pengedar
#### URL : `/pengedar/login`
#### Method : `POST`

#### Body Request
```json
{
	"email" : "pengedar1@blackmarket",
	"password" : "12345"
}
```

#### Body Response
```json
{
    "id": 1,
    "level": "Pengedar",
    "email": "pengedar1@blackmarket",
    "name": "Riska",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzE2NDQyNDcsInVzZXJJZCI6MX0.h7l4OhDR5EMWT1NG0NaJEpIAfVhf84V3pR_FSWWjUqY"
}
```

## Profile
### Profile Bos
#### URL : `/bos/profile`
#### Method : `GET`

#### Body Response
```json
{
    "id": 1,
    "level": "Bos",
    "email": "bos1@blackmarket",
    "name": "Riska"
}
```

### Profile Operasional
#### URL : `/operasional/profile`
#### Method : `GET`

#### Body Response
```json
{
    "id": 1,
    "level": "Operasional",
    "email": "operasional1@blackmarket",
    "name": "Riska"
}
```

### Profile Pengedar
#### URL : `/pengedar/profile`
#### Method : `GET`

#### Body Response
```json
{
    "id": 1,
    "level": "Pengedar",
    "email": "pengedar1@blackmarket",
    "name": "Riska"
}
```

## Logout
### Logout Bos
#### URL : `/bos/logout`
#### Method : `PUT`

#### Body Response
```json
{
    "id": 1,
    "level": "Bos",
    "email": "bos1@blackmarket",
    "name": "Riska",
    "token": ""
}
```

### Logout Operasional
#### URL : `/operasional/logout`
#### Method : `PUT`

#### Body Response
```json
{
    "id": 1,
    "level": "Operasional",
    "email": "operasional@blackmarket",
    "name": "Riska",
    "token": ""
}
```

### Logout Pengedar
#### URL : `/pengedar/logout`
#### Method : `PUT`

#### Body Response
```json
{
    "id": 1,
    "level": "Pengedar",
    "email": "pengedar@blackmarket",
    "name": "Riska",
    "token": ""
}
```

## Search all pokemon from pokeapi
### URL : `/search/pokemons`
### Method : `GET`

### Body Response
```json
[
    {
        "Name": "bulbasaur",
        "ID": 1,
        "Url": "https://pokeapi.co/api/v2/pokemon/1/",
        "Weight": 69,
        "Height": 7,
        "Types": [
            {
                "Type": {
                    "Name": "grass"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "overgrow"
                }
            },
            {
                "Ability": {
                    "Name": "chlorophyll"
                }
            }
        ]
    },
    {
        "Name": "ivysaur",
        "ID": 2,
        "Url": "https://pokeapi.co/api/v2/pokemon/2/",
        "Weight": 130,
        "Height": 10,
        "Types": [
            {
                "Type": {
                    "Name": "grass"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "overgrow"
                }
            },
            {
                "Ability": {
                    "Name": "chlorophyll"
                }
            }
        ]
    },
    {
        "Name": "venusaur",
        "ID": 3,
        "Url": "https://pokeapi.co/api/v2/pokemon/3/",
        "Weight": 1000,
        "Height": 20,
        "Types": [
            {
                "Type": {
                    "Name": "grass"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "overgrow"
                }
            },
            {
                "Ability": {
                    "Name": "chlorophyll"
                }
            }
        ]
    },
    {
        "Name": "charmander",
        "ID": 4,
        "Url": "https://pokeapi.co/api/v2/pokemon/4/",
        "Weight": 85,
        "Height": 6,
        "Types": [
            {
                "Type": {
                    "Name": "fire"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "blaze"
                }
            },
            {
                "Ability": {
                    "Name": "solar-power"
                }
            }
        ]
    },
    {
        "Name": "charmeleon",
        "ID": 5,
        "Url": "https://pokeapi.co/api/v2/pokemon/5/",
        "Weight": 190,
        "Height": 11,
        "Types": [
            {
                "Type": {
                    "Name": "fire"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "blaze"
                }
            },
            {
                "Ability": {
                    "Name": "solar-power"
                }
            }
        ]
    },
    {
        "Name": "charizard",
        "ID": 6,
        "Url": "https://pokeapi.co/api/v2/pokemon/6/",
        "Weight": 905,
        "Height": 17,
        "Types": [
            {
                "Type": {
                    "Name": "fire"
                }
            },
            {
                "Type": {
                    "Name": "flying"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "blaze"
                }
            },
            {
                "Ability": {
                    "Name": "solar-power"
                }
            }
        ]
    },
    {
        "Name": "squirtle",
        "ID": 7,
        "Url": "https://pokeapi.co/api/v2/pokemon/7/",
        "Weight": 90,
        "Height": 5,
        "Types": [
            {
                "Type": {
                    "Name": "water"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "torrent"
                }
            },
            {
                "Ability": {
                    "Name": "rain-dish"
                }
            }
        ]
    },
    {
        "Name": "wartortle",
        "ID": 8,
        "Url": "https://pokeapi.co/api/v2/pokemon/8/",
        "Weight": 225,
        "Height": 10,
        "Types": [
            {
                "Type": {
                    "Name": "water"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "torrent"
                }
            },
            {
                "Ability": {
                    "Name": "rain-dish"
                }
            }
        ]
    },
    {
        "Name": "blastoise",
        "ID": 9,
        "Url": "https://pokeapi.co/api/v2/pokemon/9/",
        "Weight": 855,
        "Height": 16,
        "Types": [
            {
                "Type": {
                    "Name": "water"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "torrent"
                }
            },
            {
                "Ability": {
                    "Name": "rain-dish"
                }
            }
        ]
    },
    {
        "Name": "caterpie",
        "ID": 10,
        "Url": "https://pokeapi.co/api/v2/pokemon/10/",
        "Weight": 29,
        "Height": 3,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "shield-dust"
                }
            },
            {
                "Ability": {
                    "Name": "run-away"
                }
            }
        ]
    },
    {
        "Name": "metapod",
        "ID": 11,
        "Url": "https://pokeapi.co/api/v2/pokemon/11/",
        "Weight": 99,
        "Height": 7,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "shed-skin"
                }
            }
        ]
    },
    {
        "Name": "butterfree",
        "ID": 12,
        "Url": "https://pokeapi.co/api/v2/pokemon/12/",
        "Weight": 320,
        "Height": 11,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            },
            {
                "Type": {
                    "Name": "flying"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "compound-eyes"
                }
            },
            {
                "Ability": {
                    "Name": "tinted-lens"
                }
            }
        ]
    },
    {
        "Name": "weedle",
        "ID": 13,
        "Url": "https://pokeapi.co/api/v2/pokemon/13/",
        "Weight": 32,
        "Height": 3,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "shield-dust"
                }
            },
            {
                "Ability": {
                    "Name": "run-away"
                }
            }
        ]
    },
    {
        "Name": "kakuna",
        "ID": 14,
        "Url": "https://pokeapi.co/api/v2/pokemon/14/",
        "Weight": 100,
        "Height": 6,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "shed-skin"
                }
            }
        ]
    },
    {
        "Name": "beedrill",
        "ID": 15,
        "Url": "https://pokeapi.co/api/v2/pokemon/15/",
        "Weight": 295,
        "Height": 10,
        "Types": [
            {
                "Type": {
                    "Name": "bug"
                }
            },
            {
                "Type": {
                    "Name": "poison"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "swarm"
                }
            },
            {
                "Ability": {
                    "Name": "sniper"
                }
            }
        ]
    },
    {
        "Name": "pidgey",
        "ID": 16,
        "Url": "https://pokeapi.co/api/v2/pokemon/16/",
        "Weight": 18,
        "Height": 3,
        "Types": [
            {
                "Type": {
                    "Name": "normal"
                }
            },
            {
                "Type": {
                    "Name": "flying"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "keen-eye"
                }
            },
            {
                "Ability": {
                    "Name": "tangled-feet"
                }
            },
            {
                "Ability": {
                    "Name": "big-pecks"
                }
            }
        ]
    },
    {
        "Name": "pidgeotto",
        "ID": 17,
        "Url": "https://pokeapi.co/api/v2/pokemon/17/",
        "Weight": 300,
        "Height": 11,
        "Types": [
            {
                "Type": {
                    "Name": "normal"
                }
            },
            {
                "Type": {
                    "Name": "flying"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "keen-eye"
                }
            },
            {
                "Ability": {
                    "Name": "tangled-feet"
                }
            },
            {
                "Ability": {
                    "Name": "big-pecks"
                }
            }
        ]
    },
    {
        "Name": "pidgeot",
        "ID": 18,
        "Url": "https://pokeapi.co/api/v2/pokemon/18/",
        "Weight": 395,
        "Height": 15,
        "Types": [
            {
                "Type": {
                    "Name": "normal"
                }
            },
            {
                "Type": {
                    "Name": "flying"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "keen-eye"
                }
            },
            {
                "Ability": {
                    "Name": "tangled-feet"
                }
            },
            {
                "Ability": {
                    "Name": "big-pecks"
                }
            }
        ]
    },
    {
        "Name": "rattata",
        "ID": 19,
        "Url": "https://pokeapi.co/api/v2/pokemon/19/",
        "Weight": 35,
        "Height": 3,
        "Types": [
            {
                "Type": {
                    "Name": "normal"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "run-away"
                }
            },
            {
                "Ability": {
                    "Name": "guts"
                }
            },
            {
                "Ability": {
                    "Name": "hustle"
                }
            }
        ]
    },
    {
        "Name": "raticate",
        "ID": 20,
        "Url": "https://pokeapi.co/api/v2/pokemon/20/",
        "Weight": 185,
        "Height": 7,
        "Types": [
            {
                "Type": {
                    "Name": "normal"
                }
            }
        ],
        "Abilities": [
            {
                "Ability": {
                    "Name": "run-away"
                }
            },
            {
                "Ability": {
                    "Name": "guts"
                }
            },
            {
                "Ability": {
                    "Name": "hustle"
                }
            }
        ]
    }
]
```

## Search pokemon by pokemon's name from pokeapi
### URL : `/search/pokemon?name=bulbasaur`
### Method : `GET`

### Body Response
```json
{
    "Name": "bulbasaur",
    "ID": 1,
    "Weight": 69,
    "Height": 7,
    "Types": [
        {
            "Name": "grass"
        },
        {
            "Name": "poison"
        }
    ],
    "Abilities": [
        {
            "Name": "overgrow"
        },
        {
            "Name": "chlorophyll"
        }
    ]
}
```

## Create pokemon
### URL : `/pokemon?name=squirtle`
### Method : `POST`

### Body Request
```json
{
	"price" : 50000,
	"stock" : 100
}
```

### Body Response
```json
{
    "Name": "squirtle",
    "ID": 4,
    "Weight": 90,
    "Height": 5,
    "Price": 50000,
    "Stock": 100,
    "Types": [
        {
            "ID": 4,
            "name": "water",
            "detail_types": null
        }
    ],
    "Abilities": [
        {
            "id": 5,
            "name": "torrent",
            "detail_abilities": null
        },
        {
            "id": 6,
            "name": "rain-dish",
            "detail_abilities": null
        }
    ]
}
```

## Update pokemon
### URL : `/pokemon?id=4`
### Method : `PUT`

### Body Request
```json
{
	"price" : 100000
}
```

### Body Response
```json
{
    "Name": "squirtle",
    "ID": 4,
    "Weight": 90,
    "Height": 5,
    "Price": 100000,
    "Stock": 100
}
```

## Delete pokemon
### URL : `/pokemon?id=4`
### Method : `DELETE`

### Body Response
```json
{
    "Name": "squirtle",
    "ID": 4,
    "Weight": 90,
    "Height": 5,
    "Price": 100000,
    "Stock": 100
}
```

## Get all pokemons
### URL : `/pokemons`
### Method : `GET`

### Body Response
```json
[
    {
        "Name": "bulbasaur",
        "ID": 1,
        "Weight": 69,
        "Height": 7,
        "Price": 100000,
        "Stock": 90
    },
    {
        "Name": "ivysaur",
        "ID": 2,
        "Weight": 130,
        "Height": 10,
        "Price": 50000,
        "Stock": 90
    },
    {
        "Name": "charmander",
        "ID": 3,
        "Weight": 85,
        "Height": 6,
        "Price": 50000,
        "Stock": 100
    },
    {
        "Name": "squirtle",
        "ID": 4,
        "Weight": 90,
        "Height": 5,
        "Price": 50000,
        "Stock": 100
    }
]
```

## Search pokemon by id
### URL : `/pokemon/id?id=2`
### Method : `GET`

### Body Response
```json
{
    "Name": "ivysaur",
    "ID": 2,
    "Weight": 130,
    "Height": 10,
    "Price": 50000,
    "Stock": 90,
    "Types": [
        {
            "ID": 1,
            "name": "grass",
            "detail_types": null
        },
        {
            "ID": 2,
            "name": "poison",
            "detail_types": null
        }
    ],
    "Abilities": [
        {
            "id": 1,
            "name": "overgrow",
            "detail_abilities": null
        },
        {
            "id": 2,
            "name": "chlorophyll",
            "detail_abilities": null
        }
    ]
}
```

## Search pokemon by name
### URL : `/pokemon/name?name=sau`
### Method : `GET`

### Body Response
```json
[
    {
        "Name": "bulbasaur",
        "ID": 1,
        "Weight": 69,
        "Height": 7,
        "Price": 100000,
        "Stock": 90
    },
    {
        "Name": "ivysaur",
        "ID": 2,
        "Weight": 130,
        "Height": 10,
        "Price": 50000,
        "Stock": 90
    }
]
```

## Create transaction
### URL : `/transaction`
### Method : `POST`

### Body Request
```json
[
	{
		"pokemon_id" : 1,
		"quantity" : 5
	},
	{
		"pokemon_id" : 2,
		"quantity" : 10
	}
]
```

### Body Response
```json
{
    "EmployeeName": "Dewi",
    "ID": 6,
    "Code": "PBM-6",
    "Date": "2021-09-14T22:06:32.8241273+07:00",
    "Total": 1000000,
    "Status": "Success",
    "Pokemon": [
        {
            "PokemonName": "bulbasaur",
            "Quantity": 5,
            "Price": 100000
        },
        {
            "PokemonName": "ivysaur",
            "Quantity": 10,
            "Price": 50000
        }
    ]
}
```

## Cancel transaction
### URL : `/transaction?id=6`
### Method : `PUT`

### Body Response
```json
{
    "EmployeeName": "Dewi",
    "ID": 6,
    "Code": "PBM-6",
    "Date": "2021-09-14T22:06:32.824+07:00",
    "Total": 1000000,
    "Status": "Cancelled",
    "Pokemon": [
        {
            "PokemonName": "bulbasaur",
            "Quantity": 5,
            "Price": 100000
        },
        {
            "PokemonName": "ivysaur",
            "Quantity": 10,
            "Price": 50000
        }
    ]
}
```

## Get transaction by id
### URL : `/transaction`
### Method : `GET`

### Body Response
```json
{
    "EmployeeName": "Riska",
    "ID": 4,
    "Code": "PBM-4",
    "Date": "2021-09-08T00:00:00+07:00",
    "Total": 500000,
    "Status": "Success",
    "Pokemon": [
        {
            "PokemonName": "bulbasaur",
            "Quantity": 5,
            "Price": 100000
        }
    ]
}
```

## Get all success transaction
### URL : `/transactions/success`
### Method : `GET`

### Body Response
```json
{
    "Total": 1500000,
    "Transaction": [
        {
            "EmployeeName": "Riska",
            "ID": 4,
            "Code": "PBM-4",
            "Date": "2021-09-08T00:00:00+07:00",
            "Total": 500000,
            "Status": "Success"
        },
        {
            "EmployeeName": "Riska",
            "ID": 5,
            "Code": "PBM-5",
            "Date": "2021-09-14T20:18:16.365+07:00",
            "Total": 1000000,
            "Status": "Success"
        }
    ]
}
```

## Get all cancelled transaction
### URL : `/transactions/cancelled`
### Method : `GET`

### Body Response
```json
{
    "Total": 2750000,
    "Transaction": [
        {
            "EmployeeName": "Riska",
            "ID": 1,
            "Code": "PBM-1",
            "Date": "2021-09-14T19:28:28.873+07:00",
            "Total": 750000,
            "Status": "Cancelled"
        },
        {
            "EmployeeName": "Riska",
            "ID": 2,
            "Code": "PBM-2",
            "Date": "2021-09-14T19:30:07.722+07:00",
            "Total": 500000,
            "Status": "Cancelled"
        },
        {
            "EmployeeName": "Riska",
            "ID": 3,
            "Code": "PBM-3",
            "Date": "2021-09-14T19:33:34.856+07:00",
            "Total": 500000,
            "Status": "Cancelled"
        },
        {
            "EmployeeName": "Riska",
            "ID": 6,
            "Code": "PBM-6",
            "Date": "2021-09-14T22:06:32.824+07:00",
            "Total": 1000000,
            "Status": "Cancelled"
        }
    ]
}
```