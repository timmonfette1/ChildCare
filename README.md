Child Care Problem
===================

Description
------------
Classic Child Care Problem.

Implemented in Golang as an example of Go Synchronization.

The premise of the problem is a Daycare center where adults look over children.
An adult is allowed by state law to only oversee at most 3 children.
A child can only enter the daycare if there are enough adults inside to look after them.
An adult can only leave the daycare if there won't be more children inside than state law allows.

In other worlds, a child cannot enter unless the current number of children inside is less than
adults * 3.

An adult cannot leave unless the current number of children inside is less than
(adults - 1) * 3.
