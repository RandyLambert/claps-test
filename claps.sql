# user
SELECT * FROM claps.user;
INSERT INTO `claps`.`user` (`id`, `name`, `display_name`, `email`, `avatar_url`) VALUES ('46085959', 'RandyLambert', 'SShouXun', 'randylambert@xiyoulinux.org', 'https://avatars1.githubusercontent.com/u/46085959?v=4');
INSERT INTO `claps`.`user` (`id`, `name`, `display_name`, `email`, `avatar_url`) VALUES ('1013557', 'janily', 'janily', 'nanhuchenjie@163.com', 'https://avatars3.githubusercontent.com/u/1013557?v=4');
INSERT INTO `claps`.`user` (`id`, `name`, `display_name`, `email`, `avatar_url`) VALUES ('67439', 'lyricat', 'Lyric Wai', '5h3ll3x@gmail.com', 'https://avatars1.githubusercontent.com/u/67439?v=4');

# project   
SELECT * FROM claps.project;
INSERT INTO `claps`.`project` (`id`, `name`, `display_name`, `description`, `avatar_url`, `donations`, `total`, `created_at`) VALUES (1, 'claps.dev', 'Claps.dev', 'abc', 'abc', '0', '0', now());

# repository
SELECT * FROM claps.repository;
INSERT INTO `claps`.`repository` (`id`, `project_id`, `type`, `name`, `slug`, `description`, `created_at`, `updated_at`) VALUES ('1', '1', 'GITHUB', 'Claps.dev', 'lyricat/claps.dev', 'Help you funding the creators and projects you appreciate with crypto currencies.', now(), now());
 SELECT * FROM `transaction`  WHERE (asset_id='815b0b1a-2764-3736-8faa-42d694fa620a' AND project.id=(SELECT project.id FROM `project`  WHERE (project.name='claps.dev')));
# member
SELECT * FROM claps.member;
INSERT INTO `claps`.`member` (`project_id`, `user_id`) VALUES ('1', '46085959');
INSERT INTO `claps`.`member` (`project_id`, `user_id`) VALUES ('1', '67439');
INSERT INTO `claps`.`member` (`project_id`, `user_id`) VALUES ('1', '1013557');

