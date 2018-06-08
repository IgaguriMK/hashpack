# hashpack

ハッシュ値計算ツール

## 概要

ファイルもしくは標準入力の内容のハッシュ値を求める。
各種アルゴリズムに対応し、出力長を指定可能である。


## 使用法

```sh
$ hashpack [-t ハッシュアルゴリズム] [-l 出力長] [ファイル名]
```

ファイル名が指定されない場合、標準入力を入力とする。

## 対応アルゴリズム

| 指定 | 概要 |
|:----:|:-----|
| crc32 | IEEE 802.3 CRC-32 |
| crc64 | ISO 3309 CRC-64 |
| md5 | Message Digest Algorithm 5 |
| sha224 | 224bit SHA-2 |
| sha256 | 256bit SHA-2 |
| sha384 | 384bit SHA-2 |
| sha512 | 512bit SHA-2 |
| sha3 | 512bit SHA-3 |
