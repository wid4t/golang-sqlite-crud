### How to run this program
To run this program, please type this command.

    > go run .
### Test API
#### Create

    > curl -X POST -H "Content-Type:application/json" -d "{\"name\":\"Mr.X\", \"jobName\":\"java developer\"}" http://localhost:3000/employee
    > curl -X POST -H "Content-Type:application/json" -d "{\"name\":\"Mr.Y\", \"jobName\":\"golang developer\"}" http://localhost:3000/employee
#### Read 
    
    > curl http://localhost:3000/employees
#### Update

    > curl -X PUT -H "Content-Type:application/json" -d "{\"name\":\"Mr.Y\", \"jobName\":\"senior golang developer\"}" http://localhost:3000/employee/2
#### Delete

    > curl -X DELETE http://localhost:3000/employee/2