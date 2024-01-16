-- Tabel untuk menyimpan users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE
);

-- Insert data contoh untuk tabel users
INSERT INTO users (username) VALUES
    ('user1'),
    ('user2'),
    ('user3'),
    ('user4'),
    ('user5'),
    ('user6');

-- Table Purchases
CREATE TABLE purchases (
    purchase_id SERIAL PRIMARY KEY,
    user_id INT,
    amount DECIMAL(10, 2),
    purchase_date DATE
);

-- Insert data contoh
INSERT INTO purchases (user_id, amount, purchase_date) VALUES
    (1, 100.50, '2024-01-01'),
    (2, 150.75, '2024-01-02'),
    (1, 75.25, '2024-01-03'),
    (3, 200.00, '2024-01-04'),
    (2, 120.30, '2024-01-05'),
    (1, 90.60, '2024-01-06'),
    (4, 120.50, '2024-01-07'),
    (4, 80.25, '2024-01-08'),
    (5, 150.00, '2024-01-07'),
    (5, 90.75, '2024-01-08'),
    (6, 200.00, '2024-01-08'),
    (6, 50.30, '2024-01-09');


-- Buat materialized view untuk peringkat top spender secara umum berdasarkan total pembelian
CREATE MATERIALIZED VIEW top_spenders_ranking AS
SELECT
    u.username,
    SUM(p.amount) AS total_spent,
    ROW_NUMBER() OVER (ORDER BY SUM(p.amount) DESC) AS ranking
FROM purchases p
JOIN users u ON
p.user_id = u.user_id
GROUP BY u.username;

-- Fungsi ROW_NUMBER() dengan klausa OVER digunakan 
-- untuk memberikan nomor urut atau peringkat pada setiap baris dari hasil query berdasarkan urutan tertentu. 
-- Pada contoh :
-- ROW_NUMBER() OVER (ORDER BY SUM(p.amount) DESC) AS ranking

-- Ini berarti kita menghitung nomor urut (peringkat) 
-- untuk setiap baris hasil query berdasarkan jumlah total pembelian (SUM(p.amount)) dari besar ke kecil (DESC). Jadi, pengguna dengan total pembelian terbanyak akan memiliki peringkat 1, yang kedua terbanyak akan memiliki peringkat 2, dan seterusnya.

-- Mari kita bahas lebih detail:

-- ROW_NUMBER(): Ini adalah fungsi analitik yang menghitung 
-- nomor urut atau peringkat setiap baris dalam kelompok hasil yang diurutkan.

-- OVER: Klausa ini memberikan informasi tambahan kepada fungsi 
-- analitik tentang bagaimana mengurutkan dan mengelompokkan data. 
-- Dalam hal ini, kita mengurutkan data berdasarkan jumlah total pembelian (SUM(p.amount)) secara descending (DESC).

-- (ORDER BY SUM(p.amount) DESC): Menentukan urutan pengurutan data 
-- yang akan digunakan untuk menghitung peringkat. 
-- Dalam hal ini, kita mengurutkan berdasarkan jumlah total pembelian dari besar ke kecil.

-- AS ranking: Memberikan alias "ranking" pada hasil perhitungan peringkat.

-- Contoh penggunaan ROW_NUMBER() seringkali muncul 
-- ketika kita ingin memberikan peringkat pada hasil query 
-- berdasarkan suatu kriteria tertentu, seperti jumlah total pembelian pada contoh di atas.

-- Buat indeks untuk meningkatkan kinerja query
CREATE INDEX idx_overall_ranking ON top_spenders_ranking(ranking);

SELECT username, total_spent
FROM top_spenders_ranking
WHERE ranking <= 3
ORDER BY ranking;

REFRESH MATERIALIZED VIEW top_spenders_ranking;