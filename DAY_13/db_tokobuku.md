CREATE DATABASE db_tokobuku;

USE db_tokobuku;

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
    stok_buku int NOT NULL,
    harga int,
    FOREIGN KEY(penerbit_kode) REFERENCES penerbit(kode_penerbit),
    FOREIGN KEY(kategori_kode) REFERENCES kategori_buku(kode_kategori));
    
CREATE TABLE suplier (
	kode_suplier varchar(3) PRIMARY KEY,
    nama_suplier varchar(30) NOT NULL,
    alamat text);

CREATE TABLE pembelian (
	id int primary key,
    nopenerimaan varchar(5) NOT NULL,
    tgl_penerimaan date,
    kode_suplier varchar(3) NOT NULL,
    keterangan varchar(50),
    total_harga int,
    FOREIGN KEY(kode_suplier) REFERENCES suplier(kode_suplier));
    
CREATE TABLE pembelian_detail (
	id int primary key,
    nopenerimaan varchar(5) NOT NULL,
    kode_buku varchar(5) NOT NULL,
    jumlah_buku int,
    harga int);
    
ALTER TABLE pembelian_detail ADD FOREIGN KEY(kode_buku) REFERENCES buku(kode_buku)
    
    
    
    