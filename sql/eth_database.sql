DROP DATABASE IF EXISTS one_chain_db;
CREATE DATABASE IF NOT EXISTS one_chain_db;

CREATE TABLE IF NOT EXISTS one_chain_db.t_coins_info (
    `id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `name_id`           VARCHAR(256)        NOT NULL DEFAULT ''     COMMENT '别名，币虎id',
    `name`              VARCHAR(256)        NOT NULL DEFAULT ''     COMMENT '代币名称',
    `symbol`            VARCHAR(256)        NOT NULL DEFAULT ''     COMMENT '代币符号',
    `decimals`          SMALLINT UNSIGNED   NOT NULL DEFAULT 1      COMMENT '代币位数',
    `contract`          CHAR(42)            NOT NULL DEFAULT ''     COMMENT '代币合约地址',
    `tx_hash`           CHAR(64)            NOT NULL DEFAULT ''     COMMENT '代币交易哈希',
    `price`             FLOAT               NOT NULL DEFAULT 0      COMMENT '代币价格，USD',
    `icon`              TEXT                NOT NULL                COMMENT '代币图标',
    `supply`            VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '代币发布总数',
    PRIMARY KEY (`id`),
    UNIQUE KEY `contract`(`contract`),
    KEY `hash`(`tx_hash`(10))    USING BTREE ,
    KEY `symbol`(`symbol`)   USING BTREE ,
    KEY `name`(`name`)   USING BTREE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
