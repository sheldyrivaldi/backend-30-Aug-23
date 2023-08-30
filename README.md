# ERD Rumah Makan

![ERD Rumah Makan](/backend-30-Aug-23/ERD%20Rrumah%20Makan.png)

## Requirements

1. Cashier bisa menambahkan pesanan pelanggan ke dalam keranjang.
2. Aplikasi bisa mendapatkan data struk dari tabel _Receipt_.
3. Laporan bisa didapatkan dengan mengambil data dari tabel _Income_ yang berisi daftar pemasukan dikurangi dengan tabel _Expense_ dan mengambil data 7 dan 30 hari berdasarkan kolom created_at.
4. Laporan stock bisa dibagi 2 dengan mengambil 2 tabel yaitu pada tabel _Menu_ kolom _stock_ dan tabel _Ingredients_ kolom _stock_.
5. Menambahkan fitur loyalitas dengan menjadi member pada kolom member.
6. Table payment untuk agar buyer dapat memilih metode pembayaran dan pendataan yang lebih detail.

## Techstack

Aplikasi ini akan saya buat menggunakan basis mobile karena saya melihat efisiensi dari mobile pada rumah makan jauh lebih baik dari versi desktop atau pun web. Dan budgeting juga tidak membengkak. Cukup dengan tablet spesifikasi standar sudah bisa menjalankan aplikasi ini. Techstack yang akan digunakan yaitu :

#### 1. Frontend

    Untuk Frontend, saya akan menggunakan react native. Karena dengan react native saya bisa membuat aplikasi android dan ios sekaligus dengan kode yang sama sehingga bisa menghemat waktu dan usaha dalam pengembangan.

#### 2. Backend

    Untuk Backend, saya akan menggunakan golang dengan framework echo. Karena golang memiliki performa dan efisiensi memory yang baik. Selain itu golang merupakan static type dan dalam pembelajarannya tidak sesulit bahasa static type yang lain seperti java dan Ruby, sehingga cocok untuk pengelolaan keuangan yang mana keuangan sangat sensitif dan jika terjadi kesalahan bisa berimbas pada keuangan bisnis. Dan Echo merupakan framework yang paling popular dan cepat menurut data yang saya cari dari beberapa website.

#### 3. DBMS

    Untuk database, saya akan mengunakan type SQL yaitu PostgreSQL. Karena core bisnis dari aplikasi ini sudah jelas yaitu rumah makan. Maka lebih baik mengunakan SQL karena lebih terstruktur dan terdefinisi dengan baik. Sehingga meminimalisir resiko kesalahan pencatatan keuangan. Dan saya memilih PostgreSQL karena memiliki banyak fitur yang tidak ada pada RDBMS lain seperti tipe data JSONB untuk menyimpan data json.

# Permainan Dadu

Untuk script permainan dadu bisa di cek pada file main.go. Jika ingin menjalankan permainan dadu, ketik pada terminal :

    go run main.go

maka akan muncul input player seperti ini :

    Masukan jumlah player : (masukan disini)

Setelah input player, maka akan diminta untuk input jumlah dadu seperti ini :

    Masukan jumlah dadu : (Masukan disini)

Selanjutnya game akan berjalan dan menentukan pemenangnya :

    Masukkan jumlah player: 3
    Masukkan jumlah dadu: 4

    Lemparan Ke - 1 [[6 2 4 3] [1 5 3 5] [6 5 6 3]]
    Setelah evaluasi [[2 4 3] [5 3 5] [1 5 3]]
    Score [1 0 2]

    Lemparan Ke - 2 [[5 2 2] [3 6 4] [2 2 2]]
    Setelah evaluasi [[5 2 2] [3 4] [2 2 2]]
    Score [1 1 2]

    Lemparan Ke - 3 [[5 5 2] [2 4] [6 2 1]]
    Setelah evaluasi [[5 5 2 1] [2 4] [2]]
    Score [1 1 3]

    Lemparan Ke - 4 [[5 1 5 2] [5 3] [4]]
    Setelah evaluasi [[5 5 2] [1 5 3] [4]]
    Score [1 1 3]

    Lemparan Ke - 5 [[5 5 4] [4 4 4] [2]]
    Setelah evaluasi [[5 5 4] [4 4 4] [2]]
    Score [1 1 3]

    Lemparan Ke - 6 [[6 2 3] [3 3 5] [4]]
    Setelah evaluasi [[2 3] [3 3 5] [4]]
    Score [2 1 3]

    Lemparan Ke - 7 [[1 4] [3 5 6] [4]]
    Setelah evaluasi [[4] [1 3 5] [4]]
    Score [2 2 3]

    Lemparan Ke - 8 [[3] [6 4 4] [2]]
    Setelah evaluasi [[3] [4 4] [2]]
    Score [2 3 3]

    Lemparan Ke - 9 [[6] [4 3] [2]]
    Setelah evaluasi [[] [4 3] [2]]
    Score [3 3 3]

    Lemparan Ke - 10 [[] [3 4] [6]]
    Setelah evaluasi [[] [3 4] []]
    Score [3 3 4]

    Pemenangnya adalah Player 3
