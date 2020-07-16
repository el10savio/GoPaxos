# GoPaxos

Paxos Made Simple, Implemented On Docker Containers

## Introduction

Paxos is a consensus algorithm used to establish consensus among several nodes in a distributed system. Here, we use Docker to containerize a Paxos cluster node, and Golang to implement the Paxos made simple protocol among every node. Using Paxos the cluster forms a very simple distributed Key-Value Store enabling the user to write and read data across any node in the cluster.

## Steps

After cloning the repo. To provision the cluster:

```
$ make provision
```    

This creates a 3 node Paxos cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

Now we can send requests to Set and Get Key-Values to any peer node using its port allocated.

```
$ curl -i localhost:<peer-port>/store/set/<key>/<value>
$ curl -i localhost:<peer-port>/store/get/<key>
```

In the logs for each peer docker container, we can see the logs of the Paxos transaction taking place.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a docker rmi to delete them.

##  Paxos

The Paxos consensus algorithm is implemented using Golang running as a Paxos server in each node. Paxos consists of 3 phases:

- **Prepare Phase**:  This is the start of the Paxos phase enabled when a client would like to write data to the cluster. Here, the Prepare process generates a round ID of its own and propagates it to all the nodes in the cluster. Once a majority of nodes accept the prepare message it then moves to the accept phase.


- **Accept Phase**:  Here the same leader node that transmitted the prepare message sends an accept request to all the nodes again to accept the given value to be chosen, thus achieving consensus among a majority of nodes and at times all the nodes.


- **Learn Phase**: Once the above two phases are complete the leader then sends a learn request which enables all the nodes to persist the agreed-upon value to its store.

## Docker

Docker enables each Paxos node to be isolated and can run anywhere. Docker network here establishes a network across all the nodes so that each node can communicate with each other and ingress/egress with the host machine.

## Simple Paxos vs Multi-Paxos

The current implementation of Paxos here is Paxos Made Simple protocol, which in a real-world production environment would fare much better. Future improvements to GoPaxos would look at upgrading the protocol to Multi-Paxos. Multi-Paxos works by running multiple Paxos rounds across the nodes, auto leader election, log replication to handle failure scenarios, and several other improvements.

## References

 - [Paxos Made Simple](https://lamport.azurewebsites.net/pubs/paxos-simple.pdf) [Leslie Lamport]
 - [Paxos lecture (Raft user study)](https://youtu.be/JEpsBg0AO6o) [Diego Ongaro & John Ousterhout]