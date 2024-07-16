CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    -- Write your PostgreSQL query statement below.
    WITH Salary AS (
    SELECT
        t.salary,
        DENSE_RANK() OVER (ORDER BY t.salary DESC) AS rank
    FROM
        Employee t
    )
    SELECT distinct(t.salary) AS "getNthHighestSalary(2)"
    FROM Salary t
    WHERE t.rank = N
  );
END;
$$ LANGUAGE plpgsql;