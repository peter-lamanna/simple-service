# requires statically linked go binary to be compiled
# to ./magalix/simple-service before docker build
FROM golang:latest

# executable folder
#RUN mkdir /magalix
#COPY simple-service /magalix
RUN mkdir /peter-lamanna
COPY simple-service /peter-lamanna

# run micro-service after boot up
#ENTRYPOINT ["/magalix/simple-service"]
ENTRYPOINT ["/peter-lamanna/simple-service"]

#expose port 8080
EXPOSE 8080