# API Alkitab Indonesia

## Demo: [https://api-alkitab.herokuapp.com](https://api-alkitab.herokuapp.com)

API Alkitab Indonesia adalah sebuah *third-party backend* Alkitab yang sumber bacaannya berasal dari [sabda.org](https://sabda.org).

[sabda.org](https://sabda.org) adalah sebuah Yayasan Kristen non-profit, non-komersial yang bergerak dalam bidang pelayanan media komputer dan internet di Indonesia.

## Penggunaan

Saat ini (24 Oktober 2020) API masih mempunyai dua fitur utama untuk penampilan pasal dan ayat Alkitab.

### Passage and Chapter

Untuk membaca bagian ayat Alkitab gunakan **passage/{passage}/{chapter}** dengan mengirim nama Ayat dan nomor bagian.

#### Base URL

Version | URL
-- | --
v1 | https://api-alkitab.herokuapp.com/v1/passage/{passage}/{chapter}
v2 | https://api-alkitab.herokuapp.com/v2/passage/{passage}/{chapter}

Variable | Keterangan | Tipe data
-- | -- | --
passage | Rentang ayat yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string`
chapter | Adalah nomor pasal/bab yang ingin ditampilkan dari bagian Alkitab | `int`

#### Example

Version | Method | URL
-- | -- | --
v1 | GET | [https://api-alkitab.herokuapp.com/v1/Yoh/1](https://api-alkitab.herokuapp.com/v1/Yoh/1)
v2 | GET | [https://api-alkitab.herokuapp.com/v2/Yoh/1](https://api-alkitab.herokuapp.com/v2/Yoh/1)

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
        },
        {
            "verse": 3,
            "content": "Segala sesuatu dijadikan oleh Dia dan tanpa Dia tidak ada suatupun yang telah jadi dari segala yang telah dijadikan."
        },
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

Untuk membaca satu ayat Alkitab gunakan **passage/{passage}/{chapter}/{verse}** dengan mengirim nama Pasal dan nomor bagian dan ayat.

#### Base URL

Version | URL
-- | --
v1 | https://api-alkitab.herokuapp.com/v1/passage/{passage}/{chapter}/{verse}
v2 | https://api-alkitab.herokuapp.com/v2/passage/{passage}/{chapter}/{verse}

Variable | Keterangan | Tipe data
-- | -- | --
passage | Rentang ayat yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string`
chapter | Adalah nomor pasal/bab yang ingin ditampilkan dari bagian Alkitab | `int`
verse | Adalah nomor ayat dari bagian pasal/bab yang ingin ditampilkan | `int`

#### Example

Version | Method | URL
-- | -- | --
v1 | GET | [https://api-alkitab.herokuapp.com/v1/Yoh/1/1](https://api-alkitab.herokuapp.com/v1/Yoh/1/1)
v2 | GET | [https://api-alkitab.herokuapp.com/v1/Yoh/1/1](https://api-alkitab.herokuapp.com/v1/Yoh/1/1)

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

Status Code | Description
-- | --
200 | OK
201 | CREATED
400 | BAD REQUEST
404 | NOT FOUND
500 | INTERNAL SERVER ERROR

### Author

API Alkitab Indonesia ini dibuat oleh **Erwindo Sianipar**. Untuk pertanyaan kritik dan saran, silahkan *drop* email ke [erwindosianipar@gmail.com](mailto:erwindosianipar@gmail.com).

---

> Kemuliaan bagi Tuhan saja - Soli Deo Gloria
