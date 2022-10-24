create table `user`
(
    `id`          bigint(20) primary key auto_increment,
    `gender`      tinyint(4) not null default 0,
    `user_id`     bigint(20) not null,
    `username`    varchar(64) not null,
    `password`    varchar(64) not null,
    `email`       varchar(64) default null,
    `create_time` timestamp   default 'current_timestamp',
    `update_time` timestamp   default 'current_timestamp',
    unique key 'idx_username' ('username') using btree
)engine=innodb default charset=utf8mb64;
