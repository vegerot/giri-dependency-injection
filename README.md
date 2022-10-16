### Dependency Injection / User Repository

#### What is dependency injection?


> Dependency injection is a programming technique that makes a class independent of its dependencies. 
It achieves that by decoupling the usage of an object from its creation. 
This helps you to follow SOLID’s dependency inversion and single responsibility principles, 
in order to write a good programme.

#### SOLID Principles

> Single Responsibility Principle: By creating small interfaces, we define obvious 
> responsibilities for implementing classes. It makes it easier to follow the SRP, 
> especially when we make our classes implement only a handful or even a single interface.

> Open/Closed Principle: With loose coupling and hidden implementations, following OCP is 
> also more straightforward. Since the client code doesn’t rely on the implementation, 
> we can introduce additional subclasses as needed.

> Liskov Substitution Principle: LSP is not directly connected to this technique. 
> However, we must take care when we’re designing our inheritance hierarchy to follow this 
> principle, too.

> Interface Segregation Principle: ISP isn’t a result but a good practice to follow when 
> we’re programming interfaces. Note that we already talked about the importance of defining 
> small, well-defined responsibilities. Those notes were hidden hints to follow ISP.

> Dependency Inversion Principle: By relying on abstractions, we already did the majority 
> of the work to follow DIP. The last thing to do is to expect dependencies from an external 
> party instead of instantiating them internally.


#### Uses gorm.io Library for PostgreSQL

#### Verify SQL Connection

```bash
psql -P expanded=auto -h postgres-instance.company.com --username=testuser --dbname=testdb
```

#### Output

```bash
[~/git/goworkspace/src/di]$ go run *.go
```

### Verify SQL Table

```sql
testdb=> select * from user_records;
 userid |      name
--------+-----------------
    100 | Jordan Peterson
(1 row)
```