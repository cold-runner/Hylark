drop database if exists hylark_user;
create database hylark_user;
use hylark_user;

create table lark
(
    id              char(36)     not null comment '自然主键'
        primary key,
    created_at      datetime     not null comment '创建时间',
    deleted_at      datetime     null comment '删除时间（软删除）',
    updated_at      datetime     null comment '更新时间',
    stu_num         char(8)      unique not null comment '学号',
    password        char(255)   not null comment '密码',
    name            varchar(30)  not null comment '姓名',
    gender          varchar(10)  default '保密' not null comment '用户性别：女，男，其他，保密',
    college         varchar(30)  not null comment '用户所在学院',
    major           varchar(30)  not null comment '用户专业',
    grade           varchar(10)  not null comment '用户年级：大一，大二，大三，大四，研究生,毕业生',
    stu_card_url    varchar(255) unique not null comment '学生证照片url',
    phone           char(11)     unique not null comment '用户手机号',
    province        varchar(10)  null comment '用户家乡省份',
    age             tinyint      null comment '用户年龄',
    photo_url       varchar(255) null comment '照片url',
    email           varchar(255) unique null comment '邮箱地址',
    introduce       text         null comment '用户个人介绍',
    avatar          varchar(255) default 'https://static.skylab.org.cn/default/avatar.png' null comment '用户头像url',
    qq_union_id     varchar(255) unique null comment 'qq社会化登录',
    wechat_union_id varchar(255) unique null comment '微信社会化登录',
    state           tinyint      default 0 not null comment '用户状态：0禁用、1审核中、2启用、3其他'
)
    comment '用户表';
create index deleted_at
    on lark (deleted_at);

insert into lark (id, created_at, deleted_at, updated_at, stu_num, password, name, gender, college, major, grade,
                  stu_card_url, phone, province, age, photo_url, email, introduce, avatar, qq_union_id, wechat_union_id,
                  state)
values ("84a392ab-4426-4f3f-b7bf-d3dbdc3f21bb","2024-03-01 10:12:45",null,null,"20999004","salkdjflkasdjfl;a","郑赫","男","软件学院", "计算机科学与技术","大四","https://static.skylab.org.cn/test.jpg","18342728255","辽宁",23,null,"zzzheng80@gmail.com","即将毕业，祝安好",null,null,null,1);

create table user_interaction
(
    id          char(36) not null comment '自然主键'
        primary key,
    created_at  datetime not null comment '创建时间',
    deleted_at  datetime null comment '删除时间（软删除）',
    updated_at  datetime null comment '更新时间',
    user_id     char(36) not null comment 'subject',
    followed_id char(36) not null comment 'object'
)
    comment '社交关系表';

create index deleted_at
    on user_interaction (deleted_at);

