FROM alpine

# Step 2: Download and install dependency

RUN apk add --update redis
RUN apk add --update gcc 
# first we run without gcc (so entire image and temp container flow was done)
# after adding gcc line (since first 2 cmds were already done , a cached image was stored on local machine and that cache image is itself used )
# changing order of above 2 RUN cmds will result into again temp image and container flow as it will not be found on local machine cache

# Step 3: Tell the image what to do when it starts as container

CMD ["redis-server"]


# # Manual image generation with docker commit # #
# we can manually create a image from a running container using docker commit
# for redis server image we can use :
# docker run -it alpine sh and run this cmd in shell (apk add --update redis)
# use docker ps to get current running container id and use this cmd to provide startup command to this container
# docker commit -c 'CMD ["redis-server"]' 0bed5c617a45
# now u will get a  image id and can run docker run image-id to run this image

# Tagging an image
# docker build -it dockerhub-id/projectname:version .

# Rebuilds with cache
# docker always checks if there is a image in its cache with same set of instructions done earlier and if present then uses it instead of going through temp container image flow
