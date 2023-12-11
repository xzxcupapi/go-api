
-- Tabel Pelanggan
CREATE TABLE IF NOT EXISTS pelanggan (
id SERIAL PRIMARY KEY,
nama VARCHAR(255) NOT NULL,
nomorTelepon VARCHAR(20) NOT NULL,
alamat TEXT NOT NULL
);

-- Tabel Produk
CREATE TABLE IF NOT EXISTS produk (
id SERIAL PRIMARY KEY,
nama VARCHAR(255) NOT NULL,
harga INT NOT NULL,
satuan VARCHAR(20) NOT NULL
);

-- Tabel Karyawan
CREATE TABLE IF NOT EXISTS karyawan (
id SERIAL PRIMARY KEY,
nama VARCHAR(255) NOT NULL,
nomorTelepon VARCHAR(20) NOT NULL,
alamat TEXT NOT NULL
);

-- Tabel Transaksi
CREATE TABLE IF NOT EXISTS transaksi (
id SERIAL PRIMARY KEY,
tanggalTagihan DATE NOT NULL,
tanggalMasuk DATE NOT NULL,
tanggalSelesai DATE NOT NULL,
idKaryawan INT REFERENCES karyawan(id) NOT NULL,
idPelanggan INT REFERENCES pelanggan(id) NOT NULL
);

-- Tabel DetailTransaksi
CREATE TABLE IF NOT EXISTS detailTransaksi (
id SERIAL PRIMARY KEY,
idTagihan INT REFERENCES transaksi(id) NOT NULL,
idProduk INT REFERENCES produk(id) NOT NULL,
jumlah INT NOT NULL CHECK (jumlah > 0),
harga INT NOT NULL
);