CREATE TABLE `shorturl`
(
    `short_link` varchar(255) NOT NULL COMMENT '短连接',
    `real_url`   varchar(255) NOT NULL COMMENT '真实连接',
    PRIMARY KEY (`short_link`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--  goctl model mysql ddl --src .\1shorturl.sql --dir .
-- [convertColumns]: The column "real_url" is recommended to add constraint `DEFAULT`
-- goctl model mysql ddl --src .\1shorturl.sql --dir . --cache 带缓存模式