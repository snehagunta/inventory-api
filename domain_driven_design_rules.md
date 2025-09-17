# Comprehensive Coding Guidelines for Domain-Driven Design

This document serves as the complete set of rules for writing code in this project, intended for both **human developers** and our **AI coding assistant (Cursor)**. All principles are derived directly from Vaughn Vernon's "Implementing Domain-Driven Design". Adherence to these guidelines is mandatory to ensure we build a robust, maintainable, and business-aligned application.

Our primary goal is to model the **Core Domain** of our business with precision, creating software that provides a competitive advantage.

***

## Part 1: Strategic Design Principles

Strategic design is our highest priority. Getting these patterns right prevents architectural decay and ensures our software investment is sound.

### 1. Domains, Subdomains, and Focus on the Core

* [cite_start]**Rule:** All development effort must be classified according to the subdomain it addresses: **Core Domain**, **Supporting Subdomain**, or **Generic Subdomain**[cite: 86, 565, 581]. [cite_start]The most talented developers and deepest design effort must be focused on the Core Domain[cite: 579].
* [cite_start]**Why?** The Core Domain is what provides a competitive advantage and is the most valuable software asset[cite: 579]. [cite_start]Supporting and Generic subdomains are necessary but don't require the same level of innovation and should be built or acquired with less effort[cite: 581].
* **For Cursor:** When generating new features, assess if the logic belongs to the **Core Domain**. If so, apply all tactical patterns with maximum rigor. If it's a **Supporting Subdomain**, apply DDD patterns but prefer simpler implementations. If it's a **Generic Subdomain**, consider if an off-the-shelf solution can be integrated instead of custom-building it.

### 2. Bounded Contexts

* [cite_start]**Rule:** Every model exists within an explicit **Bounded Context**[cite: 86, 356, 524]. [cite_start]A Bounded Context sets a linguistic boundary for the Ubiquitous Language[cite: 554, 620]. [cite_start]Code must be physically organized into modules/packages that reflect these distinct contexts[cite: 938].
* [cite_start]**Why?** This prevents model corruption and ambiguity by ensuring every domain term has a precise, context-specific meaning[cite: 554, 620]. It is the cornerstone of successful DDD.
* **For Cursor:** Place all new code within its designated Bounded Context package (e.g., `com.mycompany.collaboration`). Do not create classes within this context that belong to another (e.g., do not model a detailed `User` security object inside the `Collaboration Context`; instead, model a `Collaborator` role).

### 3. Context Maps

* [cite_start]**Rule:** A **Context Map** diagram must be maintained for the project, illustrating the relationships between our Bounded Contexts and external ones[cite: 88, 650]. [cite_start]All integrations must adhere to a formally defined relationship pattern (e.g., **Customer-Supplier**, **Anti-Corruption Layer**)[cite: 657].
* [cite_start]**Why?** The Context Map makes integration strategies explicit, highlights organizational dependencies, and prevents project failure due to misunderstood inter-team relationships[cite: 88, 652].
* **For Cursor:** Before generating code that integrates with another Bounded Context, reference the project's Context Map. If the map defines an **Anti-Corruption Layer (ACL)**, generate the ACL's interfaces (Facade, Adapter, Translator) to protect our Core Domain from the external model's concepts.

### 4. Architecture: Model-Infrastructure Independence

* [cite_start]**Rule:** The Domain Layer must be completely independent of infrastructure and application concerns[cite: 82, 691]. [cite_start]Use the **Dependency Inversion Principle** (as seen in Hexagonal/Ports and Adapters architecture) where infrastructure components (e.g., Repositories, messaging) implement interfaces defined in the Domain Layer[cite: 695, 1140].
* [cite_start]**Why?** This keeps the domain model pure, expressive, and focused only on business logic[cite: 82]. It allows us to swap out technical components (databases, UI frameworks) without impacting the core business rules.
* **For Cursor:** Never place infrastructure-specific code or annotations (e.g., `@Entity` from JPA, SQL queries, framework-specific dependencies) inside a domain model class. Repository implementations and messaging components belong in the `infrastructure` package and must implement interfaces defined in the `domain.model` package.

***

## Part 2: Tactical Design "Building Block" Patterns

These patterns are the implementation tools we use *inside* a Bounded Context.

