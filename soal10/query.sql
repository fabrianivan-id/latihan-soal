SELECT c.iban,
       b.amount,
       COALESCE(t.cnt, 0) AS transaction_count
FROM balances b
JOIN customers c ON c.id = b.customer_id
LEFT JOIN (
  SELECT customer_id, COUNT(*) AS cnt
  FROM transactions
  GROUP BY customer_id
) t ON t.customer_id = b.customer_id
WHERE b.amount < 0
ORDER BY b.amount ASC;
