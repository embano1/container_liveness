# What?
Tiny program to simulate exit &lt;> 0 failures (option -e) after some time (option -t), e.g. for desired state reconciliation testing of 
container managegement frameworks.

# Build
Modify Makefile (at least change REPO to your own repo)  
make all

# Run
docker run \<repo\>container_liveness\:\<tag_id\> or (with some arguments)  
docker run \<repo\>container_liveness\:\<tag_id\> -e 2 -t 5 // (-h for help)