### 5. Aggregates

[cite_start]Aggregates are the fundamental unit of consistency in our model[cite: 101, 171]. We will adhere strictly to the following four rules.

* [cite_start]**Rule 5.1: Model True Invariants:** The Aggregate boundary must enclose entities and value objects that represent a true business invariant, meaning a rule that must be 100% consistent within a single transaction[cite: 101, 353, 963].
    * **For Cursor:** When asked to enforce a business rule between two objects, first determine if they belong inside the same Aggregate. If a rule like "the total of order line items cannot exceed the order's credit limit" exists, `Order` and `LineItem` belong in the same Aggregate, with `Order` as the root.
* [cite_start]**Rule 5.2: Design Small Aggregates:** Aggregates should be as small as possible[cite: 101, 355]. [cite_start]Start with just a root entity and only add more entities if they are subject to a true invariant[cite: 966].
    * **For Cursor:** When creating an Aggregate, default to a single root entity with its attributes and Value Objects. Do not add other entities to the cluster unless a business invariant rule explicitly requires it.
* [cite_start]**Rule 5.3: Reference by Identity:** One Aggregate must only reference another Aggregate by its identity (e.g., `ProductId`), never by holding a direct object reference[cite: 101, 359, 970].
    * **For Cursor:** If a `BacklogItem` Aggregate needs to be associated with a `Product` Aggregate, add a `productId` property to `BacklogItem`. Do not add a `private Product product;` property. Lookups of other Aggregates must be done in Application Services or Domain Services using a Repository.
* [cite_start]**Rule 5.4: Use Eventual Consistency Outside the Boundary:** Updates to multiple aggregates in a single business process must be managed via eventual consistency[cite: 101, 364]. One transaction modifies one Aggregate, which then publishes a Domain Event. [cite_start]A separate, asynchronous process listens for that event and executes a follow-up transaction on another Aggregate[cite: 975].
    * **For Cursor:** If a user action must update both `AggregateA` and `AggregateB`, generate the code to modify `AggregateA` first. Then, have `AggregateA` publish an event (e.g., `StateChangedInA`). Generate a separate listener/handler that consumes this event and then modifies `AggregateB` in a new transaction.

### 6. Entities

* [cite_start]**Rule:** An **Entity** is defined by its thread of continuity and a stable, unique identity[cite: 93, 751]. [cite_start]Its identity must be defined upon creation and remain constant[cite: 754, 769]. [cite_start]Prefer application-generated identities (like UUIDs) over database-generated ones to ensure the identity is available before persistence[cite: 175, 755].
* [cite_start]**Why?** This models objects with a lifecycle and history, where identity matters more than attributes[cite: 752]. [cite_start]Early identity assignment prevents certain persistence bugs and allows for Domain Events to be published immediately from the constructor[cite: 765].
* **For Cursor:** When creating an Entity class (e.g., `Product`), ensure it has a mandatory identity property (e.g., `productId`). This identity should be passed into the constructor. Its `equals()` and `hashCode()` methods must be based solely on this identity.

### 7. Value Objects

* [cite_start]**Rule:** A **Value Object** measures, quantifies, or describes a domain concept[cite: 808]. [cite_start]It must be **immutable**, model a **Conceptual Whole**, be comparable via **Value Equality**, and have **Side-Effect-Free Behavior**[cite: 808]. [cite_start]Favor Value Objects over Entities wherever possible[cite: 805].
* [cite_start]**Why?** They are simple, safe, and reduce complexity by eliminating the need to track identity and lifecycle for descriptive aspects of the model[cite: 805, 806].
* **For Cursor:** When generating a class for a concept like `FullName` or `BusinessPriority`, make it a Value Object. All properties must be `final` or `readonly` and set only in the constructor. If a modification is needed, create a method like `withLastName(String newName)` that returns a *new* instance of the Value Object.

### 8. Domain Services

* [cite_start]**Rule:** Use a **Domain Service** for significant business processes that do not naturally fit on an Entity or Value Object[cite: 98, 858]. [cite_start]A Domain Service must be **stateless**[cite: 98, 860].
* [cite_start]**Why?** This prevents bloating Entities/VOs with unnatural responsibilities and avoids placing core domain logic in Application Services (which leads to an anemic model)[cite: 860].
* **For Cursor:** If a requested operation involves coordinating multiple Aggregates or depends on an external system (e.g., a currency conversion service), create a new class in the domain model named `[Concept]Service`. Ensure the service is stateless and its methods accept all necessary domain objects as parameters.

