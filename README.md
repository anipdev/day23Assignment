# day23Assignment

Sebuah backend API yang dibangun untuk assignment day 23 dengan bahasa Go untuk mengelola produk, inventori, dan pesanan.

## Project Structure

```
day23Assignment/
├── config/ # Database configuration
├── controller/ # Handler functions for each endpoint
├── model/ # Database models and queries
├── router/ # Route definitions
└── main.go # App entry point
```

## Features

### Product Management

GET /products – Mendapatkan semua produk

GET /products/:id – Mendapatkan detail produk berdasarkan ID

POST /products – Menambahkan produk baru

PUT /products/:id – Memperbarui data produk berdasarkan ID

DELETE /products/:id – Menghapus produk berdasarkan ID

### Stock Management

GET /stock – Melihat stok saat ini

PUT /stock – Memperbarui stok produk

### Order Management

GET /orders – Mendapatkan semua pesanan

POST /orders – Membuat pesanan baru

## Environment Variables

Buat file .env di root project. Contoh isinya:

```
DB_USER=user
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=Day23_Assignment
```

Pastikan nilai variabel sesuai dengan konfigurasi MySQL Anda

## Database Definition

Buat database MySQL menggunakan DDL berikut:

```
CREATE TABLE `products` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(100),
`description` TEXT,
`price` DECIMAL(10,2),
`category` VARCHAR(50),
PRIMARY KEY (`id`)
);

CREATE TABLE `inventories` (
`id` INT NOT NULL AUTO_INCREMENT,
`product_id` INT,
`quantity` INT,
`location` VARCHAR(100),
PRIMARY KEY (`id`),
FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

CREATE TABLE `orders` (
`id` INT NOT NULL AUTO_INCREMENT,
`product_id` INT,
`quantity` INT,
`datetime` DATETIME,
PRIMARY KEY (`id`),
FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);
```

## Running the App

Pastikan MySQL berjalan dan database sudah dibuat sesuai DDL.

Pastikan sudah membuat file .env di root folder.

Jalankan aplikasi dengan perintah:

```
go run main.go
```
