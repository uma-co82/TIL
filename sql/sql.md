## 中間テーブル順

```sql
SELECT *, (SELECT COUNT(post_id) FROM post_hashtag WHERE hashtag_id = id) AS rank, (SELECT COUNT(review_id) FROM review_hashtag WHERE hashtag_id = id) AS foo FROM hashtag WHERE id IN (SELECT id FROM category_hashtag WHERE category_id = 1) ORDER BY rank + foo;
```

## 緯度経度半径検索

```sql
Select("*, (6371 * acos(cos(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* cos(radians(lat))* cos(radians(lng) - radians((SELECT lng FROM tourist_spot WHERE id = ?)))+ sin(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* sin(radians(lat)))) AS distance", query.ID, query.ID, query.ID).Having("distance <= ?", 5).Order("distance")
```

## 全文インデックス

https://dev.mysql.com/doc/refman/5.6/ja/innodb-fulltext-index.html

```sql
 FULLTEXT KEY(name) WITH PARSER NGRAM
```

```sql
WHERE (MATCH(name) AGAINST('keyward'))
```

## sqlクエリ頑張りすぎるのどうなん？
```
あんまsqlに任せすぎると業務ロジックがそっちに行くしでちょっと微妙なところだよね。
```

## 中間テーブルのupdated_at順

```sql
SELECT * FROM post INNER JOIN (SELECT post_id, updated_at FROM user_favorite_post WHERE user_id = 1) hoge ON post.id = hoge.post_id ORDER BY hoge.`updated_at` DESC;
```


```sql
SELECT f.*, u.*
FROM user_favorite_post f
JOIN user u
ON f.user_id = u.id AND u.deleted_at IS NULL
LEFT JOIN user_following uf
ON u.id = uf.target_id and uf.user_id=1
WHERE f.`post_id`=1
ORDER BY uf.created_at DESC, f.created_at DESC;
```

## followしてるかどうか

```sql
LEFT JOIN user_following ON user.id = user_following.target_id AND user_following.user_id = ?
```

## 支援金額

```sql
SELECT * FROM cf_project JOIN (SELECT b.cf_project_id, SUM(a.price * b.amount) AS amount_achieved FROM cf_return_gift_snapshot a LEFT JOIN payment_cf_return_gift b ON a.id = b.cf_return_gift_snapshot_id GROUP BY b.cf_project_id) t1 ON cf_project.id = t1.cf_project_id;
```

```sql
Joins("(SELECT cf_project_id, MAX(id) AS id FROM post GROUP BY cf_project_id) p ON cf_project.id = p.cf_project_id").
		Where("IFNULL(cf_project.latest_post_id, 0) < p.id AND cf_project.id > ?", lastID).
```

## 支援したproject

```sql
SELECT * FROM cf_project INNER JOIN (SELECT cf_project_id, MAX(created_at) AS created_at FROM payment_cf_return_gift WHERE payment_id IN (SELECT id FROM payment WHERE user_id = 483) GROUP BY cf_project_id) pc ON cf_project.id = pc.cf_project_id ORDER BY pc.created_at DESC;
```