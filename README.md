## Banking-Kata

Think of your personal bank account experience. If any doubt, go for the simplest solution

Requirements
------------

Deposit and Withdrawal  
Transfer  
Account statement (date, amount, balance)  
Statement printing  
Statement filters (just deposits, withdrawal, date)

The Rules
---------
#### For more information:

-  [Object Calisthenics pdf](http://www.cs.helsinki.fi/u/luontola/tdd-2009/ext/ObjectCalisthenics.pdf)
-  Object Calisthenics (full book), Jeff Bay in: The ThoughtWorks Anthology.
Pragmatic Bookshelf 2008
-  Original idea for the kata: How Object-Oriented Are You Feeling Today? - Krzysztof Jelski (Session on the Software Craftsmanship UK 2011 conference)


### Solution

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

Do it yourself first and then compare the solutions. 
