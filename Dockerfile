FROM alpine
EXPOSE 80/tcp
ADD cart-service /cart-service
ENTRYPOINT [ "/cart-service" ]
