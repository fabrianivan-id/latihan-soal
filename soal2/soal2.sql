SELECT
  CONCAT(FullName, ' (ID: ', EmployeeID, ') has a performance rating of: ',
    CASE
      WHEN PerformanceScore < 50 THEN 'Needs Improvement'
      WHEN PerformanceScore < 75 THEN 'Meets Expectations'
      WHEN PerformanceScore < 90 THEN 'Exceeds Expectations'
      ELSE 'Outstanding'
    END
  ) AS result
FROM EMPLOYEES
ORDER BY EmployeeID ASC;