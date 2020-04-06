https://kazuhira-r.hatenablog.com/entry/2018/09/30/164930

## s3バケット作成

```sh
aws --endpoint-url=http://localhost:4572 s3 mb s3://hoge
```

## s3 upload

```sh
aws --endpoint-url=http://localhost:4572 s3 cp hoge s3://hoge/tmp/
```

## s3 download

```sh

```

## s3 show
```sh
aws --endpoint-url=http://localhost:4572 s3 ls test-bucket
```