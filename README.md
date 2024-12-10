## Conflict Manager

Table of contents
•	Introduction
•	Prerequisites
•	Project folders structure
•	Installation guide
•	Compiling code

## Introduction
• This module shall be responsible for identifying potentially overlapping or conflicting
RAN control decisions from multiple xApps and provide guidance to them.
• E.g., the user needs to be moved from one cell to another due to high traffic load,
while the other xApp at the same time may want to move the user back as the
handover boundary change due to a high handover failure rate.
• In such a case the individual user will be „ping-ponged” between the two cells if we
don’t have a conflict mitigation function that looks at what’s happening in the network
and what impact the potential action may have on the network operation if being
executed in the E2-Node. In such an example the conflict mitigation function’s job is
to align the potential actions to avoid such undesired behavior.

## Prerequisites
Make sure that following tools are properly installed and configured
	GO (golang) development and runtime tools
	Module's Grpc Service Up 
	RMR (ricplt/lib/rmr)
	Healthy kubernetes cluster
	Access to the common docker registry.

## Project folder structure
/main:contains go project's main file
/conflictCache:contains funtions to Check E2 guidance and add New Guidance Resources and  data structures
/procedure:contains handling of requests

## Installation guide 
Compiling code
Enter the project root and execute make container. This will compile the code and generate the docker image
