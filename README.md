# hashpack

ハッシュ値計算ツール

## 概要

ファイルもしくは標準入力の内容のハッシュ値を求める。
各種アルゴリズムに対応し、出力長を指定可能である。


## 使用法

```sh
$ hashpack [-t ハッシュアルゴリズム] [-l 出力文字列長] [ファイル名]
```

ファイル名が指定されない場合、標準入力を入力とする。

## 対応アルゴリズム

| 指定 | 概要 |
|:----:|:-----|
| crc32 | crc32_ieeeの省略形 |
| crc32_ieee | IEEE 802.3で使用されているCRC-32 |
| crc32_castagnoli | Castagnoliによる多項式を使用したCRC-32 |
| crc32_koopman | Koopmanによる多項式を使用したCRC-32 |
| crc64 | crc64_isoの省略形 |
| crc64_iso | ISO 3309で使用されているCRC-64 |
| crc64_ecma | ECMA 182で使用されているCRC-64 |
| md5 | Message Digest Algorithm 5 |
| sha2 | sha256の省略形 |
| sha224 | SHA-224 (SHA-2) |
| sha256 | SHA-256 (SHA-2) |
| sha384 | SHA-384 (SHA-2) |
| sha512 | SHA-512 (SHA-2) |
| sha512_224 | SHA-512/224 (SHA-2) |
| sha512_256 | SHA-512/256 (SHA-2) |
| sha3 | sha3_512の省略形 |
| sha3_224 | SHA3-224 |
| sha3_256 | SHA3-256 |
| sha3_384 | SHA3-384 |
| sha3_512 | SHA3-512 |
