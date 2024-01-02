package main

/*
...........................................Design Splitwise................................
1. SPLIT EQUALLY
	Splitting the amount equally into all members of a group
	Example
	4000 split equally among 8 persons equals 500 as due amount on each person

2. SPLIT EXACTLY
	Exact Split distributes the amount to the exact given amount.
	Example
	4000 divided exactly into two persons as 2500, 1500.
	First person owes 2500 and the second 1500 to the creditor

Terminologies
SPLIT BY PERCENTAGE
Split the given amount as per the given percentage values
Example
4000 divided among 4 in the percentage 10%, 20%, 20% and 50% would
mean the first user owe an amount of 400, second and third would owe
800 each and the fourth person would owe a total of 2000 to the creditor.


Application Requirements
1. Every user who uses the app should be registered. Only registered users can be involved
in expense metrics.
2. Expenses created should be of type EXACT, EQUAL or PERCENT only.
➢ In case of EXACT and PERCENT, verify if the figures match the given total amount
3. At any point of time, the application can print the total sum owed and individual amounts
owed to and by each user.
4. The final amount should be rounded off to two decimal places. And the total should add
up to the principal amount
*/
