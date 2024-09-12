```markdown
# Distributed Systems Course

## Course Schedule

| **Date**      | **Lecture Topic**       | **Details**                                                                                                                                                                                |
|---------------|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Tue 8/20**  | Introduction & MapReduce | - Slides <br>- Lab Setup, Homework 1, and Lab 2 Released (Lab 1 will not be done) <br> **Optional Reading:** <br> - [MapReduce paper](#) <br> - Coulouris §21.6.1 |
| **Thu 8/22**  | Networking & RPC         | - Slides <br> **Required:** Complete the [Online Go Tutorial](#) before class <br> **Optional Reading:** <br> - Coulouris §§3.4, 5.2-5.3 <br> - van Steen §§4.0-4.2, 8.3 |
| **Tue 8/27**  | RPC                      | - Homework 1 Due                                                                                                                                                                            |
| **Thu 8/29**  | Consistency: Linearizability | - Slides <br> **Required Reading:** [Testing Distributed Systems for Linearizability](#) <br> **Optional Reading:** [Linearizability](#) |
| **Fri Aug 30**| Drop/Add Deadline        |                                                                                                                                                                                            |
| **Tue 9/3**   | Time & Clocks            | - Slides <br> **Required Reading:** [Time, Clocks](#) (Read at least the first 5 pages) <br> **Optional Reading:** <br> - Coulouris §§14.1-14.4 <br> - van Steen §§6.1-6.2 (Slides follow van Steen closely) |
| **Thu 9/5**   | Logical Time, State Machine Replication, and Primary-Backup | - Slides <br> **Optional Reading:** <br> - Coulouris §§18.3.1-18.3.2 <br> - van Steen §7.5 (Remote-write protocols) |
| **Tue 9/10**  | Paxos                    | - Slides <br> - Lab 2 Due <br> **Optional Reading:** <br> - [Paxos Made Simple](#) <br> - Watch [Ousterhout's Paxos Video](#) <br> - van Steen §§8.1-8.2 |
| **Thu 9/12**  | Raft                     | - **Required:** Watch [Ousterhout's Raft Video](#) (Stop at 48:00) <br> - Read [Raft (Extended Version)](#) (Through Section 5) |
| **Tue 9/17**  | -                        |                                                                                                                                                                                            |
| **Thu 9/19**  | -                        |                                                                                                                                                                                            |
| **Tue 9/24**  | -                        |                                                                                                                                                                                            |
| **Thu 9/26**  | -                        |                                                                                                                                                                                            |
| **Tue 10/1**  | -                        | - Lab 3a Due                                                                                                                                                                                |
| **Thu 10/3**  | -                        |                                                                                                                                                                                            |
| **Tue 10/8**  | No Class - Fall Break    |                                                                                                                                                                                            |
| **Thu 10/10** | No Class - Fall Break    |                                                                                                                                                                                            |
| **Tue 10/15** | -                        |                                                                                                                                                                                            |
| **Thu 10/16** | -                        |                                                                                                                                                                                            |
| **Tue 10/22** | -                        | - Lab 3b Due                                                                                                                                                                                |
| **Thu 10/24** | -                        |                                                                                                                                                                                            |
| **Tue 10/29** | -                        |                                                                                                                                                                                            |
| **Thu 10/31** | -                        |                                                                                                                                                                                            |
| **Tue 11/5**  | -                        | - Lab 3c Due                                                                                                                                                                                |
| **Thu 11/7**  | -                        |                                                                                                                                                                                            |
| **Tue 11/12** | -                        |                                                                                                                                                                                            |
| **Thu 11/14** | -                        |                                                                                                                                                                                            |
| **Tue 11/19** | -                        |                                                                                                                                                                                            |
| **Thu 11/21** | -                        |                                                                                                                                                                                            |
| **Tue 11/26** | -                        |                                                                                                                                                                                            |
| **Thu 11/28** | No Class - Thanksgiving Break |                                                                                                                                                                                            |
| **Tue 12/3**  | -                        |                                                                                                                                                                                            |
| **Thu 12/5**  | -                        |                                                                                                                                                                                            |
| **Fri 12/10** | -                        |                                                                                                                                                                                            |

## Course Content Overview

In recent decades, large-scale distributed systems have revolutionized communication, data storage, and analysis. These systems will need to grow in size, speed, and reliability as online connectivity increases. This class introduces key concepts for designing and building distributed systems, including RPC, replication, consistency, fault tolerance, and transactions.

**Labs** will focus on developing a fault-tolerant, consensus-replicated key-value store. The course will also include lectures on research topics and a final presentation based on academic papers.

## Course Objectives

By the end of the course, students should be able to:

1. Explain modern distributed systems' mechanisms (e.g., RPC, consensus, leases, concurrency control).
2. Understand various large-scale distributed infrastructures (e.g., metadata stores, filesystems, databases).
3. Implement fault-tolerant systems using RPC and distributed algorithms.
4. Discuss and communicate technical topics in distributed systems.

## Topics Covered

- **Fundamentals** (RPC, concurrency, networking)
- **Time & Logical Clocks** (Vector Clocks, Eventual Consistency)
- **Replication and Consensus** (Paxos, Raft)
- **Consistency** (Linearizability, Sequential Consistency)
- **Transactions** (2PC, MVCC, Serializability)
- **Performance & Scalability** (Modeling, Measuring)
- **Byzantine Fault Tolerance** (PBFT, Bitcoin Consensus)

## Course Information

- **Instructor**: Ryan Stutsman
- **TAs**: Chloe Pronovost
- **Lecture Time**: Tuesdays & Thursdays, 15:40-17:00
- **Office Hours**:
    - **Ryan**: Thursdays 12:00-14:00 (In-person & MS Teams)
    - **Chloe**: Mondays 16:40-17:40, Tuesdays 17:30-18:30 (In-person & MS Teams)
- **Credits**: 3 Credit Hours

**Prerequisites**: CS 5460, CS 4480, or similar courses with substantial programming components.

## Textbooks

No official textbook. Helpful resources:

- *Distributed Systems: Concepts and Design* by Coulouris et al.
- *Distributed Systems* by van Steen & Tanenbaum (available online).

```

Feel free to adjust links and text as necessary for your specific content and audience.
