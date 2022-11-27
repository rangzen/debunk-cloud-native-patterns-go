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

<!-- TOC -->
* [Debunk Cloud Native Patterns in Go](#debunk-cloud-native-patterns-in-go)
  * [Patterns](#patterns)
    * [Stability Patterns](#stability-patterns)
      * [Circuit Breaker](#circuit-breaker)
      * [Debounce -- Work In Progress](#debounce----work-in-progress)
      * [Retry -- Work In Progress](#retry----work-in-progress)
      * [Throttle -- Work In Progress](#throttle----work-in-progress)
      * [Timeout -- Work In Progress](#timeout----work-in-progress)
    * [Concurrency Patterns](#concurrency-patterns)
      * [Fan-in -- Work In Progress](#fan-in----work-in-progress)
      * [Fan-out -- Work In Progress](#fan-out----work-in-progress)
      * [Future -- Work In Progress](#future----work-in-progress)
      * [Horizontal Sharding -- Work In Progress](#horizontal-sharding----work-in-progress)
      * [Vertical Sharding -- Work In Progress](#vertical-sharding----work-in-progress)
  * [References](#references)
<!-- TOC -->

## Patterns

The structure of the repository follow the same taxonomy.

### Stability Patterns

#### Circuit Breaker

Fail fast by returning an error if the called service send back too many errors.

#### Debounce -- Work In Progress

Cache result and send it back until a new call is done after a duration (function-first),
or after a certain duration after the last call (function-last).

#### Retry -- Work In Progress

Retry a failed call waiting for a certain duration and for a certain number of time.

#### Throttle -- Work In Progress

Limit the maximum number of time per unit of time when you can call a function.

#### Timeout -- Work In Progress

Allow a process to stop waiting after a certain period.

### Concurrency Patterns

#### Fan-in -- Work In Progress

Merge input sources.

#### Fan-out -- Work In Progress

Evenly distributes messages to multiple outputs.

#### Future -- Work In Progress

Placeholder for a future value.

#### Horizontal Sharding -- Work In Progress

Split large data structure to localize effects of read/write locks across service instances.

#### Vertical Sharding -- Work In Progress

Split large data structure to localize effects of read/write locks across a unique instance.

## References

* [Cloud Native Computing Foundation (CNCF)](https://www.cncf.io/)
