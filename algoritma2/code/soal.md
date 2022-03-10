# Soal Latihan Algoritma 2

## 1. Candles
Level : `easy`

Pada sebuah kue, terdapat beberapa lilin
dengan tinggi yang berbeda.

Saat kamu meniup lilin tersebut, maka yang kamu tiup adalah
lilin paling tinggi. Jadi yang akan padam adalah lilin paling tinggi.

Diketahui :
Tinggi lilin = `[3,2,4,4,1]`

Ditanya :
Berapa lilin yang akan padam ?

Jawab :
Pada lilin tersebut, yang paling tinggi adalah 4
Dan jumlah lilin yang mempunyai tinggi 4 adalah 2 buah

Maka, lilin yang akan padam sejumlah 2 buah

## 2. Count Balls
Level : `easy`

Dalam sebuah kantong, terdapat beberapa bola.
Ada 3 warna bola, yaitu Merah, Kuning, dan Biru.

Tampilkanlah jumlah bola paling banyak

Diketahui :
```bash
bola = [Merah, Biru, Biru, Kuning, Merah]
warna = [Merah, Biru, Kuning]
```

Ditanya :
Bola apa yang paling banyak?
Dan berapa jumlahnya?

Jawab :
```bash
Warna merah = 2
Warna biru = 2
Warna kuning = 1
```

Karena merah dan biru sama, maka outputnya adalah :
```bash
merah : 2, biru : 2
```

## 3. Jarak Terdekat [BONUS]
Level : `medium`

Lihatlah gambar berikut :
![img](https://res.cloudinary.com/noobeeid/image/upload/q_50/v1646911170/blog/courses/golang/Screen_Shot_2022-03-10_at_18.17.41_aibrwv.png)

Carilah jarak terdekat, dari kondisi berikut :
- Dari A -> E
- Dari C -> H

Jawab :
- Dari A -> E
    
    Kita perlu melihat path apa yang dilewati terlebih dahulu. Adapun path nya adalah sebagai berikut : 

    1. A -> B -> D -> E ... Total = 5 
    2. A -> C -> D -> E ... Total = 5
    3. A -> C -> F -> E ... Total = 6

    Maka, hasilnya adalah :
    ```bash
    [
        0 : {
            "path" : "A, B, D, E",
            "jarak" : 5
        },
        1 : {
            "path" : "A, C, D, E",
            "jarak" : 5
        },
    ]
    ```

- Dari C -> H

    Kita perlu melihat path apa yang dilewati terlebih dahulu. Adapun path nya adalah sebagai berikut : 

    1. C -> D -> E -> G -> H ... Total = 5 
    2. C -> F -> G -> H      ... Total = 6
    3. C -> F -> E -> G -> H ... Total = 6

    Maka, hasilnya adalah :
    ```bash
    [
        0 : {
            "path" : "C, D, E, G, H",
            "jarak" : 5
        },
    ]
    ```

> Hasil dibuat dalam bentuk `map[int]map[string]interface{}`