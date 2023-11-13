begin;
delete from products; -- 刪除所有商品
alter sequence products_id_seq restart; -- 重設id編號，從1開始
insert into products (name, price) values 
('Americano medium', 35),
('Americano large', 45),
('Latte', 65);
 commit;