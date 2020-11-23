# API Alkitab Indonesia

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/erwindosianipar/api-alkitab/Build%20Application)
![Travis (.com)](https://img.shields.io/travis/com/erwindosianipar/api-alkitab)
[![codecov](https://codecov.io/gh/erwindosianipar/api-alkitab/branch/master/graph/badge.svg?token=N04FFE5MUG)](https://codecov.io/gh/erwindosianipar/api-alkitab)
![GitHub](https://img.shields.io/github/license/erwindosianipar/api-alkitab)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/erwindosianipar/api-alkitab)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/erwindosianipar/api-alkitab)
![GitHub forks](https://img.shields.io/github/forks/erwindosianipar/api-alkitab)
![GitHub last commit](https://img.shields.io/github/last-commit/erwindosianipar/api-alkitab)
![Uptime Robot ratio (7 days)](https://img.shields.io/uptimerobot/ratio/7/m786403864-a6328db4aa0b4270dadbe851)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ferwindosianipar%2Fapi-alkitab.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ferwindosianipar%2Fapi-alkitab?ref=badge_shield)

## Demo: [https://api-alkitab.herokuapp.com](https://api-alkitab.herokuapp.com)

API Alkitab Indonesia adalah sebuah *third-party backend* Alkitab yang sumber bacaannya berasal dari [sabda.org](https://sabda.org).

[sabda.org](https://sabda.org) adalah sebuah Yayasan Kristen non-profit dan non-komersial yang bergerak dalam bidang pelayanan media komputer dan internet di Indonesia.

## Dokumentasi

Versi saat (**v2.2.1**) telah memiliki tiga fitur utama untuk menampilkan list pasal, ayat, dan pasal Alkitab.

### Passage List

Untuk menampilkan passage list perjanjian lama dan baru

#### Base URL

Version | URL
-- | --
v2 | [https://api-alkitab.herokuapp.com/v2/passage/list](https://api-alkitab.herokuapp.com/v2/passage/list)

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://api-alkitab.herokuapp.com/v2/passage/list](https://api-alkitab.herokuapp.com/v2/passage/list)

#### Response

```json
{
    "passage_list": [
        {
            "book_number": 1,
            "abbreviation": "Kej",
            "book_name": "Kejadian",
            "total_chapter": 50
        },
        {
            "book_number": 2,
            "abbreviation": "Kel",
            "book_name": "Keluaran",
            "total_chapter": 40
        }
    ]
}
```

### Passage and Chapter

Untuk menampilkan kitab dan pasal Alkitab gunakan **passage/{passage}/{chapter}** dengan mengirim *request* nama kitab dan nomor pasal.

#### Base URL

Version | URL
-- | --
v2 | [https://api-alkitab.herokuapp.com/v2/passage/{**chapter**}/{**chapter**}?ver={**ver**}](https://api-alkitab.herokuapp.com/v2/passage/Yoh/1?ver=tb)
v1 | [https://api-alkitab.herokuapp.com/v1/passage/{**chapter**}/{**chapter**}](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1)

Variabel | Keterangan | Tipe data | *Required*
-- | -- | -- | --
**passage** | Adalah kitab yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string` | Yes
**chapter** | Adalah nomor pasal yang ingin ditampilkan dari bagian Alkitab | `int` | Yes
**ver** | Adalah versi pilihan bahasa **tb** (terjemahan baru) *default*, **bis** (bahasa Indonesia sehari-hari), **net** (bahasa Inggris) | `string` | No

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://api-alkitab.herokuapp.com/v2/passage/Yoh/1?ver=tb](https://api-alkitab.herokuapp.com/v2/passage/Yoh/1?ver=tb)
v1 | GET | [https://api-alkitab.herokuapp.com/v1/passage/Yoh/1](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1)

#### Response v2

```json
{
    "title": "Yohanes 1:1-51",
    "book_number": 43,
    "chapter": 1,
    "verses": [
        {
            "verse": 1,
            "content": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
        },
        {
            "verse": 2,
            "content": "Ia pada mulanya bersama-sama dengan Allah."
        }
    ]
}
```

#### Response v1

```json
{
    "message": "success to fetching data",
    "response": {
        "title": "Yohanes 1:1-51",
        "passage": [
            {
                "book_id": 43,
                "title": "Yohanes 1:1-51",
                "chapter": {
                    "chap": 1,
                    "verses": {
                        "verse": [
                            {
                                "number": 1,
                                "text": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
                            },
                            {
                                "number": 2,
                                "text": "Ia pada mulanya bersama-sama dengan Allah."
                            }
                        ]
                    }
                }
            }
        ]
    },
    "status": 200
}
```

### Passage, Chapter, and Verse

Untuk membaca satu ayat Alkitab gunakan **passage/{passage}/{chapter}/{verse}** dengan mengirim nama kitab, nomor pasal dan nomor ayat Alkitab.

#### Base URL

Version | URL
-- | --
v2 | [https://api-alkitab.herokuapp.com/v2/passage/{**passage**}/{**chapter**}/{**verse**}?ver={**tb**}](https://api-alkitab.herokuapp.com/v2/passage/Yoh/1/1?ver=tb)
v1 | [https://api-alkitab.herokuapp.com/v1/passage/{**passage**}/{**chapter**}/{**verse**}](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1/1)

Variabel | Keterangan | Tipe data | *Required*
-- | -- | -- | --
**passage** | Adalah kitab yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string` | Yes
**chapter** | Adalah nomor pasal yang ingin ditampilkan dari bagian Alkitab | `int` | Yes
**verse** | Adalah nomor ayat dari bagian pasal/bab yang ingin ditampilkan | `int` | Yes
**ver** | Adalah versi pilihan bahasa **tb** (terjemahan baru) *default*, **bis** (bahasa Indonesia sehari-hari), **net** (bahasa Inggris) | `string` | No

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://api-alkitab.herokuapp.com/v2/passage/Yoh/1/1?ver=tb](https://api-alkitab.herokuapp.com/v2/passage/Yoh/1/1?ver=tb)
v1 | GET | [https://api-alkitab.herokuapp.com/v1/passage/Yoh/1/1](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1/1)

#### Response v2

```json
{
    "title": "Yohanes 1:1",
    "book_number": 43,
    "chapter": 1,
    "verses": [
        {
            "verse": 1,
            "content": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
        }
    ]
}
```

#### Response v1

```json
{
    "message": "success to fetching data",
    "response": {
        "title": "Yohanes 1:1",
        "passage": [
            {
                "book_id": 43,
                "title": "Yohanes 1:1",
                "chapter": {
                    "chap": 1,
                    "verses": {
                        "verse": [
                            {
                                "number": 1,
                                "text": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
                            }
                        ]
                    }
                }
            }
        ]
    },
    "status": 200
}
```

### Status Code

API mengembalikan beberapa status code sebagai berikut:

Status Code | Keterangan
-- | --
200 | OK
201 | CREATED
400 | BAD REQUEST
404 | NOT FOUND
500 | INTERNAL SERVER ERROR

### Author

API Alkitab Indonesia ini dibuat oleh **Erwindo Sianipar**. Untuk pertanyaan kritik dan saran, silahkan *drop* email ke [erwindosianipar@gmail.com](mailto:erwindosianipar@gmail.com).

### License

API Alkitab Indonesia released under the [MIT License](http://opensource.org/licenses/MIT).

```text
Copyright (c) 2020 Erwindo Sianipar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

> Kemuliaan bagi Tuhan saja - Soli Deo Gloria
