# API Alkitab Indonesia

## Demo: [https://api-alkitab.herokuapp.com](https://api-alkitab.herokuapp.com)

API Alkitab Indonesia adalah sebuah *third-party backend* Alkitab yang sumber bacaannya berasal dari [sabda.org](https://sabda.org).

[sabda.org](https://sabda.org) adalah sebuah Yayasan Kristen non-profit, non-komersial yang bergerak dalam bidang pelayanan media komputer dan internet di Indonesia.

## Penggunaan

Saat ini (24 Oktober 2020) API masih mempunyai dua fitur utama untuk penampilan pasal dan ayat Alkitab.

### Passage and Chapter

Untuk membaca bagian ayat Alkitab gunakan **passage/{passage}/{chapter}** dengan mengirim nama Ayat dan nomor bagian.

#### Base URL

```bash
https://api-alkitab.herokuapp.com/v1/passage/{passage}/{chapter}
```

Variable | Keterangan | Tipe data
-- | -- | --
passage | Rentang ayat yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string`
chapter | Adalah nomor pasal/bab yang ingin ditampilkan dari bagian Alkitab | `int`

#### Example

Method | URL
-- | --
GET | [https://api-alkitab.herokuapp.com/v1/passage/Yoh/1](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1)

#### Response

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

### Passage, Chapter, and Number

Untuk membaca satu ayat Alkitab gunakan **passage/{passage}/{chapter}/{number}** dengan mengirim nama Pasal dan nomor bagian dan ayat.

#### Base URL

```bash
https://api-alkitab.herokuapp.com/v1/passage/{passage}/{chapter}/{number}
```

Variable | Keterangan | Tipe data
-- | -- | --
passage | Rentang ayat yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string`
chapter | Adalah nomor pasal/bab yang ingin ditampilkan dari bagian Alkitab | `int`
number | Adalah nomor ayat dari bagian pasal/bab yang ingin ditampilkan | `int`

#### Example

Method | URL
-- | --
GET | [https://api-alkitab.herokuapp.com/v1/passage/Yoh/1/1](https://api-alkitab.herokuapp.com/v1/passage/Yoh/1/1)

#### Response

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
