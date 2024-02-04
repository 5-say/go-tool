# rsatool

RSA（Rivest-Shamir-Adleman）是一种非对称加密算法，使用两个密钥：公钥和私钥。

## RSA 加解密原理

**密钥生成：**

> 随机生成两个大素数 p 和 q。\
> 计算 n = p * q，并计算欧拉函数 φ(n) = (p - 1) * (q - 1)。\
> 选择一个整数 e，1 < e < φ(n)，且 e 与 φ(n) 互质，作为公钥的指数。\
> 计算 d，使得 (d * e) mod φ(n) = 1，作为私钥的指数。\
> 公钥为 (e, n)，私钥为 (d, n)。

**加密：**

> 将明文消息转换为整数 m，满足 0 <= m < n。\
> 加密操作：c = (m^e) mod n，其中 c 为密文。

**解密：**

> 使用私钥进行解密操作：m = (c^d) mod n，其中 m 为明文。

**RSA 加解密的安全性基于两个数学难题：**

> 质因数分解问题：在已知一个大数的情况下，找到其所有的质因数。\
> 模反演问题：在已知一个数的情况下，找到其模意义下的乘法逆元。

由于这两个问题的计算复杂性，使得 RSA 算法在当前计算能力下被认为是安全的。\
但是，**随着计算能力的提升和新的算法的出现，RSA算法的安全性可能会受到威胁。**\
因此，**在实际应用中，需要选择足够大的密钥长度，并定期更新密钥，以保证安全性。**

## 密钥格式区别

PKCS#8密钥格式，多用于JAVA、PHP程序加解密中，为目前用的比较多的密钥、证书格式，也用于微信RSA等程序；
PKCS#1密钥格式，多用于JS等其它程序加解密，属于比较老的格式标准。。

RSA PKCS#1公钥格式如：
开头 -----BEGIN RSA PUBLIC KEY-----
结尾 -----END RSA PUBLIC KEY-----

RSA PKCS#1私钥格式如：
开头 -----BEGIN RSA PRIVATE KEY-----
结尾 -----END RSA PRIVATE KEY-----

RSA PKCS#8公钥格式如：
开头 -----BEGIN PUBLIC KEY-----
结尾 -----END PUBLIC KEY-----

RSA PKCS#8私钥格式如：
开头 -----BEGIN PRIVATE KEY-----
结尾 -----END PRIVATE KEY-----

## 在线工具

- [在线 Base64 编解码](https://base64.us/)
- [在线 RSA 公钥导出](http://www.metools.info/code/c87.html)
- [在线 RSA 公私钥对生成](http://www.metools.info/code/c80.html)
- [在线 RSA 公私钥对校验](http://www.metools.info/code/c83.html)
- [在线 RSA PKCS#1、PKCS#8 私钥格式转换工具](http://www.metools.info/code/c84.html)
- [在线 RSA PKCS#1、PKCS#8 公钥格式转换工具](http://www.metools.info/code/c85.html)
