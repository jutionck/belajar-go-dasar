CREATE DATABASE db_akademik;

USE db_akademik;

CREATE TABLE programstudi (
	prodi_id varchar(3) PRIMARY KEY,
    namaprodi varchar(30) NOT NULL,
    jurusan_id varchar(3) NOT NULL,
    FOREIGN KEY(jurusan_id) REFERENCES jurusan(jurusan_id)
    );

CREATE TABLE jurusan (
	jurusan_id varchar(3) PRIMARY KEY,
    nama_jurusan varchar(25) NOT NULL);

CREATE TABLE mahasiswa (
	nrp varchar(8) primary key,
    nama_mhs varchar(35) NOT NULL,
    email varchar(50) NOT NULL,
    no_hp varchar(13),
    jenis_kelamin ENUM('T','F'),
    tempat_lahir varchar(50),
    tgl_lahir date,
    alamat text,
    jurusan_id varchar (3),
    programstudi_id varchar(3),
    status_mahasiswa ENUM('Aktif','Lulus','Cuti','DO') default 'Aktif',
    FOREIGN KEY(jurusan_id) REFERENCES jurusan(jurusan_id),
    FOREIGN KEY(programstudi_id) REFERENCES programstudi(prodi_id)
    )

    
    