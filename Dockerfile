FROM golang:1.20 as build
WORKDIR /yoku
COPY * /yoku
RUN go build -o /bin/yoku ./main.go

FROM scratch
COPY --from=build /bin/yoku /bin/yoku
CMD [ "/bin/yoku" ]
