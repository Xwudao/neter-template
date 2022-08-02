-- create "prefix_users" table
CREATE TABLE `prefix_users` (`id` bigint NOT NULL AUTO_INCREMENT, `username` varchar(255) NOT NULL, `role` varchar(255) NOT NULL DEFAULT 'user', PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
