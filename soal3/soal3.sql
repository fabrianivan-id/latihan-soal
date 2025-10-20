WITH usage_cte AS (
  SELECT
    r.account_id,
    r.tariff_id,
    COUNT(*) AS cnt,
    ROW_NUMBER() OVER (
      PARTITION BY r.account_id
      ORDER BY COUNT(*) DESC, r.tariff_id ASC
    ) AS rn
  FROM readings r
  GROUP BY r.account_id, r.tariff_id
),
agg AS (
  SELECT
    r.account_id,
    SUM(r.amount) AS total_consumption,
    SUM(r.amount * t.cost) AS total_cost,
    COUNT(*) AS total_readings
  FROM readings r
  JOIN tarrifs t ON t.id = r.tariff_id
  GROUP BY r.account_id
)
SELECT
  a.username,
  a.email,
  t.name AS most_frequent_tariff,
  agg.total_consumption,
  ROUND(agg.total_cost / NULLIF(agg.total_readings,0), 2) AS average_cost_per_reading
FROM accounts a
LEFT JOIN agg ON agg.account_id = a.id
LEFT JOIN usage_cte u ON u.account_id = a.id AND u.rn = 1
LEFT JOIN tarrifs t ON t.id = u.tariff_id
ORDER BY a.username ASC;