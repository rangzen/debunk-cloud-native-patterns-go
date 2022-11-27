# Debunk Cloud Native Patterns in Go

As always, after some times, and some success at launching
[Cloud Native](https://github.com/cncf/toc/blob/main/DEFINITION.md) products
engineers identified some patterns.
These patterns can seem hard to grasp the first time,
but I hope examples of this project will help you to understand them
and when and how to use them.

A huge part of this content was inspired by the book
[Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/)
by [Matthew A. Titmus](https://www.linkedin.com/in/matthew-titmus/)
so if you want to go deeper, please consider buying it.
The book is very well written.

Note on the source, if you consider participating or fixing code, please keep it stupid simple.
E.g., I try to not use the Chain of Responsibility pattern in the codebase when it is possible.
This pattern is very common to chain HTTP Func or in middlewares
like in the [Gin](https://gin-gonic.com/) or the [Echo](https://echo.labstack.com/) frameworks,
but it can make the code less understandable at the first reading for beginners. 

* [Patterns](#patterns)
  * [Stability Patterns](#stability-patterns)
    * [Circuit Breaker](#circuit-breaker)
    * [Debounce](#debounce)
    * [Retry](#retry)
    * [Throttle](#throttle)
    * [Timeout](#timeout)
  * [Concurrency Patterns](#concurrency-patterns)
    * [Fan-in](#fan-in)
    * [Fan-out](#fan-out)
    * [Future](#future)
    * [Horizontal Sharding](#horizontal-sharding)
    * [Vertical Sharding](#vertical-sharding)
* [References](#references)

## Patterns

The structure of the repository follow the same taxonomy.

### Stability Patterns

#### Circuit Breaker

Fail fast by returning an error if the called service send back too many errors.

#### Debounce

Work in Progress.

#### Retry

Work in Progress.

#### Throttle

Work in Progress.

#### Timeout

Work in Progress.

### Concurrency Patterns

#### Fan-in

Work in Progress.

#### Fan-out

Work in Progress.

#### Future

Work in Progress.

#### Horizontal Sharding

Work in Progress.

#### Vertical Sharding

Work in Progress.

## References

* [Cloud Native Computing Foundation (CNCF)](https://www.cncf.io/)