### 9. Repositories

* [cite_start]**Rule:** A **Repository** provides the illusion of an in-memory collection for a specific type of Aggregate[cite: 105, 1024]. [cite_start]Its interface must be "collection-oriented" (e.g., `add()`, `remove()`, `productOfId()`) and must not leak persistence details[cite: 1026]. [cite_start]The interface belongs in the Domain Layer; the implementation belongs in the Infrastructure Layer[cite: 105, 1035].
* [cite_start]**Why?** This decouples the domain model from data storage technology, making the model easier to test and the data store easier to change[cite: 105].
* **For Cursor:** For any given Aggregate Root (e.g., `Product`), define a `ProductRepository` interface inside the `domain.model` package. The implementation (e.g., `HibernateProductRepository`) must be located in the `infrastructure.persistence` package.

### 10. Domain Events

* [cite_start]**Rule:** A **Domain Event** is an object that represents something that happened in the domain that experts care about[cite: 99, 881]. [cite_start]Events are named in the past tense (e.g., `ProductBacklogItemCommitted`) and should be immutable Value Objects[cite: 889]. [cite_start]Aggregates publish events when their state changes[cite: 881, 897].
* [cite_start]**Why?** Events are the primary mechanism for enabling eventual consistency and decoupling Bounded Contexts[cite: 882].
* **For Cursor:** When implementing a command method on an Aggregate that changes its state, create a corresponding Domain Event class. At the end of the method, raise the event using the `DomainEventPublisher`.

### 11. Factories

* [cite_start]**Rule:** Use a **Factory** to encapsulate the logic for creating complex Aggregates or Value Objects[cite: 104, 1008]. [cite_start]This can be a **Factory Method** on an Aggregate Root (to create another related Aggregate) or a separate **Factory Service** in the domain[cite: 1008].
* [cite_start]**Why?** Factories simplify the interface for clients, protect business invariants by ensuring objects are created in a valid state, and express the Ubiquitous Language through their method names (e.g., `product.planBacklogItem(...)`)[cite: 1008].
* **For Cursor:** If the creation of an Aggregate involves complex logic, multiple steps, or requires data from another Aggregate, create a Factory Method. For example, on the `Product` Aggregate, create a method `planBacklogItem(...)` that returns a new, fully constituted `BacklogItem` instance. The `BacklogItem` constructor should be hidden (e.g., `protected`) to force clients to use the factory.

### 12. Modules

* [cite_start]**Rule:** Organize domain model objects into **Modules** (packages/namespaces) that are named according to concepts in the Ubiquitous Language, not by technical patterns[cite: 100, 939]. For example, `com.mycompany.agilepm.domain.model.product` is correct; `com.mycompany.agilepm.domain.model.aggregates` is incorrect.
* [cite_start]**Why?** Naming modules this way makes the model more expressive and tells a story about the domain[cite: 939]. It groups cohesive concepts together, improving understandability.
* **For Cursor:** When creating a new domain class, place it in a module that reflects its domain concept. For example, a `BacklogItem` class should go in the `...product.backlogitem` module, alongside its `BacklogItemId` and related Value Objects.

***

## Part 3: Application Layer

* [cite_start]**Rule:** The **Application Layer** must be kept thin[cite: 109, 1140]. [cite_start]Its methods (Application Services) are responsible only for: 1) Coordinating the retrieval of Aggregates from Repositories, 2) Invoking a single command method on one Aggregate, and 3) Managing transactions[cite: 109, 1152]. [cite_start]It **must not** contain any business logic[cite: 1152].
* **Why?** This maintains a strict separation of concerns, ensuring that all business rules reside within the Domain Layer where they can be properly modeled and protected.
* **For Cursor:** When generating an Application Service method, it should follow this template: 1. Use a Repository to find an Aggregate instance. 2. Call one method on that instance. 3. (If using a persistence-oriented repository) Save the Aggregate. Do not add conditional logic, calculations, or transformations inside the Application Service.
