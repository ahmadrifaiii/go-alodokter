CREATE TABLE IF NOT EXISTS IF NOT EXISTS users (
	user_id varchar(50) PRIMARY KEY NOT NULL,
	user_name varchar(100) NOT NULL,
	user_password varchar(100) NOT NULL,
	user_email varchar(100) NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE TABLE IF NOT EXISTS user_roles (
	user_role_id varchar(50) PRIMARY KEY NOT NULL,
	user_role_name varchar(100) NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE TABLE IF NOT EXISTS user_role_maps (
	user_role_map_id varchar(50) PRIMARY KEY NOT NULL,
	user_id varchar(100) NOT NULL,
	user_role_id varchar(100) NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE TABLE IF NOT EXISTS user_role_acceses (
	user_role_access_id varchar(50) PRIMARY KEY NOT NULL,
	user_role_access_url varchar(100) NOT NULL,
	user_role_id varchar(100) NOT NULL,
    is_access_allowed boolean NOT NULL default true,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE UNIQUE  INDEX 
index_table_user_role_access
ON user_role_acceses (user_role_access_url, user_role_access_id);


CREATE TABLE IF NOT EXISTS products (
	product_id varchar(50) PRIMARY KEY NOT NULL,
	product_sku varchar(100) NOT NULL,
	product_name varchar(100) NOT NULL,
	product_price double NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE TABLE IF NOT EXISTS orders (
	order_id varchar(50) PRIMARY KEY NOT NULL,
	order_number varchar(100) NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

CREATE TABLE IF NOT EXISTS order_items (
	order_item_id varchar(50) PRIMARY KEY NOT NULL,
	order_id varchar(100) NOT NULL,
	product_price double NOT NULL,
	product_id varchar(100) NOT NULL,
	deleted_at timestamp,
	deleted_by varchar(100),
	created_at timestamp default current_timestamp,
	created_by varchar(100),
	updated_at timestamp ON UPDATE current_timestamp,
	updated_by varchar(100)
);

INSERT INTO user_role_acceses (user_role_access_id,user_role_access_url,user_role_id,deleted_at,deleted_by,created_at,created_by,updated_at,updated_by,is_access_allowed) VALUES
	 ('605a71db-43d8-4bf9-a18e-dd01ba416751','/api/v1/product/detail/*','d490acfc-6466-4bae-989c-434625767063',NULL,NULL,'2022-07-27 23:10:29',NULL,NULL,NULL,1),
	 ('75aa15ed-1686-4f1b-83fc-b300768192a7','/api/v1/product/list','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,NULL,NULL,1),
	 ('7b62f452-fc0c-4f82-9a15-dec948ee2236','/api/v1/user/delete/*','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1),
	 ('828020ae-2cd3-4c95-a68b-f40c0943bd43','/api/v1/user/update/*','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1),
	 ('84ccc70f-c993-463b-8874-9c7d4c221ddc','/api/v1/user/detail/*','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1),
	 ('9b3ae43b-3be9-4fbb-8c21-3700077f135d','/api/v1/order/list','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1),
	 ('a2ae3f5a-b889-4ca8-9b55-6ffe17a4ccc1','/api/v1/product/list','d490acfc-6466-4bae-989c-434625767063',NULL,NULL,'2022-07-27 23:10:29',NULL,NULL,NULL,1),
	 ('b8126751-8662-4281-b56d-43012bb3c505','/api/v1/product/detail/*','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,NULL,NULL,1),
	 ('cd383327-5462-4732-b29c-2587ed1713c1','/api/v1/user/list','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1),
	 ('e0d16661-b399-438c-8db8-14a1eeb02bfb','/api/v1/order/detail/*','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 23:10:29',NULL,'2022-07-28 01:05:33',NULL,1);

INSERT INTO users (user_id,user_name,user_password,user_email,deleted_at,deleted_by,created_at,created_by,updated_at,updated_by) VALUES
	 ('e574d16c-891b-47c9-a220-84b39973eb69','user_example_satu','$2a$10$2MnyeJCuzTsIhCuAyBCGnOaRfEvKHie5SzVSQitdNUQAGmYUo762W','user_example_satu@mailinator.com',NULL,NULL,'2022-07-27 12:04:39',NULL,NULL,NULL);

INSERT INTO user_roles (user_role_id,user_role_name,deleted_at,deleted_by,created_at,created_by,updated_at,updated_by) VALUES
	 ('29b84c83-6f46-4667-9182-c04a4dbdd513','Super Role',NULL,NULL,'2022-07-27 13:12:13',NULL,NULL,NULL),
	 ('d490acfc-6466-4bae-989c-434625767063','Product Role',NULL,NULL,'2022-07-27 13:12:13',NULL,NULL,NULL);

INSERT INTO user_role_maps (user_role_map_id,user_id,user_role_id,deleted_at,deleted_by,created_at,created_by,updated_at,updated_by) VALUES
	 ('29b84c83-6f46-4667-9182-c04a4dbdd513','e574d16c-891b-47c9-a220-84b39973eb69','29b84c83-6f46-4667-9182-c04a4dbdd513',NULL,NULL,'2022-07-27 13:13:20',NULL,NULL,NULL);
