FROM golang:1.11
VOLUME ["/data/blog"]
COPY mblog /root
CMD ["/root/mblog"]
