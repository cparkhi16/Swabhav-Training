FROM node:16-alpine as builder
WORKDIR '/app'
COPY ./package.json ./
#RUN npm install
RUN npm install -g npm@9.6.2
RUN npm install react-scripts
COPY . .
RUN npm run build

FROM nginx
EXPOSE 3000
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html

# nginx default.conf settings for client :
# The listen 3000; directive specifies the port on which the Nginx server should listen for incoming connections.
# The location / {...} block is used to define a location for serving content. In this case, the location is the root directory /usr/share/nginx/html. The index directive is used to specify the default files to serve when a directory is requested. In this case, Nginx will look for index.html and index.htm files in the specified root directory.
# The try_files directive is used to handle requests that do not match an existing file. The $uri variable represents the requested URI, and the directive tries to serve the requested file, then the directory, and finally, if none of those exist, it serves the index.html file.
# Overall, this configuration ensures that any requests to the server on port 3000 will be served content from the specified root directory, with appropriate defaults and fallbacks.