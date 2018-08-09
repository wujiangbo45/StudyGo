# Use an official Python runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /StudyGo

# Copy the current directory contents into the container at /app
ADD . /StudyGO

# Install any needed packages specified in requirements.txt
RUN go get github.com/golang/protobuf
RUN go get github.com/garyburd/redigo
# Make port 80 available to the world outside this container
# EXPOSE 80

# Define environment variable
ENV NAME studygo

# Run test.go when the container launches
# CMD ["go", "run", "/awesomeProject/src/test.go"]