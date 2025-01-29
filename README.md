# Deklarasi dan Inisialisasi di Golang

Golang memiliki dua cara utama untuk membuat sesuatu: **Deklarasi** dan **Inisialisasi**. Konsep ini berlaku untuk variabel, array, slice, map, struct, dan lainnya. Berikut adalah penjelasan lengkapnya:

---

## 1. Deklarasi
Deklarasi adalah langkah untuk **menyiapkan wadah tanpa langsung mengisi nilai**. Dengan deklarasi, kita hanya membuat "tempat" di memori untuk data tertentu. Biasanya nilai yang dihasilkan akan berupa **default value** berdasarkan tipe data.

### Contoh Deklarasi:
```go
var numbers [5]int // Array dengan 5 elemen bertipe int, default nilainya 0
var name string    // String, default nilainya ""
var price float64  // Float, default nilainya 0.0
```

**Ciri-ciri Deklarasi:**
- Menggunakan keyword `var`.
- Tidak langsung mengisi nilai.
- Nilai elemen akan menggunakan default value (contoh: 0 untuk angka, "" untuk string).
- Cocok jika data belum diketahui saat deklarasi.

---

## 2. Inisialisasi
Inisialisasi adalah proses **deklarasi sekaligus memberikan nilai awal**. Pendekatan ini lebih ringkas dan sering digunakan untuk data yang sudah diketahui nilainya.

### Contoh Inisialisasi:
```go
numbers := [5]int{10, 20, 30, 40, 50} // Array langsung ada isinya
name := "Golang"                     // String langsung ada nilai
price := 9.99                        // Float langsung ada nilai
```

**Ciri-ciri Inisialisasi:**
- Bisa menggunakan pendekatan `:=` (short declaration).
- Langsung memberikan nilai saat dideklarasikan.
- Cocok untuk data yang sudah jelas nilainya sejak awal.

---

## Perbandingan Deklarasi dan Inisialisasi
| **Deklarasi**              | **Inisialisasi**                  |
|----------------------------|-----------------------------------|
| Hanya membuat wadah tanpa nilai | Membuat wadah dan langsung mengisi nilai |
| Menggunakan keyword `var`  | Bisa menggunakan `:=` atau `var` |
| Nilai default sesuai tipe data | Nilai langsung sesuai kebutuhan |

---

## Contoh Gabungan:
Berikut adalah contoh kombinasi antara deklarasi dan inisialisasi:

```go
package main

import "fmt"

func main() {
    // Deklarasi
    var numbers [3]int // Array default: [0, 0, 0]
    var name string    // String default: ""
    var price float64  // Float default: 0.0

    // Inisialisasi
    numbers = [3]int{10, 20, 30} // Assign nilai ke array
    name = "Golang"
    price = 99.99

    fmt.Println(numbers) // Output: [10 20 30]
    fmt.Println(name)    // Output: Golang
    fmt.Println(price)   // Output: 99.99
}
```

---

## Latihan
Untuk memperdalam pemahaman, cobalah latihan berikut:

1. Buat **slice** `[]string` menggunakan deklarasi terlebih dahulu, kemudian isi nilainya di langkah berikutnya.
2. Buat **map** `map[string]int` menggunakan inisialisasi langsung.

Semoga penjelasan ini membantu! 🚀 Jangan ragu untuk bereksperimen dan mencoba kode di atas. Jika ada pertanyaan, feel free untuk berdiskusi! 😊

