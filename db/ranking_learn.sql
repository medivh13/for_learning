-- Table Purchases
CREATE TABLE public.contoh_rank (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    task_selesai INT
);

insert into public.contoh_rank (name, task_selesai) values ('Jody', 20);
insert into public.contoh_rank (name, task_selesai) values ('Budi', 20);
insert into public.contoh_rank (name, task_selesai) values ('Fahrul', 17);
insert into public.contoh_rank (name, task_selesai) values ('Agung', 17);
insert into public.contoh_rank (name, task_selesai) values ('Adhan', 15);


DENSE_RANK ()

Memberi peringkat pada data yang sama pada peringkat yang sama, dan tidak melewatkan nomor peringkat apa pun;

Select name, task_selesai, DENSE_RANK() over (Order By task_selesai DESC) as dense_rank_ranking from public.contoh_rank;


RANK()

Memberi peringkat pada data yang sama dengan peringkat yang sama, dan melewatkan peringkat berikutnya

Select name, task_selesai, RANK() over (Order By task_selesai DESC) as rank_ranking from public.contoh_rank;

ROW_NUMBER()
Seperti namanya, hanya memberi nomor pada baris berdasarkan urutan yang ditentukan

Select name, task_selesai, ROW_NUMBER() over (Order By task_selesai DESC) as row_number_ranking from public.contoh_rank;

Contoh lain :

WITH RankedEmployee as (
  select name, task_selesai, [Rank_Method] over (order by task_selesai desc) as rank from public.contoh_rank
)
Select name, task_selesai, rank from RankedEmployee where rank <=3;

