DROP DATABASE IF EXISTS cosmos_db;
CREATE DATABASE IF NOT EXISTS cosmos_db;

CREATE TABLE IF NOT EXISTS cosmos_db.t_block_info (
    `id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `number`            BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '区块高度',
    `hash`              CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前区块的哈希',
    `miner`             CHAR(40)            NOT NULL DEFAULT ''     COMMENT '挖出此区块的矿工',
    `parent_hash`       CHAR(64)            NOT NULL DEFAULT ''     COMMENT '父区块的哈希值',
    `timestamp`         INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '当前区块时间戳',
    `trx_num`           SMALLINT UNSIGNED   NOT NULL DEFAULT 0      COMMENT '当前区块中的交易数量',
    `total_trx_num`     SMALLINT UNSIGNED   NOT NULL DEFAULT 0      COMMENT '链区块中的交易总数量',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash`(`hash`)    USING BTREE ,
    KEY `number`(`number`),
    KEY `miner`(`miner`(10))   USING BTREE ,
    KEY `timestamp`(`timestamp`)   USING BTREE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

#此表单用来给遗漏的区块打补丁，入库的区块会将高度写入此表，定时查询此表，如果高度缺失，执行补漏程序
CREATE TABLE IF NOT EXISTS cosmos_db.t_block_patch_info (
`id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`number`            BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '区块高度',
PRIMARY KEY (`id`),
UNIQUE KEY `number`(`number`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS cosmos_db.t_transaction_info (
    `id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `block_hash`        CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易所在区块的哈希',
    `block_number`      BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易所在区块的高度',
    `timestamp`         INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '交易发生的时间点',
    `from`              CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易发起者',
    `to`                CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易接收者',
    `gas`               BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易提供的gas量',
    `gas_used`          BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易使用的gas量',
    `gas_price`         VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '当前交易gas单价',
    `hash`              CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易哈希值',
    `input_flag`        TINYINT UNSIGNED    NOT NULL DEFAULT 1      COMMENT '标识input数据位置，1表示从数据库取，2表示从链上取',
    `input`             TEXT                NOT NULL                COMMENT '交易输入数据，如果to是合约，这里存放合约调用方法和参数，最大15K',
    `nonce`             BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '交易记数，是一个递增的值，防止交易重放',
    `transaction_index` SMALLINT UNSIGNED   NOT NULL DEFAULT 0      COMMENT '交易在区块中的偏移量',
    `tx_id`             DOUBLE(12,4)        NOT NULL DEFAULT 0      COMMENT '交易的全局ID，用区块高度和交易偏移量组成的浮点数',
    `value`             VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易金额',
    `status`            BIGINT UNSIGNED     NOT NULL DEFAULT 1      COMMENT '标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码',
    `v`                 CHAR(64)            NOT NULL DEFAULT ''     COMMENT '交易签名',
    `s`                 CHAR(64)            NOT NULL DEFAULT ''     COMMENT '交易签名',
    `r`                 CHAR(64)            NOT NULL DEFAULT ''     COMMENT '交易签名',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash` (`hash`),
    KEY `tx_id` (`tx_id`) USING BTREE ,
    KEY `status` (`status`) USING BTREE ,
    KEY `number` (`block_number`) USING BTREE ,
    KEY `from` (`from`(10)) USING BTREE ,
    KEY `to` (`to`(10)) USING BTREE ,
    KEY `block_hash` (`block_hash`(10))     USING BTREE ,
    KEY `value` (`value`)           USING BTREE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS cosmos_db.t_asserts_info (
`id`                BIGINT UNSIGNED             NOT NULL AUTO_INCREMENT     COMMENT '自增主键',
`address`           CHAR(40)                    NOT NULL DEFAULT ''         COMMENT '账户地址',
`contract`          CHAR(40)                    NOT NULL DEFAULT ''         COMMENT '资产合约地址，避免重名无法辨识问题',
PRIMARY KEY (`id`),
UNIQUE KEY (`address`, `contract`),
KEY (`contract`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
