create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));

- a implementacao dos drivers do banco nao é responsabiliade do Go
- para usar o driver basta apenas importar o pacote