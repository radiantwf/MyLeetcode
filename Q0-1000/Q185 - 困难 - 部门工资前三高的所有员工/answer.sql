WITH Salary AS (
    SELECT
        id,
        name,
        departmentId,
        salary,
        DENSE_RANK() OVER (PARTITION BY departmentId ORDER BY salary DESC) AS rank
    FROM
        employee
)
SELECT
    t2.name AS Department,
    t1.name AS Employee,
    t1.salary AS Salary
FROM
    Salary t1
INNER JOIN
    Department t2
ON
    t1.departmentId = t2.id
WHERE
    t1.rank <= 3;