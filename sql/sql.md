## 中間テーブル順

```sql
SELECT *, (SELECT COUNT(post_id) FROM post_hashtag WHERE hashtag_id = id) AS rank, (SELECT COUNT(review_id) FROM review_hashtag WHERE hashtag_id = id) AS foo FROM hashtag WHERE id IN (SELECT id FROM category_hashtag WHERE category_id = 1) ORDER BY rank + foo;
```

## 緯度経度半径検索

```sql
Select("*, (6371 * acos(cos(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* cos(radians(lat))* cos(radians(lng) - radians((SELECT lng FROM tourist_spot WHERE id = ?)))+ sin(radians((SELECT lat FROM tourist_spot WHERE id = ?)))* sin(radians(lat)))) AS distance", query.ID, query.ID, query.ID).Having("distance <= ?", 5).Order("distance")
```