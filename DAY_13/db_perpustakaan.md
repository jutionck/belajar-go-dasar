CREATE DATABASE dB_perpustakaan;

USE dB_perpustakaan;

CREATE TABLE kategori_buku (
	kode_kategori varchar(3) PRIMARY KEY,
    nama_kategori varchar(15) NOT NULL);
    
CREATE TABLE penerbit (
	kode_penerbit varchar(3) PRIMARY KEY,
    nama_penerbit varchar(25) NOT NULL);

CREATE TABLE buku (
	kode_buku varchar(5) PRIMARY KEY,
    isbn varchar(13) NOT NULL,
    judul_buku varchar(50) NULL,
    penulis varchar(35) NOT NULL,
    penerbit_kode varchar(3) NOT NULL,
    tahun varchar(4),
    kategori_kode varchar(3),
    FOREIGN KEY(penerbit_kode) REFERENCES penerbit(kode_penerbit),
    FOREIGN KEY(kategori_kode) REFERENCES kategori_buku(kode_kategori));

CREATE TABLE anggota (
	kode_anggota varchar(4) PRIMARY KEY,
    nama_anggota varchar(35) NOT NULL);
    
CREATE TABLE denda (
	kode_denda varchar(3) PRIMARY KEY,
    jenis_denda varchar(10),
    bayar_denda int);

CREATE TABLE peminjaman (
	kode_peminjaman varchar(5) PRIMARY KEY,
    kode_anggota varchar(4),
    tgl_peminjaman date,
    tgl_pengembalian date,
    FOREIGN KEY(kode_anggota) REFERENCES anggota(kode_anggota));

ALTER TABLE peminjaman ADD kode_buku varchar(5) NOT NULL;

ALTER TABLE peminjaman ADD FOREIGN KEY(kode_buku) REFERENCES buku(kode_buku);

CREATE TABLE pengembalian (
	kode_pengembalian varchar(5) PRIMARY KEY,
    kode_anggota varchar(4),
    tgl_pengembalian date,
    kode_buku varchar(5) NOT NULL,
    kode_denda varchar(3) NOT NULL,
    FOREIGN KEY(kode_anggota) REFERENCES anggota(kode_anggota),
	FOREIGN KEY(kode_buku) REFERENCES buku(kode_buku),
	FOREIGN KEY(kode_denda) REFERENCES denda(kode_denda))
