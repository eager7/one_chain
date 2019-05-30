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

# 此表单容纳几种不同交易，通过action类型区分，不同交易的from和to不一致，并且有的还需要附属信息，下面列出不同交易的from，to类型和前缀
# action=send; from=sender(cosmos); to=recipient(cosmos); appendix=nil
# action=delegate; from=delegator_address(cosmos); to=validator_address(cosmosvaloper); appendix=nil
# action=edit_validator; from=val_address(cosmosvaloper); to=nil; appendix=Description[moniker,identity,website,details]
# action=begin_unbonding; from=delegator_address(cosmos); to=validator_address(cosmosvaloper); appendix=nil
# action=set_withdraw_address; from=delegator_address(cosmos); to=withdraw_address(cosmos); appendix=nil
# action=withdraw_delegator_reward; from=delegator_address(cosmos); to=validator_address(cosmosvaloper); appendix=nil
# action=withdraw_validator_rewards_all; from=validator_address(cosmosvaloper); to=nil; appendix=nil
# action=unjail; from=address(cosmos); to=nil; appendix=nil
# action=vote; from=voter(cosmos); to=nil; appendix=proposal_id+option
# action=deposit; from=depositor(cosmos); to=nil; appendix=proposal_id+amount
# action=submit_proposal; from=proposer(cosmos);initial_deposit=amount; to=nil; appendix=title+description+proposer
CREATE TABLE IF NOT EXISTS cosmos_db.t_transaction_info (
    `id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `action`            VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易类型',
    `hash`              CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易哈希值',
    `from`              CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易发起者',
    `to`                CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易接收者',
    `value`             VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易金额',
    `memo`              TEXT                NOT NULL                COMMENT '交易备注信息',
    `appendix`          TEXT                NOT NULL                COMMENT '交易附属信息',
    `timestamp`         INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '交易发生的时间点',
    `gas_wanted`        BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易提供的gas量',
    `gas_used`          BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易使用的gas量',
    `gas_price`         VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '当前交易gas单价',
    `status`            BIGINT UNSIGNED     NOT NULL DEFAULT 1      COMMENT '标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码',
    `block_hash`        CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易所在区块的哈希',
    `block_number`      BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易所在区块的高度',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash`(`hash`),
    KEY `from` (`from`) USING BTREE ,
    KEY `to` (`to`) USING BTREE,
    KEY `number` (`block_number`) USING BTREE ,
    KEY `block_hash` (`block_hash`)USING BTREE
)ENGINE = INNODB DEFAULT CHARSET =utf8mb4;

# 多重发送需要将input和output拆开，从而形成多条记录，通过hash来组装成一笔交易
CREATE TABLE IF NOT EXISTS cosmos_db.t_multisend_info (
    `id`                BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `action`            VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易类型',
    `hash`              CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易哈希值',
    `sender`            CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易发起者',
    `recipient`         CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易接收者',
    `value`             VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易金额',
    `memo`              TEXT                NOT NULL                COMMENT '交易备注信息',
    `timestamp`         INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '交易发生的时间点',
    `gas_wanted`        BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易提供的gas量',
    `gas_used`          BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易使用的gas量',
    `gas_price`         VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '当前交易gas单价',
    `status`            BIGINT UNSIGNED     NOT NULL DEFAULT 1      COMMENT '标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码',
    `block_hash`        CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易所在区块的哈希',
    `block_number`      BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易所在区块的高度',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash_sender`(`hash`, `sender`),
    KEY `sender` (`sender`) USING BTREE ,
    KEY `recipient` (`recipient`) USING BTREE,
    KEY `number` (`block_number`) USING BTREE ,
    KEY `block_hash` (`block_hash`)USING BTREE
)ENGINE = INNODB DEFAULT CHARSET =utf8mb4;

# 重新分配抵押资源相当于两笔交易，即撤销和抵押，不过是在一个交易内完成，因此无法容纳到通用交易格式中，只能单独列出，因为三个地址都需要索引
CREATE TABLE IF NOT EXISTS cosmos_db.t_begin_redelegate_info (
    `id`                    BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `action`                VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易类型',
    `hash`                  CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易哈希值',
    `delegator_address`     CHAR(40)            NOT NULL DEFAULT ''     COMMENT '当前交易发起者(cosmos)',
    `validator_src_address` CHAR(40)            NOT NULL DEFAULT ''     COMMENT '原有验证节点(cosmosvaloper)',
    `validator_dst_address` CHAR(40)            NOT NULL DEFAULT ''     COMMENT '新的验证节点(cosmosvaloper)',
    `value`                 VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易金额',
    `memo`                  TEXT                NOT NULL                COMMENT '交易备注信息',
    `timestamp`             INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '交易发生的时间点',
    `gas_wanted`            BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易提供的gas量',
    `gas_used`              BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易使用的gas量',
    `gas_price`             VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '当前交易gas单价',
    `status`                BIGINT UNSIGNED     NOT NULL DEFAULT 1      COMMENT '标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码',
    `block_hash`            CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易所在区块的哈希',
    `block_number`          BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易所在区块的高度',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash`(`hash`),
    KEY `delegator_address` (`delegator_address`) USING BTREE ,
    KEY `validator_src_address` (`validator_src_address`) USING BTREE,
    KEY `validator_dst_address` (`validator_dst_address`) USING BTREE,
    KEY `number` (`block_number`) USING BTREE ,
    KEY `block_hash` (`block_hash`)USING BTREE
)ENGINE = INNODB DEFAULT CHARSET =utf8mb4;

# 创建验证者交易单独存储，做为验证者资料
CREATE TABLE IF NOT EXISTS cosmos_db.t_create_validator_info (
    `id`                    BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `action`                VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易类型',
    `hash`                  CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易哈希值',
    `delegator_address`     CHAR(40)            NOT NULL DEFAULT ''     COMMENT '抵押账户地址(cosmos)',
    `validator_address`     CHAR(40)            NOT NULL DEFAULT ''     COMMENT '验证账户地址(cosmosvaloper)',
    `pubkey`                CHAR(70)            NOT NULL DEFAULT ''     COMMENT '验证节点公钥(cosmosvalconspub)',
    `value`                 VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '交易金额',
    `description`           TEXT                NOT NULL                COMMENT '验证节点描述信息',
    `commission`            TEXT                NOT NULL                COMMENT '验证节点描述信息',
    `min_self_delegation`   VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '验证节点自身最低抵押量',
    `memo`                  TEXT                NOT NULL                COMMENT '交易备注信息',
    `timestamp`             INT UNSIGNED        NOT NULL DEFAULT 0      COMMENT '交易发生的时间点',
    `gas_wanted`            BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易提供的gas量',
    `gas_used`              BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易使用的gas量',
    `gas_price`             VARCHAR(64)         NOT NULL DEFAULT ''     COMMENT '当前交易gas单价',
    `status`                BIGINT UNSIGNED     NOT NULL DEFAULT 1      COMMENT '标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码',
    `block_hash`            CHAR(64)            NOT NULL DEFAULT ''     COMMENT '当前交易所在区块的哈希',
    `block_number`          BIGINT UNSIGNED     NOT NULL DEFAULT 0      COMMENT '当前交易所在区块的高度',
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash`(`hash`),
    KEY `delegator_address` (`delegator_address`) USING BTREE ,
    KEY `validator_address` (`validator_address`) USING BTREE,
    KEY `number` (`block_number`) USING BTREE ,
    KEY `block_hash` (`block_hash`)USING BTREE
)ENGINE = INNODB DEFAULT CHARSET =utf8mb4;
