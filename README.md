<!-- 
# Distributed Systems

## Course Schedule

| **Date**      | **Lecture**                     | **Details** |
|---------------|---------------------------------|-------------|
| Tue 8/20      | Introduction & MapReduce        | [Slides](https://utah.instructure.com/courses/985754/files/166481860?wrap=1), [MapReduce Paper](https://www.cs.utah.edu/~stutsman/cs6450/public/papers/mapreduce.pdf), Lab Setup, Homework 1, and Lab 2 released |
| Thu 8/22      | Networking & RPC                | [Slides](https://utah.instructure.com/courses/985754/files/166526051?wrap=1), [Go Tutorial](http://tour.golang.org/) |
| Tue 8/27      | RPC                             | Homework 1 due |
| Thu 8/29      | Consistency: Linearizability    | [Slides](https://utah.instructure.com/courses/985754/files/166991214?wrap=1), [Testing Distributed Systems for Linearizability](https://anishathalye.com/testing-distributed-systems-for-linearizability/) |
| Tue 9/3       | Time & Clocks                   | [Slides](https://utah.instructure.com/courses/985754/files/167145361?wrap=1), [Time, Clocks Paper](https://www.cs.utah.edu/~stutsman/cs6450/public/papers/time.pdf) |
| Thu 9/5       | Logical Time, State Machine Replication, and Primary-Backup | [Slides](https://utah.instructure.com/courses/985754/files/167216489?wrap=1) |
| Tue 9/10      | Paxos                           | [Slides](https://utah.instructure.com/courses/985754/files/167477801?wrap=1), Lab 2 Due, [Paxos Made Simple](https://www.cs.utah.edu/~stutsman/cs6450/public/papers/paxos-simple.pdf) |
| Thu 9/12      | Raft                            | [Ousterhout's Raft Video](https://www.youtube.com/watch?v=YbZ3zDzDnrw) (Stop at 48:00), [Raft Paper](https://raft.github.io/raft.pdf) |

...

## Course Content Overview

In the last few decades, large-scale distributed systems have revolutionized the way we live and work. They form the basis for ourcommunications platforms. They provide our only means of finding, aggregating, storing, and analyzing the massive data we collect eachday. These systems will only need to grow larger, faster, and more reliable as more and more of our lives and devices are online.  

Building distributed systems, especially those that operate at large-scale with high performance, presents special challenges. For example,scale introduces faults and the need for redundancy. Redundancy complicates consistency.  

This class introduces many of the key aspects of designing and building distributed systems such RPC, naming, routing, replication,consistency, fault tolerance, transactions, and time.
The first half of the class is driven in large part by lectures often extracted from real systems described in papers from the systems researchcommunity. Students are expected to attend lectures and take short comprehension quizzes and take home exams. Later lectures will fold insome current research topics like kernel-bypass.  

However, the bulk of the work of the class involves a progressive series of labs where students develop and debug a fault-toleranceconsensus-replicated key-value store. Students will also complete a final presentation based on academic papers.  

## Course Objectives

Students that successfully complete this course should be able to:  

* explain the underlying standard mechanisms that modern distributed systems use (RPC, consensus, leases, concurrency control).
* explain the various types of systems that comprise many large-scale distributed infrastructures (metadata stores, batch computeframeworks, filesystems, databases, caching, etc).
* use RPC and specifications of distributed algorithms to implement fault-tolerant, available systems.
* approach, discuss, and communicate about difficult and technical subject matters in the area of distributed systems.
-->
## Topic List

This list of topics serve as a road map for the class. Depending on time some topics will change, and we likely won't (quite) be able to covereverything listed here.

* Fundamentals
* Introduction
* Networking, Concurrency, and RPC
* Soft/hard State, Caching
* Time & Logical Clocks
* Vector Clocks
* Eventual Consistency & Scaling
* Distributed Snapshots
* Decentralized Systems and Distributed Hash Tables
* Dynamo
* Replicated State Machines
* Primary/Backup
* View Change & Consensus
* Paxos & Raft
* Consistency
* Strong Consistency (Linerizability, Sequential Consistency)
* Weak Consistency (Causal Consistency, Eventual Consistency)
* Transactions
* Serializability
* Atomic Commit and 2PC
* Concurrency Control: 2PL, OCC, MVCC
* Spanner
* Performance
* Modeling & Measurement
* Byzantine Fault Tolerance
* PBFT
* Bitcoin & Nakamoto Consensus

## Textbooks

- *Distributed Systems: Concepts and Design* by Coulouris et al.
- *Distributed Systems* by van Steen and Tanenbaum
