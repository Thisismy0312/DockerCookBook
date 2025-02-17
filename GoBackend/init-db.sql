-- 创建 users 表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL UNIQUE,
    user_pwd VARCHAR(255) NOT NULL
    );

-- 授予 admin 用户对 users_id_seq 的权限
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO admin;
