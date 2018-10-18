FROM golang:1.11
VOLUME ["/data/blog"]
COPY mblog /root
EXPOSE 443
CMD ["/root/mblog"]
