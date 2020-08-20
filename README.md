[![Actions Status](https://xxx.execute-api.us-west-2.amazonaws.com/production/badge/{adolsalamanca}/{banking-kata})](https://xxx.execute-api.us-west-2.amazonaws.com/production/results/{adolsalamanca}/{banking-kata})

## Banking-Kata

Think of your personal bank account experience. If any doubt, go for the simplest solution

Requirements
------------

Go installed and GOPATH properly configured

Code: <br/>
Account will expose only 3 methods: <br/>
 - Deposit, Withdrawal and PrintStatement <br/>
 - Show your TDD skills! <br/>
 - Go simple, small methods and files <br/>
 
 
Started from defining an acceptance test:

> Given a client makes a deposit of 1000 on 10-01-2012  
And a deposit of 2000 on 13-01-2012  
And a withdrawal of 500 on 14-01-2012  
If we print bank statement  

This will be the output  
date       || credit   || debit    || balance  
14/01/2012 ||          || 500.00   || 2500.00   
13/01/2012 || 2000.00  ||          || 3000.00  
10/01/2012 || 1000.00  ||          || 1000.00   

<br/>Do it yourself first and then compare the solutions. <br/>
 
http://kata-log.rocks/banking-kata<br/>



Test my solution
------------

To test my solution just clone the repo and run go test.
