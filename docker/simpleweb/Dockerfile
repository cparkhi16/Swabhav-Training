# Specify a base image
FROM node:14-alpine
# workdir will specify a dir where in all the below commands like copy or run are to be executed , even docker exec -it <image-name> sh will also start a shell with this workdir
WORKDIR /usr/app 

# Install some depenendencies
# efficient way to copy the dependencies file to workdir and run npm install ( this will run only one time if command sequence is not changed) and after that everytime it will the temp image from cache thereby reducing time to install dependencies we will ot change everytime we run docker build
COPY ./package.json ./
RUN npm install
COPY ./ ./

# Default command
CMD ["npm", "start"]