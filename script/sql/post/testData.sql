insert into post
(id, created_at, deleted_at, updated_at, title, cover_image, user_id, summary, content, category_id, temperature, like_count, view_count, star_count, comment_count, share_count, state, link_url)
values
("524cf250-0c73-4923-a4ec-7ec661d56ad0", "2024-03-02 17:26:11", null, null, "testTitle1","https://static.skylark.org/test.jpg","8c542ffc-2a86-423f-bd0f-67b2ea81bef6","summary1","some content", "8375e464-48aa-4720-8a9f-27d66942669d",0,0,0,0,0,0,0,null);

insert into post
(id, created_at, deleted_at, updated_at, title, cover_image, user_id, summary, content, category_id, temperature, like_count, view_count, star_count, comment_count, share_count, state, link_url)
values
("9ef04b43-bca5-45c7-83fb-8142972d5123", "2024-03-01 11:33:11", null, null, "testTitle2","https://static.skylark.org/test.jpg","08bbb8cb-2450-46ca-8c2a-bfa3b661b071","summary2","some content2", "a920220e-3081-41f4-b731-6674fe777ff7",0,0,0,0,0,0,0,null);

insert into post
(id, created_at, deleted_at, updated_at, title, cover_image, user_id, summary, content, category_id, temperature, like_count, view_count, star_count, comment_count, share_count, state, link_url)
values
("1d03e34e-5b3e-433f-8ae3-87fce7d08dc5", "2024-02-10 10:12:45", null, null, "testTitle3","https://static.skylark.org/test.jpg","84a392ab-4426-4f3f-b7bf-d3dbdc3f21bb","summary3","some content3", "df04770c-59e7-41bc-a18d-07c14b65d58b",0,0,0,0,0,0,0,null);



insert into category (id, created_at, deleted_at, updated_at, name, background_url, ranking, plate_id, url, icon)
values
("8375e464-48aa-4720-8a9f-27d66942669d", "2024-03-02 13:22:10" ,null,null, "软件学院", "https://static.skylab.org.cn/test.jpg", 1, "7943c4d3-2599-47ce-be82-3d009aa65b56", "/software", "<span></span>");

insert into category (id, created_at, deleted_at, updated_at, name, background_url, ranking, plate_id, url, icon)
values
("a920220e-3081-41f4-b731-6674fe777ff7", "2024-03-02 10:20:10" ,null,null, "美术学院", "https://static.skylab.org.cn/test.jpg", 2, "7943c4d3-2599-47ce-be82-3d009aa65b56", "/art", "<span></span>");

insert into category (id, created_at, deleted_at, updated_at, name, background_url, ranking, plate_id, url, icon)
values
("df04770c-59e7-41bc-a18d-07c14b65d58b", "2024-03-02 13:10:59" ,null,null, "国际商学院", "https://static.skylab.org.cn/test.jpg", 3, "7943c4d3-2599-47ce-be82-3d009aa65b56", "/business", "<span></span>");


insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
("d2f4bff2-416e-47b7-8ca1-44239231ca72", "2024-03-02 09:03:51" ,null,null,"计算机网络","8375e464-48aa-4720-8a9f-27d66942669d");
insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
    ("a95e9c1a-fe7f-4ab9-9f94-c64c26ad94ce", "2024-03-02 03:03:51" ,null,null,"操作系统","8375e464-48aa-4720-8a9f-27d66942669d");

insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
    ("710eaa1c-eda6-468d-bd56-dde9538884b2", "2024-03-02 09:22:22" ,null,null,"美术设计","a920220e-3081-41f4-b731-6674fe777ff7");
insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
    ("071efdc6-4461-415c-b507-8c524301d2cc", "2024-03-02 09:22:22" ,null,null,"美术历史","a920220e-3081-41f4-b731-6674fe777ff7");

insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
    ("aab14099-5028-4729-8d0e-ec34614aba78", "2024-03-02 06:01:22" ,null,null,"物流管理","df04770c-59e7-41bc-a18d-07c14b65d58b");
insert into tag (id, created_at, deleted_at, updated_at, name, category_id)
values
    ("440b54e5-3224-487b-b174-8d7fb3aba5b9", "2024-03-02 09:22:22" ,null,null,"商业计算","df04770c-59e7-41bc-a18d-07c14b65d58b");
