## Distribiuted File System
In today's rapidly expanding technological landscape, distributed systems are gaining prominence. They represent a complex and extensive area of study within computer science. At its core, a distributed system is a collection of computers working together to present themselves as a unified entity to end-users. These machines share a state, operate concurrently, and can experience independent failures without disrupting the entire system's uptime.

In this project, we aim to develop a simple distributed file system capable of reading and writing mp4 files while ensuring fault tolerance through file replication.

## Architecture:
The Distributed File System (DFS) comprises two types of machine nodes: the Master Tracker node and the Data Keeper nodes. The Master Tracker maintains a lookup table containing information about files, including their names, associated Data Keeper nodes, file paths on those nodes, and the status of each data node. The Data Keeper nodes store the actual data files. Both the Master Tracker and Data Keeper nodes are designed to be multi-threaded to handle concurrent requests efficiently.

## Communication:
All communication between the Master Tracker, Data Keepers, and Clients occurs over gRPC. File transfers take place over TCP connections.

## Heartbeats:
To ensure the system's health and availability, each Data Keeper node sends a keepalive ping to the Master Tracker every second. The Master Tracker updates the lookup table accordingly. If a Data Keeper node becomes unresponsive, the corresponding entry in the lookup table is updated to reflect its status.

## Uploading A File:
The protocol for a client to upload a file to the cluster follows these steps:

1. The client communicates with the Master Tracker node.

2. The Master Tracker responds with the port number of one of the available Data Keeper nodes.

3. The client establishes a communication channel with the specified Data Keeper node and transfers the file.

4. Upon completion of the transfer, the Data Keeper node notifies the Master Tracker.

5. The Master Tracker updates the lookup table with the file record.

6. The Master Tracker notifies the client of the successful upload.

7. The Master Tracker selects two other nodes to replicate the uploaded file.

## Replication:
To maintain fault tolerance and data redundancy, a separate thread on the Master Tracker periodically checks for file replication. The replication algorithm ensures that each file exists on at least three active Data Keeper nodes, mitigating the risk of data loss.

## Downloading A File:
Clients can request to download mp4 files from the Distributed File System following these steps:

1. The client sends a download request to the Master Tracker, specifying the desired file name.
2. The Master Tracker responds with a list of IP addresses and ports from which the file can be downloaded.
3. The client initiates parallel requests to download the file from each provided port, ensuring efficient data retrieval.


## Contributoes:
[Donia Gameel](https://github.com/DoniaGameel)

[Heba Ashraf](https://github.com/hebaashraf21)

[Shaza Mohammed](https://github.com/ShazaMohamed)