# bot
SELECT * FROM claps.bot;
INSERT INTO `bot` (`id`,`project_id`,`distribution`,`session_id`,`pin`,`pin_token`,`private_key`) VALUES ('f1d0ee2d-af22-3022-aa08-de3657455ce0',1,1,'8243f4cc-adfd-4f57-b081-52c84f13398d','184475','Hlm8zGAIM66/dFrrCnJt13sc8ns26wDwu2ka6neb9yHSRpQXQz/hYsIZdJjAMxBWM6O/fB2ziTWIuSMdh78nHlQFznHnX9cCwkislAJvNvUh0XWtp7z1t+t4sgXHP4qbrPzLkFtAGu/nn50l/ThU1NItM7UbBXBsC+rMAKALy+c=','-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC+IzJ162ZrjP1wJcJqN3GWcRoGoncGo2QJlbvrPB88+Yn0nUxB
XLn6RvBNh+7BurRnZPLdcOD7FXXJwI0q2xEf2zsucFOuJBqbb6xbi+4+IzAHRFBF
cWHocDEZJZfL/4lv6jdems2DrDT0wYqpBlSSm127WP6O5Oq5Z+hwmpkUYQIDAQAB
AoGAJH3Y12zgcU/T7Ewy7fPKJxd56UARHAML1hMYx+L4E9nIslvmIL1NLE2lHRwz
pJbSvO1Q91MMuuO5gYklDs6QkHnuUIlQphbvhik5j4bPbVDXOgFQNrIXBoHLcUdl
9lL8eie/FGjcwjq9uw+F02/6l7z9bJQB/xskGzf5fpb4TYECQQDZFJR9KsTHDrbB
HsruTpkmNv6JTu0tdY341m6J6MvL5CZwzifBbvSYai/sVuJj2955RoL9RC+/WPZ7
i0Og4Vd5AkEA4DoCNJdHE74pyQ7py4VDsgWiwEB6xx/Vlc2jlu+6nasgHMebuEBq
C4eRO5m5km7G8qUCBwHUJ0EdkusbIwQiKQJBAKudhIKrpCOGc16bnGznwFWg1nvw
5Lqym8mkpIDshOks9mLp6C4ZLM+t6zMZwSKW+PvBjd7x4BmTGFG1WILAg2kCQFUt
tCjGTvnxA26de7MUrOKzwV/HHt0F+t0tgTeVWg8LMue77CvSTHaUyVcazqQR8QG8
LUj8KNvAoLtvFJ/4sgECQQDS0oZ7u+mdpqh4ZcLmyzp2doF620GUZBsmsyi9EulX
PvgExXfS3llfTqxyzoqKXcU6VkMTunMwDpJuu9Lxz0BM
-----END RSA PRIVATE KEY-----
');  
 INSERT INTO `bot` (`id`,`project_id`,`distribution`,`session_id`,`pin`,`pin_token`,`private_key`) VALUES ('aac43e6e-0a28-35b8-b3de-c9ff84256451',1,2,'f6e13083-e1a6-480f-8f41-4e9b22d55008','184475','cZW3Ggd5YGjb3Q7RW3RC6sAVtko1jLcS8ylZY7G9m/P2esQZfbusDWA8612TjOLUWBGZWTExFtGSBAaZ0h7XIDT4HEHm7QA17+Wti+1OX48EQ/0AE774tebDNLiH/NzDIPzrnKUIU8zby4qNjv/kmSMSZZELgeJlDy7oL0/k0k0=','-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDIzP27eHLAig6fJSyCy9VE0Q6DQvnPsGcArAybN5+ynGQoDpd3
cYd8QL7OmGsHVf8Ikl4cz3yrrBnx1XSNX7k5oa0MI6UWlLFLV+ri15zlkWtBRVR8
CHG8V66KaoYxW+gbbqKa0NSjPAc0XblSg5mkE6hJ0wUulWyiyMFjDpx//QIDAQAB
AoGAAhPM1DGszj0fZZoW2cuOC0Y2Zjk9KF7k0eb1wm1S46AmkRuFiaDNDAYHc0+0
W8ESAF6zRo0G9yeypQPWTtgcrG1CgWszKzj/cIb3vZge9V5q6YhDjZs4v2sN2Ml+
4fEUT1002JAJy7V4z5bFWOssuAZJkq753I2c0vqgU6j/7QECQQDVFW/lnzNqapf4
3yORx+N2LdX0oFwcXDVDIGIKB4dX0ukkfT0m6bihwtGftMLasmbDiwwiumlGGC94
CvYFktshAkEA8T4+tqCJPOe6xOREO85niKHf69txL+o954XpMQtWNRLDdp8IwG/r
omCuapiwsJd+odLeVLYuedqlXQ6r99NFXQJBAJcrFC1VGkbexF38/+EGbCqFLgrU
UUSVbfvnV2ZCHRSDPn9ykhWvLhskeU7SEILSmfEUDlH86X6e3d5N+GfP3cECQQDq
pKBbAcp6cuo2l8/GW/xX6RrjTY3KDQwpJRarnVs8RAPaXNUmV7XZOjBrfhhdqvyA
aZnWy1xpKUGuQZcCdXwBAkBLs/osW+BFCr5E4Z2mrjFhpbnMHZuTBUzZP/Cr0aUK
bl6cPLjEPWyBbuNS9gixiI339NnxJaEibVyJ1V8S+Nlh
-----END RSA PRIVATE KEY-----
'); 
INSERT INTO `bot` (`id`,`project_id`,`distribution`,`session_id`,`pin`,`pin_token`,`private_key`) VALUES ('dd10b157-9d0f-31e5-bfa3-39c442f3a9b8',1,3,'fb808202-7b10-48e6-805a-9bb5d3ad1fdd','184475','Q9iFjzvUkoVUTaO2qPlSL/Dg0Ua/yMz/tnfBqVVCnPd5kTJVgGN8ka/dZawSgAYSiMpIo93D/o5cfT9TchbgT4Bdw6e21a6M8D9hj6eQT7udK/0xek3l3DxTNESTIfpmgA7UlHDJ89RvMSiOP80Xw/rxQA+vsGBkK0MDqTdms48=','-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDSBB8uHWWXVTmvDcUTjkQkL6XPgMCm6/dxePFKuodF0fNvrNc6
3MsMzBk4cyqhEIjmJ5IIvKCtbJ8CigC+9KniPlPaAv86XfPHnd9NdQr7sxambtXF
QK3qlu6nXDIQ/fvFm+YWEaCr89xzSBhkxYkaGkhMWGYABNgp3ZawJwn3MQIDAQAB
AoGBAINpcNk6K8d13JJc22RRMPIznl0pA2NvY3XtZ46LCPn3VYwbatG6NpPbYiyg
Y5xE7GSXfhlZbnEV9qlwEOdr0KBNR/Ws9oPh5Vr1mvpJrh8rtO8nqWKr366y/WXC
mVBqAlwdC1ziLrSr5rKp0f+I33FNFM84Mif7IQykccb8GVbBAkEA/smr9+LycIzb
V5o5H9+Zsi1xpXc5Bq8oNXtORAVDvWFXjTj4a3FDI7lA5zLxcUl0kbeLh8welcJw
zejj5Ny5dwJBANMD6zDj14GrqXv4RSV5CViVIwdX5Bc4oBqobl6ZJ+ziTrKjKlmA
q94/+Ne6V7GLsaQYUUOoXM5qrL3I5QrwfpcCQQD5HpkNjBI+qAsDMaEvIAL4a2SW
Q+c3OOYYvNK+wWMFdXsUcyK6cwkRkd368R2QBiF7JLrB8XvqNC71tgO1z3drAkA3
YIzlbLXOyu0UoqgK2IPSYnkp4S/zxCGIPXGRk+H9cbqzeMyRZoo0Llewza9b4cxB
wzv4ZIPOjAI/YCxzvX+LAkEAsTQ5uF0vIk9CB5BqiU3HKlI30L9Oe+mB6JID/lkU
KlZ2hKKNZUwDMHTvcp9q5XhHnZcuJS1aBo1ueNGj2lVZcQ==
-----END RSA PRIVATE KEY-----
');
INSERT INTO `bot` (`id`,`project_id`,`distribution`,`session_id`,`pin`,`pin_token`,`private_key`) VALUES ('1ccd828e-7639-3c60-90f8-c133414bf210',1,4,'2cf27900-9d44-48d2-9532-13f000ff2803','184475','v2/oU4Dc3L5ts67YsvgMNdTE174dztBDHtR4sPRIAg5f7AlgJ5aUgbKSIPKIm5Bq4zKro+pHgEO6dP78XUVygQijAJdipffmeq93rkeH+GXU79SQ5UftWmuUilE5s/x0Vu+NTU1o5ksc7kcjPZLLYCjBYY3HlPr0CKXrkgaH73I=','-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDdtA+9E/kmPZqH0VxGseXYB8j/PLROoK/OiIaTLIsZR7nKJdlP
GSRMuFJpqQRIKc/ma2G3OyOAjiqPWowYtX9izKSDX2fdzn7e2DmxfT0Vb8ynfk10
Q5UyewgmG2+QyhJ3cm5oOLF1tcG4MKVMXlVJE4NDxed6qYdFQo6NBR6wgwIDAQAB
AoGBAMOdWXmqQt5T4qJNvs59ruAy1k0mYa0yqIxh9+OpnL3chHhxdtEMzPOIeubR
g36srcyQGLLUUlnelnzQFubCVbbRUQF3ZORV7CQBBAxE3SEQikEcwa0zAed7tJrr
L+lCztvVfEmaaCZ8IHVwzEA8Ehp1K1GVC0CdRhkQZM7OCUR5AkEA9rlTdbIjzYCV
CGdlZW6zttHCw21F91iucJXcJLdJfLa6/onH1req2TKHM8gmE3Y0GS9IWwLt+gMx
7Vi50dkUNQJBAOYJ6r7rfJY6/4hJjxpNI+YdCqhHNtEpQEaw0vefrMPOZYGhRp6B
PdjC1yV9H+FqSXTB6e2/LnXRcWKv47f62NcCQQCtxrrSGza8d+SAltMELoTGL9hO
bZjqLrwu8F6uPaq0/L+YqNLomVTsxnoULtUiwA7R7ku4TsfvYHC6C9RVyeBBAkAu
KvVbBeYGPKcGUkUPMUYwK8n0xf4hclb9GQXuPrSsw6KpppWGwEeKVmhZlMguNGez
sCtj1MfdS4CnHsfkJ8sXAkEAiXm7Cjx0kJ9QC2b0mjZvShrUWSeqAhOR0rLF18M7
2Fvz3jM9TZH+Snr4lyOrbmXDR24fKwQUBEm+g0HrzOihGA==
-----END RSA PRIVATE KEY-----
');

# wallet 
SELECT * FROM claps.wallet;
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('1ccd828e-7639-3c60-90f8-c133414bf210', '815b0b1a-2764-3736-8faa-42d694fa620a', '1', '0', now(), now(), now());
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '815b0b1a-2764-3736-8faa-42d694fa620a', '1', '0', now(), now(), now());
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('aac43e6e-0a28-35b8-b3de-c9ff84256451', '815b0b1a-2764-3736-8faa-42d694fa620a', '1', '0', now(), now(),now());
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('f1d0ee2d-af22-3022-aa08-de3657455ce0', '815b0b1a-2764-3736-8faa-42d694fa620a', '1', '0', now(), now(), now());
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('1ccd828e-7639-3c60-90f8-c133414bf210', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '1', '0', '2020-07-24 17:07:59', '2020-07-24 17:07:59', '2020-07-24 17:07:59');
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('aac43e6e-0a28-35b8-b3de-c9ff84256451', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '1', '0', '2020-07-24 17:07:59', '2020-07-24 17:07:59', '2020-07-24 17:07:59');
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '1', '0', '2020-07-24 17:07:59', '2020-07-24 17:07:59', '2020-07-24 17:07:59');
INSERT INTO `claps`.`wallet` (`bot_id`, `asset_id`, `project_id`, `total`, `created_at`, `updated_at`, `synced_at`) VALUES ('f1d0ee2d-af22-3022-aa08-de3657455ce0', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '1', '0', '2020-07-24 17:07:59', '2020-07-24 17:07:59', '2020-07-24 17:07:59');

# member_wallet
SELECT * FROM claps.member_wallet;
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'f1d0ee2d-af22-3022-aa08-de3657455ce0', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', '1ccd828e-7639-3c60-90f8-c133414bf210', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'aac43e6e-0a28-35b8-b3de-c9ff84256451', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', '1ccd828e-7639-3c60-90f8-c133414bf210', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'aac43e6e-0a28-35b8-b3de-c9ff84256451', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085959', 'f1d0ee2d-af22-3022-aa08-de3657455ce0', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');

INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'f1d0ee2d-af22-3022-aa08-de3657455ce0', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', '1ccd828e-7639-3c60-90f8-c133414bf210', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'aac43e6e-0a28-35b8-b3de-c9ff84256451', '815b0b1a-2764-3736-8faa-42d694fa620a', now(), now(), '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', '1ccd828e-7639-3c60-90f8-c133414bf210', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'aac43e6e-0a28-35b8-b3de-c9ff84256451', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'dd10b157-9d0f-31e5-bfa3-39c442f3a9b8', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');
INSERT INTO `claps`.`member_wallet` (`project_id`, `user_id`, `bot_id`, `asset_id`, `created_at`, `updated_at`, `total`, `balance`) VALUES ('1', '46085963', 'f1d0ee2d-af22-3022-aa08-de3657455ce0', '6770a1e5-6086-44d5-b60f-545f9d9e8ffd', '2020-07-24 17:20:10', '2020-07-24 17:20:10', '0', '0');


# transaction
SELECT * FROM claps.transaction;
SELECT id,donations,total FROM `project`  WHERE (id=(SELECT project_id FROM `bot`  WHERE (id='1ccd828e-7639-3c60-90f8-c133414bf210'))); 
SELECT project_id FROM `bot`  WHERE (id='1ccd828e-7639-3c60-90f8-c133414bf210');

 SELECT * FROM `user`  WHERE (user.id IN ((SELECT user_id FROM `member`  WHERE (project_id=1)))) ;