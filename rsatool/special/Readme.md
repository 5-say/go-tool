# 特殊的 RSA 工具封装

## 私钥加密，公钥解密

> 私钥加密数据、公钥进行解密 在 RSA 加密中是不常见的。\
> **这违反了 RSA 加密的基本原则。**
>
> openssl 对应的 c 函数为：\
> RSA_private_encrypt 和 RSA_public_decrypt\
> 事实上 openssl 将其的描述为 `low-level signature operations（低级别的签名操作）` **并且从 OpenSSL3.0 起不再推荐使用**

详见 [OpenSSL 官方文档](https://www.openssl.org/docs/manmaster/man3/RSA_private_encrypt.html)
参考 [stackoverflow](https://stackoverflow.com/questions/18011708/encrypt-message-with-rsa-private-key-as-in-openssls-rsa-private-encrypt)
