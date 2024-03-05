drop database if exists hylark_post;
create database hylark_post;
use hylark_post;

create table post
(
    id            char(36)                    not null comment '自然主键'
        primary key,
    created_at    datetime                    not null comment '创建时间',
    deleted_at    datetime                    null comment '删除时间（软删除）',
    updated_at    datetime                    null comment '更新时间',
    title         varchar(200)                null comment '博文标题',
    cover_image   varchar(255)                null comment '博文标题配图',
    user_id       char(36)                    not null comment '作者id',
    summary       text                        not null comment '博文概览',
    content       text                        not null comment '博文内容',
    category_id  char(36)                    not null comment '隶属哪个归档',
    temperature   bigint unsigned default '0' not null comment '博文热度（排序文章时用）',
    like_count    bigint unsigned default '0' not null comment '博文点赞量',
    view_count    bigint unsigned default '0' not null comment '观看量',
    star_count    bigint unsigned default '0' not null comment '收藏数量',
    comment_count int             default 0   not null,
    share_count   int             default 0   not null comment '分享数量',
    state         tinyint         default 0   not null comment '文章状态：0审核中、1通过、2被举报、3热点文章',
    link_url      varchar(255)                null comment '文章外部链接'
)
    comment '博文';

create index deleted_at
    on post (deleted_at);
create table post_like
(
    id         char(36) not null comment '自然主键'
        primary key,
    created_at datetime not null comment '创建时间',
    deleted_at datetime null comment '删除时间（软删除）',
    updated_at datetime null comment '更新时间',
    comment_id char(36) not null comment '评论id（考虑性能，不加外键）',
    user_id    char(36) not null comment '点赞人（考虑性能，不加外键）'
)
    comment '评论-点赞表';

create index deleted_at
    on post_like (deleted_at);

create table category
(
    id             char(36)     not null comment '自然主键'
        primary key,
    created_at     datetime     not null comment '创建时间',
    deleted_at     datetime     null comment '删除时间（软删除）',
    updated_at     datetime     null comment '更新时间',
    name           varchar(50)  not null comment '归档名称',
    background_url varchar(255) not null comment '背景图片url',
    ranking         tinyint      not null comment '板块排序权重',
    plate_id       char(36)     null comment '板块id',
    url            varchar(255) not null comment '跳转url地址',
    icon           varchar(255) null comment 'icon图标'

)
    comment '归档表（板块）';
create index deleted_at
    on category (deleted_at);

create table tag
(
    id           char(36)    not null comment '自然主键'
        primary key,
    created_at   datetime    not null comment '创建时间',
    deleted_at   datetime    null comment '删除时间（软删除）',
    updated_at   datetime    null comment '更新时间',
    name         varchar(20) not null comment '标签名',
    category_id char(36)    null comment '该标签隶属的归档'

)
    comment '标签表';
create index deleted_at
    on tag (deleted_at);


create table post_tag
(
    id         char(36) not null comment '自然主键'
        primary key,
    created_at datetime not null comment '创建时间',
    deleted_at datetime null comment '删除时间（软删除）',
    updated_at datetime null comment '更新时间',
    tag_id     char(36) null comment '文章所属标签',
    post_id    char(36) null comment '博文id'
);

create index deleted_at
    on post_tag (deleted_at);

