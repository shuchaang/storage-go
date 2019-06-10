create table tbl_file
(
  id        int not null primary key auto_increment,
  file_sha1 varchar(40) DEFAULT '',
  file_name varchar(40) DEFAULT '',
  file_size varchar(40) DEFAULT '',
  file_addr varchar(40) DEFAULT '',
  status    int         DEFAULT 0,
  create_at datetime    default now(),
  update_at datetime    default now(),
  ext1      text,
  ext2      int         default 0
)
create table tbl_user
(
  id          int          not null primary key auto_increment,
  user_name   varchar(64)  not null comment '用户名',
  user_pwd    varchar(256) not null comment '密码',
  email       varchar(128)          default '' comment '邮箱',
  phone       varchar(20)           default '' comment '手机号码',
  email_valid tinyint(1) default 0 comment '邮箱是否验证',
  phone_valid tinyint(1) default 0 comment '手机是否验证',
  signup_at   datetime              default current_timestamp comment '注册日期',
  last_active datetime              default current_timestamp on update current_timestamp comment '最后活跃时间',
  profile     text comment '用户属性',
  status      int          not null default 0 comment '账户状态',
  unique key `uk_phone`(phone),
  key         `idx_status`(status)
)

create table tbl_user_token
(
  id         int         not null primary key auto_increment,
  user_name  varchar(64) not null comment '用户名',
  user_token varchar(40) not null comment 'token',
  unique key `uk_token`(user_token),
  key        `idx_user`(user_name)
)