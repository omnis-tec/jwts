# Jwts

generate key files:
```
ssh-keygen -t rsa -b 4096 -m PEM -f private.pem
openssl rsa -in private.pem -pubout -outform PEM -out public.pem
rm private.pem.pub
```
