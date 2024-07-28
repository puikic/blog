create table if not exists user(
  id int auto_increment comment '用户id,自增',
  name varchar(20) not null comment '用户名',
  password char(32) not null comment '密码的md5值',
  primary key (id),
  unique key idx_name (name)
) default charset = utf8mb4 comment '用户登陆表';

insert into
  user(name, password)
values
  ('cpq', 'cbff36039c3d0212b3e34c23dcde1456'),
  ('dqq', 'e10adc3949ba59abbe56e057f20f883e');

create table if not exists blog(
  id int auto_increment comment '博客id,自增',
  user_id int not null comment '作者id',
  title varchar(100) not null comment '标题',
  article text not null comment '正文',
  create_time datetime default current_timestamp comment '创建时间',
  update_time datetime default current_timestamp on update current_timestamp comment '最后修改时间',
  primary key (id),
  key idx_user (user_id)
) default charset = utf8mb4 comment '博客内容';